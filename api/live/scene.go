package live

import (
	"github.com/bububa/kwai-openapi/core"
	"github.com/bububa/kwai-openapi/model/live"
)

func GetLiveScene(clt *core.Client, accessToken string) (*live.LiveSceneGames, error) {
	req := &live.GetLiveSceneRequest{
		AppID:       clt.AppID,
		AccessToken: accessToken,
	}
	var ret live.GetLiveSceneResponse
	err := clt.Get(req, &ret)
	if err != nil {
		return nil, err
	}
	if ret.Content != nil {
		return ret.Content.Games, nil
	}
	return nil, nil
}
