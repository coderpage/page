package controller

import (
	"encoding/json"
	"page/constant/rsp"
)

// Http 请求返回数据
type Response map[string]interface{}

// 创建一个 Response
func NewResponse() (resp Response) {
	return make(Response)
}

// Http 返回数据的 json 格式字符串
func (resp Response) JsonString() string {
	respBytes, _ := json.Marshal(resp)
	return string(respBytes)
}

// 设置 response 的 status 值
func (resp Response) SetStatus(status interface{}) {
	resp[rsp.BodyStatus] = status
}

// 设置 response 的 message 值
func (resp Response) SetMessage(message interface{}) {
	resp[rsp.BodyMsg] = message
}

// 设置 response 的 data 值
func (resp Response) SetData(name string, data interface{}) {
	resp[name] = data
}
