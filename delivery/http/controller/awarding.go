package controller

import (
	"net/http"

	"github.com/Ubidy/Ubidy_EmployerNotificationAPI/delivery/usecase"
	"github.com/gin-gonic/gin"
)

// AwardingController implementation of awarding controller
type AwardingController struct{}

// AddAwardAgencyActivity api for creating activity when employer awarded an agency
func (ac *AwardingController) AddAwardAgencyActivity(ai usecase.AwardingInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		clientID := c.PostForm("clientID")
		supplierID := c.PostForm("supplierID")
		auctionID := c.PostForm("auctionID")

		activity, err := ai.AddAwardAgencyActivity(clientID, supplierID, auctionID)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"error":           "Unable to create activity for employer awarded agency",
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

// AddDeclinedAgencyActivity api for creating activity when employer declined agency
func (ac *AwardingController) AddDeclinedAgencyActivity(ai usecase.AwardingInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		clientID := c.PostForm("clientID")
		supplierID := c.PostForm("supplierID")
		auctionID := c.PostForm("auctionID")

		activity, err := ai.AddDeclinedAgencyActivity(clientID, supplierID, auctionID)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"error":           "Unable to create activity for employer declined agency",
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
