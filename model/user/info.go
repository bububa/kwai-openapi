package user

import (
	"net/url"

	"github.com/bububa/kwai-openapi/model"
)

type GetUserInfoRequest struct {
	model.BaseRequest
	AppID       string
	AccessToken string
}

func (r GetUserInfoRequest) GetScope() string {
	return "openapi/user_info"
}

func (r GetUserInfoRequest) GetParams() url.Values {
	ret := url.Values{}
	ret.Set("app_id", r.AppID)
	ret.Set("access_token", r.AccessToken)
	return ret
}

type GetUserInfoResponse struct {
	model.BaseResponse
	UserInfo *UserInfo `json:"user_info,omitempty"`
}

type UserInfo struct {
	Name    string `json:"name,omitempty"`
	Sex     string `json:"sex,omitempty"`
	Head    string `json:"head,omitempty"`
	BigHead string `json:"bigHead,omitempty"`
	City    string `json:"city,omitempty"`
	Fan     int64  `json:"fan,omitempty"`
	Follow  int64  `json:"follow,omitempty"`
}
