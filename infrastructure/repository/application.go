package repository

import (
	"fmt"
	"net/url"
	"time"

	"database/sql"
	"strings"

	"github.com/Microsoft/ApplicationInsights-Go/appinsights"
	stream "gopkg.in/GetStream/stream-go2.v1"
)

const (
	candidateSubmission3DaysIdle            = "/api/v1.0.0/agency/candidatesubmission/3daysidle"
	candidateSubmission10DaysIdle           = "/api/v1.0.0/agency/candidatesubmission/10daysidle"
	candidateSubmission14DaysIdle           = "/api/v1.0.0/agency/candidatesubmission/14aysidle"
	candidateSubmission3DaysIdleSuccessFee  = "/api/v1.0.0/agency/candidatesubmission/3daysidle/successfee"
	candidateSubmission10DaysIdleSuccessFee = "/api/v1.0.0/agency/candidatesubmission/10daysidle/successfee"
	candidateSubmission14DaysIdleSuccessFee = "/api/v1.0.0/agency/candidatesubmission/14daysidle/successfee"
)

// Application implementation of application repository
type Application struct {
	client          *stream.Client
	agencyDB        *sql.DB
	auctionDB       *sql.DB
	telemetryClient appinsights.TelemetryClient
}

// NewApplicationRepository create new instance of application repository
func NewApplicationRepository(client *stream.Client, agencyDB *sql.DB, auctionDB *sql.DB, telemetryClient appinsights.TelemetryClient) *Application {
	return &Application{
		client:          client,
		agencyDB:        agencyDB,
		auctionDB:       auctionDB,
		telemetryClient: telemetryClient,
	}
}

// AddApprovedApplicationActivity  create an activity for approved agency
func (a *Application) AddApprovedApplicationActivity(clientID string, clientName string, supplierID string, supplierName string, auctionID string, auctionNumber string, employerTenantID string, agencyTenantID string) (stream.Activity, error) {
	helper := NewHelper(a.telemetryClient)

	clientData, _ := helper.createJSONMarshal(clientID, clientName, "client")
	auctionData, _ := helper.createJSONMarshal(auctionID, auctionNumber, "auction")

	assignmentRepository, settingsRepository := a.createAssignmentAndSettingsRepository()
	assignments, _ := assignmentRepository.GetAgencyAssignmentsByAuctionID(auctionID, agencyTenantID)

	if assignmentRepository.IsApprovedAuctionStatus(auctionID) {

		for _, assignment := range assignments {
			userID := strings.ToLower(assignment.UserID)
			settings, _ := settingsRepository.GetSettingsBySupplierID(userID)

			if settings.Applications {
				verb := "approve"
				object := fmt.Sprintf("auction:%s", auctionID)
				content := fmt.Sprintf("%s approved your application to Engagement #%s", clientData, auctionData)
				category := "Application"
				subcategory := map[string]string{
					"type":   "Application",
					"status": "Approved",
				}

				a.sendNotification(userID, verb, object, clientID, employerTenantID, content, category, subcategory)
			}
		}
	}

	return stream.Activity{}, nil
}

// AddDeclinedApplicationActivity create an activity for declined activity
func (a *Application) AddDeclinedApplicationActivity(clientID string, clientName string, supplierID string, supplierName string, auctionID string, auctionNumber string, employerTenantID string, agencyTenantID string) (stream.Activity, error) {
	helper := NewHelper(a.telemetryClient)

	clientData, _ := helper.createJSONMarshal(clientID, clientName, "client")
	auctionData, _ := helper.createJSONMarshal(auctionID, auctionNumber, "auction")

	assignmentRepository, settingsRepository := a.createAssignmentAndSettingsRepository()
	assignments, _ := assignmentRepository.GetAgencyAssignmentsByAuctionID(auctionID, agencyTenantID)

	if assignmentRepository.IsApprovedAuctionStatus(auctionID) {
		for _, assignment := range assignments {
			userID := strings.ToLower(assignment.UserID)
			settings, _ := settingsRepository.GetSettingsBySupplierID(userID)

			if settings.Applications {
				verb := "decline"
				object := fmt.Sprintf("auction:%s", auctionID)
				content := fmt.Sprintf("%s declined your application to Engagement #%s", clientData, auctionData)
				category := "Application"
				subcategory := map[string]string{
					"type":   "Application",
					"status": "Declined",
				}

				a.sendNotification(userID, verb, object, clientID, employerTenantID, content, category, subcategory)
			}
		}
	}

	return stream.Activity{}, nil
}

// AddRevokedApplicationActivity create an activity for revoked application
func (a *Application) AddRevokedApplicationActivity(clientID string, clientName string, supplierID string, supplierName string, auctionID string, auctionNumber string, employerTenantID string, agencyTenantID string) (stream.Activity, error) {
	helper := NewHelper(a.telemetryClient)

	clientData, _ := helper.createJSONMarshal(clientID, clientName, "client")
	auctionData, _ := helper.createJSONMarshal(auctionID, auctionNumber, "auction")

	assignmentRepository, settingsRepository := a.createAssignmentAndSettingsRepository()
	assignments, _ := assignmentRepository.GetAgencyAssignmentsByAuctionID(auctionID, agencyTenantID)

	if assignmentRepository.IsApprovedAuctionStatus(auctionID) {
		for _, assignment := range assignments {
			userID := strings.ToLower(assignment.UserID)
			settings, _ := settingsRepository.GetSettingsBySupplierID(userID)

			if settings.Applications {
				verb := "revoke"
				object := fmt.Sprintf("auction:%s", auctionID)
				content := fmt.Sprintf("Please note, %s revoked your application approval to their Auction #%s (Competitive).", clientData, auctionData)
				category := "Application"
				subcategory := map[string]string{
					"type":   "Application",
					"status": "Revoked",
				}

				a.sendNotification(userID, verb, object, clientID, employerTenantID, content, category, subcategory)
			}
		}
	}

	return stream.Activity{}, nil
}

// AddApprovedApplicationSuccessFeeActivity create an activity for approved agency
func (a *Application) AddApprovedApplicationSuccessFeeActivity(clientID string, clientName string, supplierID string, supplierName string, successFeeID string, successFeeNumber string, employerTenantID string, agencyTenantID string, token string) (stream.Activity, error) {
	helper := NewHelper(a.telemetryClient)

	clientData, _ := helper.createJSONMarshal(clientID, clientName, "client")
	successFeeData, _ := helper.createJSONMarshal(successFeeID, successFeeNumber, "successFee")

	assignmentRepository, settingsRepository := a.createAssignmentAndSettingsRepository()
	assignments, _ := assignmentRepository.GetAgencyAssignmentsBySuccessFeeID(successFeeID, agencyTenantID)

	if assignmentRepository.IsApprovedSuccessFeeStatus(successFeeID) {
		for _, assignment := range assignments {
			userID := strings.ToLower(assignment.UserID)
			settings, _ := settingsRepository.GetSettingsBySupplierID(userID)

			if settings.Applications {
				verb := "approve"
				object := fmt.Sprintf("successfee:%s", successFeeID)
				content := fmt.Sprintf("%s approved your application to Engagement #%s", clientData, successFeeData)
				category := "Application"
				subcategory := map[string]string{
					"type":   "Application",
					"status": "Approved",
				}

				a.sendNotification(userID, verb, object, clientID, employerTenantID, content, category, subcategory)
			}
		}

		formData := url.Values{
			"clientID":         {clientID},
			"clientName":       {clientName},
			"supplierID":       {supplierID},
			"supplierName":     {supplierName},
			"successFeeID":     {successFeeID},
			"successFeeNumber": {successFeeNumber},
			"agencyTenantID":   {agencyTenantID},
			"employerTenantID": {employerTenantID},
		}
		helper.post(candidateSubmission3DaysIdleSuccessFee, formData, token)
		helper.post(candidateSubmission10DaysIdleSuccessFee, formData, token)
		helper.post(candidateSubmission14DaysIdleSuccessFee, formData, token)
	}

	return stream.Activity{}, nil
}

// AddDeclinedApplicationSuccessFeeActivity create an activity for declined activity
func (a *Application) AddDeclinedApplicationSuccessFeeActivity(clientID string, clientName string, supplierID string, supplierName string, successFeeID string, successFeeNumber string, employerTenantID string, agencyTenantID string) (stream.Activity, error) {
	helper := NewHelper(a.telemetryClient)

	clientData, _ := helper.createJSONMarshal(clientID, clientName, "client")
	successFeeData, _ := helper.createJSONMarshal(successFeeID, successFeeNumber, "successFee")

	assignmentRepository, settingsRepository := a.createAssignmentAndSettingsRepository()
	assignments, _ := assignmentRepository.GetAgencyAssignmentsBySuccessFeeID(successFeeID, agencyTenantID)

	if assignmentRepository.IsApprovedSuccessFeeStatus(successFeeID) {
		for _, assignment := range assignments {
			userID := strings.ToLower(assignment.UserID)
			settings, _ := settingsRepository.GetSettingsBySupplierID(userID)

			if settings.Applications {
				verb := "decline"
				object := fmt.Sprintf("successfee:%s", successFeeID)
				content := fmt.Sprintf("%s declined your application to Engagement #%s", clientData, successFeeData)
				category := "Application"
				subcategory := map[string]string{
					"type":   "Application",
					"status": "Declined",
				}

				a.sendNotification(userID, verb, object, clientID, employerTenantID, content, category, subcategory)
			}
		}
	}

	return stream.Activity{}, nil
}

// AddRevokedApplicationSuccessFeeActivity create an activity for revoked application
func (a *Application) AddRevokedApplicationSuccessFeeActivity(clientID string, clientName string, supplierID string, supplierName string, successFeeID string, successFeeNumber string, employerTenantID string, agencyTenantID string) (stream.Activity, error) {
	helper := NewHelper(a.telemetryClient)

	clientData, _ := helper.createJSONMarshal(clientID, clientName, "client")
	successFeeData, _ := helper.createJSONMarshal(successFeeID, successFeeNumber, "successFee")

	assignmentRepository, settingsRepository := a.createAssignmentAndSettingsRepository()
	assignments, _ := assignmentRepository.GetAgencyAssignmentsBySuccessFeeID(successFeeID, agencyTenantID)

	if assignmentRepository.IsApprovedSuccessFeeStatus(successFeeID) {
		for _, assignment := range assignments {
			userID := strings.ToLower(assignment.UserID)
			settings, _ := settingsRepository.GetSettingsBySupplierID(userID)

			if settings.Applications {
				verb := "revoke"
				object := fmt.Sprintf("successfee:%s", successFeeID)
				content := fmt.Sprintf("%s revoked your approval to Engagement #%s", clientData, successFeeData)
				category := "Application"
				subcategory := map[string]string{
					"type":   "Application",
					"status": "Revoked",
				}

				a.sendNotification(userID, verb, object, clientID, employerTenantID, content, category, subcategory)
			}
		}
	}

	return stream.Activity{}, nil
}

func (a *Application) createAssignmentAndSettingsRepository() (*Assignment, *Settings) {
	assignmentRepository := NewAssignmentRepository(a.auctionDB, a.telemetryClient)
	settingsRepository := NewSettingsRepository(nil, a.agencyDB, a.telemetryClient)

	return assignmentRepository, settingsRepository
}

func (a *Application) sendNotification(userID string, verb string, object string, clientID string, employerTenantID string, content string, category string, subcategory map[string]string) {
	agencyNotificationFeed := a.client.NotificationFeed("agencynotification", userID)
	_, err := agencyNotificationFeed.AddActivity(stream.Activity{
		Actor:     agencyNotificationFeed.ID(),
		Verb:      verb,
		Object:    object,
		ForeignID: agencyNotificationFeed.ID(),
		Time:      stream.Time{time.Now().UTC()},
		Extra: map[string]interface{}{
			"employer":         fmt.Sprintf("employer:%s", clientID),
			"agency":           fmt.Sprintf("agency:%s", userID),
			"employerTenantID": employerTenantID,
			"content":          content,
			"category":         category,
			"subcategory":      subcategory,
		},
	})

	if err != nil {
		a.telemetryClient.TrackException(err)
		panic(err)
	}
}
