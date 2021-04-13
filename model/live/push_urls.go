package live

import (
	"net/url"
	"os"
	"strconv"

	"github.com/bububa/kwai-openapi/model"
)

type GetPushUrlsRequest struct {
	model.BaseRequest
	AppID         string
	AccessToken   string
	File          *os.File
	Caption       string
	Panoramic     int
	ShopLive      bool
	DeviceName    string
	LiveSceneType int
}

func (r GetPushUrlsRequest) GetScope() string {
	return "openapi/kwaiUser/live/pushUrls"
}

func (r GetPushUrlsRequest) GetParams() url.Values {
	ret := url.Values{}
	ret.Set("app_id", r.AppID)
	ret.Set("access_token", r.AccessToken)
	return ret
}

func (r GetPushUrlsRequest) GetData() map[string]string {
	ret := make(map[string]string, 5)
	if r.Caption != "" {
		ret["caption"] = r.Caption
	}
	if r.Panoramic > 0 {
		ret["panoramic"] = strconv.Itoa(r.Panoramic)
	}
	if r.ShopLive {
		ret["shopLive"] = "true"
	}
	if r.DeviceName != "" {
		ret["deviceName"] = r.DeviceName
	}
	if r.LiveSceneType > 0 {
		ret["liveSceneType"] = strconv.Itoa(r.LiveSceneType)
	}
	return ret
}

func (r GetPushUrlsRequest) GetFile() *os.File {
	return r.File
}

type GetPushUrlsResponse struct {
	model.BaseResponse
	HostName string   `json:"hostName,omitempty"`
	Content  *PushUrl `json:"content,omitempty"`
}

type PushUrl struct {
	PushURL        string `json:"pushUrl,omitempty"`
	LiveStreamName string `json:"liveStreamName,omitempty"`
}
