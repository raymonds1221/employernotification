package usecase

import (
	stream "gopkg.in/GetStream/stream-go2.v1"
)

// AuctionSchedulingInteractor interface for auction scheduling usecase
type AuctionSchedulingInteractor interface {
	AddAuctionCreatedActivity(tenantID string, auctionID string) (stream.Activity, error)
	AddAuctionCancelledActivity(tenantID string, auctionID string) (stream.Activity, error)
	AddAuctionUpdatedActivity(tenantID string, auctionID string) (stream.Activity, error)
	AddAuctionDiscontinuedActivity(tenantID string, auctionID string) (stream.Activity, error)
}
