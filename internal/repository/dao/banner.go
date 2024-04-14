package dao

import (
	"encoding/json"
	"time"

	"github.com/mvp-mogila/avito-intership-backend-2024/internal/models"
	"github.com/mvp-mogila/avito-intership-backend-2024/internal/pkg/utils"
)

type BannerTable struct {
	ID        int       `db:"id"`
	Content   string    `db:"content"`
	IsActive  bool      `db:"is_active"`
	FeatureID int       `db:"feature_id"`
	TagIDs    string    `db:"tag_ids"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func ConvertTablesToModels(tt []BannerTable) []models.Banner {
	banners := make([]models.Banner, 0)
	for _, t := range tt {
		content := make(map[string]interface{})
		json.Unmarshal([]byte(t.Content), &content)
		tagIDs := utils.ConvertTagIDs(t.TagIDs)
		banners = append(banners, models.Banner{
			ID:        t.ID,
			TagIDs:    tagIDs,
			FeatureID: t.FeatureID,
			Content:   content,
			IsActive:  t.IsActive,
			CreatedAt: t.CreatedAt,
			UpdatedAt: t.UpdatedAt,
		})
	}
	return banners
}
