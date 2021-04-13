package live

import (
	"github.com/bububa/kwai-openapi/core"
	"github.com/bububa/kwai-openapi/model/live"
)

func GetPushUrls(clt *core.Client, req *live.GetPushUrlsRequest) (*live.PushUrl, error) {
	var ret live.GetPushUrlsResponse
	err := clt.UploadFile(req, &ret)
	if err != nil {
		return nil, err
	}
	return ret.Content, nil
}
