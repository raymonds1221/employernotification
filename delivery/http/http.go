package http

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Ubidy/Ubidy_EmployerNotificationAPI/delivery/http/controller"
	"github.com/Ubidy/Ubidy_EmployerNotificationAPI/delivery/usecase"
	"github.com/Ubidy/Ubidy_EmployerNotificationAPI/infrastructure/repository"
	"github.com/auth0-community/auth0"
	"github.com/gin-gonic/gin"
	jose "gopkg.in/square/go-jose.v2"
)

const apiVersion = "v2.1"
const apiVersion2dot2 = "v2.2"

// App for creating http
type App struct {
	Router *gin.Engine
}

// New create new instance of Http app
func New(
	tokenInteractor usecase.TokenInteractor,
	settingsInteractor usecase.SettingsInteractor,
	auctionSchedulingInteractor usecase.AuctionSchedulingInteractor,
	applicationInteractor usecase.ApplicationInteractor,
	clarificationInteractor usecase.ClarificationInteractor,
	biddingInteractor usecase.BiddingInteractor,
	awardingInteractor usecase.AwardingInteractor,
	fulfillmentInteractor usecase.FulfillmentInteractor,
	activityInteractor usecase.ActivityInteractor,
	notificationInteractor usecase.NotificationInteractor,
	commentInteractor usecase.CommentInteractor,
	marketplaceInteractor usecase.MarketplaceInteractor,
	talentRequestInteractor usecase.TalentRequestInteractor,
) *App {
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,Token,Authorization,X-Requested-With,X-PINGOTHER,Access-Control-Allow-Origin,Accept,x-http-method-override,Referer")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,DELETE,PUT,PATCH,OPTIONS")
		c.Writer.Header().Set("Access-Control-Request-Headers", "authorization,content-type,x-requested-with")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Content-Type", "text/plain")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Pragma", "no-cache")
		c.Writer.Header().Set("Accept", "text/html,application/xhtml+xml,application/x-www-form-urlencoded,application/json,application/xml;q=0.9,*/*;q=0.8")
		c.Next()
	})
	router.OPTIONS("/*cors", func(c *gin.Context) {
	})

	tokenController := controller.TokenController{}
	settingsController := controller.SettingsController{}
	auctionSchedulingController := controller.AuctionSchedulingController{}
	applicationController := controller.ApplicationController{}
	clarificationController := controller.ClarificationController{}
	biddingController := controller.BiddingController{}
	awardingController := controller.AwardingController{}
	fulfillmentController := controller.FulfillmentController{}
	activityController := controller.ActivityController{}
	notificationController := controller.NotificationController{}
	commentController := controller.CommentController{}
	marketplaceController := controller.MarketplaceController{}
	talentRequestController := controller.TalentRequestController{}

	v2 := router.Group(fmt.Sprintf("/api/%s", apiVersion)) //.Use(AuthMiddleware())
	{
		v2.GET("/clients/activities/:clientID", AuthMiddleware(activityController.GetActivities(activityInteractor)))

		v2.GET("/activities/tokens/:userID", AuthMiddleware(tokenController.GetToken(tokenInteractor)))
		v2.GET("/activities/settings/clients/:clientID", AuthMiddleware(settingsController.GetSettingsByClientID(settingsInteractor)))
		v2.GET("/activities/settings/suppliers/:supplierID", AuthMiddleware(settingsController.GetSettingsBySupplierID(settingsInteractor)))
		v2.PATCH("/activities/settings", AuthMiddleware(settingsController.CreateOrUpdateSettings(settingsInteractor)))
		v2.POST("/activities/auctionscheduling/create", AuthMiddleware(auctionSchedulingController.AddAuctionCreatedActivity(auctionSchedulingInteractor)))
		v2.POST("/activities/auctionscheduling/cancel", AuthMiddleware(auctionSchedulingController.AddAuctionCancelledActivity(auctionSchedulingInteractor)))
		v2.POST("/activities/auctionscheduling/update", AuthMiddleware(auctionSchedulingController.AddAuctionUpdatedActivity(auctionSchedulingInteractor)))
		v2.POST("/activities/auctionscheduling/discontinue", AuthMiddleware(auctionSchedulingController.AddAuctionDiscontinuedActivity(auctionSchedulingInteractor)))

		v2.POST("/activities/applications/approve", AuthMiddleware(applicationController.AddApprovedApplicationActivity(applicationInteractor)))
		v2.POST("/activities/applications/decline", AuthMiddleware(applicationController.AddDeclinedApplicationActivity(applicationInteractor)))
		v2.POST("/activities/applications/revoke", AuthMiddleware(applicationController.AddRevokedApplicationActivity(applicationInteractor)))
		v2.POST("/activities/applications/approve/successfee", AuthMiddleware(applicationController.AddApprovedApplicationSuccessFeeActivity(applicationInteractor)))
		v2.POST("/activities/applications/decline/successfee", AuthMiddleware(applicationController.AddDeclinedApplicationSuccessFeeActivity(applicationInteractor)))
		v2.POST("/activities/applications/revoke/successfee", AuthMiddleware(applicationController.AddRevokedApplicationSuccessFeeActivity(applicationInteractor)))

		v2.POST("/activities/clarifications/reply", AuthMiddleware(clarificationController.AddReplyClarificationActivity(clarificationInteractor)))
		v2.POST("/activities/clarifications/post", AuthMiddleware(clarificationController.AddPostTopicActivity(clarificationInteractor)))
		v2.POST("/activities/clarifications/reply/successfee", AuthMiddleware(clarificationController.AddReplyClarificationSuccessFeeActivity(clarificationInteractor)))
		v2.POST("/activities/clarifications/post/successfee", AuthMiddleware(clarificationController.AddPostTopicSuccessFeeActivity(clarificationInteractor)))

		v2.POST("/activities/bidding/positionchange", AuthMiddleware(biddingController.AddBidChangedPositionActivity(biddingInteractor)))

		v2.POST("/activities/awarding/award", AuthMiddleware(awardingController.AddAwardAgencyActivity(awardingInteractor)))
		v2.POST("/activities/awarding/decline", AuthMiddleware(awardingController.AddDeclinedAgencyActivity(awardingInteractor)))

		v2.POST("/activities/fulfillment/shortlist", AuthMiddleware(fulfillmentController.AddCandidateShortlistActivity(fulfillmentInteractor)))
		v2.POST("/activities/fulfillment/decline", AuthMiddleware(fulfillmentController.AddCandidateDeclineActivity(fulfillmentInteractor)))
		v2.POST("/activities/fulfillment/hired", AuthMiddleware(fulfillmentController.AddCandidateHiredActivity(fulfillmentInteractor)))
		v2.POST("/activities/fulfillment/update", AuthMiddleware(fulfillmentController.AddCandidateUpdateActivity(fulfillmentInteractor)))
		v2.POST("/activities/fulfillment/pendingstatus", AuthMiddlewareServiceToService(fulfillmentController.AddCandidatePendingStatusActivity(fulfillmentInteractor)))
		v2.POST("/activities/fulfillment/shortliststatus", AuthMiddlewareServiceToService(fulfillmentController.AddCandidateShortlistStatusActivity(fulfillmentInteractor)))
		v2.POST("/activities/fulfillment/shortlist/successfee", AuthMiddleware(fulfillmentController.AddCandidateShortlistSuccessFeeActivity(fulfillmentInteractor)))
		v2.POST("/activities/fulfillment/decline/successfee", AuthMiddleware(fulfillmentController.AddCandidateDeclineSuccessFeeActivity(fulfillmentInteractor)))
		v2.POST("/activities/fulfillment/hired/successfee", AuthMiddleware(fulfillmentController.AddCandidateHiredSuccessFeeActivity(fulfillmentInteractor)))
		v2.POST("/activities/fulfillment/update/successfee", AuthMiddleware(fulfillmentController.AddCandidateUpdateSuccessFeeActivity(fulfillmentInteractor)))
		v2.POST("/activities/fulfillment/pendingstatus/successfee", AuthMiddlewareServiceToService(fulfillmentController.AddCandidatePendingStatusSuccessFeeActivity(fulfillmentInteractor)))
		v2.POST("/activities/fulfillment/shortliststatus/successfee", AuthMiddlewareServiceToService(fulfillmentController.AddCandidateShortlistStatusSuccessFeeActivity(fulfillmentInteractor)))

		v2.GET("/clients/notifications/:userID", AuthMiddleware(notificationController.GetNotifications(notificationInteractor)))
		v2.GET("/clients/notifications/:userID/:limit/:offset", AuthMiddleware(notificationController.GetNotificationsWithLimitAndOffset(notificationInteractor)))
		v2.PUT("/clients/notifications/archive", AuthMiddleware(notificationController.UpdateNotificationArchive(notificationInteractor)))
		v2.PUT("/clients/notifications/viewed", AuthMiddleware(notificationController.UpdateNotificationViewed(notificationInteractor)))

		v2.POST("/activities/comment", AuthMiddleware(commentController.AddCommentToCandidateActivity(commentInteractor)))
		v2.POST("/activities/comment/successfee", AuthMiddleware(commentController.AddCommentToCandidateSuccessFeeActivity(commentInteractor)))

		v2.POST("/activities/talentrequest/engagement", AuthMiddleware(talentRequestController.AddTalentRequestCancelActivity(talentRequestInteractor)))

		v2.POST("/marketplace/auctioncreated", AuthMiddleware(marketplaceController.AddAuctionCreatedActivity(marketplaceInteractor)))
		v2.POST("/marketplace/auctioncreated/successfee", AuthMiddleware(marketplaceController.AddAuctionCreatedSuccessFeeActivity(marketplaceInteractor)))
	}

	router.GET("/health/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	router.GET("/accesstoken", func(c *gin.Context) {
		helper := repository.NewHelper(nil)
		accessToken := helper.GetPublicAccessToken()
		c.JSON(http.StatusOK, gin.H{"accessToken": accessToken})
	})

	app := App{
		Router: router,
	}

	return &app
}

// Run start the server
func (a *App) Run(addr ...string) error {
	requireTLS := os.Getenv("GO_ENV") == "production" || os.Getenv("GO_ENV") == "uat"

	if requireTLS {
		certFile := "./cert.pem"
		privateKeyFile := "./server.key"
		return a.Router.RunTLS(addr[0], certFile, privateKeyFile)
	}

	return a.Router.Run(addr[0])
}

// AuthMiddleware apply middleware for the Auth0
func AuthMiddleware(f func(*gin.Context)) func(*gin.Context) {
	t := func(c *gin.Context) {
		const JWKSURI = "https://accounts.ubidyapp.com/.well-known/jwks.json"
		const issuer = "https://accounts.ubidyapp.com/"
		var audience = []string{"https://ubidy-api-endpoint/"}

		client := auth0.NewJWKClient(auth0.JWKClientOptions{URI: JWKSURI}, nil)
		config := auth0.NewConfiguration(client, audience, issuer, jose.RS256)

		validator := auth0.NewValidator(config, nil)
		token, err := validator.ValidateRequest(c.Request)

		if err != nil {
			log.Println("Token is not valid or missing token: ", token, err.Error())
			c.JSON(http.StatusForbidden, gin.H{
				"errors:": map[string]string{
					"userMessage":     "Token is not valid or missing token",
					"internalMessage": err.Error(),
				},
			})
			c.Abort()
		} else {
			c.Next()
			f(c)
		}
	}

	return t
}

// AuthMiddlewareServiceToService apply middleware for the Auth0 on service to service microservices
func AuthMiddlewareServiceToService(f func(*gin.Context)) func(*gin.Context) {
	t := func(c *gin.Context) {
		const JWKSURI = "https://accounts.ubidyapp.com/.well-known/jwks.json"
		const issuer = "https://ubidy.au.auth0.com/"
		var audience = []string{"https://ubidy-api-endpoint/"}

		client := auth0.NewJWKClient(auth0.JWKClientOptions{URI: JWKSURI}, nil)
		config := auth0.NewConfiguration(client, audience, issuer, jose.RS256)

		validator := auth0.NewValidator(config, nil)
		token, err := validator.ValidateRequest(c.Request)

		if err != nil {
			log.Println("Token is not valid or missing token: ", token, err.Error())
			c.JSON(http.StatusForbidden, gin.H{
				"errors:": map[string]string{
					"userMessage":     "Token is not valid or missing token",
					"internalMessage": err.Error(),
				},
			})
			c.Abort()
		} else {
			c.Next()
			f(c)
		}
	}

	return t
}
