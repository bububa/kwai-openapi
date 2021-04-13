package live

import (
	"net/url"

	"github.com/bububa/kwai-openapi/model"
)

type GetLiveSceneRequest struct {
	model.BaseRequest
	AppID       string
	AccessToken string
}

func (r GetLiveSceneRequest) GetScope() string {
	return "openapi/kwaiUser/live/scene"
}

func (r GetLiveSceneRequest) GetParams() url.Values {
	ret := url.Values{}
	ret.Set("app_id", r.AppID)
	ret.Set("access_token", r.AccessToken)
	return ret
}

type GetLiveSceneResponse struct {
	model.BaseResponse
	HostName string     `json:"hostName,omitempty"`
	Content  *LiveScene `json:"content,omitempty"`
}

type LiveScene struct {
	Games *LiveSceneGames `json:"games,omitempty"`
}

type LiveSceneGames struct {
	PC     []LiveSceneGame `json:"pc,omitempty"`
	PCShow []LiveSceneGame `json:"pcShow,omitempty"`
	Event  []LiveSceneGame `json:"event,omitempty"`
	PCHand []LiveSceneGame `json:"pcHand,omitempty"`
}

type LiveSceneGame struct {
	ID            uint64   `json:"id,omitempty"`
	Name          string   `json:"name,omitempty"`
	Icon          string   `json:"icon,omitempty"`
	Category      string   `json:"category,omitempty"`
	ScreenType    int      `json:"screenType,omitempty"`
	LaunchURLs    []string `json:"launchUrls,omitempty"`
	Offline       bool     `json:"offline,omitempty"`
	BackgroundURL string   `json:"backgroundUrl,omitempty"`
	Game          bool     `json:"game,omitempty"`
}
