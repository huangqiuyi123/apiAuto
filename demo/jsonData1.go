package main

import (
	"os"

	"github.com/antchfx/jsonquery"
)

const file = "testData/DiscountCodeActivity.json"

// 读取json数据
func ReadJson1() (request interface{}) {
	f, _ := os.Open(file)
	doc, _ := jsonquery.Parse(f)
	//// 获取指定数据
	//ID = doc.SelectElement("ID").Value()
	//method = doc.SelectElement("method").Value()
	//url = doc.SelectElement("url").Value()
	//resBody = doc.SelectElement("resBody").Value()

	nodes := doc.ChildNodes()
	var listRequest = make([]interface{}, 0)
	for i := 0; i < len(nodes); i++ {
		request = nodes[i].SelectElement("add").Value()
		listRequest = append(listRequest, request)
		//fmt.Println("ly:", request.(string))
		//
		//var p AddActivity
		//err := json.Unmarshal([]byte(request.(string)), &p)
		//if err != nil {
		//	fmt.Println("转换失败：", err)
		//}
	}
	//fmt.Println(listRequest)

	//request = doc.FirstChild.SelectElement("add").Value()

	//// 遍厉json结构
	//for _, n := range doc.ChildNodes() {
	//    fmt.Printf("data: %v", n.Value())
	//}

	// 获取json全部字段数据
	//list := jsonquery.Find(doc, "*")
	//for _, n := range list {
	//	fmt.Println(n.Value())
	//}

	//获取json指定字段的值
	//ID := jsonquery.FindOne(doc, "ID")
	//url := jsonquery.FindOne(doc, "url")
	//method := jsonquery.FindOne(doc, "method")
	//resBody := jsonquery.FindOne(doc, "resBody")
	//list, _ := jsonquery.QueryAll(doc, "resBody")
	//fmt.Println(list)

	//fmt.Printf("%#v\n", ID)
	//fmt.Printf("%#v\n", url)
	//fmt.Printf("%#v\n", method)
	//fmt.Printf("%#v\n", resBody)
	//fmt.Println("格式化后的数据： ", request)
	//fmt.Printf("数据类型%T： ", resBody)

	//ID, url, method,
	return listRequest

}

//
//func main() {
//	ReadJson1()
//}
