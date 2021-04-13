package model

type Response interface {
	IsError() bool
	Error() string
}

type BaseResponse struct {
	Result int    `json:"result,omitempty"`
	ErrMsg string `json:"error_msg,omitempty"`
}

func (r BaseResponse) IsError() bool {
	return r.Result != 1
}

func (r BaseResponse) Error() string {
	return r.ErrMsg
}
