package repository

import (
	stream "gopkg.in/GetStream/stream-go2.v1"
)

// Clarification interface for clarification repository
type Clarification interface {
	AddReplyClarificationActivity(clientID string, clientName string, supplierID string, supplierName string, auctionID string, auctionNumber string, employerTenantID string, agencyTenantID string) (stream.Activity, error)
	AddPostTopicActivity(clientID string, clientName string, supplierID string, supplierName string, auctionID string, auctionNumber string, employerTenantID string, agencyTenantID string) (stream.Activity, error)

	AddReplyClarificationSuccessFeeActivity(clientID string, clientName string, supplierID string, supplierName string, successFeeID string, successFeeNumber string, employerTenantID string, agencyTenantID string) (stream.Activity, error)
	AddPostTopicSuccessFeeActivity(clientID string, clientName string, supplierID string, supplierName string, successFeeID string, successFeeNumber string, employerTenantID string, agencyTenantID string) (stream.Activity, error)
}
