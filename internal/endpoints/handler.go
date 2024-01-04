package endpoints

import "learning-golang/internal/domain/campaign"

type Handler struct {
	CampaignService campaign.Service
}
