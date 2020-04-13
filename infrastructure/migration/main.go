package main

import (
	"flag"
	"fmt"
	"log"

	mssql "github.com/Ubidy/Ubidy_EmployerNotificationAPI/infrastructure/repository/mssql"
	mssql2 "github.com/denisenkom/go-mssqldb"
	stream "gopkg.in/GetStream/stream-go2.v1"
)

func main() {
	env := flag.String("env", "uat", "determine the current environment")
	auction := flag.String("type", "competitive", "determine the auction type")
	flag.Parse()
	log.Printf("environment: %s", *env)
	var auctionConfig, queryAgency, queryEmployer string

	if *env == "uat" {
		auctionConfig = uatAuctionConfig()
	} else if *env == "staging" {
		auctionConfig = stagingAuctionConfig()
	} else if *env == "production" {
		auctionConfig = productionAuctionConfig()
	}

	if *auction == "competitive" {
		queryAgency = "select AuctionId, SupplierId from AuctionAgencyAssignments"
		queryEmployer = "select AuctionId, ClientId from AuctionEmployerAssignments"
	} else if *auction == "successfee" {
		queryAgency = "select SuccessFeeId, SupplierId from SuccessFeeAgencyAssignments"
		queryEmployer = "select SuccessFeeId, ClientId from SuccessFeeEmployerAssignments"
	}

	auctionDB, _ := mssql.NewDBConnection(auctionConfig)
	rowsAgency, _ := auctionDB.Query(queryAgency)
	rowsEmployer, _ := auctionDB.Query(queryEmployer)

	defer rowsAgency.Close()
	defer rowsEmployer.Close()

	var auctionID, supplierID, clientID mssql2.UniqueIdentifier
	client := createActivityStreamClient(*env)

	for rowsAgency.Next() {
		err := rowsAgency.Scan(&auctionID, &supplierID)

		if err != nil {
			panic(err)
		}

		supplierFeed := client.FlatFeed("agency", supplierID.String())

		if *auction == "competitive" {
			auctionFeed := client.FlatFeed("auction", auctionID.String())

			supplierFeed.Follow(auctionFeed)
		} else if *auction == "successfee" {
			successFeeFeed := client.FlatFeed("successfee", auctionID.String())

			supplierFeed.Follow(successFeeFeed)
		}

		log.Printf("supplierFeed: %s", supplierFeed.ID())
	}

	for rowsEmployer.Next() {
		err := rowsEmployer.Scan(&auctionID, &clientID)

		if err != nil {
			panic(err)
		}

		clientFeed := client.FlatFeed("employer", auctionID.String())

		if *auction == "competitive" {
			auctionFeed := client.FlatFeed("auction", auctionID.String())

			clientFeed.Follow(auctionFeed)
		} else if *auction == "successfee" {
			successFeeFeed := client.FlatFeed("successfee", auctionID.String())

			clientFeed.Follow(successFeeFeed)
		}

		log.Printf("clientFeed: %s", clientFeed.ID())
	}
}

func uatAuctionConfig() string {
	host := "210.4.126.35"
	db := "UbidyServicesAuctionsDatabase"
	usr := "sa"
	pwd := "Cl@rk2017"

	return fmt.Sprintf("sqlserver://%s:%s@%s?database=%s&connection+timeout=30", usr, pwd, host, db)
}

func stagingAuctionConfig() string {
	host := "ubidyaustraliaeaststaging.database.windows.net"
	db := "UbidyServicesAuctionsDatabase"
	usr := "ubidyuserstaging"
	pwd := "ZZC7jBSIR3oafNrB2Tkt"

	return fmt.Sprintf("sqlserver://%s:%s@%s?database=%s&connection+timeout=30", usr, pwd, host, db)
}

func productionAuctionConfig() string {
	host := "ubidyaustraliaeastprod.database.windows.net"
	db := "UbidyServicesAuctionsDatabase"
	usr := "ubidyuserae"
	pwd := "aiFcjhe5KGsrrJkcQ0Ah"

	return fmt.Sprintf("sqlserver://%s:%s@%s?database=%s&connection+timeout=30", usr, pwd, host, db)
}

func createActivityStreamClient(env string) *stream.Client {
	switch env {
	case "staging":
		key := "azn9639dzp8u"
		secret := "pgw3e7qpvqfmzypj72vw9eucs32vbwqgmn87h7tsmaypkkrnktk5zte8twd8w5m7"
		client, _ := stream.NewClient(key, secret)
		return client
	case "production":
		key := "6zrevvcyq96v"
		secret := "cu6p7scrzwc3tyddaxgek6uykvnrr9cgpe3tjqhrktq5swd6pzbxxdpteudd2ysm"
		client, _ := stream.NewClient(key, secret)
		return client
	default:
		key := "j9uzyqp6cyzq"
		secret := "5y485r8nq9jre4fk6anpu59sqdcpq8xdkuqbd5jxqpvw455gek3aw27ysx4uq7tz"
		client, _ := stream.NewClient(key, secret)
		return client
	}
}
