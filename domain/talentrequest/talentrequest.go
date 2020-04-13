package talentrequest

// TalentRequest base struct for talentrequest
type TalentRequest struct {
	EmployerUserID      string
	EmployerTenantID    string
	EmployerName        string
	TalentRequestID     string
	TalentRequestNumber string
	JobTitle            string
	AgencyTenantIDs     []string
}
