package live

import (
	"github.com/bububa/kwai-openapi/core"
	"github.com/bububa/kwai-openapi/model/live"
)

func GetStatus(clt *core.Client, accessToken string, liveStreamName string) (*live.LiveStream, error) {
	req := &live.StatusRequest{
		AppID:          clt.AppID,
		AccessToken:    accessToken,
		LiveStreamName: liveStreamName,
	}
	var ret live.StatusResponse
	err := clt.Get(req, &ret)
	if err != nil {
		return nil, err
	}
	return ret.Content, nil
}
