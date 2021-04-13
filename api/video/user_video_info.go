package video

import (
	"github.com/bububa/kwai-openapi/core"
	"github.com/bububa/kwai-openapi/model/video"
)

func GetUserVideoInfo(clt *core.Client, accessToken string, cursor string, count int64) ([]video.UserVideo, error) {
	req := &video.GetUserVideoInfoRequest{
		AppID:       clt.AppID,
		AccessToken: accessToken,
		Cursor:      cursor,
		Count:       count,
	}
	var ret video.GetUserVideoInfoResponse
	err := clt.Get(req, &ret)
	if err != nil {
		return nil, err
	}
	return ret.VideoList, nil
}
