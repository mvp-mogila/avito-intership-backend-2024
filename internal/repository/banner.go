package repository

import (
	st "github.com/mvp-mogila/avito-intership-backend-2024/internal/pkg/storage"
)

type BannerRepo struct {
	bannerStorage st.Storage
}

func NewBannerRepo(s st.Storage) *BannerRepo {
	return &BannerRepo{
		bannerStorage: s,
	}
}
