package oauth2

import (
	"net/url"

	"github.com/bububa/kwai-openapi/model"
)

type AccessTokenRequest struct {
	model.BaseRequest
	AppID  string
	Secret string
	Code   string
}

func (r AccessTokenRequest) GetScope() string {
	return "oauth2/access_token"
}

func (r AccessTokenRequest) GetParams() url.Values {
	ret := url.Values{}
	ret.Set("app_id", r.AppID)
	ret.Set("app_secret", r.Secret)
	ret.Set("code", r.Code)
	ret.Set("grant_type", "authorization_code")
	return ret
}

type AccessTokenResponse struct {
	model.BaseResponse
	AccessToken           string   `json:"access_token,omitempty"`
	ExpiresIn             int64    `json:"expires_in,omitempty"`
	RefreshToken          string   `json:"refresh_token,omitempty"`
	RefreshTokenExpiresIn int64    `json:"refresh_token_expires_in,omitempty"`
	OpenID                string   `json:"open_id,omitempty"`
	Scopes                []string `json:"scopes,omitempty"`
}
