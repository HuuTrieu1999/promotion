package entity

import "time"

type CampaignEntity struct {
	CampaignID  string    `bson:"campaignID"`
	UserID      string    `bson:"userID"`
	CreatedDate time.Time `bson:"createdDate"`
}
