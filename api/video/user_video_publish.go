package video

import (
	"github.com/bububa/kwai-openapi/core"
	"github.com/bububa/kwai-openapi/model"
	"github.com/bububa/kwai-openapi/model/video"
)

func UserVideoStartUpload(clt *core.Client, accessToken string) (*video.UserVideoStartUploadResponse, error) {
	req := &video.UserVideoStartUploadRequest{
		AppID:       clt.AppID,
		AccessToken: accessToken,
	}
	var ret video.UserVideoStartUploadResponse
	err := clt.Post(req, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

func UserVideoUpload(clt *core.Client, req *video.UserVideoUploadRequest) error {
	var ret model.BaseResponse
	return clt.UploadData(req, &ret)
}

func UserVideoUploadFile(clt *core.Client, req *video.UserVideoUploadFileRequest) error {
	var ret model.BaseResponse
	return clt.UploadFile(req, &ret)
}

func UserVideoUploadFragment(clt *core.Client, req *video.UserVideoUploadFragmentRequest) error {
	var ret model.BaseResponse
	return clt.UploadData(req, &ret)
}

func UserVideoUploadResume(clt *core.Client, req *video.UserVideoUploadResumeRequest) (*video.UserVideoUploadResumeResponse, error) {
	var ret video.UserVideoUploadResumeResponse
	err := clt.Get(req, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

func UserVideoUploadComplete(clt *core.Client, req *video.UserVideoUploadCompleteRequest) error {
	var ret model.BaseResponse
	return clt.Post(req, &ret)
}

func UserVideoPublish(clt *core.Client, req *video.UserVideoPublishRequest) (*video.UserVideo, error) {
	req.AppID = clt.AppID
	var ret video.UserVideoPublishResponse
	err := clt.PostForm(req, &ret)
	if err != nil {
		return nil, err
	}
	return ret.VideoInfo, nil
}
