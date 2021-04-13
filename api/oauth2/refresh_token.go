package oauth2

import (
	"github.com/bububa/kwai-openapi/core"
	"github.com/bububa/kwai-openapi/model/oauth2"
)

func RefreshToken(clt *core.Client, refreshToken string) (*oauth2.RefreshTokenResponse, error) {
	req := &oauth2.RefreshTokenRequest{
		AppID:        clt.AppID,
		Secret:       clt.Secret,
		RefreshToken: refreshToken,
	}
	var ret oauth2.RefreshTokenResponse
	err := clt.Get(req, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
