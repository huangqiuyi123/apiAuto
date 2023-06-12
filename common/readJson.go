package common

import (
	"encoding/json"
	"log"
	"os"
	"sync"

	"apiAuto/config"
)

// Params 定义结构体，和json文件的字段一致
type Params struct {
	ID      int         `json:"id"`
	Type    string      `json:"type"`
	Name    string      `json:"name"`
	Url     string      `json:"url"`
	Method  string      `json:"method"`
	ResBody interface{} `json:"resBody"`
}

var fileLocker sync.Mutex //config file locker
var configData = config.GetConfig()

// ReadJson 读取json数据
func ReadJson(environment, filePath string) []Params {

	var params []Params

	fileLocker.Lock()                  //锁定文件，用于实现文件读写操作的并发控制
	data, err := os.ReadFile(filePath) //read config file
	fileLocker.Unlock()                //解锁文件
	if err != nil {
		log.Println("read json file error")
		return params
	}

	//反序列化成结构体格式
	err = json.Unmarshal(data, &params)
	if err != nil {
		log.Println("Unmarshal 解析JSON失败：", err)
		return params
	}

	//拼接并替换环境URL
	params = replaceUrl(environment, params)
	return params

}

// 拼接URL,并替换结构体中的URL
func replaceUrl(environment string, params []Params) []Params {
	// 循环结构体
	for i := range params {
		switch environment {
		case "test":
			// 拼接测试环境URL
			params[i].Url = "https://" + configData.Handle.Test + configData.Environment.Test + params[i].Url
			// 将替换后的Params结构体重新放入切片中
			params[i] = params[i]
		case "pre":
			// 拼接预发布环境URL
			params[i].Url = "https://" + configData.Handle.Pre + configData.Environment.Pre + params[i].Url
			// 将替换后的Params结构体重新放入切片中
			params[i] = params[i]
		case "pro":
			// 拼接正式环境URL
			params[i].Url = "https://" + configData.Handle.Pro + configData.Environment.Pro + params[i].Url
			// 将替换后的Params结构体重新放入切片中
			params[i] = params[i]
		}

	}
	return params
}

//func main() {

//replaceCode(Params{})

//conf := ReadJson()
//// 序列化成json格式
//jsonData, err := json.Marshal(conf)
//if err != nil {
//	fmt.Println("序列化失败!")
//}
////fmt.Println("======================================================")
//
//fmt.Printf("序列化成json格式结果: %v", string(jsonData))
//}
