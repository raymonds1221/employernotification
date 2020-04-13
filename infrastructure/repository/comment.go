package repository

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/Microsoft/ApplicationInsights-Go/appinsights"
	commentModel "github.com/Ubidy/Ubidy_EmployerNotificationAPI/domain/comment"
	stream "gopkg.in/GetStream/stream-go2.v1"
)

// Comment implementation of comment repository
type Comment struct {
	client          *stream.Client
	agencyDB        *sql.DB
	auctionDB       *sql.DB
	engagementDB    *sql.DB
	telemetryClient appinsights.TelemetryClient
}

// NewCommentRepository create new instance of comment repository
func NewCommentRepository(client *stream.Client, agencyDB *sql.DB, auctionDB *sql.DB, engagementDB *sql.DB, telemetryClient appinsights.TelemetryClient) *Comment {
	return &Comment{
		client:          client,
		agencyDB:        agencyDB,
		auctionDB:       auctionDB,
		engagementDB:    engagementDB,
		telemetryClient: telemetryClient,
	}
}

// AddCommentToCandidateActivity create an activity when the employer add comment on the candidate viewer
func (c *Comment) AddCommentToCandidateActivity(comment commentModel.AuctionComment) (stream.Activity, error) {
	assignmentRepository, settingsRepository := c.createAssignmentAndSettingsRepository()
	assignments, _ := assignmentRepository.GetAgencyAssignmentsByAuctionID(comment.AuctionID, comment.AgencyTenantID)

	if assignmentRepository.IsApprovedAuctionStatus(comment.AuctionID) {
		for _, assignment := range assignments {
			userID := strings.ToLower(assignment.UserID)
			settings, _ := settingsRepository.GetSettingsBySupplierID(userID)

			if settings.Fulfillment {
				msg := "Please note, %s commented to candidate %s, for Talent Request #%s: %s, in Auction #%s (Competitive)."
				c.sendNotificationToAgency(comment.Comment, userID, comment.AuctionID, comment.AuctionNumber, "auction", msg, comment.AuctionID)
			}
		}
	}

	return stream.Activity{}, nil
}

// AddCommentToCandidateSuccessFeeActivity create an activity when the employer add comment on the candidate viewer for success fee
func (c *Comment) AddCommentToCandidateSuccessFeeActivity(comment commentModel.SuccessFeeComment) (stream.Activity, error) {
	assignmentRepository, settingsRepository := c.createAssignmentAndSettingsRepository()
	engagementID := comment.SuccessFeeID // the SuccessFeeID has value of engagement id, it's a legacy code that's why it hasn't change yet.

	assignments, _ := assignmentRepository.GetAgencyAssignmentsByEngagementID(engagementID, comment.AgencyTenantID)

	if assignmentRepository.IsApprovedEngagementStatus(engagementID) {
		for _, assignment := range assignments {
			userID := strings.ToLower(assignment.UserID)
			settings, _ := settingsRepository.GetSettingsBySupplierID(userID)

			if settings.Fulfillment {
				// msg := "%s commented on candidate %s, for %s, in Engagement #%s"
				msg := "%s commented on candidate %s, for %s, in Engagement #%s"
				c.sendNotificationToAgency(comment.Comment, userID, comment.SuccessFeeID, comment.SuccessFeeNumber, "successFee", msg, engagementID)
			}
		}
	}

	return stream.Activity{}, nil
}

func (c *Comment) createAssignmentAndSettingsRepository() (*Assignment, *Settings) {
	settingsRepository := NewSettingsRepository(nil, c.agencyDB, c.telemetryClient)
	assignmentRepository := NewAssignmentRepository(c.engagementDB, c.telemetryClient)

	return assignmentRepository, settingsRepository
}

func (c *Comment) sendNotificationToAgency(comment commentModel.Comment, userID string, id string, number string, t string, msg string, engagementID string) error {
	helper := NewHelper(c.telemetryClient)

	employerData, _ := helper.createJSONMarshal(comment.EmployerUserID, comment.EmployerName, "client")
	// candidateData, _ := helper.createJSONMarshal(comment.CandidateID, comment.CandidateName, "candidate")
	talentRequestData, _ := helper.createJSONMarshal(comment.TalentRequestID, comment.JobTitle, "talentRequest")
	data, _ := helper.createJSONMarshal(id, number, t)

	agencyNotificationFeed := c.client.NotificationFeed("agencynotification", userID)

	_, err := agencyNotificationFeed.AddActivity(stream.Activity{
		Actor:     agencyNotificationFeed.ID(),
		Verb:      "create",
		Object:    fmt.Sprintf("%s:%s", t, id),
		ForeignID: agencyNotificationFeed.ID(),
		Time:      stream.Time{time.Now().UTC()},
		Extra: map[string]interface{}{
			"employer":         fmt.Sprintf("employer %s", comment.EmployerTenantID),
			"agency":           fmt.Sprintf("agency:%s", userID),
			"employerTenantID": comment.EmployerTenantID,
			// "content":          fmt.Sprintf(msg, employerData, candidateData, comment.JobTitle, data),
			"content":  fmt.Sprintf(msg, employerData, comment.CandidateName, talentRequestData, data),
			"object":   fmt.Sprintf("candidate: %s", comment.CandidateID),
			"target":   fmt.Sprintf("engagement: %s", engagementID),
			"category": "Fulfillment",
			"subcategory": map[string]string{
				"type":   "Candidate",
				"status": "New Comment",
				"data":   fmt.Sprintf("%s", talentRequestData),
			},
		},
	})

	if err != nil {
		c.telemetryClient.TrackException(err)
		return err
	}

	return nil
}
