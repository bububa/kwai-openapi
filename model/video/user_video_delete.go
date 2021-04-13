package video

import (
	"net/url"

	"github.com/bububa/kwai-openapi/model"
)

type UserVideoDeleteRequest struct {
	model.BaseRequest
	AppID       string
	AccessToken string
	PhotoID     string
}

func (r UserVideoDeleteRequest) GetScope() string {
	return "openapi/photo/delete"
}

func (r UserVideoDeleteRequest) GetParams() url.Values {
	ret := url.Values{}
	ret.Set("app_id", r.AppID)
	ret.Set("access_token", r.AccessToken)
	ret.Set("photo_id", r.PhotoID)
	return ret
}

func (r UserVideoDeleteRequest) GetData() []byte {
	return nil
}
