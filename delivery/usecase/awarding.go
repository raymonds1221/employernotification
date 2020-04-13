package usecase

import (
	stream "gopkg.in/GetStream/stream-go2.v1"
)

// AwardingInteractor interface for awarding usecase
type AwardingInteractor interface {
	AddAwardAgencyActivity(clientID string, supplierID string, auctionID string) (stream.Activity, error)
	AddDeclinedAgencyActivity(clientID string, supplierID string, auctionID string) (stream.Activity, error)
}
