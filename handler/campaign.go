package handler

import (
	"bwastartup/campaign"
	"bwastartup/helper"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaign, err := h.service.GetCampaigns(userID)
	if err != nil {
		response := helper.APIResponse("error to get campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("success to get campaigns", http.StatusOK, "success", campaign)
	c.JSON(http.StatusOK, response)

}
