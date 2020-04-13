package usecase

import (
	stream "gopkg.in/GetStream/stream-go2.v1"
)

// FulfillmentInteractor interface for fulfillment usecase
type FulfillmentInteractor interface {
	AddCandidateShortlistActivity(clientID string, clientName string, supplierID string, supplierName string, candidateID string, candidateName string, talentRequestID string, talentRequestNumber string, auctionID string, auctionNumber string, jobTitle string, employerTenantID string, agencyTenantID string, token string) (stream.Activity, error)
	AddCandidateDeclineActivity(clientID string, clientName string, supplierID string, supplierName string, candidateID string, candidateName string, talentRequestID string, talentRequestNumber string, auctionID string, auctionNumber string, jobTitle string, employerTenantID string, agencyTenantID string) (stream.Activity, error)
	AddCandidateHiredActivity(clientID string, clientName string, supplierID string, supplierName string, candidateID string, candidateName string, talentRequestID string, talentRequestNumber string, auctionID string, auctionNumber string, jobTitle string, employerTenantID string, agencyTenantID string) (stream.Activity, error)
	AddCandidateUpdateActivity(clientID string, clientName string, supplierID string, supplierName string, candidateID string, candidateName string, talentRequestID string, talentRequestNumber string, auctionID string, auctionNumber string, jobTitle string, employerTenantID string, agencyTenantID string, status string) (stream.Activity, error)
	AddCandidatePendingStatusActivity(clientID string, clientName string, supplierID string, supplierName string, candidateID string, candidateName string, talentRequestID string, talentRequestNumber string, auctionID string, auctionNumber string, jobTitle string, employerTenantID string, agencyTenantID string) (stream.Activity, error)
	AddCandidateShortlistStatusActivity(clientID string, clientName string, supplierID string, supplierName string, candidateID string, candidateName string, talentRequestID string, talentRequestNumber string, auctionID string, auctionNumber string, jobTitle string, employerTenantID string, agencyTenantID string) (stream.Activity, error)

	AddCandidateShortlistSuccessFeeActivity(clientID string, clientName string, supplierID string, supplierName string, candidateID string, candidateName string, talentRequestID string, talentRequestNumber string, successFeeID string, successFeeNumber string, jobTitle string, employerTenantID string, agencyTenantID string, token string) (stream.Activity, error)
	AddCandidateDeclineSuccessFeeActivity(clientID string, clientName string, supplierID string, supplierName string, candidateID string, candidateName string, talentRequestID string, talentRequestNumber string, successFeeID string, successFeeNumber string, jobTitle string, employerTenantID string, agencyTenantID string) (stream.Activity, error)
	AddCandidateHiredSuccessFeeActivity(clientID string, clientName string, supplierID string, supplierName string, candidateID string, candidateName string, talentRequestID string, talentRequestNumber string, successFeeID string, successFeeNumber string, jobTitle string, employerTenantID string, agencyTenantID string) (stream.Activity, error)
	AddCandidateUpdateSuccessFeeActivity(clientID string, clientName string, supplierID string, supplierName string, candidateID string, candidateName string, talentRequestID string, talentRequestNumber string, successFeeID string, successFeeNumber string, jobTitle string, employerTenantID string, agencyTenantID string, status string) (stream.Activity, error)
	AddCandidatePendingStatusSuccessFeeActivity(clientID string, clientName string, supplierID string, supplierName string, candidateID string, candidateName string, talentRequestID string, talentRequestNumber string, successFeeID string, successFeeNumber string, jobTitle string, employerTenantID string, agencyTenantID string) (stream.Activity, error)
	AddCandidateShortlistStatusSuccessFeeActivity(clientID string, clientName string, supplierID string, supplierName string, candidateID string, candidateName string, talentRequestID string, talentRequestNumber string, successFeeID string, successFeeNumber string, jobTitle string, employerTenantID string, agencyTenantID string) (stream.Activity, error)
}
