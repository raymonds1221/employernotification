package repository

import (
	"database/sql"

	"github.com/Ubidy/Ubidy_EmployerNotificationAPI/domain/engagement"
)

// Engagement implementation of Auction repository
type Engagement struct {
	db *sql.DB
}

// NewEngagementRepository create new instance of auction repository
func NewEngagementRepository(engagementDB *sql.DB) *Engagement {
	return &Engagement{
		db: engagementDB,
	}
}

func (a *Engagement) getEngagementByID(engagementID string) *engagement.Engagement {
	query := "SELECT ClientID, ClientName, EngagementNumber, EngagementStatusId, EngagementStatus FROM Engagements where EngagementId=?"

	a.db.Ping()

	rows, err := a.db.Query(query, engagementID)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if n := rows.Next(); !n {
		return nil
	}

	var engagementNumber, engagementStatusID int64
	var clientID, engagementStatus string
	var clientName sql.NullString

	err = rows.Scan(&clientID, &clientName, &engagementNumber, &engagementStatusID, &engagementStatus)

	if err != nil {
		panic(err)
	}

	return engagement.NewEngagement(engagementID, clientID, clientName, engagementNumber, engagementStatusID, engagementStatus)
}
