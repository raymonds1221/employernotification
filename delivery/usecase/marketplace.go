package usecase

import "github.com/Ubidy/Ubidy_EmployerNotificationAPI/domain/marketplace"

// MarketplaceInteractor interface for marketplace usecase
type MarketplaceInteractor interface {
	AddAuctionCreatedActivity(competitive marketplace.Competitive) error
	AddAuctionCreatedSuccessFeeActivity(successFee marketplace.SuccessFee) error
}
