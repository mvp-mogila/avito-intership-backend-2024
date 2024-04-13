package usecase

import (
	"context"

	"github.com/mvp-mogila/avito-intership-backend-2024/internal/models"
	"github.com/mvp-mogila/avito-intership-backend-2024/internal/pkg/utils"
)

type Banners interface {
	Create(ctx context.Context, bannerData models.Banner) (int, error)
	Get(ctx context.Context, params models.BannerDefinition) (models.Banner, error)
	GetAll(ctx context.Context, limit, offset int) ([]models.Banner, error)
	GetFiltered(ctx context.Context, params models.BannerDefinition, limit, offset int) ([]models.Banner, error)
	GetFilteredByTag(ctx context.Context, tagID, limit, offset int) ([]models.Banner, error)
	GetFilteredByFeature(ctx context.Context, featureID, limit, offset int) ([]models.Banner, error)
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

func (u *BannerUsecase) Create(ctx context.Context, bannerData models.Banner) (int, error) {
	if err := utils.ValidatePositive(bannerData.FeatureID, true); err != nil {
		return 0, utils.NewValidationError("invalid feature id")
	}
	for _, tagID := range bannerData.TagIDs {
		if err := utils.ValidatePositive(tagID, true); err != nil {
			return 0, utils.NewValidationError("invalid tag id")
		}
	}

	return u.bannerRepo.Create(ctx, bannerData)
}

func (u *BannerUsecase) Get(ctx context.Context, params models.BannerDefinition) (models.Banner, error) {
	if err := utils.ValidatePositive(params.FeatureID, true); err != nil {
		return models.Banner{}, utils.NewValidationError("invalid feature id")
	}
	if err := utils.ValidatePositive(params.TagID, true); err != nil {
		return models.Banner{}, utils.NewValidationError("invalid tag id")
	}

	return u.bannerRepo.Get(ctx, params)
}

func (u *BannerUsecase) GetFiltered(ctx context.Context, params models.BannerDefinition, limit, offset int) ([]models.Banner, error) {
	if err := utils.ValidatePositive(params.FeatureID, false); err != nil {
		return nil, utils.NewValidationError("invalid feature id")
	}
	if err := utils.ValidatePositive(params.TagID, false); err != nil {
		return nil, utils.NewValidationError("invalid tag id")
	}
	if err := utils.ValidatePositive(limit, true); err != nil {
		return nil, utils.NewValidationError("invalid limit")
	}
	if err := utils.ValidatePositive(offset, false); err != nil {
		return nil, utils.NewValidationError("invalid offset")
	}

	switch {
	case params.FeatureID != 0 && params.TagID != 0:
		return u.bannerRepo.GetFiltered(ctx, params, limit, offset)
	case params.FeatureID != 0 && params.TagID == 0:
		return u.bannerRepo.GetFilteredByFeature(ctx, params.FeatureID, limit, offset)
	case params.FeatureID == 0 && params.TagID != 0:
		return u.bannerRepo.GetFilteredByFeature(ctx, params.FeatureID, limit, offset)
	default:
		return u.bannerRepo.GetAll(ctx, limit, offset)
	}
}

func (u *BannerUsecase) Update(ctx context.Context, bannerID int, bannerData models.Banner) error {
	if err := utils.ValidatePositive(bannerData.FeatureID, true); err != nil {
		return utils.NewValidationError("invalid feature id")
	}
	for _, tagID := range bannerData.TagIDs {
		if err := utils.ValidatePositive(tagID, true); err != nil {
			return utils.NewValidationError("invalid tag id")
		}
	}

	return u.bannerRepo.Update(ctx, bannerID, bannerData)
}

func (u *BannerUsecase) Delete(ctx context.Context, bannerID int) error {
	return u.bannerRepo.Delete(ctx, bannerID)
}
