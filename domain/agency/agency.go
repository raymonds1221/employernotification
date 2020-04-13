package agency

// Agency model for agency
type Agency struct {
	TenantID                string `json:"tenantID"`
	OrganizationName        string `json:"organizationName"`
	PrimaryContactFirstName string `json:"primaryContactFirstName"`
	PrimaryContactLastName  string `json:"primaryContactLastName"`
	PrimaryContactEmail     string `json:"primaryContactEmail"`
}

// NewAgency create new instance of agency model
func NewAgency(tenantID string, organizationName string, primaryContactFirstName string, primaryContactLastName string, primaryContactEmail string) *Agency {
	return &Agency{
		TenantID:                tenantID,
		OrganizationName:        organizationName,
		PrimaryContactFirstName: primaryContactFirstName,
		PrimaryContactLastName:  primaryContactLastName,
		PrimaryContactEmail:     primaryContactEmail,
	}
}
