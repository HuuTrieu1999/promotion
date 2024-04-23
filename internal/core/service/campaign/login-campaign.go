package campaign

import (
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"promotion/internal/controller/http/request"
	response2 "promotion/internal/controller/http/response"
	"promotion/internal/core/entity/error_code"
)

func (c campaignService) LoginCampaign(request request.LoginCampaignRequest) *response2.Response {
	// skip login campaign logics
	// this project focus on dispatching vouchers

	// insert userID and campaignID into database
	// if insert successful, it means this is the first time campaign login this campaign
	// if insert fail due to duplicate, do nothing after that
	err := c.campaignRepo.Insert(request.CampaignID, request.UserID)
	if err != nil && mongo.IsDuplicateKeyError(err) {
		return c.createSuccessResponse()
	} else if err != nil {
		log.Printf("fail to insert user login campaign info %s\n", err)
		return c.createFailedResponse(error_code.InternalError, error_code.InternalErrMsg)
	}

	// use atomic integer in redis as a counter
	count, err := c.campaignCounter.GetAndIncrease(request.CampaignID)
	if err != nil {
		log.Printf("fail to GetAndIncrease counter %s\n", err)
		return c.createFailedResponse(error_code.InternalError, error_code.InternalErrMsg)
	}

	if count <= c.voucherQuota {
		err = c.voucherService.CreateDiscountVoucherForTopupFee(request.UserID, 30)
		if err != nil {
			log.Printf("fail to Create Discount Voucher %s\n", err)
			return c.createFailedResponse(error_code.InternalError, error_code.InternalErrMsg)
		}
		log.Printf("Create Discount Voucher, userID %s\n", request.UserID)
	}
	// return success response
	return c.createSuccessResponse()
}
