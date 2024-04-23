package http

import (
	"log"
	"promotion/internal/controller/http/request"
	"promotion/internal/controller/http/response"
	"promotion/internal/core/common/utils"
	"promotion/internal/core/entity/error_code"
	"promotion/internal/core/port/service"

	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	invalidRequestResponse = &response.Response{
		ErrorCode:    error_code.InvalidRequest,
		ErrorMessage: error_code.InvalidRequestErrMsg,
		Status:       false,
	}
)

type CampaignController struct {
	gin             *gin.Engine
	campaignService service.CampaignService
}

func NewCampaignController(
	gin *gin.Engine,
	campaignService service.CampaignService,
) CampaignController {
	return CampaignController{
		gin:             gin,
		campaignService: campaignService,
	}
}

func (u CampaignController) InitRouter() {
	api := u.gin.Group("/api")
	api.POST("/login-campaign", u.loginCampaign)
}

func (u CampaignController) loginCampaign(c *gin.Context) {
	var req request.LoginCampaignRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &invalidRequestResponse)
		return
	}
	resp := u.campaignService.LoginCampaign(req)
	log.Printf("login, request %s response %s\n", utils.GetJsonString(req), utils.GetJsonString(resp))
	c.JSON(http.StatusOK, resp)
}
