package usecase

import (
	"github.com/Ubidy/Ubidy_EmployerNotificationAPI/usecase/repository"
	stream "gopkg.in/GetStream/stream-go2.v1"
)

// AuctionSchedulingInteractor implementation for auction scheduling usecase
type AuctionSchedulingInteractor struct {
	repository repository.AuctionScheduling
}

// NewAuctionSchedulingInteractor create new instance of auction scheduling repository
func NewAuctionSchedulingInteractor(r repository.AuctionScheduling) *AuctionSchedulingInteractor {
	return &AuctionSchedulingInteractor{
		repository: r,
	}
}

// AddAuctionCreatedActivity create an activity stream when the employer created an auctionn
func (asi *AuctionSchedulingInteractor) AddAuctionCreatedActivity(tenantID string, auctionID string) (stream.Activity, error) {
	activity, err := asi.repository.AddAuctionCreatedActivity(tenantID, auctionID)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}

// AddAuctionCancelledActivity create an activity stream when the employer cancelled an auction
func (asi *AuctionSchedulingInteractor) AddAuctionCancelledActivity(tenantID string, auctionID string) (stream.Activity, error) {
	activity, err := asi.repository.AddAuctionCancelledActivity(tenantID, auctionID)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}

// AddAuctionUpdatedActivity create an activity stream when the employerr updated an auction
func (asi *AuctionSchedulingInteractor) AddAuctionUpdatedActivity(tenantID string, auctionID string) (stream.Activity, error) {
	activity, err := asi.repository.AddAuctionUpdatedActivity(tenantID, auctionID)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}

// AddAuctionDiscontinuedActivity create an activity stream when the employer discontinued an auction
func (asi *AuctionSchedulingInteractor) AddAuctionDiscontinuedActivity(tenantID string, auctionID string) (stream.Activity, error) {
	activity, err := asi.repository.AddAuctionDiscontinuedActivity(tenantID, auctionID)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}
