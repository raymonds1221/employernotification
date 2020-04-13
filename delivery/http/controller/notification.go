package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Ubidy/Ubidy_EmployerNotificationAPI/delivery/usecase"
	"github.com/gin-gonic/gin"
)

// NotificationController controller for notification
type NotificationController struct{}

// GetNotifications api for getting list of notifications
func (nc *NotificationController) GetNotifications(ni usecase.NotificationInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		userID := c.Param("userID")

		activities, err := ni.GetNotifications(userID)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"error":           "Unable to get list of notifications",
				"internalMessage": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status:":       "ok",
			"notifications": activities,
		})
	}
}

// GetNotificationsWithLimitAndOffset api for getting list of notifications with limit and offset
func (nc *NotificationController) GetNotificationsWithLimitAndOffset(ni usecase.NotificationInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		userID := c.Param("userID")
		limit, _ := strconv.Atoi(c.Param("limit"))
		offset, _ := strconv.Atoi(c.Param("offset"))

		activities, err := ni.GetNotificationsWithLimitAndOffset(userID, limit, offset)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"error":           "Unable to get list of notifications",
				"internalMessage": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status:":       "ok",
			"notifications": activities,
		})
	}
}

// UpdateNotificationArchive api for updating notification archive
func (nc *NotificationController) UpdateNotificationArchive(ni usecase.NotificationInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		userID := c.PostForm("userID")
		feedID := c.PostForm("feedID")
		isArchive, _ := strconv.ParseBool(c.PostForm("isArchive"))

		activity, err := ni.UpdateNotificationArchive(userID, feedID, isArchive)

		if err != nil {
			log.Print(err)
			c.JSON(http.StatusNoContent, gin.H{
				"error:":          "Unable to update notification archive",
				"internalMessage": err.Error(),
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"status:":  "ok",
			"activity": activity,
		})
	}
}

// UpdateNotificationViewed api for updating notification viewed
func (nc *NotificationController) UpdateNotificationViewed(ni usecase.NotificationInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		userID := c.PostForm("userID")
		feedID := c.PostForm("feedID")
		isViewed, _ := strconv.ParseBool(c.PostForm("isViewed"))

		activity, err := ni.UpdateNotificationViewed(userID, feedID, isViewed)

		if err != nil {
			log.Print(err)
			c.JSON(http.StatusNoContent, gin.H{
				"error:":          "Unable to update notification archive",
				"internalMessage": err.Error(),
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"status:":  "ok",
			"activity": activity,
		})
	}
}
