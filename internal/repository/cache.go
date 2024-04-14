package repository

import (
	"encoding/json"

	"github.com/mvp-mogila/avito-intership-backend-2024/internal/models"
	st "github.com/mvp-mogila/avito-intership-backend-2024/internal/pkg/storage"
	"github.com/mvp-mogila/avito-intership-backend-2024/internal/pkg/utils"
)

type BannerCache struct {
	cache st.Cache
}

func NewBannerCache(c st.Cache) *BannerCache {
	return &BannerCache{
		cache: c,
	}
}

func (c *BannerCache) Set(featID, tagID int, banner models.Banner) error {
	key := utils.MakeUniqueHash(featID, tagID)
	return c.cache.Set(key, banner)
}

func (c *BannerCache) Get(featID, tagID int) (models.Banner, error) {
	key := utils.MakeUniqueHash(featID, tagID)
	res, err := c.cache.Get(key)
	if err != nil {
		return models.Banner{}, err
	}
	var bannerContent map[string]interface{}
	err = json.Unmarshal(res, &bannerContent)
	if err != nil {
		return models.Banner{}, err
	}
	return models.Banner{
		Content: bannerContent,
	}, nil
}
