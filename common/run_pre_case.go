package common

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/huangqiuyi123/apiAuto/config"
	"github.com/tidwall/sjson"
)

func GetLocationSeq(environment string) string {
	/*
	   获取地址库ID
	*/
	conf := config.GetConfig()
	preCodePath := conf.FilePath.PreCasePath
	request := ReadJson(environment, preCodePath)

	url := request[0].Url
	resBody := request[0].ResBody

	response, _ := Post(url, resBody, nil, nil)
	// 提取并返回地址库ID
	locationSeq := response["data"].(map[string]interface{})["list"].([]interface{})[0].(map[string]interface{})["locationSeq"].(string)

	return locationSeq
}

func AddProduct(environment string) (string, string) {
	/*
	   添加商品
	*/
	conf := config.GetConfig()
	preCodePath := conf.FilePath.PreCasePath
	request := ReadJson(environment, preCodePath)
	url := request[1].Url
	resBody := request[1].ResBody

	jsonData, err := json.Marshal(resBody)
	if err != nil {
		log.Println(err)
	}

	// 替换地址库ID
	locationId := GetLocationSeq(environment)
	newBody, _ := sjson.Set(string(jsonData), "skuList.0.locationStockList.0.locationId", locationId)

	// 提取并返回商品ID
	response, _ := Post(url, newBody, nil, nil)
	spu := response["data"].(map[string]interface{})["spuSeq"].(string)
	sku := response["data"].(map[string]interface{})["skuList"].([]interface{})[0].(map[string]interface{})["skuSeq"].(string)

	return spu, sku

}

func AddCustomer(environment string) string {
	/*
	   添加客户
	*/
	conf := config.GetConfig()
	preCodePath := conf.FilePath.PreCasePath
	request := ReadJson(environment, preCodePath)

	url := request[2].Url
	resBody := request[2].ResBody

	jsonData, err := json.Marshal(resBody)
	if err != nil {
		log.Println(err)
	}

	// 设置随机邮件和电话号码
	randomEmail := fmt.Sprintf("%d%s", time.Now().Unix(), "@test.com")
	randomMobile := "00509904" + fmt.Sprintf("%d%s", time.Now().Unix(), "")

	// 替换json中的邮件和电话号码字段
	newEmail, _ := sjson.Set(string(jsonData), "userInfo.email", randomEmail)
	newBody, _ := sjson.Set(newEmail, "userInfo.mobile", randomMobile)

	response, _ := Post(url, newBody, nil, nil)

	// 提取客户ID并返回
	memberId, _ := response["data"].(map[string]interface{})["uid"].(string)

	return memberId

}

func EnableCustomStyles(environment, activitySeq string) map[string]any {
	/*
	   后置用例：开启自定义样式
	*/
	conf := config.GetConfig()
	preCodePath := conf.FilePath.PreCasePath
	request := ReadJson(environment, preCodePath)

	url := request[3].Url
	resBody := request[3].ResBody

	jsonData, err := json.Marshal(resBody)
	if err != nil {
		log.Println(err)
	}

	// 替换json中的activitySeq
	newBody, _ := sjson.Set(string(jsonData), "activitySeq", activitySeq)

	response, _ := Post(url, newBody, nil, nil)

	return response

}
