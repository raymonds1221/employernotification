package controller

import (
	"log"
	"net/http"

	"github.com/Ubidy/Ubidy_EmployerNotificationAPI/delivery/usecase"
	"github.com/gin-gonic/gin"
)

// FulfillmentController implementation of fulfillment controller
type FulfillmentController struct{}

// AddCandidateShortlistActivity api for creating activity when employer shortlist a candidate
func (fc *FulfillmentController) AddCandidateShortlistActivity(fi usecase.FulfillmentInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		clientID := c.PostForm("clientID")
		clientName := c.PostForm("clientName")
		supplierID := c.PostForm("supplierID")
		supplierName := c.PostForm("supplierName")
		candidateID := c.PostForm("candidateID")
		candidateName := c.PostForm("candidateName")
		talentRequestID := c.PostForm("talentRequestID")
		talentRequestNumber := c.PostForm("talentRequestNumber")
		auctionID := c.PostForm("auctionID")
		auctionNumber := c.PostForm("auctionNumber")
		jobTitle := c.PostForm("jobTitle")
		employerTenantID := c.PostForm("employerTenantID")
		agencyTenantID := c.PostForm("agencyTenantID")
		token := c.GetHeader("Authorization")

		activity, err := fi.AddCandidateShortlistActivity(clientID, clientName, supplierID, supplierName, candidateID, candidateName, talentRequestID, talentRequestNumber, auctionID, auctionNumber, jobTitle, employerTenantID, agencyTenantID, token)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"error":           "Unable to create activity for employer shortlist a candidate",
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

// AddCandidateDeclineActivity api for creating activity when employer declined a candidate
func (fc *FulfillmentController) AddCandidateDeclineActivity(fi usecase.FulfillmentInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		clientID := c.PostForm("clientID")
		clientName := c.PostForm("clientName")
		supplierID := c.PostForm("supplierID")
		supplierName := c.PostForm("supplierName")
		candidateID := c.PostForm("candidateID")
		candidateName := c.PostForm("candidateName")
		talentRequestID := c.PostForm("talentRequestID")
		talentRequestNumber := c.PostForm("talentRequestNumber")
		auctionID := c.PostForm("auctionID")
		auctionNumber := c.PostForm("auctionNumber")
		jobTitle := c.PostForm("jobTitle")
		employerTenantID := c.PostForm("employerTenantID")
		agencyTenantID := c.PostForm("agencyTenantID")

		activity, err := fi.AddCandidateDeclineActivity(clientID, clientName, supplierID, supplierName, candidateID, candidateName, talentRequestID, talentRequestNumber, auctionID, auctionNumber, jobTitle, employerTenantID, agencyTenantID)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"error":           "Unable to create activity for employer decline a candidate",
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

// AddCandidateHiredActivity api for creating activity when employer accepted a candidate
func (fc *FulfillmentController) AddCandidateHiredActivity(fi usecase.FulfillmentInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		clientID := c.PostForm("clientID")
		clientName := c.PostForm("clientName")
		supplierID := c.PostForm("supplierID")
		supplierName := c.PostForm("supplierName")
		candidateID := c.PostForm("candidateID")
		candidateName := c.PostForm("candidateName")
		talentRequestID := c.PostForm("talentRequestID")
		talentRequestNumber := c.PostForm("talentRequestNumber")
		auctionID := c.PostForm("auctionID")
		auctionNumber := c.PostForm("auctionNumber")
		jobTitle := c.PostForm("jobTitle")
		employerTenantID := c.PostForm("employerTenantID")
		agencyTenantID := c.PostForm("agencyTenantID")

		activity, err := fi.AddCandidateHiredActivity(clientID, clientName, supplierID, supplierName, candidateID, candidateName, talentRequestID, talentRequestNumber, auctionID, auctionNumber, jobTitle, employerTenantID, agencyTenantID)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"error":           "Unable to create activity for employer accept a candidate",
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

// AddCandidateUpdateActivity api for creating activity when employer updated a candidate
func (fc *FulfillmentController) AddCandidateUpdateActivity(fi usecase.FulfillmentInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		clientID := c.PostForm("clientID")
		clientName := c.PostForm("clientName")
		supplierID := c.PostForm("supplierID")
		supplierName := c.PostForm("supplierName")
		candidateID := c.PostForm("candidateID")
		candidateName := c.PostForm("candidateName")
		talentRequestID := c.PostForm("talentRequestID")
		talentRequestNumber := c.PostForm("talentRequestNumber")
		auctionID := c.PostForm("auctionID")
		auctionNumber := c.PostForm("auctionNumber")
		jobTitle := c.PostForm("jobTitle")
		employerTenantID := c.PostForm("employerTenantID")
		agencyTenantID := c.PostForm("agencyTenantID")
		status := c.PostForm("status")

		activity, err := fi.AddCandidateUpdateActivity(clientID, clientName, supplierID, supplierName, candidateID, candidateName, talentRequestID, talentRequestNumber, auctionID, auctionNumber, jobTitle, employerTenantID, agencyTenantID, status)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"error":           "Unable to create activity for employer updated a candidate",
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

// AddCandidatePendingStatusActivity api for creating activity when employer did not respond to pending candidates
func (fc *FulfillmentController) AddCandidatePendingStatusActivity(fi usecase.FulfillmentInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		clientID := c.PostForm("clientID")
		clientName := c.PostForm("clientName")
		supplierID := c.PostForm("supplierID")
		supplierName := c.PostForm("supplierName")
		candidateID := c.PostForm("candidateID")
		candidateName := c.PostForm("candidateName")
		talentRequestID := c.PostForm("talentRequestID")
		talentRequestNumber := c.PostForm("talentRequestNumber")
		auctionID := c.PostForm("auctionID")
		auctionNumber := c.PostForm("auctionNumber")
		jobTitle := c.PostForm("jobTitle")
		employerTenantID := c.PostForm("employerTenantID")
		agencyTenantID := c.PostForm("agencyTenantID")

		activity, err := fi.AddCandidatePendingStatusActivity(clientID, clientName, supplierID, supplierName, candidateID, candidateName, talentRequestID, talentRequestNumber, auctionID, auctionNumber, jobTitle, employerTenantID, agencyTenantID)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"error":           "Unable to create activity for employer not responding the pending candidate",
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

// AddCandidateShortlistStatusActivity api for creating activity when employer did not respond to shortlisted candidates
func (fc *FulfillmentController) AddCandidateShortlistStatusActivity(fi usecase.FulfillmentInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		clientID := c.PostForm("clientID")
		clientName := c.PostForm("clientName")
		supplierID := c.PostForm("supplierID")
		supplierName := c.PostForm("supplierName")
		candidateID := c.PostForm("candidateID")
		candidateName := c.PostForm("candidateName")
		talentRequestID := c.PostForm("talentRequestID")
		talentRequestNumber := c.PostForm("talentRequestNumber")
		auctionID := c.PostForm("auctionID")
		auctionNumber := c.PostForm("auctionNumber")
		jobTitle := c.PostForm("jobTitle")
		employerTenantID := c.PostForm("employerTenantID")
		agencyTenantID := c.PostForm("agencyTenantID")

		activity, err := fi.AddCandidateShortlistStatusActivity(clientID, clientName, supplierID, supplierName, candidateID, candidateName, talentRequestID, talentRequestNumber, auctionID, auctionNumber, jobTitle, employerTenantID, agencyTenantID)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"error":           "Unable to create activity for employer not responding the shortlisted candidate",
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

// AddCandidateShortlistSuccessFeeActivity api for creating activity when employer shortlist a candidate
func (fc *FulfillmentController) AddCandidateShortlistSuccessFeeActivity(fi usecase.FulfillmentInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		clientID := c.PostForm("clientID")
		clientName := c.PostForm("clientName")
		supplierID := c.PostForm("supplierID")
		supplierName := c.PostForm("supplierName")
		candidateID := c.PostForm("candidateID")
		candidateName := c.PostForm("candidateName")
		talentRequestID := c.PostForm("talentRequestId")
		talentRequestNumber := c.PostForm("talentRequestNumber")
		successFeeID := c.PostForm("successFeeID")
		successFeeNumber := c.PostForm("successFeeNumber")
		jobTitle := c.PostForm("jobTitle")
		employerTenantID := c.PostForm("employerTenantID")
		agencyTenantID := c.PostForm("agencyTenantID")
		token := c.GetHeader("Authorization")
		log.Printf("header: %s", token)

		activity, err := fi.AddCandidateShortlistSuccessFeeActivity(clientID, clientName, supplierID, supplierName, candidateID, candidateName, talentRequestID, talentRequestNumber, successFeeID, successFeeNumber, jobTitle, employerTenantID, agencyTenantID, token)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"error":           "Unable to create activity for employer shortlist a candidate",
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

// AddCandidateDeclineSuccessFeeActivity api for creating activity when employer declined a candidate
func (fc *FulfillmentController) AddCandidateDeclineSuccessFeeActivity(fi usecase.FulfillmentInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		clientID := c.PostForm("clientID")
		clientName := c.PostForm("clientName")
		supplierID := c.PostForm("supplierID")
		supplierName := c.PostForm("supplierName")
		candidateID := c.PostForm("candidateID")
		candidateName := c.PostForm("candidateName")
		talentRequestID := c.PostForm("talentRequestID")
		talentRequestNumber := c.PostForm("talentRequestNumber")
		successFeeID := c.PostForm("successFeeID")
		successFeeNumber := c.PostForm("successFeeNumber")
		jobTitle := c.PostForm("jobTitle")
		employerTenantID := c.PostForm("employerTenantID")
		agencyTenantID := c.PostForm("agencyTenantID")

		activity, err := fi.AddCandidateDeclineSuccessFeeActivity(clientID, clientName, supplierID, supplierName, candidateID, candidateName, talentRequestID, talentRequestNumber, successFeeID, successFeeNumber, jobTitle, employerTenantID, agencyTenantID)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"error":           "Unable to create activity for employer decline a candidate",
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

// AddCandidateHiredSuccessFeeActivity api for creating activity when employer accepted a candidate
func (fc *FulfillmentController) AddCandidateHiredSuccessFeeActivity(fi usecase.FulfillmentInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		clientID := c.PostForm("clientID")
		clientName := c.PostForm("clientName")
		supplierID := c.PostForm("supplierID")
		supplierName := c.PostForm("supplierName")
		candidateID := c.PostForm("candidateID")
		candidateName := c.PostForm("candidateName")
		talentRequestID := c.PostForm("talentRequestID")
		talentRequestNumber := c.PostForm("talentRequestNumber")
		successFeeID := c.PostForm("successFeeID")
		successFeeNumber := c.PostForm("successFeeNumber")
		jobTitle := c.PostForm("jobTitle")
		employerTenantID := c.PostForm("employerTenantID")
		agencyTenantID := c.PostForm("agencyTenantID")

		activity, err := fi.AddCandidateHiredSuccessFeeActivity(clientID, clientName, supplierID, supplierName, candidateID, candidateName, talentRequestID, talentRequestNumber, successFeeID, successFeeNumber, jobTitle, employerTenantID, agencyTenantID)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"error":           "Unable to create activity for employer accept a candidate",
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

// AddCandidateUpdateSuccessFeeActivity api for creating activity when employer updated a candidate
func (fc *FulfillmentController) AddCandidateUpdateSuccessFeeActivity(fi usecase.FulfillmentInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		clientID := c.PostForm("clientID")
		clientName := c.PostForm("clientName")
		supplierID := c.PostForm("supplierID")
		supplierName := c.PostForm("supplierName")
		candidateID := c.PostForm("candidateID")
		candidateName := c.PostForm("candidateName")
		talentRequestID := c.PostForm("talentRequestID")
		talentRequestNumber := c.PostForm("talentRequestNumber")
		successFeeID := c.PostForm("successFeeID")
		successFeeNumber := c.PostForm("successFeeNumber")
		jobTitle := c.PostForm("jobTitle")
		employerTenantID := c.PostForm("employerTenantID")
		agencyTenantID := c.PostForm("agencyTenantID")
		status := c.PostForm("status")

		activity, err := fi.AddCandidateUpdateSuccessFeeActivity(clientID, clientName, supplierID, supplierName, candidateID, candidateName, talentRequestID, talentRequestNumber, successFeeID, successFeeNumber, jobTitle, employerTenantID, agencyTenantID, status)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"error":           "Unable to create activity for employer updated a candidate",
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

// AddCandidatePendingStatusSuccessFeeActivity api for creating activity when employer did not respond to pending candidate
func (fc *FulfillmentController) AddCandidatePendingStatusSuccessFeeActivity(fi usecase.FulfillmentInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		clientID := c.PostForm("clientID")
		clientName := c.PostForm("clientName")
		supplierID := c.PostForm("supplierID")
		supplierName := c.PostForm("supplierName")
		candidateID := c.PostForm("candidateID")
		candidateName := c.PostForm("candidateName")
		talentRequestID := c.PostForm("talentRequestID")
		talentRequestNumber := c.PostForm("talentRequestNumber")
		successFeeID := c.PostForm("successFeeID")
		successFeeNumber := c.PostForm("successFeeNumber")
		jobTitle := c.PostForm("jobTitle")
		employerTenantID := c.PostForm("employerTenantID")
		agencyTenantID := c.PostForm("agencyTenantID")

		activity, err := fi.AddCandidatePendingStatusSuccessFeeActivity(clientID, clientName, supplierID, supplierName, candidateID, candidateName, talentRequestID, talentRequestNumber, successFeeID, successFeeNumber, jobTitle, employerTenantID, agencyTenantID)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"error":           "Unable to create activity for employer did not respond to pending candidate",
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

// AddCandidateShortlistStatusSuccessFeeActivity api for creating activity when employer did not respond to shortlisted candidate
func (fc *FulfillmentController) AddCandidateShortlistStatusSuccessFeeActivity(fi usecase.FulfillmentInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		clientID := c.PostForm("clientID")
		clientName := c.PostForm("clientName")
		supplierID := c.PostForm("supplierID")
		supplierName := c.PostForm("supplierName")
		candidateID := c.PostForm("candidateID")
		candidateName := c.PostForm("candidateName")
		talentRequestID := c.PostForm("talentRequestID")
		talentRequestNumber := c.PostForm("talentRequestNumber")
		successFeeID := c.PostForm("successFeeID")
		successFeeNumber := c.PostForm("successFeeNumber")
		jobTitle := c.PostForm("jobTitle")
		employerTenantID := c.PostForm("employerTenantID")
		agencyTenantID := c.PostForm("agencyTenantID")

		activity, err := fi.AddCandidateShortlistStatusSuccessFeeActivity(clientID, clientName, supplierID, supplierName, candidateID, candidateName, talentRequestID, talentRequestNumber, successFeeID, successFeeNumber, jobTitle, employerTenantID, agencyTenantID)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"error":           "Unable to create activity for employer did not respond to shortlisted candidate",
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
