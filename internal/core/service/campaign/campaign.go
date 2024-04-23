package campaign

import (
	response2 "promotion/internal/controller/http/response"
	"promotion/internal/core/entity/error_code"
	"promotion/internal/core/port/repository"
	"promotion/internal/core/port/service"
	"promotion/internal/core/port/service/counter"
)

type campaignService struct {
	campaignRepo    repository.CampaignRepository
	campaignCounter counter.CampaignCounter
	voucherService  service.VoucherService
	voucherQuota    int64
}

func NewCampaignService(campaignRepo repository.CampaignRepository, campaignCounter counter.CampaignCounter, voucherService service.VoucherService, voucherQuota int64) service.CampaignService {
	return &campaignService{
		campaignRepo:    campaignRepo,
		campaignCounter: campaignCounter,
		voucherService:  voucherService,
		voucherQuota:    voucherQuota,
	}
}

func (u campaignService) createFailedResponse(
	code error_code.ErrorCode, message string,
) *response2.Response {
	return &response2.Response{
		ErrorCode:    code,
		ErrorMessage: message,
		Status:       false,
	}
}

func (u campaignService) createSuccessResponse() *response2.Response {
	return &response2.Response{
		ErrorCode:    error_code.Success,
		ErrorMessage: error_code.SuccessErrMsg,
		Status:       true,
	}
}
