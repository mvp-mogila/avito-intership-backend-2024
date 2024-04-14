package repository

import (
	"encoding/json"
	"fmt"

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
	key := fmt.Sprintf("banner: %s", utils.MakeUniqueHash(featID, tagID))
	data, err := json.Marshal(banner)
	if err != nil {
		return nil
	}
	return c.cache.Set(key, data)
}

func (c *BannerCache) Get(featID, tagID int) (models.Banner, error) {
	key := fmt.Sprintf("banner: %s", utils.MakeUniqueHash(featID, tagID))
	data, err := c.cache.Get(key)
	if err != nil {
		return models.Banner{}, err
	}
	var banner models.Banner
	if err := json.Unmarshal(data, &banner); err != nil {
		return models.Banner{}, err
	}
	return banner, nil
}
