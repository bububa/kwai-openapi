package core

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/bububa/kwai-openapi/core/internal/debug"
	"github.com/bububa/kwai-openapi/model"
)

type Client struct {
	AppID  string
	Secret string
	debug  bool
}

func NewClient(appID string, secret string) *Client {
	return &Client{
		AppID:  appID,
		Secret: secret,
	}
}

func (c *Client) SetDebug(debug bool) {
	c.debug = debug
}

func (c *Client) Get(req model.GetRequest, ret model.Response) error {
	endpoint := req.GetEndpoint()
	if endpoint == "" {
		endpoint = BASE_URL
	}
	reqUrl := fmt.Sprintf("%s/%s?%s", endpoint, req.GetScope(), req.GetParams().Encode())
	debug.PrintGetRequest(reqUrl, c.debug)
	httpResp, err := http.DefaultClient.Get(reqUrl)
	if err != nil {
		debug.PrintError(err, c.debug)
		return err
	}
	defer httpResp.Body.Close()
	err = debug.DecodeJSONHttpResponse(httpResp.Body, ret, c.debug)
	if err != nil {
		debug.PrintError(err, c.debug)
		return err
	}
	if ret.IsError() {
		return ret
	}
	return nil
}

func (c *Client) Post(req model.PostRequest, ret model.Response) error {
	return c.PostRaw(req, ret, "application/json")
}

func (c *Client) PostRaw(req model.PostRequest, ret model.Response, contentType string) error {
	endpoint := req.GetEndpoint()
	if endpoint == "" {
		endpoint = BASE_URL
	}
	reqUrl := fmt.Sprintf("%s/%s", endpoint, req.GetScope())
	if req.GetParams() != nil {
		reqUrl = fmt.Sprintf("%s?%s", reqUrl, req.GetParams().Encode())
	}
	debug.PrintGetRequest(reqUrl, c.debug)
	httpReq, err := http.NewRequest("POST", reqUrl, bytes.NewReader(req.GetData()))
	httpReq.Header.Add("Content-Type", contentType)
	if err != nil {
		debug.PrintError(err, c.debug)
	}
	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()
	err = debug.DecodeJSONHttpResponse(httpResp.Body, ret, c.debug)
	if err != nil {
		debug.PrintError(err, c.debug)
		return err
	}
	if ret.IsError() {
		return ret
	}
	return nil
}

func (c *Client) UploadData(req model.UploadDataRequest, ret model.Response) error {
	endpoint := req.GetEndpoint()
	if endpoint == "" {
		endpoint = BASE_URL
	}
	reqUrl := fmt.Sprintf("%s/%s", endpoint, req.GetScope())
	if req.GetParams() != nil {
		reqUrl = fmt.Sprintf("%s?%s", reqUrl, req.GetParams().Encode())
	}
	debug.PrintGetRequest(reqUrl, c.debug)
	httpReq, err := http.NewRequest("POST", reqUrl, bytes.NewReader(req.GetData()))
	httpReq.Header.Add("Content-Type", req.GetContentType())
	if err != nil {
		debug.PrintError(err, c.debug)
	}
	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()
	err = debug.DecodeJSONHttpResponse(httpResp.Body, ret, c.debug)
	if err != nil {
		debug.PrintError(err, c.debug)
		return err
	}
	if ret.IsError() {
		return ret
	}
	return nil
}

func (c *Client) PostForm(req model.FormRequest, ret model.Response) error {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for k, v := range req.GetData() {
		writer.WriteField(k, v)
	}
	writer.Close()
	endpoint := req.GetEndpoint()
	if endpoint == "" {
		endpoint = BASE_URL
	}
	reqUrl := fmt.Sprintf("%s/%s", endpoint, req.GetScope())
	if req.GetParams() != nil {
		reqUrl = fmt.Sprintf("%s?%s", reqUrl, req.GetParams().Encode())
	}
	debug.PrintGetRequest(reqUrl, c.debug)
	httpReq, err := http.NewRequest("POST", reqUrl, body)
	httpReq.Header.Add("Content-Type", "multipart/form-data")
	if err != nil {
		debug.PrintError(err, c.debug)
	}
	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()
	err = debug.DecodeJSONHttpResponse(httpResp.Body, ret, c.debug)
	if err != nil {
		debug.PrintError(err, c.debug)
		return err
	}
	if ret.IsError() {
		return ret
	}
	return nil
}

func (c *Client) UploadFile(req model.UploadRequest, ret model.Response) error {
	file := req.GetFile()
	fileField := req.GetFileFieldName()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(fileField, filepath.Base(file.Name()))
	if err != nil {
		debug.PrintError(err, c.debug)
		return err
	}
	io.Copy(part, file)
	for k, v := range req.GetData() {
		writer.WriteField(k, v)
	}
	writer.Close()
	endpoint := req.GetEndpoint()
	if endpoint == "" {
		endpoint = BASE_URL
	}
	reqUrl := fmt.Sprintf("%s/%s", endpoint, req.GetScope())
	if req.GetParams() != nil {
		reqUrl = fmt.Sprintf("%s?%s", reqUrl, req.GetParams().Encode())
	}
	debug.PrintGetRequest(reqUrl, c.debug)
	httpReq, err := http.NewRequest("POST", reqUrl, body)
	httpReq.Header.Add("Content-Type", writer.FormDataContentType())
	if err != nil {
		debug.PrintError(err, c.debug)
	}
	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()
	err = debug.DecodeJSONHttpResponse(httpResp.Body, ret, c.debug)
	if err != nil {
		debug.PrintError(err, c.debug)
		return err
	}
	if ret.IsError() {
		return ret
	}
	return nil
}
