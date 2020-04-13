package marketplace

// Competitive model for marketplace competitive
type Competitive struct {
	Marketplace
	AuctionID     string `json:"auctionID"`
	AuctionNumber string `json:"auctionNumber"`
}
