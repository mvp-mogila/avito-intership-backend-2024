package utils

import (
	"encoding/json"
	"io"
	"net/url"
	"strconv"

	"github.com/mvp-mogila/avito-intership-backend-2024/internal/delivery/dto"
)

func GetFeatureIDFromQuery(args url.Values, required bool) (int, error) {
	strID := args.Get("feature_id")
	if strID == "" {
		if required {
			return 0, NewValidationError("feature_id is required in query params")
		}
		return 0, nil
	}
	intID, err := strconv.Atoi(strID)
	if err != nil {
		return 0, NewValidationError("feature_id is invalid")
	}
	return int(intID), nil
}

func GetTagIDFromQuery(args url.Values, required bool) (int, error) {
	strID := args.Get("tag_id")
	if strID == "" {
		if required {
			return 0, NewValidationError("tag_id is required in query params")
		}
		return 0, nil
	}
	intID, err := strconv.Atoi(strID)
	if err != nil {
		return 0, NewValidationError("tag_id is invalid")
	}
	return int(intID), nil
}

func GetLastRevisionFromQuery(args url.Values) bool {
	param := args.Get("use_last_revision")
	if param == "" {
		return false
	}
	useLast, _ := strconv.ParseBool(param)
	return useLast
}

func GetLimitFromQuery(args url.Values) int {
	param := args.Get("limit")
	limit, _ := strconv.Atoi(param)
	return limit
}

func GetOffsetFromQuery(args url.Values) int {
	param := args.Get("offset")
	offset, _ := strconv.Atoi(param)
	return offset
}

func GetBannerData(body io.ReadCloser, data *dto.BannerDetails) error {
	return json.NewDecoder(body).Decode(data)
}
