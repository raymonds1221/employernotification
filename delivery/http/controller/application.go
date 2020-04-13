package controller

import (
	"net/http"

	"github.com/Ubidy/Ubidy_EmployerNotificationAPI/delivery/usecase"
	"github.com/gin-gonic/gin"
)

// ApplicationController implementation of application controller
type ApplicationController struct{}

// AddApprovedApplicationActivity api for creating activity stream when employer approved agency to the auction
func (ac *ApplicationController) AddApprovedApplicationActivity(ai usecase.ApplicationInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		clientID := c.PostForm("clientID")
		clientName := c.PostForm("clientName")
		supplierID := c.PostForm("supplierID")
		supplierName := c.PostForm("supplierName")
		auctionID := c.PostForm("auctionID")
		auctionNumber := c.PostForm("auctionNumber")
		employerTenantID := c.PostForm("employerTenantID")
		agencyTenantID := c.PostForm("agencyTenantID")

		activity, err := ai.AddApprovedApplicationActivity(clientID, clientName, supplierID, supplierName, auctionID, auctionNumber, employerTenantID, agencyTenantID)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"message":         "Unable to create activity stream for employer approved agency",
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

// AddDeclinedApplicationActivity api for creating activity stram when employer declined agency to the auction
func (ac *ApplicationController) AddDeclinedApplicationActivity(ai usecase.ApplicationInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		clientID := c.PostForm("clientID")
		clientName := c.PostForm("clientName")
		supplierID := c.PostForm("supplierID")
		supplierName := c.PostForm("supplierName")
		auctionID := c.PostForm("auctionID")
		auctionNumber := c.PostForm("auctionNumber")
		employerTenantID := c.PostForm("employerTenantID")
		agencyTenantID := c.PostForm("agencyTenantID")

		activity, err := ai.AddDeclinedApplicationActivity(clientID, clientName, supplierID, supplierName, auctionID, auctionNumber, employerTenantID, agencyTenantID)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"message":         "Unable to create activity stream for employer declined agency",
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

// AddRevokedApplicationActivity api fo creating activity stream when employer revoked agency to the auction
func (ac *ApplicationController) AddRevokedApplicationActivity(ai usecase.ApplicationInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		clientID := c.PostForm("clientID")
		clientName := c.PostForm("clientName")
		supplierID := c.PostForm("supplierID")
		supplierName := c.PostForm("supplierName")
		auctionID := c.PostForm("auctionID")
		auctionNumber := c.PostForm("auctionNumber")
		employerTenantID := c.PostForm("employerTenantID")
		agencyTenantID := c.PostForm("agencyTenantID")

		activity, err := ai.AddRevokedApplicationActivity(clientID, clientName, supplierID, supplierName, auctionID, auctionNumber, employerTenantID, agencyTenantID)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"message":         "Unable to create activity stream for employer revoked agency",
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

// AddApprovedApplicationSuccessFeeActivity api for creating activity stream when employer approved agency to the success fee
func (ac *ApplicationController) AddApprovedApplicationSuccessFeeActivity(ai usecase.ApplicationInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		clientID := c.PostForm("clientID")
		clientName := c.PostForm("clientName")
		supplierID := c.PostForm("supplierID")
		supplierName := c.PostForm("supplierName")
		successFeeID := c.PostForm("successFeeID")
		successFeeNumber := c.PostForm("successFeeNumber")
		employerTenantID := c.PostForm("employerTenantID")
		agencyTenantID := c.PostForm("agencyTenantID")
		token := c.GetHeader("Authorization")

		activity, err := ai.AddApprovedApplicationSuccessFeeActivity(clientID, clientName, supplierID, supplierName, successFeeID, successFeeNumber, employerTenantID, agencyTenantID, token)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"message":         "Unable to create activity stream for employer approved agency",
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

// AddDeclinedApplicationSuccessFeeActivity api for creating activity stram when employer declined agency to the success fee
func (ac *ApplicationController) AddDeclinedApplicationSuccessFeeActivity(ai usecase.ApplicationInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		clientID := c.PostForm("clientID")
		clientName := c.PostForm("clientName")
		supplierID := c.PostForm("supplierID")
		supplierName := c.PostForm("supplierName")
		successFeeID := c.PostForm("successFeeID")
		successFeeNumber := c.PostForm("successFeeNumber")
		employerTenantID := c.PostForm("employerTenantID")
		agencyTenantID := c.PostForm("agencyTenantID")

		activity, err := ai.AddDeclinedApplicationSuccessFeeActivity(clientID, clientName, supplierID, supplierName, successFeeID, successFeeNumber, employerTenantID, agencyTenantID)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"message":         "Unable to create activity stream for employer declined agency",
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

// AddRevokedApplicationSuccessFeeActivity api fo creating activity stream when employer revoked agency to the success fee
func (ac *ApplicationController) AddRevokedApplicationSuccessFeeActivity(ai usecase.ApplicationInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		clientID := c.PostForm("clientID")
		clientName := c.PostForm("clientName")
		supplierID := c.PostForm("supplierID")
		supplierName := c.PostForm("supplierName")
		successFeeID := c.PostForm("successFeeID")
		successFeeNumber := c.PostForm("successFeeNumber")
		employerTenantID := c.PostForm("employerTenantID")
		agencyTenantID := c.PostForm("agencyTenantID")

		activity, err := ai.AddRevokedApplicationSuccessFeeActivity(clientID, clientName, supplierID, supplierName, successFeeID, successFeeNumber, employerTenantID, agencyTenantID)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"message":         "Unable to create activity stream for employer revoked agency",
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
