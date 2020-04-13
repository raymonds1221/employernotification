package usecase

import stream "gopkg.in/GetStream/stream-go2.v1"

// BiddingInteractor interface for bidding usecase
type BiddingInteractor interface {
	AddBidChangedPositionActivity(clientID string, supplierID string, auctionID string) (stream.Activity, error)
}
