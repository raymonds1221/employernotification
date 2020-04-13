package notification

import stream "gopkg.in/GetStream/stream-go2.v1"

// Notification model for notification
type Notification struct {
	Activities []stream.Activity `json:"activities"`
	FeedID     string            `json:"feedID"`
}
