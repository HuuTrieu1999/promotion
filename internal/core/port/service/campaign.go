package service

import (
	"promotion/internal/controller/http/request"
	"promotion/internal/controller/http/response"
)

type CampaignService interface {
	LoginCampaign(request request.LoginCampaignRequest) *response.Response
}
