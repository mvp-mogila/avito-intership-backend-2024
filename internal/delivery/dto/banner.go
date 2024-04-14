package dto

import (
	"time"

	"github.com/mvp-mogila/avito-intership-backend-2024/internal/models"
)

type BannerContent struct {
	Content map[string]interface{}
}

type BannerDetails struct {
	TagIDs    []int                  `json:"tag_ids"`
	FeatureID int                    `json:"feature_id"`
	Content   map[string]interface{} `json:"content"`
	IsActive  bool                   `json:"is_active"`
}

func ConvertBannerDetailsToModel(d BannerDetails) models.Banner {
	return models.Banner{
		TagIDs:    d.TagIDs,
		FeatureID: d.FeatureID,
		Content:   d.Content,
		IsActive:  d.IsActive,
	}
}

type CreateBannerPayload struct {
	BannerID int `json:"banner_id"`
}

type Banner struct {
	BannerID int `json:"banner_id"`
	BannerDetails
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ConvertModelsToDto(mm []models.Banner) []Banner {
	banners := make([]Banner, 0)
	for _, m := range mm {
		banners = append(banners, Banner{
			BannerID: m.ID,
			BannerDetails: BannerDetails{
				TagIDs:    m.TagIDs,
				FeatureID: m.FeatureID,
				Content:   m.Content,
				IsActive:  m.IsActive,
			},
			CreatedAt: m.CreatedAt,
			UpdatedAt: m.UpdatedAt,
		})
	}
	return banners
}
