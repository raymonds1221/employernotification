package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Ubidy/Ubidy_EmployerNotificationAPI/delivery/usecase"
	talentRequestModel "github.com/Ubidy/Ubidy_EmployerNotificationAPI/domain/talentrequest"
	"github.com/gin-gonic/gin"
)

// TalentRequestController implementation of talentrequest interactor
type TalentRequestController struct{}

// AddTalentRequestCancelActivity api for sending notification to agency when employer cancelled talentrequest for success fee
func (tc *TalentRequestController) AddTalentRequestCancelActivity(ti usecase.TalentRequestInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		talentRequest := tc.buildTalentRequestFromContext(c)
		successFeeID := c.PostForm("successFeeID")
		successFeeNumber := c.PostForm("successFeeNumber")

		successFeeTalentRequest := talentRequestModel.SuccessFeeTalentRequest{
			TalentRequest:    talentRequest,
			SuccessFeeID:     successFeeID,
			SuccessFeeNumber: successFeeNumber,
		}

		activity, err := ti.AddTalentRequestCancelActivity(successFeeTalentRequest)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":         "Unable to send notification to agency for cancelled talent request",
				"internalError": err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"status":   "ok",
			"activity": activity,
		})
	}
}

func (tc *TalentRequestController) buildTalentRequestFromContext(c *gin.Context) talentRequestModel.TalentRequest {
	employerUserID := c.PostForm("employerUserID")
	employerTenantID := c.PostForm("employerTenantID")
	employerName := c.PostForm("employerName")
	talentRequestID := c.PostForm("talentRequestID")
	talentRequestNumber, _ := strconv.Atoi(c.PostForm("talentRequestNumber"))
	jobTitle := c.PostForm("jobTitle")
	agencyTenantIDs := c.PostFormArray("agencyTenantIDs")

	return talentRequestModel.TalentRequest{
		EmployerUserID:      employerUserID,
		EmployerTenantID:    employerTenantID,
		EmployerName:        employerName,
		TalentRequestID:     talentRequestID,
		TalentRequestNumber: fmt.Sprintf("%06d", talentRequestNumber),
		JobTitle:            jobTitle,
		AgencyTenantIDs:     agencyTenantIDs,
	}
}
