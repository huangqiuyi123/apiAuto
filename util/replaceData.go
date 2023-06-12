package util

import (
	"encoding/json"
	"log"
)

func ReplaceData(field string, ReplaceValue []string, resBody interface{}) map[string]interface{} {

	data, _ := json.Marshal(resBody)

	var jsonData map[string]interface{}

	err := json.Unmarshal(data, &jsonData)
	if err != nil {
		log.Println("Unmarshal解析失败： ", err)
	}

	jsonData[field] = ReplaceValue

	return jsonData

}

func PreReplaceData(field string, ReplaceValue string, resBody interface{}) map[string]interface{} {

	//switch resBody.(type) {
	//case []any:
	//	resBody = resBody.([]any)[0]
	//}

	data, _ := json.Marshal(resBody)

	var jsonData map[string]interface{}

	err := json.Unmarshal(data, &jsonData)
	if err != nil {
		log.Println("Unmarshal解析失败： ", err)
	}

	jsonData[field] = ReplaceValue

	return jsonData

}

func Interface2Map(resBody interface{}) map[string]interface{} {
	var jsonData map[string]interface{}
	data, _ := json.Marshal(resBody)
	err := json.Unmarshal(data, &jsonData)
	if err != nil {
		log.Println("Unmarshal解析失败： ", err)
	}
	return jsonData
}
