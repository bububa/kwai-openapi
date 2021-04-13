package oauth2

import (
	"net/url"

	"github.com/bububa/kwai-openapi/model"
)

type RefreshTokenRequest struct {
	model.BaseRequest
	AppID        string
	Secret       string
	RefreshToken string
}

func (r RefreshTokenRequest) GetScope() string {
	return "oauth2/refresh_token"
}

func (r RefreshTokenRequest) GetParams() url.Values {
	ret := url.Values{}
	ret.Set("app_id", r.AppID)
	ret.Set("app_secret", r.Secret)
	ret.Set("refresh_token", r.RefreshToken)
	ret.Set("grant_type", "refresh_token")
	return ret
}

type RefreshTokenResponse struct {
	model.BaseResponse
	AccessToken           string   `json:"access_token,omitempty"`
	ExpiresIn             int64    `json:"expires_in,omitempty"`
	RefreshToken          string   `json:"refresh_token,omitempty"`
	RefreshTokenExpiresIn int64    `json:"refresh_token_expires_in,omitempty"`
	Scopes                []string `json:"scopes,omitempty"`
}
