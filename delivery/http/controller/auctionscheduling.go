package controller

import (
	"net/http"

	"github.com/Ubidy/Ubidy_EmployerNotificationAPI/delivery/usecase"
	"github.com/gin-gonic/gin"
)

// AuctionSchedulingController implementation of auction scheduling controller
type AuctionSchedulingController struct{}

// AddAuctionCreatedActivity api for creating activity stream when employer created an auction
func (asc *AuctionSchedulingController) AddAuctionCreatedActivity(asi usecase.AuctionSchedulingInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		tenantID := c.PostForm("tenantID")
		auctionID := c.PostForm("auctionID")

		activity, err := asi.AddAuctionCreatedActivity(tenantID, auctionID)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"error":           "Unable to create activity stream",
				"internalMessage": "Unable to create activity stream",
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"status":   "ok",
			"activity": activity,
		})
	}
}

// AddAuctionCancelledActivity api for creating activity stream when employer cancelled an auction
func (asc *AuctionSchedulingController) AddAuctionCancelledActivity(asi usecase.AuctionSchedulingInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		tenantID := c.PostForm("tenantID")
		auctionID := c.PostForm("auctionID")

		activity, err := asi.AddAuctionCancelledActivity(tenantID, auctionID)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"error":           "Unable to create activity stream",
				"internalMessage": "Unable to create activity stream",
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"status":   "ok",
			"activity": activity,
		})
	}
}

// AddAuctionUpdatedActivity api for creating activity stream when employer updated an auction
func (asc *AuctionSchedulingController) AddAuctionUpdatedActivity(asi usecase.AuctionSchedulingInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		tenantID := c.PostForm("tenantID")
		auctionID := c.PostForm("auctionID")

		activity, err := asi.AddAuctionUpdatedActivity(tenantID, auctionID)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"error":           "Unable to create activity stream",
				"internalMessage": "Unable to create activity stream",
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"status":     "ok",
			"activityID": activity,
		})
	}
}

// AddAuctionDiscontinuedActivity api for creating activity stream when employer discontinued an auction
func (asc *AuctionSchedulingController) AddAuctionDiscontinuedActivity(asi usecase.AuctionSchedulingInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		tenantID := c.PostForm("tenantID")
		auctionID := c.PostForm("auctionID")

		activity, err := asi.AddAuctionDiscontinuedActivity(tenantID, auctionID)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"error":           "Unable to create activity stream",
				"internalMessage": "Unable to create activity stream",
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"status":     "ok",
			"activityID": activity,
		})
	}
}
