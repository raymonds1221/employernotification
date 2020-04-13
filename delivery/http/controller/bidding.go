package controller

import (
	"net/http"

	"github.com/Ubidy/Ubidy_EmployerNotificationAPI/delivery/usecase"
	"github.com/gin-gonic/gin"
)

// BiddingController implementation of bidding controller
type BiddingController struct{}

// AddBidChangedPositionActivity api for creating activity when agency position is changed
func (bc *BiddingController) AddBidChangedPositionActivity(bi usecase.BiddingInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		clientID := c.PostForm("clientID")
		supplierID := c.PostForm("supplierID")
		auctionID := c.PostForm("auctionID")

		activity, err := bi.AddBidChangedPositionActivity(clientID, supplierID, auctionID)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"error":           "Unable to create activity when agency position is changed",
				"internalMessage": err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"status":   "ok",
			"activity": activity,
		})
	}
}
