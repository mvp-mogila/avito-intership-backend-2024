package delivery

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mvp-mogila/avito-intership-backend-2024/internal/delivery/dto"
	"github.com/mvp-mogila/avito-intership-backend-2024/internal/models"
	"github.com/mvp-mogila/avito-intership-backend-2024/internal/pkg/middleware"
	"github.com/mvp-mogila/avito-intership-backend-2024/internal/pkg/utils"
)

type Banners interface {
	Create(ctx context.Context, bannerData models.Banner) (int, error)
	Get(ctx context.Context, params models.BannerDefinition) (models.Banner, error)
	GetFiltered(ctx context.Context, params models.BannerDefinition, limit, offset int) ([]models.Banner, error)
	Update(ctx context.Context, bannerID int, bannerData models.Banner) error
	Delete(ctx context.Context, bannerID int) error
}

type BannerHandler struct {
	router        *mux.Router
	bannerUsecase Banners
}

func NewBannerHandler(u Banners) *BannerHandler {
	return &BannerHandler{
		bannerUsecase: u,
	}
}

func (h *BannerHandler) SetupRouting(r *mux.Router) {
	h.router = r.NewRoute().Subrouter()
	h.router.HandleFunc("/user_banner", h.GetBanner).Methods(http.MethodGet)
	h.router.HandleFunc("/banner", middleware.Authorization(h.GetBanners)).Methods(http.MethodGet)
	h.router.HandleFunc("/banner", middleware.Authorization(h.CreateBanner)).Methods(http.MethodPost)
	h.router.HandleFunc("/banner/{id:[1-9]+[0-9]*}", middleware.Authorization(h.UpdateBanner)).Methods(http.MethodPatch)
	h.router.HandleFunc("/banner/{id:[1-9]+[0-9]*}", middleware.Authorization(h.DeleteBanner)).Methods(http.MethodDelete)
}

func (h *BannerHandler) GetBanner(w http.ResponseWriter, r *http.Request) {
	queryArgs := r.URL.Query()
	featureID, err := utils.GetFeatureIDFromQuery(queryArgs, true)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	tagID, err := utils.GetTagIDFromQuery(queryArgs, true)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	useLastRevision := utils.GetLastRevisionFromQuery(queryArgs)

	params := models.BannerDefinition{
		FeatureID:    featureID,
		TagID:        tagID,
		LastRevision: useLastRevision,
	}
	banner, err := h.bannerUsecase.Get(r.Context(), params)

	switch {
	case errors.Is(err, models.ErrNoBanner):
		utils.SendErrorResponse(w, http.StatusNotFound, "")
		return
	case errors.Is(err, models.ErrInternal):
		utils.SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	case err != nil:
		utils.SendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.SendSuccessResponse(w, http.StatusOK, banner.Content)
}

func (h *BannerHandler) GetBanners(w http.ResponseWriter, r *http.Request) {
	queryArgs := r.URL.Query()
	featureID, _ := utils.GetFeatureIDFromQuery(queryArgs, false)
	tagID, _ := utils.GetTagIDFromQuery(queryArgs, false)
	limit := utils.GetLimitFromQuery(queryArgs)
	offset := utils.GetOffsetFromQuery(queryArgs)

	bannerDef := models.BannerDefinition{
		FeatureID: featureID,
		TagID:     tagID,
	}
	banners, err := h.bannerUsecase.GetFiltered(r.Context(), bannerDef, limit, offset)

	if errors.Is(err, models.ErrInternal) {
		utils.SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	bannersPayload := dto.ConvertModelsToDto(banners)
	utils.SendSuccessResponse(w, http.StatusOK, bannersPayload)
}

func (h *BannerHandler) CreateBanner(w http.ResponseWriter, r *http.Request) {
	var requestData dto.BannerDetails
	err := utils.GetBannerData(r.Body, &requestData)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	banner := dto.ConvertBannerDetailsToModel(requestData)
	ID, err := h.bannerUsecase.Create(r.Context(), banner)

	if err != nil {
		if errors.Is(err, models.ErrInternal) {
			utils.SendErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		utils.SendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	responseData := dto.CreateBannerPayload{
		BannerID: ID,
	}
	utils.SendSuccessResponse(w, http.StatusCreated, responseData)
}

func (h *BannerHandler) UpdateBanner(w http.ResponseWriter, r *http.Request) {
	var requestData dto.BannerDetails
	err := utils.GetBannerData(r.Body, &requestData)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	bannerID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	banner := dto.ConvertBannerDetailsToModel(requestData)
	err = h.bannerUsecase.Update(r.Context(), bannerID, banner)

	switch {
	case errors.Is(err, models.ErrNoBanner):
		utils.SendErrorResponse(w, http.StatusNotFound, "")
		return
	case errors.Is(err, models.ErrInternal):
		utils.SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	case err != nil:
		utils.SendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *BannerHandler) DeleteBanner(w http.ResponseWriter, r *http.Request) {
	bannerID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	err = h.bannerUsecase.Delete(r.Context(), bannerID)

	switch {
	case errors.Is(err, models.ErrNoBanner):
		utils.SendErrorResponse(w, http.StatusNotFound, "")
		return
	case errors.Is(err, models.ErrInternal):
		utils.SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
