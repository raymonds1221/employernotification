package repository

import (
	stream "gopkg.in/GetStream/stream-go2.v1"
)

// Application interface for application repository
type Application interface {
	AddApprovedApplicationActivity(clientID string, clientName string, supplierID string, supplierName string, auctionID string, auctionNumber string, employerTenantID string, agencyTenantID string) (stream.Activity, error)
	AddDeclinedApplicationActivity(clientID string, clientName string, supplierID string, supplierName string, auctionID string, auctionNumber string, employerTenantID string, agencyTenantID string) (stream.Activity, error)
	AddRevokedApplicationActivity(clientID string, clientName string, supplierID string, supplierName string, auctionID string, auctionNumber string, employerTenantID string, agencyTenantID string) (stream.Activity, error)

	AddApprovedApplicationSuccessFeeActivity(clientID string, clientName string, supplierID string, supplierName string, successFeeID string, successFeeNumber string, employerTenantID string, agencyTenantID string, token string) (stream.Activity, error)
	AddDeclinedApplicationSuccessFeeActivity(clientID string, clientName string, supplierID string, supplierName string, successFeeID string, successFeeNumber string, employerTenantID string, agencyTenantID string) (stream.Activity, error)
	AddRevokedApplicationSuccessFeeActivity(clientID string, clientName string, supplierID string, supplierName string, successFeeID string, successFeeNumber string, employerTenantID string, agencyTenantID string) (stream.Activity, error)
}
