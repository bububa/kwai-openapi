package model

import (
	"fmt"
	"net/url"
	"os"
)

type GetRequest interface {
	GetScope() string
	GetParams() url.Values
	GetEndpoint() string
}

type PostRequest interface {
	GetScope() string
	GetParams() url.Values
	GetEndpoint() string
	GetData() []byte
}

type UploadDataRequest interface {
	GetScope() string
	GetParams() url.Values
	GetData() []byte
	GetContentType() string
	GetEndpoint() string
}

type FormRequest interface {
	GetScope() string
	GetParams() url.Values
	GetData() map[string]string
	GetEndpoint() string
}

type UploadRequest interface {
	GetScope() string
	GetParams() url.Values
	GetData() map[string]string
	GetFile() *os.File
	GetFileFieldName() string
	GetEndpoint() string
}

type BaseRequest struct {
	Endpoint string
}

func (r BaseRequest) GetEndpoint() string {
	if r.Endpoint != "" {
		return fmt.Sprintf("https://%s", r.Endpoint)
	}
	return ""
}

func (r BaseRequest) GetFileFieldName() string {
	return "file"
}
