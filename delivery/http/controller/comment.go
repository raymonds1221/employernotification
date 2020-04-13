package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Ubidy/Ubidy_EmployerNotificationAPI/delivery/usecase"
	commentModel "github.com/Ubidy/Ubidy_EmployerNotificationAPI/domain/comment"
	"github.com/gin-gonic/gin"
)

// CommentController implementation of comment interactor
type CommentController struct{}

// AddCommentToCandidateActivity api for sending notification to agency when employer commented on the candidate
func (cc *CommentController) AddCommentToCandidateActivity(ci usecase.CommentInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		comment := cc.buildCommentFromContext(c)
		auctionID := c.PostForm("auctionID")
		auctionNumber := c.PostForm("auctionNumber")

		auctionComment := commentModel.AuctionComment{
			Comment:       comment,
			AuctionID:     auctionID,
			AuctionNumber: auctionNumber,
		}

		activity, err := ci.AddCommentToCandidateActivity(auctionComment)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":         "Unable to send notification to agency for employer commented to candidate",
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

// AddCommentToCandidateSuccessFeeActivity api for sending notification to agency when employer commented on the candidate for success fee
func (cc *CommentController) AddCommentToCandidateSuccessFeeActivity(ci usecase.CommentInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		comment := cc.buildCommentFromContext(c)
		successFeeID := c.PostForm("successFeeID")
		successFeeNumber := c.PostForm("successFeeNumber")

		successFeeComment := commentModel.SuccessFeeComment{
			Comment:          comment,
			SuccessFeeID:     successFeeID,
			SuccessFeeNumber: successFeeNumber,
		}

		activity, err := ci.AddCommentToCandidateSuccessFeeActivity(successFeeComment)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":         "Unable to send notification to agency for employer commented to candidate",
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

func (cc *CommentController) buildCommentFromContext(c *gin.Context) commentModel.Comment {
	employerUserID := c.PostForm("employerUserID")
	employerTenantID := c.PostForm("employerTenantID")
	employerName := c.PostForm("employerName")
	candidateID := c.PostForm("candidateID")
	candidateName := c.PostForm("candidateName")
	talentRequestID := c.PostForm("talentRequestID")
	talentRequestNumber, _ := strconv.Atoi(c.PostForm("talentRequestNumber"))
	jobTitle := c.PostForm("jobTitle")
	agencyTenantID := c.PostForm("agencyTenantID")

	return commentModel.Comment{
		EmployerUserID:      employerUserID,
		EmployerTenantID:    employerTenantID,
		EmployerName:        employerName,
		CandidateID:         candidateID,
		CandidateName:       candidateName,
		TalentRequestID:     talentRequestID,
		TalentRequestNumber: fmt.Sprintf("%06d", talentRequestNumber),
		JobTitle:            jobTitle,
		AgencyTenantID:      agencyTenantID,
	}
}
