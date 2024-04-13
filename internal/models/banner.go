package models

import "time"

type Banner struct {
	ID        int
	FeatureID int
	TagIDs    []int
	Content   map[string]interface{}
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type BannerDefinition struct {
	FeatureID int
	TagIDs    []int
}
