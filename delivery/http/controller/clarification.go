package controller

import (
	"net/http"

	"github.com/Ubidy/Ubidy_EmployerNotificationAPI/delivery/usecase"
	"github.com/gin-gonic/gin"
)

// ClarificationController implementation of clarification controller
type ClarificationController struct{}

// AddReplyClarificationActivity api for creating activity stream when employer reply to the clarification discussion
func (cc *ClarificationController) AddReplyClarificationActivity(ci usecase.ClarificationInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		clientID := c.PostForm("clientID")
		clientName := c.PostForm("clientName")
		supplierID := c.PostForm("supplierID")
		supplierName := c.PostForm("supplierName")
		auctionID := c.PostForm("auctionID")
		auctionNumber := c.PostForm("auctionNumber")
		employerTenantID := c.PostForm("employerTenantID")
		agencyTenantID := c.PostForm("agencyTenantID")

		activity, err := ci.AddReplyClarificationActivity(clientID, clientName, supplierID, supplierName, auctionID, auctionNumber, employerTenantID, agencyTenantID)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"error":           "Unable to create activity for employer replied to discussion",
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

// AddPostTopicActivity api for creating activity stream when employer reply to the clarification discussion
func (cc *ClarificationController) AddPostTopicActivity(ci usecase.ClarificationInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		clientID := c.PostForm("clientID")
		clientName := c.PostForm("clientName")
		supplierID := c.PostForm("supplierID")
		supplierName := c.PostForm("supplierName")
		auctionID := c.PostForm("auctionID")
		auctionNumber := c.PostForm("auctionNumber")
		employerTenantID := c.PostForm("employerTenantID")
		agencyTenantID := c.PostForm("agencyTenantID")

		activity, err := ci.AddPostTopicActivity(clientID, clientName, supplierID, supplierName, auctionID, auctionNumber, employerTenantID, agencyTenantID)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"error":           "Unable to create activity for employer replied to discussion",
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

// AddReplyClarificationSuccessFeeActivity api for creating activity stream when employer reply to the clarification discussion
func (cc *ClarificationController) AddReplyClarificationSuccessFeeActivity(ci usecase.ClarificationInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		clientID := c.PostForm("clientID")
		clientName := c.PostForm("clientName")
		supplierID := c.PostForm("supplierID")
		supplierName := c.PostForm("supplierName")
		successFeeID := c.PostForm("successFeeID")
		successFeeNumber := c.PostForm("successFeeNumber")
		employerTenantID := c.PostForm("employerTenantID")
		agencyTenantID := c.PostForm("agencyTenantID")

		activity, err := ci.AddReplyClarificationSuccessFeeActivity(clientID, clientName, supplierID, supplierName, successFeeID, successFeeNumber, employerTenantID, agencyTenantID)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"error":           "Unable to create activity for employer replied to discussion",
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

// AddPostTopicSuccessFeeActivity api for creating activity stream when employer reply to the clarification discussion
func (cc *ClarificationController) AddPostTopicSuccessFeeActivity(ci usecase.ClarificationInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		clientID := c.PostForm("clientID")
		clientName := c.PostForm("clientName")
		supplierID := c.PostForm("supplierID")
		supplierName := c.PostForm("supplierName")
		successFeeID := c.PostForm("successFeeID")
		successFeeNumber := c.PostForm("successFeeNumber")
		employerTenantID := c.PostForm("employerTenantID")
		agencyTenantID := c.PostForm("agencyTenantID")

		activity, err := ci.AddPostTopicSuccessFeeActivity(clientID, clientName, supplierID, supplierName, successFeeID, successFeeNumber, employerTenantID, agencyTenantID)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"error":           "Unable to create activity for employer replied to discussion",
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
