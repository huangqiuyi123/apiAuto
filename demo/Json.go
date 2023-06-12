package main

import (
	"os"

	"github.com/antchfx/jsonquery"
)

func jsonData() (request interface{}) {
	f, _ := os.Open("testData/DiscountCodeActivity.json")
	doc, _ := jsonquery.Parse(f)
	// 获取指定数据
	req := doc.SelectElement("add").Value()
	//method = doc.SelectElement("method").Value()
	//url = doc.SelectElement("url").Value()
	//resBody = doc.SelectElement("resBody").Value()

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
	//fmt.Println("格式化后的数据： ", resBody.Value())

	return req

}
