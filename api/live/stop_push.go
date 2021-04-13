package live

import (
	"github.com/bububa/kwai-openapi/core"
	"github.com/bububa/kwai-openapi/model/live"
)

func StopPush(clt *core.Client, accessToken string, liveStreamName string) (string, error) {
	req := &live.StopPushRequest{
		AppID:          clt.AppID,
		AccessToken:    accessToken,
		LiveStreamName: liveStreamName,
	}
	var ret live.StopPushResponse
	err := clt.PostRaw(req, &ret, "application/x-www-form-urlencoded")
	if err != nil {
		return "", err
	}
	return ret.HostName, nil
}
