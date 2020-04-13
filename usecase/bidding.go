package usecase

import (
	"github.com/Ubidy/Ubidy_EmployerNotificationAPI/usecase/repository"
	stream "gopkg.in/GetStream/stream-go2.v1"
)

// BiddingInteractor implementation of bidding usecase
type BiddingInteractor struct {
	repository repository.Bidding
}

// NewBiddingInteractor create new instance of bidding usecase
func NewBiddingInteractor(r repository.Bidding) *BiddingInteractor {
	return &BiddingInteractor{
		repository: r,
	}
}

// AddBidChangedPositionActivity creat an activity when agency position is changed
func (bi *BiddingInteractor) AddBidChangedPositionActivity(clientID string, supplierID string, auctionID string) (stream.Activity, error) {
	activity, err := bi.repository.AddBidChangedPositionActivity(clientID, supplierID, auctionID)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}
