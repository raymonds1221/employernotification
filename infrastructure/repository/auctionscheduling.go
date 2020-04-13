package repository

import (
	stream "gopkg.in/GetStream/stream-go2.v1"
)

// AuctionScheduling implementation of auction scheduling repository
type AuctionScheduling struct {
	client *stream.Client
}

// NewAuctionSchedulingRepository create new instance of auction scheduling repository
func NewAuctionSchedulingRepository(client *stream.Client) *AuctionScheduling {
	return &AuctionScheduling{
		client: client,
	}
}

// AddAuctionCreatedActivity create an activity stream when employer created an auction
func (as *AuctionScheduling) AddAuctionCreatedActivity(tenantID string, auctionID string) (stream.Activity, error) {
	// auction := a.getAuctionByID(auctionID)

	// if auction == nil {
	// 	return stream.Activity{}, fmt.Errorf("No record found for auctionID: %s", auctionID)
	// }

	// employerFeed := as.client.FlatFeed("employer", tenantID)
	// auctionData, _ := json.Marshal(map[string]string{
	// 	"id":   auction.AuctionID,
	// 	"name": fmt.Sprintf("%06d", auction.AuctionNumber),
	// 	"type": "auction",
	// })

	// resp, err := employerFeed.AddActivity(stream.Activity{
	// 	Actor:     tenantID,
	// 	Verb:      "create",
	// 	Object:    auctionID,
	// 	ForeignID: tenantID,
	// 	Extra: map[string]interface{}{
	// 		"action":  "auctioncreated",
	// 		"message": fmt.Sprintf("Auction #%s: Auction created.", string(auctionData)),
	// 	},
	// })

	// if err != nil {
	// 	return stream.Activity{}, err
	// }

	// return resp.Activity, nil
	return stream.Activity{}, nil
}

// AddAuctionCancelledActivity create an actvity stream when employer cancelled an auction
func (as *AuctionScheduling) AddAuctionCancelledActivity(tenantID string, auctionID string) (stream.Activity, error) {
	// auction := a.getAuctionByID(auctionID)

	// if auction == nil {
	// 	return stream.Activity{}, fmt.Errorf("No record found for auctionID: %s", auctionID)
	// }

	// employerFeed := as.client.FlatFeed("employer", tenantID)
	// auctionData, _ := json.Marshal(map[string]string{
	// 	"id":   auction.AuctionID,
	// 	"name": fmt.Sprintf("%06d", auction.AuctionNumber),
	// 	"type": "auction",
	// })

	// resp, err := employerFeed.AddActivity(stream.Activity{
	// 	Actor:     tenantID,
	// 	Verb:      "cancel",
	// 	Object:    auctionID,
	// 	ForeignID: tenantID,
	// 	Extra: map[string]interface{}{
	// 		"action":  "auctioncancelled",
	// 		"message": fmt.Sprintf("Auction #%s: Auction cancelled.", string(auctionData)),
	// 	},
	// })

	// if err != nil {
	// 	return stream.Activity{}, err
	// }

	// return resp.Activity, nil
	return stream.Activity{}, nil
}

// AddAuctionUpdatedActivity create an activity stream when employer updated an auction
func (as *AuctionScheduling) AddAuctionUpdatedActivity(tenantID string, auctionID string) (stream.Activity, error) {
	// auction := a.getAuctionByID(auctionID)

	// if auction == nil {
	// 	return stream.Activity{}, fmt.Errorf("No record found for auctionID: %s", auctionID)
	// }

	// employerFeed := as.client.FlatFeed("employer", tenantID)
	// auctionData, _ := json.Marshal(map[string]string{
	// 	"id":   auction.AuctionID,
	// 	"name": fmt.Sprintf("%06d", auction.AuctionNumber),
	// 	"type": "auction",
	// })

	// resp, err := employerFeed.AddActivity(stream.Activity{
	// 	Actor:     tenantID,
	// 	Verb:      "update",
	// 	Object:    auctionID,
	// 	ForeignID: tenantID,
	// 	Extra: map[string]interface{}{
	// 		"action":  "auctionupdated",
	// 		"message": fmt.Sprintf("Auction #%s: Auction updated.", string(auctionData)),
	// 	},
	// })

	// if err != nil {
	// 	return stream.Activity{}, err
	// }

	// return resp.Activity, nil
	return stream.Activity{}, nil
}

// AddAuctionDiscontinuedActivity create an activity stream when employer discontinued an auction
func (as *AuctionScheduling) AddAuctionDiscontinuedActivity(tenantID string, auctionID string) (stream.Activity, error) {
	// auction := a.getAuctionByID(auctionID)

	// if auction == nil {
	// 	return stream.Activity{}, fmt.Errorf("No record found for auctionID: %s", auctionID)
	// }

	// employerFeed := as.client.FlatFeed("employer", tenantID)
	// auctionData, _ := json.Marshal(map[string]string{
	// 	"id":   auction.AuctionID,
	// 	"name": fmt.Sprintf("%06d", auction.AuctionNumber),
	// 	"type": "auction",
	// })

	// resp, err := employerFeed.AddActivity(stream.Activity{
	// 	Actor:     tenantID,
	// 	Verb:      "discontinued",
	// 	Object:    auctionID,
	// 	ForeignID: tenantID,
	// 	Extra: map[string]interface{}{
	// 		"action":  "auctiondiscontinued",
	// 		"message": fmt.Sprintf("Auction #%s: Auction discontinued.", string(auctionData)),
	// 	},
	// })

	// if err != nil {
	// 	return stream.Activity{}, err
	// }

	// return resp.Activity, nil
	return stream.Activity{}, nil
}
