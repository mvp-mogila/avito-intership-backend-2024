package usecase

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

type BannerUsecase struct {
	bannerRepo Banners
}

func NewBannerUsecase(r Banners) *BannerUsecase {
	return &BannerUsecase{
		bannerRepo: r,
	}
}
