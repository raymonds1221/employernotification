package usecase

import (
	"github.com/Ubidy/Ubidy_EmployerNotificationAPI/domain/marketplace"
	"github.com/Ubidy/Ubidy_EmployerNotificationAPI/usecase/repository"
)

// MarketplaceInteractor implementation of marketplace usecase
type MarketplaceInteractor struct {
	repository repository.Marketplace
}

// NewMarketplaceInteractor create new instance of marketplace interactor
func NewMarketplaceInteractor(repository repository.Marketplace) *MarketplaceInteractor {
	return &MarketplaceInteractor{
		repository: repository,
	}
}

// AddAuctionCreatedActivity usecase implementation for notification when employer created auction to competitive
func (mi *MarketplaceInteractor) AddAuctionCreatedActivity(competitive marketplace.Competitive) error {
	err := mi.repository.AddAuctionCreatedActivity(competitive)

	if err != nil {
		return err
	}

	return nil
}

// AddAuctionCreatedSuccessFeeActivity usecase implementation for notification when employer created auction to succes fee
func (mi *MarketplaceInteractor) AddAuctionCreatedSuccessFeeActivity(successFee marketplace.SuccessFee) error {
	err := mi.repository.AddAuctionCreatedSuccessFeeActivity(successFee)

	if err != nil {
		return err
	}

	return nil
}
