package utils

import (
	"bytes"
	"encoding/json"
)

// UnmarshalJson 的原生Unmarshal会将int64转换成float64解析，造成精度丢失
func UnmarshalJson(data []byte, v interface{}) error {
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()
	return decoder.Decode(&v)
}
