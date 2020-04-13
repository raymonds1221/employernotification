package usecase

import (
	stream "gopkg.in/GetStream/stream-go2.v1"

	"github.com/Ubidy/Ubidy_EmployerNotificationAPI/usecase/repository"
)

// AwardingInteractor implementation of awarding usecase
type AwardingInteractor struct {
	repository repository.Awarding
}

// NewAwardingInteractor create new instance of awarding repository
func NewAwardingInteractor(r repository.Awarding) *AwardingInteractor {
	return &AwardingInteractor{
		repository: r,
	}
}

// AddAwardAgencyActivity create an activity when the employer awarded the agency
func (ai *AwardingInteractor) AddAwardAgencyActivity(clientID string, supplierID string, auctionID string) (stream.Activity, error) {
	activity, err := ai.repository.AddAwardAgencyActivity(clientID, supplierID, auctionID)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}

// AddDeclinedAgencyActivity create an activity when the employer declined the agency
func (ai *AwardingInteractor) AddDeclinedAgencyActivity(clientID string, supplierID string, auctionID string) (stream.Activity, error) {
	activity, err := ai.repository.AddDeclinedAgencyActivity(clientID, supplierID, auctionID)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}
