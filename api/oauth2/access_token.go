package oauth2

import (
	"github.com/bububa/kwai-openapi/core"
	"github.com/bububa/kwai-openapi/model/oauth2"
)

func GetAccessToken(clt *core.Client, code string) (*oauth2.AccessTokenResponse, error) {
	req := &oauth2.AccessTokenRequest{
		AppID:  clt.AppID,
		Secret: clt.Secret,
		Code:   code,
	}
	var ret oauth2.AccessTokenResponse
	err := clt.Get(req, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
