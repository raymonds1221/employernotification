package usecase

import (
	"github.com/Ubidy/Ubidy_EmployerNotificationAPI/usecase/repository"
	stream "gopkg.in/GetStream/stream-go2.v1"
)

// ClarificationInteractor implementation of clarification usecase
type ClarificationInteractor struct {
	repository repository.Clarification
}

// NewClarificationInteractor create new instance of clarification interactor
func NewClarificationInteractor(r repository.Clarification) *ClarificationInteractor {
	return &ClarificationInteractor{
		repository: r,
	}
}

// AddReplyClarificationActivity create an activity when emplyer replied to clarification
func (ci *ClarificationInteractor) AddReplyClarificationActivity(clientID string, clientName string, supplierID string, supplierName string, auctionID string, auctionNumber string, employerTenantID string, agencyTenantID string) (stream.Activity, error) {
	activity, err := ci.repository.AddReplyClarificationActivity(clientID, clientName, supplierID, supplierName, auctionID, auctionNumber, employerTenantID, agencyTenantID)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}

// AddPostTopicActivity create an activity when employer post a new topic to clarification
func (ci *ClarificationInteractor) AddPostTopicActivity(clientID string, clientName string, supplierID string, supplierName string, auctionID string, auctionNumber string, employerTenantID string, agencyTenantID string) (stream.Activity, error) {
	activity, err := ci.repository.AddPostTopicActivity(clientID, clientName, supplierID, supplierName, auctionID, auctionNumber, employerTenantID, agencyTenantID)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}

// AddReplyClarificationSuccessFeeActivity create an activity when emplyer replied to clarification
func (ci *ClarificationInteractor) AddReplyClarificationSuccessFeeActivity(clientID string, clientName string, supplierID string, supplierName string, successFeeID string, successFeeNumber string, employerTenantID string, agencyTenantID string) (stream.Activity, error) {
	activity, err := ci.repository.AddReplyClarificationSuccessFeeActivity(clientID, clientName, supplierID, supplierName, successFeeID, successFeeNumber, employerTenantID, agencyTenantID)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}

// AddPostTopicSuccessFeeActivity create an activity when employer post a new topic to clarification
func (ci *ClarificationInteractor) AddPostTopicSuccessFeeActivity(clientID string, clientName string, supplierID string, supplierName string, successFeeID string, successFeeNumber string, employerTenantID string, agencyTenantID string) (stream.Activity, error) {
	activity, err := ci.repository.AddPostTopicSuccessFeeActivity(clientID, clientName, supplierID, supplierName, successFeeID, successFeeNumber, employerTenantID, agencyTenantID)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}
