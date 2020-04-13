package repository

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Microsoft/ApplicationInsights-Go/appinsights"
	talentRequestModel "github.com/Ubidy/Ubidy_EmployerNotificationAPI/domain/talentrequest"
	stream "gopkg.in/GetStream/stream-go2.v1"
)

// TalentRequest implementation of talentrequest repository
type TalentRequest struct {
	client          *stream.Client
	agencyDB        *sql.DB
	auctionDB       *sql.DB
	telemetryClient appinsights.TelemetryClient
}

// NewTalentRequestRepository create new instance of talentrequest repository
func NewTalentRequestRepository(client *stream.Client, agencyDB *sql.DB, auctionDB *sql.DB, telemetryClient appinsights.TelemetryClient) *TalentRequest {
	return &TalentRequest{
		client:          client,
		agencyDB:        agencyDB,
		auctionDB:       auctionDB,
		telemetryClient: telemetryClient,
	}
}

// AddTalentRequestCancelActivity create an activity when the employer cancel a talentRequest for a successfee
func (tr *TalentRequest) AddTalentRequestCancelActivity(talentRequest talentRequestModel.SuccessFeeTalentRequest) (stream.Activity, error) {
	var userIDs []string
	assignmentRepository, settingsRepository := tr.createAssignmentAndSettingsRepository()
	tenantIDs := strings.Split(talentRequest.TalentRequest.AgencyTenantIDs[0], ",")

	for _, agencyTenantID := range tenantIDs {
		assignments, _ := assignmentRepository.GetAgencyAssignmentsBySuccessFeeID(talentRequest.SuccessFeeID, agencyTenantID)
		if assignmentRepository.IsApprovedSuccessFeeStatus(talentRequest.SuccessFeeID) {
			for _, assignment := range assignments {
				userID := strings.ToLower(assignment.UserID)
				settings, _ := settingsRepository.GetSettingsBySupplierID(userID)
				if settings.Fulfillment {
					userIDs = append(userIDs, userID)
				}
			}
		}
	}
	msg := "%s cancelled TR #%s: %s"
	tr.sendNotificationToAgencies(talentRequest.TalentRequest, userIDs, talentRequest.SuccessFeeID, talentRequest.SuccessFeeNumber, "successFee", msg)
	return stream.Activity{}, nil
}

func (tr *TalentRequest) createAssignmentAndSettingsRepository() (*Assignment, *Settings) {
	settingsRepository := NewSettingsRepository(nil, tr.agencyDB, tr.telemetryClient)
	assignmentRepository := NewAssignmentRepository(tr.auctionDB, tr.telemetryClient)

	return assignmentRepository, settingsRepository
}

func (tr *TalentRequest) sendNotificationToAgencies(talentrequest talentRequestModel.TalentRequest, userIDs []string, id string, number string, t string, msg string) error {
	helper := NewHelper(tr.telemetryClient)

	employerData, _ := helper.createJSONMarshal(talentrequest.EmployerUserID, talentrequest.EmployerName, "client")
	talentRequestData, _ := helper.createJSONMarshal(talentrequest.TalentRequestID, talentrequest.TalentRequestNumber, "talentRequest")
	data, _ := helper.createJSONMarshal(id, number, t)

	for _, userID := range userIDs {
		log.Printf("%s", userID)
		agencyNotificationFeed := tr.client.NotificationFeed("agencynotification", userID)
		_, err := agencyNotificationFeed.AddActivity(stream.Activity{
			Actor:     agencyNotificationFeed.ID(),
			Verb:      "create",
			Object:    fmt.Sprintf("%s:%s", t, id),
			ForeignID: agencyNotificationFeed.ID(),
			Time:      stream.Time{time.Now().UTC()},
			Extra: map[string]interface{}{
				"employer":         fmt.Sprintf("employer:%s", talentrequest.EmployerUserID),
				"agency":           fmt.Sprintf("agency:%s", userID),
				"employerTenantID": talentrequest.EmployerTenantID,
				"content":          fmt.Sprintf(msg, employerData, talentRequestData, talentrequest.JobTitle),
				"category":         "Fulfillment",
				"subcategory": map[string]string{
					"type":   "Talent Request",
					"status": "Cancelled",
					"data":   fmt.Sprintf("%s", data),
				},
			},
		})
		if err != nil {
			tr.telemetryClient.TrackException(err)
			return err
		}
	}

	return nil
}
