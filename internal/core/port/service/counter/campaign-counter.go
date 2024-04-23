package counter

type CampaignCounter interface {
	GetAndIncrease(campaignID string) (int64, error)
}
