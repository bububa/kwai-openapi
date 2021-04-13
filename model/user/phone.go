package user

import (
	"net/url"

	"github.com/bububa/kwai-openapi/model"
)

type GetUserPhoneRequest struct {
	model.BaseRequest
	AppID       string
	AccessToken string
}

func (r GetUserPhoneRequest) GetScope() string {
	return "openapi/user_phone"
}

func (r GetUserPhoneRequest) GetParams() url.Values {
	ret := url.Values{}
	ret.Set("app_id", r.AppID)
	ret.Set("access_token", r.AccessToken)
	return ret
}

type GetUserPhoneResponse struct {
	model.BaseResponse
	EncryptedPhone string `json:"encrypted_phone,omitempty"`
}
