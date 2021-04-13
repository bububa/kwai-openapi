package live

import (
	"net/url"

	"github.com/bububa/kwai-openapi/model"
)

type StopPushRequest struct {
	model.BaseRequest
	AppID          string
	AccessToken    string
	LiveStreamName string
}

func (r StopPushRequest) GetScope() string {
	return "openapi/kwaiUser/live/stopPush"
}

func (r StopPushRequest) GetParams() url.Values {
	ret := url.Values{}
	ret.Set("app_id", r.AppID)
	ret.Set("access_token", r.AccessToken)
	ret.Set("liveSteamName", r.LiveStreamName)
	return ret
}

func (r StopPushRequest) GetData() []byte {
	return nil
}

type StopPushResponse struct {
	model.BaseResponse
	HostName string `json:"hostName,omitempty"`
}
