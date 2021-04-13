package oauth2

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/bububa/kwai-openapi/core"
	"github.com/bububa/kwai-openapi/model/oauth2"
)

func AuthorizeUrl(clt *core.Client, req *oauth2.AuthorizeRequest) string {
	values := &url.Values{}
	values.Set("app_id", clt.AppID)
	values.Set("scope", req.Scope)
	values.Set("response_type", "code")
	values.Set("redirect_uri", req.RedirectUri)
	if req.State != "" {
		values.Set("state", req.State)
	}
	return fmt.Sprintf("%s/oauth2/authorize?%s", core.BASE_URL, values.Encode())
}

func ConnectUrl(clt *core.Client, req *oauth2.AuthorizeRequest) string {
	values := &url.Values{}
	values.Set("app_id", clt.AppID)
	values.Set("scope", req.Scope)
	values.Set("response_type", "code")
	values.Set("redirect_uri", req.RedirectUri)
	if req.State != "" {
		values.Set("state", req.State)
	}
	return fmt.Sprintf("%s/oauth2/connect?%s", core.BASE_URL, values.Encode())
}

func QrcodeUrl(clt *core.Client, req *oauth2.QrcodeUrlRequest) string {
	values := &url.Values{}
	values.Set("app_id", clt.AppID)
	values.Set("scope", req.Scope)
	values.Set("response_type", "code")
	values.Set("redirect_uri", req.RedirectUri)
	if req.State != "" {
		values.Set("state", req.State)
	}
	if req.QrcodeSize > 0 {
		values.Set("qrcode_size", strconv.Itoa(req.QrcodeSize))
	}
	if req.Type != "" {
		values.Set("type", req.Type)
	}
	if req.Debug {
		values.Set("debug", "true")
	}
	return fmt.Sprintf("%s/oauth2/qr_code?%s", core.BASE_URL, values.Encode())
}
