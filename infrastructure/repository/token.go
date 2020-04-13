package repository

import (
	"github.com/Microsoft/ApplicationInsights-Go/appinsights"
	stream "gopkg.in/GetStream/stream-go2.v1"
)

// Token implementation of token repository
type Token struct {
	client          *stream.Client
	telemetryClient appinsights.TelemetryClient
}

// NewTokenRepository create new instance of token repository
func NewTokenRepository(c *stream.Client, telemetryClient appinsights.TelemetryClient) *Token {
	return &Token{
		client:          c,
		telemetryClient: telemetryClient,
	}
}

// GetToken get the token from activity stream
func (t *Token) GetToken(userID string) string {
	employerFeed := t.client.NotificationFeed("employernotification", userID)
	return employerFeed.RealtimeToken(true)
}

// GetUnreadNotificationCount get unread notification count from activity stream
func (t *Token) GetUnreadNotificationCount(userID string) int {
	employerFeed := t.client.NotificationFeed("employernotification", userID)
	resp, err := employerFeed.GetActivities()

	if err != nil {
		t.telemetryClient.TrackException(err)
	}

	return resp.Unread
}
