package request

type LoginCampaignRequest struct {
	UserID     string `json:"userID"`
	CampaignID string `json:"campaignID"`
}
