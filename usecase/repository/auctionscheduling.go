package repository

import (
	stream "gopkg.in/GetStream/stream-go2.v1"
)

// AuctionScheduling interface for auction scheduling repository
type AuctionScheduling interface {
	AddAuctionCreatedActivity(tenantID string, auctionID string) (stream.Activity, error)
	AddAuctionCancelledActivity(tenantID string, auctionID string) (stream.Activity, error)
	AddAuctionUpdatedActivity(tenantID string, auctionID string) (stream.Activity, error)
	AddAuctionDiscontinuedActivity(tenantID string, auctionID string) (stream.Activity, error)
}
