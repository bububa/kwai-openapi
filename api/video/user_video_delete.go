package video

import (
	"github.com/bububa/kwai-openapi/core"
	"github.com/bububa/kwai-openapi/model"
	"github.com/bububa/kwai-openapi/model/video"
)

func UserVideoDelete(clt *core.Client, accessToken string, photoID string) error {
	req := &video.UserVideoDeleteRequest{
		AppID:       clt.AppID,
		AccessToken: accessToken,
		PhotoID:     photoID,
	}
	var ret model.BaseResponse
	return clt.Post(req, &ret)
}
