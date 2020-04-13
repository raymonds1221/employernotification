package repository

import (
	"database/sql"
	"fmt"

	"github.com/Ubidy/Ubidy_EmployerNotificationAPI/domain/agency"
)

// Agency implementation of agency repository
type Agency struct {
	agencyDB *sql.DB
}

func (a *Agency) getAgencyByID(agencyID string) (*agency.Agency, error) {
	query := "SELECT TenantID, OrganizationName, PrimaryContactFirstName, PrimaryContactLastName, PrimaryContactEmail FROM Tenants WHERE TenantId=?"

	a.agencyDB.Ping()

	rows, err := a.agencyDB.Query(query, agencyID)

	if err != nil {
		return nil, err
	}

	if n := rows.Next(); !n {
		return nil, fmt.Errorf("Unable to find agency with agencyID: %s", agencyID)
	}

	defer rows.Close()

	var tenantID, organizationName, primaryContactFirstName, primaryContactLastName, primaryContactEmail string

	err = rows.Scan(&tenantID, &organizationName, &primaryContactFirstName, &primaryContactLastName, &primaryContactEmail)

	if err != nil {
		return nil, err
	}

	return agency.NewAgency(tenantID, organizationName, primaryContactFirstName, primaryContactLastName, primaryContactEmail), nil
}

func (a *Agency) getAllAgencies() ([]*agency.Agency, error) {
	query := "SELECT TenantID, OrganizationName, PrimaryContactFirstName, PrimaryContactLastName, PrimaryContactEmail FROM Tenants"

	a.agencyDB.Ping()

	rows, err := a.agencyDB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tenantID, organizationName, primaryContactFirstName, primaryContactLastName, primaryContactEmail string
	var agencies []*agency.Agency

	for n := true; n; n = rows.Next() {
		err = rows.Scan(&tenantID, &organizationName, &primaryContactFirstName, &primaryContactLastName, &primaryContactEmail)

		if err != nil {
			return nil, err
		}

		agencyItem := agency.NewAgency(tenantID, organizationName, primaryContactFirstName, primaryContactLastName, primaryContactEmail)
		agencies = append(agencies, agencyItem)
	}

	return agencies, nil
}
