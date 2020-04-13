package usecase

import (
	"github.com/Ubidy/Ubidy_EmployerNotificationAPI/usecase/repository"
	stream "gopkg.in/GetStream/stream-go2.v1"
)

// ApplicationInteractor implementation of application usecase
type ApplicationInteractor struct {
	repository repository.Application
}

// NewApplicationInterator create new instance of application interactor
func NewApplicationInterator(r repository.Application) *ApplicationInteractor {
	return &ApplicationInteractor{
		repository: r,
	}
}

// AddApprovedApplicationActivity create an activity when the employer approved agency to the auction
func (ai *ApplicationInteractor) AddApprovedApplicationActivity(clientID string, clientName string, supplierID string, supplierName string, auctionID string, auctionNumber string, employerTenantID string, agencyTenantID string) (stream.Activity, error) {
	activity, err := ai.repository.AddApprovedApplicationActivity(clientID, clientName, supplierID, supplierName, auctionID, auctionNumber, employerTenantID, agencyTenantID)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}

// AddDeclinedApplicationActivity create an activity when the employer declined agency to the auction
func (ai *ApplicationInteractor) AddDeclinedApplicationActivity(clientID string, clientName string, supplierID string, supplierName string, auctionID string, auctionNumber string, employerTenantID string, agencyTenantID string) (stream.Activity, error) {
	activity, err := ai.repository.AddDeclinedApplicationActivity(clientID, clientName, supplierID, supplierName, auctionID, auctionNumber, employerTenantID, agencyTenantID)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}

// AddRevokedApplicationActivity create an activity when the employer revoked agency to the auction
func (ai *ApplicationInteractor) AddRevokedApplicationActivity(clientID string, clientName string, supplierID string, supplierName string, auctionID string, auctionNumber string, employerTenantID string, agencyTenantID string) (stream.Activity, error) {
	activity, err := ai.repository.AddRevokedApplicationActivity(clientID, clientName, supplierID, supplierName, auctionID, auctionNumber, employerTenantID, agencyTenantID)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}

// AddApprovedApplicationSuccessFeeActivity create an activity when the employer approved agency to the success fee
func (ai *ApplicationInteractor) AddApprovedApplicationSuccessFeeActivity(clientID string, clientName string, supplierID string, supplierName string, successFeeID string, successFeeNumber string, employerTenantID string, agencyTenantID string, token string) (stream.Activity, error) {
	activity, err := ai.repository.AddApprovedApplicationSuccessFeeActivity(clientID, clientName, supplierID, supplierName, successFeeID, successFeeNumber, employerTenantID, agencyTenantID, token)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}

// AddDeclinedApplicationSuccessFeeActivity create an activity when the employer declined agency to the success fee
func (ai *ApplicationInteractor) AddDeclinedApplicationSuccessFeeActivity(clientID string, clientName string, supplierID string, supplierName string, successFeeID string, successFeeNumber string, employerTenantID string, agencyTenantID string) (stream.Activity, error) {
	activity, err := ai.repository.AddDeclinedApplicationSuccessFeeActivity(clientID, clientName, supplierID, supplierName, successFeeID, successFeeNumber, employerTenantID, agencyTenantID)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}

// AddRevokedApplicationSuccessFeeActivity create an activity when the employer revoked agency to the success fee
func (ai *ApplicationInteractor) AddRevokedApplicationSuccessFeeActivity(clientID string, clientName string, supplierID string, supplierName string, successFeeID string, successFeeNumber string, employerTenantID string, agencyTenantID string) (stream.Activity, error) {
	activity, err := ai.repository.AddRevokedApplicationSuccessFeeActivity(clientID, clientName, supplierID, supplierName, successFeeID, successFeeNumber, employerTenantID, agencyTenantID)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}
