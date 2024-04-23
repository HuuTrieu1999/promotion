package repository

type CampaignRepository interface {
	Insert(campaignID string, userID string) error
}
