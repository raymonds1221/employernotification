package repository

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"database/sql"

	"github.com/Microsoft/ApplicationInsights-Go/appinsights"
	"github.com/google/uuid"
	stream "gopkg.in/GetStream/stream-go2.v1"
)

const (
	candidateShortlisted           = "/api/v1.0.0/employer/candidateshortlisted"
	candidateShortlistedSuccessFee = "/api/v1.0.0/employer/candidateshortlisted/successfee"
)

// Fulfillment implementation of fulfillment repository
type Fulfillment struct {
	client          *stream.Client
	agencyDB        *sql.DB
	auctionDB       *sql.DB
	telemetryClient appinsights.TelemetryClient
}

// NewFulfillmentRepository create new instance of fulfillment repository
func NewFulfillmentRepository(client *stream.Client, agencyDB *sql.DB, auctionDB *sql.DB, telemetryClient appinsights.TelemetryClient) *Fulfillment {
	return &Fulfillment{
		client:          client,
		agencyDB:        agencyDB,
		auctionDB:       auctionDB,
		telemetryClient: telemetryClient,
	}
}

// AddCandidateShortlistActivity create an activity when employer shortlisted a candidate
func (f *Fulfillment) AddCandidateShortlistActivity(clientID string, clientName string, supplierID string, supplierName string, candidateID string, candidateName string, talentRequestID string, talentRequestNumber string, auctionID string, auctionNumber string, jobTitle string, employerTenantID string, agencyTenantID string, token string) (stream.Activity, error) {
	helper := NewHelper(f.telemetryClient)

	clientData, _ := helper.createJSONMarshal(clientID, clientName, "client")
	candidateData, _ := helper.createJSONMarshal(candidateID, candidateName, "candidate")
	talentRequestData, _ := helper.createJSONMarshal(talentRequestID, fmt.Sprintf("%06s", talentRequestNumber), "talentRequest")
	auctionData, _ := helper.createJSONMarshal(auctionID, auctionNumber, "auction")

	assignmentRepository, settingsRepository := f.createAssignmentAndSettingsRepository()
	assignments, _ := assignmentRepository.GetAgencyAssignmentsByAuctionID(auctionID, agencyTenantID)

	if assignmentRepository.IsApprovedAuctionStatus(auctionID) {
		for _, assignment := range assignments {
			userID := strings.ToLower(assignment.UserID)
			settings, _ := settingsRepository.GetSettingsBySupplierID(userID)

			if settings.Fulfillment {
				verb := "shortlist"
				object := fmt.Sprintf("candidate:%s", candidateID)
				target := fmt.Sprintf("auction:%s", auctionID)
				content := fmt.Sprintf("Congratulations, %s shortlisted your candidate %s, for Talent Request #%s: %s, in Auction #%s (Competitive).", string(clientData), string(candidateData), string(talentRequestData), jobTitle, string(auctionData))
				category := "Fulfillment"
				subcategory := map[string]string{
					"type":   "Candidate",
					"status": "Shortlisted",
				}

				f.sendNotificationToAgency(userID, verb, object, target, clientID, employerTenantID, content, category, subcategory)
			}
		}

		formData := url.Values{
			"clientID":            {clientID},
			"clientName":          {clientName},
			"supplierID":          {supplierID},
			"supplierName":        {supplierName},
			"candidateID":         {candidateID},
			"candidateName":       {candidateName},
			"talentRequestID":     {talentRequestID},
			"talentRequestNumber": {talentRequestNumber},
			"auctionID":           {auctionID},
			"auctionNumber":       {auctionNumber},
			"jobTitle":            {jobTitle},
			"employerTenantID":    {employerTenantID},
			"agencyTenantID":      {agencyTenantID},
		}

		helper.post(candidateShortlisted, formData, token)
	}

	return stream.Activity{}, nil
}

// AddCandidateDeclineActivity create an activity when employer declined a candidate
func (f *Fulfillment) AddCandidateDeclineActivity(clientID string, clientName string, supplierID string, supplierName string, candidateID string, candidateName string, talentRequestID string, talentRequestNumber string, auctionID string, auctionNumber string, jobTitle string, employerTenantID string, agencyTenantID string) (stream.Activity, error) {
	helper := NewHelper(f.telemetryClient)

	clientData, _ := helper.createJSONMarshal(clientID, clientName, "client")
	candidateData, _ := helper.createJSONMarshal(candidateID, candidateName, "candidate")
	talentRequestData, _ := helper.createJSONMarshal(talentRequestID, fmt.Sprintf("%06s", talentRequestNumber), "talentRequest")
	auctionData, _ := helper.createJSONMarshal(auctionID, auctionNumber, "auction")

	assignmentRepository, settingsRepository := f.createAssignmentAndSettingsRepository()
	assignments, _ := assignmentRepository.GetAgencyAssignmentsByAuctionID(auctionID, agencyTenantID)

	if assignmentRepository.IsApprovedAuctionStatus(auctionID) {
		for _, assignment := range assignments {
			userID := strings.ToLower(assignment.UserID)
			settings, _ := settingsRepository.GetSettingsBySupplierID(userID)

			if settings.Fulfillment {
				verb := "reject"
				object := fmt.Sprintf("candidate:%s", candidateID)
				target := fmt.Sprintf("auction:%s", auctionID)
				content := fmt.Sprintf("Please note, %s rejected your candidate %s, for Talent Request #%s: %s in Auction #%s (Competitive). Please ensure you submit qualified candidates.", string(clientData), string(candidateData), string(talentRequestData), jobTitle, string(auctionData))
				category := "Fulfillment"
				subcategory := map[string]string{
					"type":   "Candidate",
					"status": "Rejected",
				}

				f.sendNotificationToAgency(userID, verb, object, target, clientID, employerTenantID, content, category, subcategory)
			}
		}
	}

	return stream.Activity{}, nil
}

// AddCandidateHiredActivity create an activity when employer accepted a candidate
func (f *Fulfillment) AddCandidateHiredActivity(clientID string, clientName string, supplierID string, supplierName string, candidateID string, candidateName string, talentRequestID string, talentRequestNumber string, auctionID string, auctionNumber string, jobTitle string, employerTenantID string, agencyTenantID string) (stream.Activity, error) {
	helper := NewHelper(f.telemetryClient)

	clientData, _ := helper.createJSONMarshal(clientID, clientName, "client")
	candidateData, _ := helper.createJSONMarshal(candidateID, candidateName, "candidate")
	talentRequestData, _ := helper.createJSONMarshal(talentRequestID, fmt.Sprintf("%06s", talentRequestNumber), "talentRequest")
	auctionData, _ := helper.createJSONMarshal(auctionID, auctionNumber, "auction")

	assignmentRepository, settingsRepository := f.createAssignmentAndSettingsRepository()
	assignments, _ := assignmentRepository.GetAgencyAssignmentsByAuctionID(auctionID, agencyTenantID)

	if assignmentRepository.IsApprovedAuctionStatus(auctionID) {
		for _, assignment := range assignments {
			userID := strings.ToLower(assignment.UserID)
			settings, _ := settingsRepository.GetSettingsBySupplierID(userID)

			if settings.Fulfillment {
				verb := "hire"
				object := fmt.Sprintf("candidate:%s", candidateID)
				target := fmt.Sprintf("auction:%s", auctionID)
				content := fmt.Sprintf("Congratulations, %s hired your candidate %s, for Talent Request #%s: %s in Auction #%s (Competitive). A separate email confirming your earnings will be sent.", string(clientData), string(candidateData), string(talentRequestData), jobTitle, string(auctionData))
				category := "Fulfillment"
				subcategory := map[string]string{
					"type":   "Candidate",
					"status": "Hired",
				}

				f.sendNotificationToAgency(userID, verb, object, target, clientID, employerTenantID, content, category, subcategory)
			}
		}
	}

	return stream.Activity{}, nil
}

// AddCandidateUpdateActivity create an activity when employer updated a candidate
func (f *Fulfillment) AddCandidateUpdateActivity(clientID string, clientName string, supplierID string, supplierName string, candidateID string, candidateName string, talentRequestID string, talentRequestNumber string, auctionID string, auctionNumber string, jobTitle string, employerTenantID string, agencyTenantID string, status string) (stream.Activity, error) {
	helper := NewHelper(f.telemetryClient)

	clientData, _ := helper.createJSONMarshal(clientID, clientName, "client")
	candidateData, _ := helper.createJSONMarshal(candidateID, candidateName, "candidate")
	talentRequestData, _ := helper.createJSONMarshal(talentRequestID, fmt.Sprintf("%06s", talentRequestNumber), "talentRequest")
	auctionData, _ := helper.createJSONMarshal(auctionID, auctionNumber, "auction")

	assignmentRepository, settingsRepository := f.createAssignmentAndSettingsRepository()
	assignments, _ := assignmentRepository.GetAgencyAssignmentsByAuctionID(auctionID, agencyTenantID)

	if assignmentRepository.IsApprovedAuctionStatus(auctionID) {
		for _, assignment := range assignments {
			userID := strings.ToLower(assignment.UserID)
			settings, _ := settingsRepository.GetSettingsBySupplierID(userID)

			if settings.Fulfillment {
				verb := "hire"
				object := fmt.Sprintf("candidate:%s", candidateID)
				target := fmt.Sprintf("auction:%s", auctionID)
				content := fmt.Sprintf("Please note, %s updated the status of candidate %s to %s, for Talent Request #%s: %s, in Auction #%s (Competitive).", string(clientData), string(candidateData), strings.Title(status), string(talentRequestData), jobTitle, string(auctionData))
				category := "Fulfillment"
				subcategory := map[string]string{
					"type":   "Candidate",
					"status": strings.Title(status),
				}

				f.sendNotificationToAgency(userID, verb, object, target, clientID, employerTenantID, content, category, subcategory)
			}
		}
	}

	return stream.Activity{}, nil
}

// AddCandidatePendingStatusActivity create an activity when employer doesn't updated the candidate within 5 days with pending status
func (f *Fulfillment) AddCandidatePendingStatusActivity(clientID string, clientName string, supplierID string, supplierName string, candidateID string, candidateName string, talentRequestID string, talentRequestNumber string, auctionID string, auctionNumber string, jobTitle string, employerTenantID string, agencyTenantID string) (stream.Activity, error) {
	helper := NewHelper(f.telemetryClient)

	supplierData, _ := helper.createJSONMarshal(supplierID, supplierName, "supplier")
	candidateData, _ := helper.createJSONMarshal(candidateID, candidateName, "candidate")
	talentRequestData, _ := helper.createJSONMarshal(talentRequestID, fmt.Sprintf("%06s", talentRequestNumber), "talentRequest")
	auctionData, _ := helper.createJSONMarshal(auctionID, auctionNumber, "auction")

	assignmentRepository := NewAssignmentRepository(f.auctionDB, f.telemetryClient)
	assignments, _ := assignmentRepository.GetEmployerAssignmentsByAuctionID(auctionID)

	if assignmentRepository.IsApprovedAuctionStatus(auctionID) {
		for _, assignment := range assignments {
			userID := strings.ToLower(assignment.UserID)

			verb := "pending"
			object := fmt.Sprintf("candidate:%s", candidateID)
			target := fmt.Sprintf("auction:%s", auctionID)
			content := fmt.Sprintf("Please note, %s candidate %s, uploaded for Competitive Auction #%s for Talent Request #%s: %s, has been Pending for over 5 Days! Please Shortlist or Reject this Candidate.", string(supplierData), string(candidateData), string(auctionData), string(talentRequestData), jobTitle)
			category := "Fulfillment"
			subcategory := map[string]string{
				"type":   "Candidate",
				"status": "Pending",
			}

			f.sendNotificationToEmployer(userID, verb, object, target, clientID, supplierID, employerTenantID, content, category, subcategory)
		}
	}

	return stream.Activity{}, nil
}

// AddCandidateShortlistStatusActivity create an activity when employer doesn't updated the candidate within 5 days with shortlist status
func (f *Fulfillment) AddCandidateShortlistStatusActivity(clientID string, clientName string, supplierID string, supplierName string, candidateID string, candidateName string, talentRequestID string, talentRequestNumber string, auctionID string, auctionNumber string, jobTitle string, employerTenantID string, agencyTenantID string) (stream.Activity, error) {
	helper := NewHelper(f.telemetryClient)

	supplierData, _ := helper.createJSONMarshal(supplierID, supplierName, "supplier")
	candidateData, _ := helper.createJSONMarshal(candidateID, candidateName, "candidate")
	talentRequestData, _ := helper.createJSONMarshal(talentRequestID, fmt.Sprintf("%06s", talentRequestNumber), "talentRequest")
	auctionData, _ := helper.createJSONMarshal(auctionID, auctionNumber, "auction")

	assignmentRepository := NewAssignmentRepository(f.auctionDB, f.telemetryClient)
	assignments, _ := assignmentRepository.GetEmployerAssignmentsByAuctionID(auctionID)

	if assignmentRepository.IsApprovedAuctionStatus(auctionID) {
		for _, assignment := range assignments {
			userID := strings.ToLower(assignment.UserID)

			verb := "shortlist"
			object := fmt.Sprintf("candidate:%s", candidateID)
			target := fmt.Sprintf("auction:%s", auctionID)
			content := fmt.Sprintf("Please note, %s's candidate %s, uploaded for Competitive Auction #%s for Talent Request #%s: %s, has been Shortlisted for over 14 Days! Please update the status of this Candidate.", string(supplierData), string(candidateData), string(auctionData), string(talentRequestData), jobTitle)
			category := "Fulfillment"
			subcategory := map[string]string{
				"type":   "Candidate",
				"status": "Shortlisted",
			}

			f.sendNotificationToEmployer(userID, verb, object, target, content, supplierID, employerTenantID, content, category, subcategory)
		}
	}

	return stream.Activity{}, nil
}

// AddCandidateShortlistSuccessFeeActivity create an activity when employer shortlisted a candidate
func (f *Fulfillment) AddCandidateShortlistSuccessFeeActivity(clientID string, clientName string, supplierID string, supplierName string, candidateID string, candidateName string, talentRequestID string, talentRequestNumber string, successFeeID string, successFeeNumber string, jobTitle string, employerTenantID string, agencyTenantID string, token string) (stream.Activity, error) {
	helper := NewHelper(f.telemetryClient)

	clientData, _ := helper.createJSONMarshal(clientID, clientName, "client")
	candidateData, _ := helper.createJSONMarshal(candidateID, candidateName, "candidate")
	talentRequestData, _ := helper.createJSONMarshal(talentRequestID, fmt.Sprintf("%06s", talentRequestNumber), "talentRequest")
	successFeeData, _ := helper.createJSONMarshal(successFeeID, successFeeNumber, "successFee")

	assignmentRepository, settingsRepository := f.createAssignmentAndSettingsRepository()
	assignments, _ := assignmentRepository.GetAgencyAssignmentsBySuccessFeeID(successFeeID, agencyTenantID)

	if assignmentRepository.IsApprovedSuccessFeeStatus(successFeeID) {
		for _, assignment := range assignments {
			userID := strings.ToLower(assignment.UserID)
			settings, _ := settingsRepository.GetSettingsBySupplierID(userID)

			if settings.Fulfillment {
				verb := "shortlist"
				object := fmt.Sprintf("candidate:%s", candidateID)
				target := fmt.Sprintf("successfee:%s", successFeeID)
				content := fmt.Sprintf("%s updated %s to SHORTLISTED, for TR #%s: %s", string(clientData), string(candidateData), string(talentRequestData), jobTitle)
				category := "Fulfillment"
				subcategory := map[string]string{
					"type":   "Candidate",
					"status": "Shortlisted",
					"data":   fmt.Sprintf("%s", successFeeData),
				}

				f.sendNotificationToAgency(userID, verb, object, target, clientID, employerTenantID, content, category, subcategory)
			}
		}

		formData := url.Values{
			"clientID":            {clientID},
			"clientName":          {clientName},
			"supplierID":          {supplierID},
			"supplierName":        {supplierName},
			"candidateID":         {candidateID},
			"candidateName":       {candidateName},
			"talentRequestID":     {talentRequestID},
			"talentRequestNumber": {talentRequestNumber},
			"successFeeID":        {successFeeID},
			"successFeeNumber":    {successFeeNumber},
			"jobTitle":            {jobTitle},
			"employerTenantID":    {employerTenantID},
			"agencyTenantID":      {agencyTenantID},
		}

		helper.post(candidateShortlistedSuccessFee, formData, token)
	}

	return stream.Activity{}, nil
}

// AddCandidateDeclineSuccessFeeActivity create an activity when employer declined a candidate
func (f *Fulfillment) AddCandidateDeclineSuccessFeeActivity(clientID string, clientName string, supplierID string, supplierName string, candidateID string, candidateName string, talentRequestID string, talentRequestNumber string, successFeeID string, successFeeNumber string, jobTitle string, employerTenantID string, agencyTenantID string) (stream.Activity, error) {
	helper := NewHelper(f.telemetryClient)

	clientData, _ := helper.createJSONMarshal(clientID, clientName, "client")
	candidateData, _ := helper.createJSONMarshal(candidateID, candidateName, "candidate")
	talentRequestData, _ := helper.createJSONMarshal(talentRequestID, fmt.Sprintf("%06s", talentRequestNumber), "talentRequest")
	successFeeData, _ := helper.createJSONMarshal(successFeeID, successFeeNumber, "successFee")

	assignmentRepository, settingsRepository := f.createAssignmentAndSettingsRepository()
	assignments, _ := assignmentRepository.GetAgencyAssignmentsBySuccessFeeID(successFeeID, agencyTenantID)

	if assignmentRepository.IsApprovedSuccessFeeStatus(successFeeID) {
		for _, assignment := range assignments {
			userID := strings.ToLower(assignment.UserID)
			settings, _ := settingsRepository.GetSettingsBySupplierID(userID)

			if settings.Fulfillment {
				verb := "reject"
				object := fmt.Sprintf("candidate:%s", candidateID)
				target := fmt.Sprintf("successfee:%s", successFeeID)
				content := fmt.Sprintf("%s REJECTED candidate %s, for Talent Request #%s: %s", string(clientData), string(candidateData), string(talentRequestData), jobTitle)
				category := "Fulfillment"
				subcategory := map[string]string{
					"type":   "Candidate",
					"status": "Rejected",
					"data":   fmt.Sprintf("%s", successFeeData),
				}

				f.sendNotificationToAgency(userID, verb, object, target, clientID, employerTenantID, content, category, subcategory)
			}
		}
	}

	return stream.Activity{}, nil
}

// AddCandidateHiredSuccessFeeActivity create an activity when employer accepted a candidate
func (f *Fulfillment) AddCandidateHiredSuccessFeeActivity(clientID string, clientName string, supplierID string, supplierName string, candidateID string, candidateName string, talentRequestID string, talentRequestNumber string, successFeeID string, successFeeNumber string, jobTitle string, employerTenantID string, agencyTenantID string) (stream.Activity, error) {
	helper := NewHelper(f.telemetryClient)

	clientData, _ := helper.createJSONMarshal(clientID, clientName, "client")
	candidateData, _ := helper.createJSONMarshal(candidateID, candidateName, "candidate")
	talentRequestData, _ := helper.createJSONMarshal(talentRequestID, fmt.Sprintf("%06s", talentRequestNumber), "talentRequest")
	successFeeData, _ := helper.createJSONMarshal(successFeeID, successFeeNumber, "successFee")

	assignmentRepository, settingsRepository := f.createAssignmentAndSettingsRepository()
	assignments, _ := assignmentRepository.GetAgencyAssignmentsBySuccessFeeID(successFeeID, agencyTenantID)

	if assignmentRepository.IsApprovedSuccessFeeStatus(successFeeID) {
		for _, assignment := range assignments {
			userID := strings.ToLower(assignment.UserID)
			settings, _ := settingsRepository.GetSettingsBySupplierID(userID)

			if settings.Fulfillment {
				verb := "hire"
				object := fmt.Sprintf("candidate:%s", candidateID)
				target := fmt.Sprintf("successfee:%s", successFeeID)
				content := fmt.Sprintf("%s HIRED %s for TR #%s: %s", string(clientData), string(candidateData), string(talentRequestData), jobTitle)
				category := "Fulfillment"
				subcategory := map[string]string{
					"type":   "Candidate",
					"status": "Hired",
					"data":   fmt.Sprintf("%s", successFeeData),
				}

				f.sendNotificationToAgency(userID, verb, object, target, clientID, employerTenantID, content, category, subcategory)
			}
		}
	}

	return stream.Activity{}, nil
}

// AddCandidateUpdateSuccessFeeActivity create an activity when employer updated a candidate
func (f *Fulfillment) AddCandidateUpdateSuccessFeeActivity(clientID string, clientName string, supplierID string, supplierName string, candidateID string, candidateName string, talentRequestID string, talentRequestNumber string, successFeeID string, successFeeNumber string, jobTitle string, employerTenantID string, agencyTenantID string, status string) (stream.Activity, error) {
	helper := NewHelper(f.telemetryClient)

	clientData, _ := helper.createJSONMarshal(clientID, clientName, "client")
	candidateData, _ := helper.createJSONMarshal(candidateID, candidateName, "candidate")
	talentRequestData, _ := helper.createJSONMarshal(talentRequestID, fmt.Sprintf("%06s", talentRequestNumber), "talentRequest")
	successFeeData, _ := helper.createJSONMarshal(successFeeID, successFeeNumber, "successFee")
	assignmentRepository, settingsRepository := f.createAssignmentAndSettingsRepository()
	assignments, _ := assignmentRepository.GetAgencyAssignmentsBySuccessFeeID(successFeeID, agencyTenantID)

	if assignmentRepository.IsApprovedSuccessFeeStatus(successFeeID) {
		for _, assignment := range assignments {
			userID := strings.ToLower(assignment.UserID)
			settings, _ := settingsRepository.GetSettingsBySupplierID(userID)

			if settings.Fulfillment {
				verb := "hire"
				object := fmt.Sprintf("candidate:%s", candidateID)
				target := fmt.Sprintf("successfee:%s", successFeeID)
				content := fmt.Sprintf("%s updated %s to %s, for TR #%s: %s", string(clientData), string(candidateData), strings.ToUpper(status), string(talentRequestData), jobTitle)
				category := "Fulfillment"
				subcategory := map[string]string{
					"type":   "Candidate",
					"status": strings.Title(status),
					"data":   fmt.Sprintf("%s", successFeeData),
				}

				f.sendNotificationToAgency(userID, verb, object, target, clientID, employerTenantID, content, category, subcategory)
			}
		}
	}

	return stream.Activity{}, nil
}

// AddCandidatePendingStatusSuccessFeeActivity create an activity when employer doesn't updated the candidate within 5 days with pending status
func (f *Fulfillment) AddCandidatePendingStatusSuccessFeeActivity(clientID string, clientName string, supplierID string, supplierName string, candidateID string, candidateName string, talentRequestID string, talentRequestNumber string, successFeeID string, successFeeNumber string, jobTitle string, employerTenantID string, agencyTenantID string) (stream.Activity, error) {
	helper := NewHelper(f.telemetryClient)

	supplierData, _ := helper.createJSONMarshal(supplierID, supplierName, "supplier")
	candidateData, _ := helper.createJSONMarshal(candidateID, candidateName, "candidate")
	talentRequestData, _ := helper.createJSONMarshal(talentRequestID, fmt.Sprintf("%06s", talentRequestNumber), "talentRequest")
	successFeeData, _ := helper.createJSONMarshal(successFeeID, successFeeNumber, "successFee")

	assignmentRepository := NewAssignmentRepository(f.auctionDB, f.telemetryClient)
	assignments, _ := assignmentRepository.GetEmployerAssignmentsBySuccessFeeID(successFeeID)

	if assignmentRepository.IsApprovedSuccessFeeStatus(successFeeID) {
		for _, assignment := range assignments {
			userID := strings.ToLower(assignment.UserID)

			verb := "pending"
			object := fmt.Sprintf("candidate:%s", candidateID)
			target := fmt.Sprintf("successfee:%s", successFeeID)
			content := fmt.Sprintf("%s candidate %s, uploaded for TR #%s: %s, has been Pending for over 5 Days! Please Shortlist or Reject this Candidate.", string(supplierData), string(candidateData), string(talentRequestData), jobTitle)
			category := "Fulfillment"
			subcategory := map[string]string{
				"type":   "Candidate",
				"status": "Pending",
				"data":   fmt.Sprintf("%s", successFeeData),
			}

			f.sendNotificationToEmployer(userID, verb, object, target, clientID, supplierID, employerTenantID, content, category, subcategory)
		}
	}

	return stream.Activity{}, nil
}

// AddCandidateShortlistStatusSuccessFeeActivity create an activity when employer doesn't updated the candidate within 5 days with shortlist status
func (f *Fulfillment) AddCandidateShortlistStatusSuccessFeeActivity(clientID string, clientName string, supplierID string, supplierName string, candidateID string, candidateName string, talentRequestID string, talentRequestNumber string, successFeeID string, successFeeNumber string, jobTitle string, employerTenantID string, agencyTenantID string) (stream.Activity, error) {
	helper := NewHelper(f.telemetryClient)

	supplierData, _ := helper.createJSONMarshal(supplierID, supplierName, "supplier")
	candidateData, _ := helper.createJSONMarshal(candidateID, candidateName, "candidate")
	talentRequestData, _ := helper.createJSONMarshal(talentRequestID, fmt.Sprintf("%06s", talentRequestNumber), "talentRequest")
	successFeeData, _ := helper.createJSONMarshal(successFeeID, successFeeNumber, "successFee")

	assignmentRepository := NewAssignmentRepository(f.auctionDB, f.telemetryClient)
	assignments, _ := assignmentRepository.GetEmployerAssignmentsBySuccessFeeID(successFeeID)

	if assignmentRepository.IsApprovedSuccessFeeStatus(successFeeID) {
		for _, assignment := range assignments {
			userID := strings.ToLower(assignment.UserID)

			verb := "hire"
			object := fmt.Sprintf("candidate:%s", candidateID)
			target := fmt.Sprintf("successfee:%s", successFeeID)
			content := fmt.Sprintf("%s's candidate %s, uploaded for Talent Request #%s: %s, has been Shortlisted for over 14 Days! Please update the status of this Candidate.", string(supplierData), string(candidateData), string(talentRequestData), jobTitle)
			category := "Fulfillment"
			subcategory := map[string]string{
				"type":   "Candidate",
				"status": "Shortlisted",
				"data":   fmt.Sprintf("%s", successFeeData),
			}

			f.sendNotificationToEmployer(userID, verb, object, target, clientID, supplierID, employerTenantID, content, category, subcategory)
		}
	}

	return stream.Activity{}, nil
}

func (f *Fulfillment) createAssignmentAndSettingsRepository() (*Assignment, *Settings) {
	assignmentRepository := NewAssignmentRepository(f.auctionDB, f.telemetryClient)
	settingsRepository := NewSettingsRepository(nil, f.agencyDB, f.telemetryClient)

	return assignmentRepository, settingsRepository
}

func (f *Fulfillment) sendNotificationToAgency(userID string, verb string, object string, target string, clientID string, employerTenantID string, content string, category string, subcategory map[string]string) {
	agencyNotificationFeed := f.client.NotificationFeed("agencynotification", userID)
	_, err := agencyNotificationFeed.AddActivity(stream.Activity{
		Actor:     agencyNotificationFeed.ID(),
		Verb:      verb,
		Object:    object,
		Target:    target,
		ForeignID: uuid.New().String(),
		Time:      stream.Time{time.Now().UTC()},
		Extra: map[string]interface{}{
			"employer":         fmt.Sprintf("employer:%s", clientID),
			"agency":           fmt.Sprintf("agency:%s", userID),
			"employerTenantID": employerTenantID,
			"content":          content,
			"category":         category,
			"subcategory":      subcategory,
		},
	},
	)

	if err != nil {
		f.telemetryClient.TrackException(err)
	}
}

func (f *Fulfillment) sendNotificationToEmployer(userID string, verb string, object string, target string, clientID string, supplierID string, employerTenantID string, content string, category string, subcategory map[string]string) {
	employerNotificationFeed := f.client.NotificationFeed("employernotification", userID)
	_, err := employerNotificationFeed.AddActivity(stream.Activity{
		Actor:     employerNotificationFeed.ID(),
		Verb:      verb,
		Object:    object,
		Target:    target,
		ForeignID: uuid.New().String(),
		Time:      stream.Time{time.Now().UTC()},
		Extra: map[string]interface{}{
			"employer":         fmt.Sprintf("employer:%s", clientID),
			"agency":           fmt.Sprintf("agency:%s", supplierID),
			"employerTenantID": employerTenantID,
			"content":          content,
			"category":         category,
			"subcategory":      subcategory,
		},
	})

	if err != nil {
		f.telemetryClient.TrackException(err)
	}
}
