package repository

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/Microsoft/ApplicationInsights-Go/appinsights"
	"github.com/Ubidy/Ubidy_EmployerNotificationAPI/domain/marketplace"
	mssql "github.com/denisenkom/go-mssqldb"
	stream "gopkg.in/GetStream/stream-go2.v1"
)

// Marketplace implementation of marketplace repository
type Marketplace struct {
	client          *stream.Client
	agencyDB        *sql.DB
	auctionDB       *sql.DB
	telemetryClient appinsights.TelemetryClient
}

type talentRequestData struct {
	ClassificationID    int
	Classification      string
	SubClassificationID int
	SubClassification   string
	CountryID           int
	Country             string
}

type agencyData struct {
	UserID   string
	UserName string
}

// NewMarketplaceRepository create new instance of marketplace repository
func NewMarketplaceRepository(client *stream.Client, agencyDB *sql.DB, auctionDB *sql.DB, telemetryClient appinsights.TelemetryClient) *Marketplace {
	return &Marketplace{
		client:          client,
		agencyDB:        agencyDB,
		auctionDB:       auctionDB,
		telemetryClient: telemetryClient,
	}
}

// AddAuctionCreatedActivity send notification when employer created new auction in competitive
func (m *Marketplace) AddAuctionCreatedActivity(competitive marketplace.Competitive) error {
	talentRequests := m.getTalentRequestsByAuctionID(competitive.AuctionID)
	agencies := m.getFilteredAgencies(talentRequests)
	agency := make(chan agencyData)

	for _, a := range agencies {
		go func(data agencyData) {
			agency <- data
		}(a)
		go m.sendNotificationToAgency(competitive, talentRequests, agency)
	}
	return nil
}

// AddAuctionCreatedSuccessFeeActivity send notification when employer created new auction in success fee
func (m *Marketplace) AddAuctionCreatedSuccessFeeActivity(successFee marketplace.SuccessFee) error {
	talentRequests := m.getTalentRequestsBySuccessFeeID(successFee.SuccessFeeID)
	agencies := m.getFilteredAgencies(talentRequests)
	agency := make(chan agencyData)

	for _, a := range agencies {
		go func(data agencyData) {
			agency <- data
		}(a)
		go m.sendNotificationToAgencySuccessFee(successFee, talentRequests, agency)
	}

	return nil
}

func (m *Marketplace) getClassificationsAndCountries(query string, id string) []talentRequestData {
	m.auctionDB.Ping()
	rows, err := m.auctionDB.Query(query, id)

	if err != nil {
		m.telemetryClient.TrackException(err)
		return nil
	}

	var classificationID, subClassificationID, countryID int
	var classification, subClassification, country string
	var values []talentRequestData

	for rows.Next() {
		rows.Scan(&classificationID, &classification, &subClassificationID, &subClassification, &countryID, &country)

		data := talentRequestData{
			ClassificationID:    classificationID,
			Classification:      classification,
			SubClassificationID: subClassificationID,
			SubClassification:   subClassification,
			CountryID:           countryID,
			Country:             country,
		}

		values = append(values, data)
	}

	return values
}

func (m *Marketplace) getTalentRequestsByAuctionID(auctionID string) []talentRequestData {
	query := "select ClassificationId, Classification, SubClassificationId, SubClassification, CountryId, Country from TalentRequests where AuctionId=?"
	return m.getClassificationsAndCountries(query, auctionID)
}

func (m *Marketplace) getTalentRequestsBySuccessFeeID(successFeeID string) []talentRequestData {
	query := `select t.ClassificationId, t.Classification, t.SubClassificationId, t.SubClassification, t.CountryId, t.Country 
							from SuccessFeeTalentRequests sf
							inner join TalentRequests t on sf.TalentRequestId=t.TalentRequestId
							where sf.SuccessFeeId=?`
	return m.getClassificationsAndCountries(query, successFeeID)
}

func (m *Marketplace) getAgenciesByClassificationAndCountries(classificationIDs []string, countryIDs []string) []agencyData {
	query := fmt.Sprintf(`select UserId, UserName from Users u
							inner join TenantClassifications  tc on u.TenantId=tc.TenantId
							inner join Areas a on u.TenantId=a.TenantId
							where tc.ClassificationId in (%s)
							and CountryId in (%s)
							and u.IsDeleted=0 and tc.IsDeleted=0 and u.IsActive=1
							group by UserId, UserName`, strings.Join(classificationIDs, ","), strings.Join(countryIDs, ","))
	m.agencyDB.Ping()
	rows, err := m.agencyDB.Query(query)

	if err != nil {
		m.telemetryClient.TrackException(err)
		return nil
	}

	var userID *mssql.UniqueIdentifier
	var userName string
	var values []agencyData

	for rows.Next() {
		rows.Scan(&userID, &userName)

		data := agencyData{
			UserID:   userID.String(),
			UserName: userName,
		}

		values = append(values, data)
	}

	return values
}

func (m *Marketplace) getAgenciesByClassificationAndCountry(classificationID int, countryID int) []agencyData {
	query := `select u.UserId, u.UserName from Tenants t
							inner join Users u on t.TenantId=u.TenantId
							inner join TenantClassifications  tc on t.TenantId=tc.TenantId
							inner join Areas a on t.TenantId=a.TenantId
							where tc.ClassificationId=?
							and a.CountryId=?
							and u.IsDeleted=0 and tc.IsDeleted=0 and u.IsActive=1 and t.isDeleted=0 and t.isDemo=0
							group by UserId, UserName`

	m.agencyDB.Ping()
	rows, err := m.agencyDB.Query(query, classificationID, countryID)

	if err != nil {
		m.telemetryClient.TrackException(err)
		return nil
	}

	var userID *mssql.UniqueIdentifier
	var userName string
	var values []agencyData

	for rows.Next() {
		rows.Scan(&userID, &userName)

		data := agencyData{
			UserID:   userID.String(),
			UserName: userName,
		}

		values = append(values, data)
	}

	return values
}

func (m *Marketplace) getFilteredAgencies(talentRequestData []talentRequestData) []agencyData {
	var agencies []agencyData

	for _, value := range talentRequestData {
		agenciesByClassificationAndCountry := m.getAgenciesByClassificationAndCountry(value.ClassificationID, value.CountryID)

		for _, agency := range agenciesByClassificationAndCountry {
			agencies = append(agencies, agency)
		}
	}

	return agencies
}

func (m *Marketplace) sendNotificationToAgency(competitive marketplace.Competitive, talentRequests []talentRequestData, agency <-chan agencyData) {
	select {
	case a := <-agency:
		helper := NewHelper(m.telemetryClient)
		clientData, _ := helper.createJSONMarshal(competitive.EmployerUserID, competitive.EmployerName, "client")
		verb := "created"
		object := fmt.Sprintf("auction:%s", competitive.AuctionID)
		content := fmt.Sprintf("New Opportunity! %s has just published a new auction which includes a Talent Request that matches your recruitment classification(s). Click View Now to view the application details.", clientData)
		category := "Application"
		subcategory := map[string]string{
			"type":   "Application",
			"status": "Created",
		}
		extra := map[string]interface{}{
			"employer":         fmt.Sprintf("employer:%s", competitive.EmployerUserID),
			"agency":           fmt.Sprintf("agency:%s", a.UserID),
			"employerTenantID": competitive.EmployerTenantID,
			"content":          content,
			"category":         category,
			"subcategory":      subcategory,
		}

		m.sendNotification(a.UserID, verb, object, extra)
	}
}

func (m *Marketplace) sendNotificationToAgencySuccessFee(successFee marketplace.SuccessFee, talentRequests []talentRequestData, agency <-chan agencyData) {
	select {
	case a := <-agency:
		helper := NewHelper(m.telemetryClient)

		clientData, _ := helper.createJSONMarshal(successFee.EmployerUserID, successFee.EmployerName, "client")
		verb := "created"
		object := fmt.Sprintf("successfee:%s", successFee.SuccessFeeID)
		content := fmt.Sprintf("New Opportunity! %s published a new Engagement that matches your recruitment classification(s)", clientData)
		category := "Application"
		subcategory := map[string]string{
			"type":   "Application",
			"status": "Created",
		}
		extra := map[string]interface{}{
			"employer":         fmt.Sprintf("employer:%s", successFee.EmployerUserID),
			"agency":           fmt.Sprintf("agency:%s", a.UserID),
			"employerTenantID": successFee.EmployerTenantID,
			"content":          content,
			"category":         category,
			"subcategory":      subcategory,
		}

		m.sendNotification(a.UserID, verb, object, extra)
	}
}

func (m *Marketplace) sendNotification(userID string, verb string, object string, extra map[string]interface{}) {
	agencyNotificationFeed := m.client.NotificationFeed("agencynotification", strings.ToLower(userID))

	_, err := agencyNotificationFeed.AddActivity(stream.Activity{
		Actor:     agencyNotificationFeed.ID(),
		Verb:      verb,
		Object:    object,
		ForeignID: agencyNotificationFeed.ID(),
		Time:      stream.Time{time.Now().UTC()},
		Extra:     extra,
	})

	if err != nil {
		m.telemetryClient.TrackException(err)
		panic(err)
	}
}
