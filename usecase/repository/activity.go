package repository

import (
	stream "gopkg.in/GetStream/stream-go2.v1"
)

// Activity interface for activity repository
type Activity interface {
	GetActivities(clientID string) ([]stream.Activity, error)
}
