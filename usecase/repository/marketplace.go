package repository

import "github.com/Ubidy/Ubidy_EmployerNotificationAPI/domain/marketplace"

// Marketplace interface for marketplace repository
type Marketplace interface {
	AddAuctionCreatedActivity(competitive marketplace.Competitive) error
	AddAuctionCreatedSuccessFeeActivity(successFee marketplace.SuccessFee) error
}
