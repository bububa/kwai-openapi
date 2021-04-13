package live

import (
	"fmt"
	"net/url"

	"github.com/bububa/kwai-openapi/model"
)

type StatusRequest struct {
	model.BaseRequest
	AppID          string
	AccessToken    string
	LiveStreamName string
}

func (r StatusRequest) GetScope() string {
	return fmt.Sprintf("openapi/kwaiUser/live/status/%s", r.LiveStreamName)
}

func (r StatusRequest) GetParams() url.Values {
	ret := url.Values{}
	ret.Set("app_id", r.AppID)
	ret.Set("access_token", r.AccessToken)
	return ret
}

type StatusResponse struct {
	model.BaseResponse
	HostName string      `json:"hostName,omitempty"`
	Content  *LiveStream `json:"content,omitempty"`
}

type LiveStreamStatus = string

const (
	LIVING         LiveStreamStatus = "LIVING"
	LIVE_NOT_START LiveStreamStatus = "LIVE_NOT_START"
	LIVE_END       LiveStreamStatus = "LIVE_END"
)

type LiveStream struct {
	Name      string           `json:"liveStreamName,omitempty"`
	Status    LiveStreamStatus `json:"status,omitempty"`
	StatusMsg string           `json:"statusMsg,omitempty"`
}
