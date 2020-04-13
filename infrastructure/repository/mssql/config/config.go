package config

import (
	"fmt"
	"io/ioutil"
	"os"

	rep "github.com/Ubidy/Ubidy_EmployerNotificationAPI/infrastructure/repository"
)

// NewSQLAgencyConfig create new sql config based on ENV environment variable
func NewSQLAgencyConfig() string {
	configuration := rep.NewConfig()

	host := configuration.DatabaseInfo.Host
	db := configuration.DatabaseInfo.EmployerDatabaseName
	usr := configuration.DatabaseInfo.Username
	pwd := configuration.DatabaseInfo.Password

	connectionString := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s&connection+timeout=30", string(usr), string(pwd), string(host), string(db))

	return connectionString
}

// NewSQLEmployerConfig creat new sql config based on ENV environment variable
func NewSQLEmployerConfig() string {
	configuration := rep.NewConfig()

	host := configuration.DatabaseInfo.Host
	db := configuration.DatabaseInfo.EmployerDatabaseName
	usr := configuration.DatabaseInfo.Username
	pwd := configuration.DatabaseInfo.Password

	connectionString := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s&connection+timeout=30", string(usr), string(pwd), string(host), string(db))

	return connectionString
}

// NewSQLAuctionConfig create new sql config based on ENV environment variable
func NewSQLAuctionConfig() string {
	configuration := rep.NewConfig()

	host := configuration.DatabaseInfo.Host
	db := configuration.DatabaseInfo.AuctionDatabaseName
	usr := configuration.DatabaseInfo.Username
	pwd := configuration.DatabaseInfo.Password

	connectionString := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s&connection+timeout=30", string(usr), string(pwd), string(host), string(db))

	return connectionString
}

// NewSQLEngagementConfig create new sql config based on ENV environment variable
func NewSQLEngagementConfig() string {
	configuration := rep.NewConfig()

	host := configuration.DatabaseInfo.Host
	db := configuration.DatabaseInfo.EngagementDatabaseName
	usr := configuration.DatabaseInfo.Username
	pwd := configuration.DatabaseInfo.Password
	port := configuration.DatabaseInfo.Port

	connectionString := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s&connection+timeout=30", string(usr), string(pwd), string(host), string(port), string(db))

	return connectionString
}

func developmentAgencyConfig() string {
	host := "210.4.126.35"
	db := "UbidyServicesAgenciesDatabase"
	usr := "sa"
	pwd := "Cl@rk2017"

	cs := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s&connection+timeout=30", usr, pwd, host, db)

	return cs
}

func uatAgencyConfig() string {
	host, _ := ioutil.ReadFile(os.Getenv("AZURE_UAT_DB_HOST_FILE"))
	db, _ := ioutil.ReadFile(os.Getenv("AGENCIES_UAT_DB_NAME_FILE"))
	usr, _ := ioutil.ReadFile(os.Getenv("AGENCIES_UAT_DB_USER_FILE"))
	pwd, _ := ioutil.ReadFile(os.Getenv("AGENCIES_UAT_DB_PASSWORD_FILE"))

	cs := fmt.Sprintf("sqlserver://%s:%s@%s:55107?database=%s&connection+timeout=30", string(usr), string(pwd), string(host), string(db))

	return cs
}

func productionAgencyConfig() string {
	host, _ := ioutil.ReadFile(os.Getenv("AZURE_DB_HOST_FILE"))
	db, _ := ioutil.ReadFile(os.Getenv("AGENCIES_DB_NAME_FILE"))
	usr, _ := ioutil.ReadFile(os.Getenv("AGENCIES_DB_USER_FILE"))
	pwd, _ := ioutil.ReadFile(os.Getenv("AGENCIES_DB_PASSWORD_FILE"))

	cs := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s&connection+timeout=30", string(usr), string(pwd), string(host), string(db))

	return cs
}

func developmetEmployerConfig() string {
	host := "210.4.126.35"
	db := "UbidyServicesEmployersDatabase"
	usr := "sa"
	pwd := "Cl@rk2017"

	cs := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s&connection+timeout=30", usr, pwd, host, db)

	return cs
}

func uatEmployerConfig() string {
	host, _ := ioutil.ReadFile(os.Getenv("AZURE_UAT_DB_HOST_FILE"))
	db, _ := ioutil.ReadFile(os.Getenv("EMPLOYERS_UAT_DB_NAME_FILE"))
	usr, _ := ioutil.ReadFile(os.Getenv("EMPLOYERS_UAT_DB_USER_FILE"))
	pwd, _ := ioutil.ReadFile(os.Getenv("EMPLOYERS_UAT_DB_PASSWORD_FILE"))

	cs := fmt.Sprintf("sqlserver://%s:%s@%s:55107?database=%s&connection+timeout=30", usr, pwd, host, db)

	return cs
}

func productionEmployerConfig() string {
	host, _ := ioutil.ReadFile(os.Getenv("AZURE_DB_HOST_FILE"))
	db, _ := ioutil.ReadFile(os.Getenv("EMPLOYERS_DB_NAME_FILE"))
	usr, _ := ioutil.ReadFile(os.Getenv("EMPLOYERS_DB_USER_FILE"))
	pwd, _ := ioutil.ReadFile(os.Getenv("EMPLOYERS_DB_PASSWORD_FILE"))

	cs := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s&connection+timeout=30", usr, pwd, host, db)

	return cs
}

func developmentAuctionConfig() string {
	host := "210.4.126.35"
	db := "UbidyServicesAuctionsDatabase"
	usr := "sa"
	pwd := "Cl@rk2017"

	cs := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s&connection+timeout=30", usr, pwd, host, db)

	return cs
}

func uatAuctionConfig() string {
	host, _ := ioutil.ReadFile(os.Getenv("AZURE_UAT_DB_HOST_FILE"))
	db, _ := ioutil.ReadFile(os.Getenv("AUCTIONS_UAT_DB_NAME_FILE"))
	usr, _ := ioutil.ReadFile(os.Getenv("AUCTIONS_UAT_DB_USER_FILE"))
	pwd, _ := ioutil.ReadFile(os.Getenv("AUCTIONS_UAT_DB_PASSWORD_FILE"))

	cs := fmt.Sprintf("sqlserver://%s:%s@%s:55107?database=%s&connection+timeout=30", usr, pwd, host, db)

	return cs
}

func productionAuctionConfig() string {
	host, _ := ioutil.ReadFile(os.Getenv("AZURE_DB_HOST_FILE"))
	db, _ := ioutil.ReadFile(os.Getenv("AUCTIONS_DB_NAME_FILE"))
	usr, _ := ioutil.ReadFile(os.Getenv("AUCTIONS_DB_USER_FILE"))
	pwd, _ := ioutil.ReadFile(os.Getenv("AUCTIONS_DB_PASSWORD_FILE"))

	cs := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s&connection+timeout=30", usr, pwd, host, db)

	return cs
}
