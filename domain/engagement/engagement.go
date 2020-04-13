package engagement

import "database/sql"

const (
	// ApprovedEngagementStatusID status id for approved success fee
	ApprovedEngagementStatusID = 2
	// ApprovedEngagementFeeStatus status id for approved success fee
	ApprovedEngagementFeeStatus = "Approved"
	// OpenEngagementStatusID status id is open
	OpenEngagementStatusID = 1
	// OpenEngagementStatus status is open
	OpenEngagementStatus = "Open"
)

// Engagement model for success fee
type Engagement struct {
	EngagementID       string         `json:"engagementID"`
	ClientID           string         `json:"clientID"`
	ClientName         sql.NullString `json:"clientName"`
	EngagementNumber   int64          `json:"engagementNumber"`
	EngagementStatusID int64          `json:"engagementStatusID"`
	EngagementStatus   string         `json:"engagementStatus"`
}

// NewEngagement create new instanc of success fee model
func NewEngagement(engagementID string, clientID string, clientName sql.NullString, engagementNumber int64, engagementStatusID int64, engagementStatus string) *Engagement {
	return &Engagement{
		EngagementID:       engagementID,
		ClientID:           clientID,
		ClientName:         clientName,
		EngagementNumber:   engagementNumber,
		EngagementStatusID: engagementStatusID,
		EngagementStatus:   engagementStatus,
	}
}

// IsApprovedEngagement check if the success fee is approved
func (e *Engagement) IsApprovedEngagement() bool {
	isApprovedEngagementStatus := e != nil &&
		(e.EngagementStatusID == ApprovedEngagementStatusID && e.EngagementStatus == ApprovedEngagementFeeStatus) ||
		(e.EngagementStatusID == OpenEngagementStatusID && e.EngagementStatus == OpenEngagementStatus)

	if isApprovedEngagementStatus {
		return true
	}
	return false
}
