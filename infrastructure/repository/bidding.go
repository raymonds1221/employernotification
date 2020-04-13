package repository

import (
	"fmt"

	stream "gopkg.in/GetStream/stream-go2.v1"
)

// Bidding implementation of bidding repository
type Bidding struct {
	client *stream.Client
}

// NewBiddingRepository create new instance of bidding repository
func NewBiddingRepository(client *stream.Client) *Bidding {
	return &Bidding{
		client: client,
	}
}

// AddBidChangedPositionActivity create an activity when agency change position
func (b *Bidding) AddBidChangedPositionActivity(clientID string, supplierID string, auctionID string) (stream.Activity, error) {
	employerFeed := b.client.FlatFeed("employer", clientID)

	resp, err := employerFeed.AddActivity(stream.Activity{
		Actor:     employerFeed.ID(),
		Verb:      "change",
		Object:    fmt.Sprintf("agency:%s", supplierID),
		Target:    fmt.Sprintf("auction:%s", auctionID),
		ForeignID: clientID,
	})

	if err != nil {
		return stream.Activity{}, err
	}

	return resp.Activity, nil
}
