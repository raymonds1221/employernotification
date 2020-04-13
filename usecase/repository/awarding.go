package repository

import (
	stream "gopkg.in/GetStream/stream-go2.v1"
)

// Awarding interface for awarding repository
type Awarding interface {
	AddAwardAgencyActivity(clientID string, supplierID string, auctionID string) (stream.Activity, error)
	AddDeclinedAgencyActivity(clientID string, supplierID string, auctionID string) (stream.Activity, error)
}
