package repository

import (
	stream "gopkg.in/GetStream/stream-go2.v1"
)

// Activity implementation of activity repository
type Activity struct {
	client *stream.Client
}

// NewActivityRepository create new instance of activity repository
func NewActivityRepository(client *stream.Client) *Activity {
	return &Activity{
		client: client,
	}
}

// GetActivities retrieve list of activities
func (a *Activity) GetActivities(clientID string) ([]stream.Activity, error) {
	employerFeed := a.client.FlatFeed("employer", clientID)

	resp, err := employerFeed.GetActivities()

	if err != nil {
		return nil, err
	}

	return resp.Results, nil
}
