package marketplace

// Marketplace model for marketplace
type Marketplace struct {
	EmployerUserID   string `json:"employerUserID"`
	EmployerTenantID string `json:"employerTenantID"`
	EmployerName     string `json:"employerName"`
}
