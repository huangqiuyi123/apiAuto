package main

//
//import (
//	"fmt"
//	"os"
//
//	"github.com/antchfx/jsonquery"
//)
//
//func main() {
//	f, _ := os.Open("testData/Activity.json")
//	doc, _ := jsonquery.Parse(f)
//
//	// 遍厉json结构
//	for _, n := range doc.ChildNodes() {
//		fmt.Printf("data: %v", n.Value())
//	}
//
//	// 获取json指定字段的值
//	//data := jsonquery.FindOne(doc, "*")
//	//fmt.Printf("%#v[%T]\\n", data.Value(), data.Value())
//
//	// 获取json全部字段数据
//	//list := jsonquery.Find(doc, "*")
//	//for _, n := range list {
//	//	fmt.Println(n.Value())
//	//}
//
//}
