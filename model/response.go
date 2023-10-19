package model

type Response struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
}

const (
	PARSE_BODY_ERROR = "parse body err"
	ADD_GRID_ERROR   = "parse body err"
)

func (res *Response) AddContent(key string, content interface{}) *map[string]interface{} {
	ret := make(map[string]interface{})
	ret["code"] = res.Code
	ret["msg"] = res.Msg
	ret[key] = content
	return &ret
}
