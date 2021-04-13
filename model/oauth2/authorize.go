package oauth2

type QrcodeType = string

const (
	IMG_QrcodeType  QrcodeType = "img"
	PATH_QrcodeType QrcodeType = "path"
)

type AuthorizeRequest struct {
	RedirectUri string
	Scope       string
	State       string
}

type QrcodeUrlRequest struct {
	AuthorizeRequest
	QrcodeSize int        // 二维码大小,最小不小于64,最大不大于1000,默认450。
	Type       QrcodeType // 返回值类型,可选img和path,选择img会返回二维码图片,选择path会返回二维码内容的地址,开发者可以自己生成二维码 。默认 img 。
	Debug      bool       // 时候开启调试模式,开启后,在用户授权失败时,会在快手APP内显示调试信息,可以根据调试信息来调试.默认false
}
