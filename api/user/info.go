package user

import (
	"github.com/bububa/kwai-openapi/core"
	"github.com/bububa/kwai-openapi/model/user"
)

func GetUserInfo(clt *core.Client, accessToken string) (*user.UserInfo, error) {
	req := &user.GetUserInfoRequest{
		AppID:       clt.AppID,
		AccessToken: accessToken,
	}
	var ret user.GetUserInfoResponse
	err := clt.Get(req, &ret)
	if err != nil {
		return nil, err
	}
	return ret.UserInfo, nil
}
