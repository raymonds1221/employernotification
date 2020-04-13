package controller

import (
	"net/http"

	"github.com/Ubidy/Ubidy_EmployerNotificationAPI/delivery/usecase"
	"github.com/gin-gonic/gin"
)

// ActivityController controller for activity
type ActivityController struct{}

// GetActivities get list of activities
func (ac *ActivityController) GetActivities(ai usecase.ActivityInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		clientID := c.Param("clientID")

		activities, err := ai.GetActivities(clientID)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"error":           "Unable to get list of activities",
				"internalMessage": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status:":    "ok",
			"activities": activities,
		})
	}
}
