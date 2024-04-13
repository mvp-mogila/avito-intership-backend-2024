package delivery

import (
	"context"

	"github.com/mvp-mogila/avito-intership-backend-2024/internal/models"
)

type Banners interface {
	Create(ctx context.Context, bannerData models.Banner) (int, error)
	Get(ctx context.Context, params models.BannerDefinition) (models.Banner, error)
	GetFiltered(ctx context.Context, params models.BannerDefinition, limit, offset int) ([]models.Banner, error)
	Update(ctx context.Context, bannerID int, bannerData models.Banner) error
	Delete(ctx context.Context, bannerID int) error
}

type BannerHandler struct {
	bannerUsecase Banners
}

func NewBannerHandler(u Banners) *BannerHandler {
	return &BannerHandler{
		bannerUsecase: u,
	}
}
