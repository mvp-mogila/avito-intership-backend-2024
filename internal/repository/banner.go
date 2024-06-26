package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/mvp-mogila/avito-intership-backend-2024/internal/models"
	st "github.com/mvp-mogila/avito-intership-backend-2024/internal/pkg/storage"
	"github.com/mvp-mogila/avito-intership-backend-2024/internal/repository/dao"
)

type BannerRepo struct {
	bannerStorage st.Database
}

func NewBannerRepo(d st.Database) *BannerRepo {
	return &BannerRepo{
		bannerStorage: d,
	}
}

func (r *BannerRepo) Create(ctx context.Context, bannerData models.Banner) (int, error) {
	tx, err := r.bannerStorage.Begin(ctx, &sql.TxOptions{})
	if err != nil {
		log.Println(err.Error())
		return 0, models.ErrInternal
	}

	q := `INSERT INTO banner(content, is_active) VALUES($1, $2) RETURNING id;`

	var bannerID int
	err = tx.Get(&bannerID, q, bannerData.Content, bannerData.IsActive)
	if err != nil {
		log.Println(err.Error())
		tx.Rollback()
		return 0, models.ErrDuplicateBanners
	}

	for _, tagID := range bannerData.TagIDs {
		q = `INSERT INTO banner_definition(banner_id, feature_id, tag_id) VALUES($1, $2, $3) RETURNING banner_id;`

		_, err := tx.Exec(q, bannerID, bannerData.FeatureID, tagID)
		if err != nil {
			log.Println(err.Error())
			tx.Rollback()
			return 0, models.ErrDuplicateBanners
		}
	}
	if err = tx.Commit(); err != nil {
		log.Println(err.Error())
		return 0, models.ErrInternal
	}

	return bannerID, nil
}

func (r *BannerRepo) Get(ctx context.Context, params models.BannerDefinition) (models.Banner, error) {
	q := `SELECT b.content
	FROM banner_definition d JOIN banner b ON d.banner_id = b.id
	WHERE d.feature_id = $1 AND d.tag_id = $2 AND b.is_active = true;`

	bannerRow := dao.BannerTable{}
	err := r.bannerStorage.Get(ctx, &bannerRow, q, params.FeatureID, params.TagID)
	if err != nil {
		fmt.Println("Error:" + err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return models.Banner{}, models.ErrNoBanner
		}
		return models.Banner{}, models.ErrInternal
	}

	result := make(map[string]interface{})
	json.Unmarshal([]byte(bannerRow.Content), &result)
	return models.Banner{
		Content: result,
	}, nil
}

func (r *BannerRepo) GetAll(ctx context.Context, limit, offset int) ([]models.Banner, error) {
	q := `SELECT DISTINCT b.id, (SELECT ARRAY_AGG(tag_id) FROM banner_definition d WHERE d.banner_id = b.id) as tag_ids,
	d.feature_id, b.content, b.is_active, b.created_at,	b.updated_at
	FROM banner b JOIN banner_definition d ON b.id = d.banner_id AND b.is_active = true
	ORDER BY b.created_at DESC LIMIT $1 OFFSET $2;`

	bannerRows := []dao.BannerTable{}
	err := r.bannerStorage.Select(ctx, &bannerRows, q, limit, offset)

	if err != nil {
		log.Println(err.Error())
		if len(bannerRows) == 0 {
			return nil, models.ErrNoBanner
		}
		return nil, models.ErrInternal
	}

	banners := dao.ConvertTablesToModels(bannerRows)
	return banners, nil
}

func (r *BannerRepo) GetFiltered(ctx context.Context, params models.BannerDefinition, limit, offset int) ([]models.Banner, error) {
	q := `SELECT DISTINCT b.id, (SELECT ARRAY_AGG(tag_id) FROM banner_definition d WHERE d.banner_id = b.id) as tag_ids,
	d.feature_id, b.content, b.is_active, b.created_at,	b.updated_at
	FROM banner b JOIN banner_definition d ON b.id = d.banner_id
	WHERE d.tag_id = $1 AND d.feature_id = $2 AND b.is_active = true
	ORDER BY b.created_at DESC
	LIMIT $3 OFFSET $4;`

	bannerRows := []dao.BannerTable{}
	err := r.bannerStorage.Select(ctx, &bannerRows, q, params.FeatureID, params.TagID, limit, offset)

	if err != nil {
		if len(bannerRows) == 0 {
			return nil, models.ErrNoBanner
		}
		return nil, models.ErrInternal
	}

	banners := dao.ConvertTablesToModels(bannerRows)
	return banners, nil
}

func (r *BannerRepo) GetFilteredByFeature(ctx context.Context, featureID, limit, offset int) ([]models.Banner, error) {
	q := `SELECT DISTINCT b.id, (SELECT ARRAY_AGG(tag_id) FROM banner_definition d WHERE d.banner_id = b.id) as tag_ids,
	d.feature_id, b.content, b.is_active, b.created_at,	b.updated_at
	FROM banner b JOIN banner_definition d on b.id = d.banner_id
	WHERE d.feature_id = $1 AND b.is_active = true
	ORDER BY b.created_at DESC
	LIMIT $2 OFFSET $3;`

	bannerRows := []dao.BannerTable{}
	err := r.bannerStorage.Select(ctx, &bannerRows, q, featureID, limit, offset)

	if err != nil {
		if len(bannerRows) == 0 {
			return nil, models.ErrNoBanner
		}
		return nil, models.ErrInternal
	}

	banners := dao.ConvertTablesToModels(bannerRows)
	return banners, nil
}

func (r *BannerRepo) GetFilteredByTag(ctx context.Context, tagID, limit, offset int) ([]models.Banner, error) {
	q := `SELECT DISTINCT b.id, (SELECT ARRAY_AGG(tag_id) FROM banner_definition d WHERE d.banner_id = b.id) as tag_ids,
	d.feature_id, b.content, b.is_active, b.created_at,	b.updated_at
	FROM banner b join banner_definition d on b.id = d.banner_id
	where d.tag_id = $1 AND b.is_active = true
	ORDER BY b.created_at DESC
	LIMIT $2 OFFSET $3;`

	bannerRows := []dao.BannerTable{}
	err := r.bannerStorage.Select(ctx, &bannerRows, q, tagID, limit, offset)

	if err != nil {
		if len(bannerRows) == 0 {
			return nil, models.ErrNoBanner
		}
		return nil, models.ErrInternal
	}

	banners := dao.ConvertTablesToModels(bannerRows)
	return banners, nil
}

func (r *BannerRepo) Update(ctx context.Context, bannerID int, bannerData models.Banner) error {
	tx, err := r.bannerStorage.Begin(ctx, &sql.TxOptions{})
	if err != nil {
		log.Println(err.Error())
		return models.ErrInternal
	}

	q := `UPDATE banner SET content = $1, is_active = $2 WHERE id = $3 RETURNING id;`

	var bannerId int
	err = tx.Get(&bannerId, q, bannerData.Content, bannerData.IsActive, bannerID)
	if err != nil {
		log.Println(err.Error())
		tx.Rollback()
		if errors.Is(err, sql.ErrNoRows) {
			return models.ErrNoBanner
		}
		return models.ErrInternal
	}

	q = `DELETE FROM banner_definition WHERE banner_id = $1 RETURNING banner_id;`
	err = tx.Get(&bannerId, q, bannerID)
	if err != nil {
		log.Println(err.Error())
		tx.Rollback()
		return models.ErrInternal
	}

	for _, tagID := range bannerData.TagIDs {
		q = `INSERT INTO banner_definition (banner_id, feature_id, tag_id) VALUES ($1, $2, $3)RETURNING banner_id;`

		_, err := tx.Exec(q, bannerID, bannerData.FeatureID, tagID)
		if err != nil {
			log.Println(err.Error())
			tx.Rollback()
			return models.ErrValidation
		}
	}
	if err = tx.Commit(); err != nil {
		log.Println(err.Error())
		return models.ErrInternal
	}

	return nil
}

func (r *BannerRepo) Delete(ctx context.Context, bannerID int) error {
	q := `DELETE FROM banner * WHERE id = $1 RETURNING id;`

	var oldID int
	err := r.bannerStorage.Get(ctx, &oldID, q, bannerID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.ErrNoBanner
		}
		return models.ErrInternal
	}
	return nil
}
