package main

import (
	"os"

	"github.com/Microsoft/ApplicationInsights-Go/appinsights"
	"github.com/Ubidy/Ubidy_EmployerNotificationAPI/delivery/http"
	"github.com/Ubidy/Ubidy_EmployerNotificationAPI/infrastructure/repository"
	"github.com/Ubidy/Ubidy_EmployerNotificationAPI/infrastructure/repository/mssql"
	mssqlconfig "github.com/Ubidy/Ubidy_EmployerNotificationAPI/infrastructure/repository/mssql/config"
	"github.com/Ubidy/Ubidy_EmployerNotificationAPI/usecase"
	stream "gopkg.in/GetStream/stream-go2.v1"
)

var (
	tokenInteractor             *usecase.TokenInteractor
	settingsInteractor          *usecase.SettingsInteractor
	auctionSchedulingInteractor *usecase.AuctionSchedulingInteractor
	applicationInteractor       *usecase.ApplicationInteractor
	clarificationInteractor     *usecase.ClarificationInteractor
	biddingInteractor           *usecase.BiddingInteractor
	awardingInterator           *usecase.AwardingInteractor
	fulfillmentInterator        *usecase.FulfillmentInteractor
	activityInteractor          *usecase.ActivityInteractor
	notificationInteractor      *usecase.NotificationInteractor
	commentInteractor           *usecase.CommentInteractor
	marketplaceInteractor       *usecase.MarketplaceInteractor
	talentRequestInteractor     *usecase.TalentRequestInteractor
)

func init() {
	client := createActivityStreamClient()
	telemetryClient := appinsights.NewTelemetryClient("0a23dc0f-477e-4a6d-b04b-06ad77695a89")

	employerDB, _ := mssql.NewDBConnection(mssqlconfig.NewSQLEmployerConfig())
	agencyDB, _ := mssql.NewDBConnection(mssqlconfig.NewSQLAgencyConfig())
	auctionDB, _ := mssql.NewDBConnection(mssqlconfig.NewSQLAuctionConfig())
	engagementDB, _ := mssql.NewDBConnection(mssqlconfig.NewSQLEngagementConfig())

	tokenRepository := repository.NewTokenRepository(client, telemetryClient)
	settingsRepository := repository.NewSettingsRepository(employerDB, agencyDB, telemetryClient)
	auctionSchedulingRepository := repository.NewAuctionSchedulingRepository(client)
	applicationRepository := repository.NewApplicationRepository(client, agencyDB, auctionDB, telemetryClient)
	clarificationRepository := repository.NewClarificationRepository(client, agencyDB, auctionDB, telemetryClient)
	biddingRepository := repository.NewBiddingRepository(client)
	awardingRepository := repository.NewAwardingRepository(client)
	fulfillmentRepository := repository.NewFulfillmentRepository(client, agencyDB, auctionDB, telemetryClient)
	activityRepository := repository.NewActivityRepository(client)
	notificationRepository := repository.NewNotificationRepository(client)
	commentRepository := repository.NewCommentRepository(client, agencyDB, auctionDB, engagementDB, telemetryClient)
	marketplaceRepository := repository.NewMarketplaceRepository(client, agencyDB, auctionDB, telemetryClient)
	talentRequestRepository := repository.NewTalentRequestRepository(client, agencyDB, auctionDB, telemetryClient)

	tokenInteractor = usecase.NewTokenInteractor(tokenRepository)
	settingsInteractor = usecase.NewSettingsInteractor(settingsRepository)
	auctionSchedulingInteractor = usecase.NewAuctionSchedulingInteractor(auctionSchedulingRepository)
	applicationInteractor = usecase.NewApplicationInterator(applicationRepository)
	clarificationInteractor = usecase.NewClarificationInteractor(clarificationRepository)
	biddingInteractor = usecase.NewBiddingInteractor(biddingRepository)
	awardingInterator = usecase.NewAwardingInteractor(awardingRepository)
	fulfillmentInterator = usecase.NewFulfillmentInteractor(fulfillmentRepository)
	activityInteractor = usecase.NewActivityInteractor(activityRepository)
	notificationInteractor = usecase.NewNotificationInteractor(notificationRepository)
	commentInteractor = usecase.NewCommentInteractor(commentRepository)
	marketplaceInteractor = usecase.NewMarketplaceInteractor(marketplaceRepository)
	talentRequestInteractor = usecase.NewTalentRequestInteractor(talentRequestRepository)
}

func main() {
	http := http.New(
		tokenInteractor,
		settingsInteractor,
		auctionSchedulingInteractor,
		applicationInteractor,
		clarificationInteractor,
		biddingInteractor,
		awardingInterator,
		fulfillmentInterator,
		activityInteractor,
		notificationInteractor,
		commentInteractor,
		marketplaceInteractor,
		talentRequestInteractor,
	)
	http.Run(":5020")
}

func createActivityStreamClient() *stream.Client {
	switch env := os.Getenv("ENV"); env {
	case "staging", "live":
		client, _ := stream.NewClientFromEnv()
		return client
	default:
		key := "3qvb6784xdde"
		secret := "c77yxbqmssdctscbzeckfpv96yk6nmmnyrkq4bab9xr7f4hgm85nkcjsg4gz59jd"
		client, _ := stream.NewClient(key, secret)
		return client
	}
}
