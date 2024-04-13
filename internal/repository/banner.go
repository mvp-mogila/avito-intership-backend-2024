package repository

type Storage interface {
}

type BannerRepo struct {
	bannerStorage Storage
}

func NewBannerRepo(s Storage) *BannerRepo {
	return &BannerRepo{
		bannerStorage: s,
	}
}
