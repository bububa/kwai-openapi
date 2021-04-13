package video

import (
	"net/url"
	"strconv"

	"github.com/bububa/kwai-openapi/model"
)

type GetUserVideoInfoRequest struct {
	model.BaseRequest
	AppID       string
	AccessToken string
	Cursor      string
	Count       int64
}

func (r GetUserVideoInfoRequest) GetScope() string {
	return "openapi/photo/list"
}

func (r GetUserVideoInfoRequest) GetParams() url.Values {
	ret := url.Values{}
	ret.Set("app_id", r.AppID)
	ret.Set("access_token", r.AccessToken)
	if r.Cursor != "" {
		ret.Set("cursor", r.Cursor)
	}
	if r.Count > 0 {
		ret.Set("count", strconv.FormatInt(r.Count, 10))
	}
	return ret
}

type GetUserVideoInfoResponse struct {
	model.BaseResponse
	VideoList []UserVideo `json:"video_list,omitempty"`
}

type UserVideo struct {
	PhotoID      string `json:"photo_id,omitempty"`
	Caption      string `json:"caption,omitempty"`
	Cover        string `json:"cover,omitempty"`
	PlayURL      string `json:"play_url,omitempty"`
	CreateTime   int64  `json:"create_time,omitempty"`
	LikeCount    int64  `json:"like_count,omitempty"`
	CommentCount int64  `json:"comment_count,omitempty"`
	ViewCount    int64  `json:"view_count,omitempty"`
	Pending      bool   `json:"pending,omitempty"`
}
