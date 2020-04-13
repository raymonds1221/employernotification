package controller

import (
	"net/http"

	"github.com/Ubidy/Ubidy_EmployerNotificationAPI/delivery/usecase"
	marketplaceModel "github.com/Ubidy/Ubidy_EmployerNotificationAPI/domain/marketplace"
	"github.com/gin-gonic/gin"
)

// MarketplaceController implementation of marketplace controller
type MarketplaceController struct{}

// AddAuctionCreatedActivity api for sending notification to agency when employer created auction to competitive
func (mc *MarketplaceController) AddAuctionCreatedActivity(mi usecase.MarketplaceInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		auctionID := c.PostForm("auctionID")
		auctionNumber := c.PostForm("auctionNumber")
		employerUserID := c.PostForm("employerUserID")
		employerTenantID := c.PostForm("employerTenantID")
		employerName := c.PostForm("employerName")

		marketplace := marketplaceModel.Marketplace{
			EmployerUserID:   employerUserID,
			EmployerTenantID: employerTenantID,
			EmployerName:     employerName,
		}

		competitive := marketplaceModel.Competitive{
			Marketplace:   marketplace,
			AuctionID:     auctionID,
			AuctionNumber: auctionNumber,
		}

		err := mi.AddAuctionCreatedActivity(competitive)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":         "Unable to send notification to agency",
				"internalError": err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"status": "ok",
		})
	}
}

// AddAuctionCreatedSuccessFeeActivity api for sending notification when employer created auction to success fee
func (mc *MarketplaceController) AddAuctionCreatedSuccessFeeActivity(mi usecase.MarketplaceInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		successFeeID := c.PostForm("successFeeID")
		successFeeNumber := c.PostForm("successFeeNumber")
		employerUserID := c.PostForm("employerUserID")
		employerTenantID := c.PostForm("employerTenantID")
		employerName := c.PostForm("employerName")

		marketplace := marketplaceModel.Marketplace{
			EmployerUserID:   employerUserID,
			EmployerTenantID: employerTenantID,
			EmployerName:     employerName,
		}

		successFee := marketplaceModel.SuccessFee{
			Marketplace:      marketplace,
			SuccessFeeID:     successFeeID,
			SuccessFeeNumber: successFeeNumber,
		}

		err := mi.AddAuctionCreatedSuccessFeeActivity(successFee)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":         "Unable to send notification to agency",
				"internalError": err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"status": "ok",
		})
	}
}
