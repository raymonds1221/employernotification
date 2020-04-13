package marketplace

// SuccessFee model for marketplace success fee
type SuccessFee struct {
	Marketplace
	SuccessFeeID     string `json:"successFeeID"`
	SuccessFeeNumber string `json:"successFeeNumber"`
}
