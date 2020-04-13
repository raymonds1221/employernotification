package repository

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Microsoft/ApplicationInsights-Go/appinsights"
	stream "gopkg.in/GetStream/stream-go2.v1"
)

// Clarification implementation of clarification repository
type Clarification struct {
	client          *stream.Client
	agencyDB        *sql.DB
	auctionDB       *sql.DB
	telemetryClient appinsights.TelemetryClient
}

// NewClarificationRepository create new instance of clarification repository
func NewClarificationRepository(client *stream.Client, agencyDB *sql.DB, auctionDB *sql.DB, telemetryClient appinsights.TelemetryClient) *Clarification {
	return &Clarification{
		client:          client,
		agencyDB:        agencyDB,
		auctionDB:       auctionDB,
		telemetryClient: telemetryClient,
	}
}

// AddReplyClarificationActivity create an activity when employer replied to the clarification discussion
func (c *Clarification) AddReplyClarificationActivity(clientID string, clientName string, supplierID string, supplierName string, auctionID string, auctionNumber string, employerTenantID string, agencyTenantID string) (stream.Activity, error) {
	helper := NewHelper(c.telemetryClient)

	clientData, _ := helper.createJSONMarshal(clientID, clientName, "client")
	auctionData, _ := helper.createJSONMarshal(auctionID, auctionNumber, "auction")

	assignmentRepository, settingsRepository := c.createAssignmentAndSettingsRepository()
	assignments, _ := assignmentRepository.GetAgencyAssignmentsByAuctionID(auctionID, agencyTenantID)

	if assignmentRepository.IsApprovedAuctionStatus(auctionID) {
		for _, assignment := range assignments {
			userID := strings.ToLower(assignment.UserID)
			settings, _ := settingsRepository.GetSettingsBySupplierID(userID)

			if settings.Clarifications {
				verb := "reply"
				object := fmt.Sprintf("auction:%s", auctionID)
				content := fmt.Sprintf("Please note, %s created a new clarification reply in their Auction #%s (Competitive).", clientData, auctionData)
				category := "Clarification"
				subcategory := map[string]string{
					"type":   "Clarification",
					"status": "Replied",
				}

				c.sendNotification(userID, verb, object, clientID, employerTenantID, content, category, subcategory)
			}
		}
	}

	return stream.Activity{}, nil
}

// AddPostTopicActivity create an activity when employer posted new discussion
func (c *Clarification) AddPostTopicActivity(clientID string, clientName string, supplierID string, supplierName string, auctionID string, auctionNumber string, employerTenantID string, agencyTenantID string) (stream.Activity, error) {
	helper := NewHelper(c.telemetryClient)

	clientData, _ := helper.createJSONMarshal(clientID, clientName, "client")
	auctionData, _ := helper.createJSONMarshal(auctionID, auctionNumber, "auction")

	assignmentRepository, settingsRepository := c.createAssignmentAndSettingsRepository()
	assignments, _ := assignmentRepository.GetAgencyAssignmentsByAuctionID(auctionID, agencyTenantID)

	if assignmentRepository.IsApprovedAuctionStatus(auctionID) {
		for _, assignment := range assignments {
			userID := strings.ToLower(assignment.UserID)
			settings, _ := settingsRepository.GetSettingsBySupplierID(userID)

			log.Printf("assignment userid: %s\n", userID)

			if settings.Clarifications {
				verb := "post"
				object := fmt.Sprintf("auction:%s", auctionID)
				content := fmt.Sprintf("Please note, %s created a new clarification topic for their Auction #%s (Competitive).", clientData, auctionData)
				category := "Clarification"
				subcategory := map[string]string{
					"type":   "Clarification",
					"status": "Topic",
				}

				c.sendNotification(userID, verb, object, clientID, employerTenantID, content, category, subcategory)
			}
		}
	}

	return stream.Activity{}, nil
}

// AddReplyClarificationSuccessFeeActivity create an activity when employer replied to the clarification discussion
func (c *Clarification) AddReplyClarificationSuccessFeeActivity(clientID string, clientName string, supplierID string, supplierName string, successFeeID string, successFeeNumber string, employerTenantID string, agencyTenantID string) (stream.Activity, error) {
	helper := NewHelper(c.telemetryClient)

	clientData, _ := helper.createJSONMarshal(clientID, clientName, "client")
	successFeeData, _ := helper.createJSONMarshal(successFeeID, successFeeNumber, "successFee")

	assignmentRepository, settingsRepository := c.createAssignmentAndSettingsRepository()
	assignments, _ := assignmentRepository.GetAgencyAssignmentsBySuccessFeeID(successFeeID, agencyTenantID)

	if assignmentRepository.IsApprovedSuccessFeeStatus(successFeeID) {
		for _, assignment := range assignments {
			userID := strings.ToLower(assignment.UserID)
			settings, _ := settingsRepository.GetSettingsBySupplierID(userID)

			if settings.Clarifications {
				verb := "reply"
				object := fmt.Sprintf("successfee:%s", successFeeID)
				content := fmt.Sprintf("%s replied to a clarification topic in Engagement #%s", clientData, successFeeData)
				category := "Clarification"
				subcategory := map[string]string{
					"type":   "Clarification",
					"status": "Replied",
				}

				c.sendNotification(userID, verb, object, clientID, employerTenantID, content, category, subcategory)
			}
		}
	}

	return stream.Activity{}, nil
}

// AddPostTopicSuccessFeeActivity create an activity when employer posted new discussion
func (c *Clarification) AddPostTopicSuccessFeeActivity(clientID string, clientName string, supplierID string, supplierName string, successFeeID string, successFeeNumber string, employerTenantID string, agencyTenantID string) (stream.Activity, error) {
	helper := NewHelper(c.telemetryClient)

	clientData, _ := helper.createJSONMarshal(clientID, clientName, "client")
	successFeeData, _ := helper.createJSONMarshal(successFeeID, successFeeNumber, "successFee")

	assignmentRepository, settingsRepository := c.createAssignmentAndSettingsRepository()
	assignments, _ := assignmentRepository.GetAgencyAssignmentsBySuccessFeeID(successFeeID, agencyTenantID)

	if assignmentRepository.IsApprovedSuccessFeeStatus(successFeeID) {
		for _, assignment := range assignments {
			userID := strings.ToLower(assignment.UserID)
			settings, _ := settingsRepository.GetSettingsBySupplierID(userID)

			if settings.Clarifications {
				verb := "post"
				object := fmt.Sprintf("successfee:%s", successFeeID)
				content := fmt.Sprintf("%s created a clarification topic for Engagement #%s", clientData, successFeeData)
				category := "Clarification"
				subcategory := map[string]string{
					"type":   "Clarification",
					"status": "Topic",
				}

				c.sendNotification(userID, verb, object, clientID, employerTenantID, content, category, subcategory)
			}
		}
	}

	return stream.Activity{}, nil
}

func (c *Clarification) createAssignmentAndSettingsRepository() (*Assignment, *Settings) {
	assignmentRepository := NewAssignmentRepository(c.auctionDB, c.telemetryClient)
	settingsRepository := NewSettingsRepository(nil, c.agencyDB, c.telemetryClient)

	return assignmentRepository, settingsRepository
}

func (c *Clarification) sendNotification(userID string, verb string, object string, clientID string, employerTenantID string, content string, category string, subcategory map[string]string) {
	agencyNotificationFeed := c.client.NotificationFeed("agencynotification", userID)
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
		c.telemetryClient.TrackException(err)
	}
}
