package util

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// ParseResponse 请求结果转化为Map
func ParseResponse(response *http.Response) (map[string]any, error) {
	var result map[string]any
	body, err := ioutil.ReadAll(response.Body)
	if err == nil {
		err = json.Unmarshal(body, &result)
	}
	//log.Printf("响应数据：%+v", result)

	return result, err
}
