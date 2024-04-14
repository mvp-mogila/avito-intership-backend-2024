package models

import "errors"

var (
	ErrValidation       = errors.New("validation error")
	ErrInternal         = errors.New("internal server error")
	ErrNoBanner         = errors.New("banners not found")
	ErrNoAuth           = errors.New("no authentication")
	ErrEmptyCache       = errors.New("not found in cache")
	ErrDuplicateBanners = errors.New("duplicate banners")
)
