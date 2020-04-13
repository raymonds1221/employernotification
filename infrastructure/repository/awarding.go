package repository

import (
	"fmt"

	stream "gopkg.in/GetStream/stream-go2.v1"
)

// Awarding implementation of awarding repository
type Awarding struct {
	client *stream.Client
}

// NewAwardingRepository create new instance of awarding repository
func NewAwardingRepository(client *stream.Client) *Awarding {
	return &Awarding{
		client: client,
	}
}

// AddAwardAgencyActivity create an activity when the employer award an agency
func (a *Awarding) AddAwardAgencyActivity(clientID string, supplierID string, auctionID string) (stream.Activity, error) {
	employerFeed := a.client.FlatFeed("employer", clientID)

	resp, err := employerFeed.AddActivity(stream.Activity{
		Actor:     employerFeed.ID(),
		Verb:      "award",
		Object:    fmt.Sprintf("agency:%s", supplierID),
		Target:    fmt.Sprintf("auction:%s", auctionID),
		ForeignID: clientID,
	})

	if err != nil {
		return stream.Activity{}, err
	}

	return resp.Activity, nil
}

// AddDeclinedAgencyActivity create an activity when the employer declined an agency
func (a *Awarding) AddDeclinedAgencyActivity(clientID string, supplierID string, auctionID string) (stream.Activity, error) {
	employerFeed := a.client.FlatFeed("employer", clientID)

	resp, err := employerFeed.AddActivity(stream.Activity{
		Actor:     employerFeed.ID(),
		Verb:      "decline",
		Object:    fmt.Sprintf("agency:%s", supplierID),
		Target:    fmt.Sprintf("auction:%s", auctionID),
		ForeignID: clientID,
	})

	if err != nil {
		return stream.Activity{}, err
	}

	return resp.Activity, nil
}
