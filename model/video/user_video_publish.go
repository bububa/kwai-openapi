package video

import (
	"net/url"
	"os"
	"strconv"

	"github.com/bububa/kwai-openapi/model"
)

type UserVideoStartUploadRequest struct {
	model.BaseRequest
	AppID       string
	AccessToken string
}

func (r UserVideoStartUploadRequest) GetScope() string {
	return "openapi/photo/start_upload"
}

func (r UserVideoStartUploadRequest) GetParams() url.Values {
	ret := url.Values{}
	ret.Set("app_id", r.AppID)
	ret.Set("access_token", r.AccessToken)
	return ret
}

func (r UserVideoStartUploadRequest) GetData() []byte {
	return nil
}

type UserVideoStartUploadResponse struct {
	model.BaseResponse
	UploadToken string `json:"upload_token,omitempty"`
	Endpoint    string `json:"endpoint,omitempty"`
}

type UserVideoUploadRequest struct {
	model.BaseRequest
	UploadToken string
	ContentType string
	Data        []byte
}

func (r UserVideoUploadRequest) GetScope() string {
	return "api/upload"
}

func (r UserVideoUploadRequest) GetParams() url.Values {
	ret := url.Values{}
	ret.Set("upload_token", r.UploadToken)
	return ret
}

func (r UserVideoUploadRequest) GetContentType() string {
	return r.ContentType
}

func (r UserVideoUploadRequest) GetData() []byte {
	return r.Data
}

type UserVideoUploadFileRequest struct {
	model.BaseRequest
	UploadToken string
	File        *os.File
}

func (r UserVideoUploadFileRequest) GetScope() string {
	return "api/upload/multipart"
}

func (r UserVideoUploadFileRequest) GetParams() url.Values {
	ret := url.Values{}
	ret.Set("upload_token", r.UploadToken)
	return ret
}

func (r UserVideoUploadFileRequest) GetData() map[string]string {
	return nil
}

func (r UserVideoUploadFileRequest) GetFile() *os.File {
	return r.File
}

type UserVideoUploadFragmentRequest struct {
	model.BaseRequest
	UploadToken string
	FragmentID  int
	ContentType string
	Data        []byte
}

func (r UserVideoUploadFragmentRequest) GetScope() string {
	return "api/upload/fragment"
}

func (r UserVideoUploadFragmentRequest) GetParams() url.Values {
	ret := url.Values{}
	ret.Set("upload_token", r.UploadToken)
	ret.Set("fragment_id", strconv.Itoa(r.FragmentID))
	return ret
}

func (r UserVideoUploadFragmentRequest) GetContentType() string {
	return r.ContentType
}

func (r UserVideoUploadFragmentRequest) GetData() []byte {
	return r.Data
}

type UserVideoUploadResumeRequest struct {
	model.BaseRequest
	UploadToken string
}

func (r UserVideoUploadResumeRequest) GetScope() string {
	return "api/upload/resume"
}

func (r UserVideoUploadResumeRequest) GetParams() url.Values {
	ret := url.Values{}
	ret.Set("upload_token", r.UploadToken)
	return ret
}

type UserVideoUploadResumeResponse struct {
	model.BaseResponse
	Existed            bool                      `json:"existed,omitempty"`
	FragmentIndex      int                       `json:"fragment_index,omitempty"`
	FragmentList       []int                     `json:"fragment_list,omitempty"`
	Endpoint           []UserVideoUploadEndpoint `json:"endpoint,omitempty"`
	FragmentIndexBytes int64                     `json:"fragment_index_bytes,omitempty"`
	TokenID            string                    `json:"token_id,omitempty"`
}

type UserVideoUploadEndpoint struct {
	Protocol string `json:"protocol,omitempty"`
	Host     string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
}

type UserVideoUploadCompleteRequest struct {
	model.BaseRequest
	UploadToken   string
	FragmentCount int
}

func (r UserVideoUploadCompleteRequest) GetScope() string {
	return "api/upload/complete"
}

func (r UserVideoUploadCompleteRequest) GetParams() url.Values {
	ret := url.Values{}
	ret.Set("upload_token", r.UploadToken)
	ret.Set("fragment_count", strconv.Itoa(r.FragmentCount))
	return ret
}

func (r UserVideoUploadCompleteRequest) GetData() []byte {
	return nil
}

type SteroType = string

const (
	NOT_SPHERICAL_VIDEO SteroType = "NOT_SPHERICAL_VIDEO"
	SPHERICAL_VIDEO_360 SteroType = "SPHERICAL_VIDEO_360"
	SPHERICAL_VIDEO_180 SteroType = "SPHERICAL_VIDEO_180"
)

type UserVideoPublishRequest struct {
	model.BaseRequest
	AppID       string
	AccessToken string
	UploadToken string
	Cover       string
	Caption     string
	SteroType   SteroType
}

func (r UserVideoPublishRequest) GetScope() string {
	return "openapi/photo/publish"
}

func (r UserVideoPublishRequest) GetParams() url.Values {
	ret := url.Values{}
	ret.Set("app_id", r.AppID)
	ret.Set("access_token", r.AccessToken)
	ret.Set("upload_token", r.UploadToken)
	return ret
}

func (r UserVideoPublishRequest) GetData() map[string]string {
	ret := make(map[string]string, 3)
	ret["cover"] = r.Cover
	ret["caption"] = r.Caption
	ret["stero_type"] = r.SteroType
	return ret
}

type UserVideoPublishResponse struct {
	model.BaseResponse
	VideoInfo *UserVideo `json:"video_info,omitempty"`
}
