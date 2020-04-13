package usecase

import (
	"github.com/Ubidy/Ubidy_EmployerNotificationAPI/usecase/repository"
	stream "gopkg.in/GetStream/stream-go2.v1"
)

// FulfillmentInteractor implementation of fulfillment usecase
type FulfillmentInteractor struct {
	repository repository.Fulfillment
}

// NewFulfillmentInteractor create new instance of fulfillment interactor
func NewFulfillmentInteractor(r repository.Fulfillment) *FulfillmentInteractor {
	return &FulfillmentInteractor{
		repository: r,
	}
}

// AddCandidateShortlistActivity create an activity when employer shortlisted a candidate
func (fi *FulfillmentInteractor) AddCandidateShortlistActivity(clientID string, clientName string, supplierID string, supplierName string, candidateID string, candidateName string, talentRequestID string, talentRequestNumber string, auctionID string, auctionNumber string, jobTitle string, employerTenantID string, agencyTenantID string, token string) (stream.Activity, error) {
	activity, err := fi.repository.AddCandidateShortlistActivity(clientID, clientName, supplierID, supplierName, candidateID, candidateName, talentRequestID, talentRequestNumber, auctionID, auctionNumber, jobTitle, employerTenantID, agencyTenantID, token)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}

// AddCandidateDeclineActivity create an activity when employer declined a candidate
func (fi *FulfillmentInteractor) AddCandidateDeclineActivity(clientID string, clientName string, supplierID string, supplierName string, candidateID string, candidateName string, talentRequestID string, talentRequestNumber string, auctionID string, auctionNumber string, jobTitle string, employerTenantID string, agencyTenantID string) (stream.Activity, error) {
	activity, err := fi.repository.AddCandidateDeclineActivity(clientID, clientName, supplierID, supplierName, candidateID, candidateName, talentRequestID, talentRequestNumber, auctionID, auctionNumber, jobTitle, employerTenantID, agencyTenantID)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}

// AddCandidateHiredActivity create an activity when employer accepted a candidate
func (fi *FulfillmentInteractor) AddCandidateHiredActivity(clientID string, clientName string, supplierID string, supplierName string, candidateID string, candidateName string, talentRequestID string, talentRequestNumber string, auctionID string, auctionNumber string, jobTitle string, employerTenantID string, agencyTenantID string) (stream.Activity, error) {
	activity, err := fi.repository.AddCandidateHiredActivity(clientID, clientName, supplierID, supplierName, candidateID, candidateName, talentRequestID, talentRequestNumber, auctionID, auctionNumber, jobTitle, employerTenantID, agencyTenantID)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}

// AddCandidateUpdateActivity create an activity when employer updated a candidate
func (fi *FulfillmentInteractor) AddCandidateUpdateActivity(clientID string, clientName string, supplierID string, supplierName string, candidateID string, candidateName string, talentRequestID string, talentRequestNumber string, auctionID string, auctionNumber string, jobTitle string, employerTenantID string, agencyTenantID string, status string) (stream.Activity, error) {
	activity, err := fi.repository.AddCandidateUpdateActivity(clientID, clientName, supplierID, supplierName, candidateID, candidateName, talentRequestID, talentRequestNumber, auctionID, auctionNumber, jobTitle, employerTenantID, agencyTenantID, status)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}

// AddCandidatePendingStatusActivity create an activity when employer doesn't updated the candidate within 5 days with pending status
func (fi *FulfillmentInteractor) AddCandidatePendingStatusActivity(clientID string, clientName string, supplierID string, supplierName string, candidateID string, candidateName string, talentRequestID string, talentRequestNumber string, auctionID string, auctionNumber string, jobTitle string, employerTenantID string, agencyTenantID string) (stream.Activity, error) {
	activity, err := fi.repository.AddCandidatePendingStatusActivity(clientID, clientName, supplierID, supplierName, candidateID, candidateName, talentRequestID, talentRequestNumber, auctionID, auctionNumber, jobTitle, employerTenantID, agencyTenantID)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}

// AddCandidateShortlistStatusActivity create an activity when employer doesn't updated the candidate within 5 days with shortlist status
func (fi *FulfillmentInteractor) AddCandidateShortlistStatusActivity(clientID string, clientName string, supplierID string, supplierName string, candidateID string, candidateName string, talentRequestID string, talentRequestNumber string, auctionID string, auctionNumber string, jobTitle string, employerTenantID string, agencyTenantID string) (stream.Activity, error) {
	activity, err := fi.repository.AddCandidateShortlistStatusActivity(clientID, clientName, supplierID, supplierName, candidateID, candidateName, talentRequestID, talentRequestNumber, auctionID, auctionNumber, jobTitle, employerTenantID, agencyTenantID)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}

// AddCandidateShortlistSuccessFeeActivity create an activity when employer shortlisted a candidate
func (fi *FulfillmentInteractor) AddCandidateShortlistSuccessFeeActivity(clientID string, clientName string, supplierID string, supplierName string, candidateID string, candidateName string, talentRequestID string, talentRequestNumber string, successFeeID string, successFeeNumber string, jobTitle string, employerTenantID string, agencyTenantID string, token string) (stream.Activity, error) {
	activity, err := fi.repository.AddCandidateShortlistSuccessFeeActivity(clientID, clientName, supplierID, supplierName, candidateID, candidateName, talentRequestID, talentRequestNumber, successFeeID, successFeeNumber, jobTitle, employerTenantID, agencyTenantID, token)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}

// AddCandidateDeclineSuccessFeeActivity create an activity when employer declined a candidate
func (fi *FulfillmentInteractor) AddCandidateDeclineSuccessFeeActivity(clientID string, clientName string, supplierID string, supplierName string, candidateID string, candidateName string, talentRequestID string, talentRequestNumber string, successFeeID string, successFeeNumber string, jobTitle string, employerTenantID string, agencyTenantID string) (stream.Activity, error) {
	activity, err := fi.repository.AddCandidateDeclineSuccessFeeActivity(clientID, clientName, supplierID, supplierName, candidateID, candidateName, talentRequestID, talentRequestNumber, successFeeID, successFeeNumber, jobTitle, employerTenantID, agencyTenantID)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}

// AddCandidateHiredSuccessFeeActivity create an activity when employer accepted a candidate
func (fi *FulfillmentInteractor) AddCandidateHiredSuccessFeeActivity(clientID string, clientName string, supplierID string, supplierName string, candidateID string, candidateName string, talentRequestID string, talentRequestNumber string, successFeeID string, successFeeNumber string, jobTitle string, employerTenantID string, agencyTenantID string) (stream.Activity, error) {
	activity, err := fi.repository.AddCandidateHiredSuccessFeeActivity(clientID, clientName, supplierID, supplierName, candidateID, candidateName, talentRequestID, talentRequestNumber, successFeeID, successFeeNumber, jobTitle, employerTenantID, agencyTenantID)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}

// AddCandidateUpdateSuccessFeeActivity create an activity when employer updated a candidate
func (fi *FulfillmentInteractor) AddCandidateUpdateSuccessFeeActivity(clientID string, clientName string, supplierID string, supplierName string, candidateID string, candidateName string, talentRequestID string, talentRequestNumber string, successFeeID string, successFeeNumber string, jobTitle string, employerTenantID string, agencyTenantID string, status string) (stream.Activity, error) {
	activity, err := fi.repository.AddCandidateUpdateSuccessFeeActivity(clientID, clientName, supplierID, supplierName, candidateID, candidateName, talentRequestID, talentRequestNumber, successFeeID, successFeeNumber, jobTitle, employerTenantID, agencyTenantID, status)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}

// AddCandidatePendingStatusSuccessFeeActivity create an activity when employer doesn't updated the candidate within 5 days with pending status
func (fi *FulfillmentInteractor) AddCandidatePendingStatusSuccessFeeActivity(clientID string, clientName string, supplierID string, supplierName string, candidateID string, candidateName string, talentRequestID string, talentRequestNumber string, successFeeID string, successFeeNumber string, jobTitle string, employerTenantID string, agencyTenantID string) (stream.Activity, error) {
	activity, err := fi.repository.AddCandidatePendingStatusSuccessFeeActivity(clientID, clientName, supplierID, supplierName, candidateID, candidateName, talentRequestID, talentRequestNumber, successFeeID, successFeeNumber, jobTitle, employerTenantID, agencyTenantID)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}

// AddCandidateShortlistStatusSuccessFeeActivity create an activity when employer doesn't updated the candidate within 5 days with shortlist status
func (fi *FulfillmentInteractor) AddCandidateShortlistStatusSuccessFeeActivity(clientID string, clientName string, supplierID string, supplierName string, candidateID string, candidateName string, talentRequestID string, talentRequestNumber string, successFeeID string, successFeeNumber string, jobTitle string, employerTenantID string, agencyTenantID string) (stream.Activity, error) {
	activity, err := fi.repository.AddCandidateShortlistStatusSuccessFeeActivity(clientID, clientName, supplierID, supplierName, candidateID, candidateName, talentRequestID, talentRequestNumber, successFeeID, successFeeNumber, jobTitle, employerTenantID, agencyTenantID)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}
