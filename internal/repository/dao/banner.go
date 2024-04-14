package dao

import (
	"time"

	"github.com/mvp-mogila/avito-intership-backend-2024/internal/models"
)

type BannerTable struct {
	ID        int                    `db:"id"`
	Content   map[string]interface{} `db:"content"`
	IsActive  bool                   `db:"is_active"`
	FeatureID int                    `db:"feature_id"`
	TagIDs    []int                  `db:"tag_ids"`
	CreatedAt time.Time              `db:"created_at"`
	UpdatedAt time.Time              `db:"updated_at"`
}

func ConvertTablesToModels(tt []BannerTable) []models.Banner {
	banners := make([]models.Banner, 0)
	for _, t := range tt {
		banners = append(banners, models.Banner{
			ID:        t.ID,
			TagIDs:    t.TagIDs,
			FeatureID: t.FeatureID,
			Content:   t.Content,
			IsActive:  t.IsActive,
			CreatedAt: t.CreatedAt,
			UpdatedAt: t.UpdatedAt,
		})
	}
	return banners
}
