package common

import (
	"encoding/json"
	"log"

	"apiAuto/config"
	"apiAuto/util"
	"github.com/tidwall/sjson"
)

func ExecuteDiscountActivity(environment string) {
	//获取数据
	conf := config.GetConfig()
	discountCodePath := conf.FilePath.DiscountCodePath
	request := ReadJson(environment, discountCodePath)

	// 循环发起请求json文件的数据
	for i := 0; i < len(request); i++ {
		id := request[i].ID
		dataType := request[i].Type
		activityType := request[i].ActivityType
		name := request[i].Name
		url := request[i].Url
		method := request[i].Method
		resBody := request[i].ResBody
		log.Printf("执行用例 %d ｜ ", id)
		log.Print(name, " | ")
		// 为了使用sjson这个工具类的替换，解析为json格式
		jsonData, err := json.Marshal(resBody)
		if err != nil {
			log.Println(err)
			return
		}

		if activityType == "coupon" {
			// 生成随机折扣码
			randomCodes := util.RandomString(1, 12)

			//替换折扣码
			newBody, _ := sjson.Set(string(jsonData), "promotionCodes", randomCodes)

			switch {
			case dataType == "all":
				///*
				//   直接创建，无限制
				//*/
				{
					if method == "POST" {
						response, _ := Post(url, newBody, nil, nil)
						activitySeq := response["data"].(string)
						EnableCustomStyles(environment, activitySeq) //后置操作，开启自定义样式
					}

				}

			case dataType == "specifyCustomer":
				/*
				   指定客户需要前置新增客户
				*/
				{
					memberId := AddCustomer(environment)

					//替换客户ID
					newBody1, _ := sjson.Set(newBody, "targetUserInfoList.0.memberId", memberId)

					if method == "POST" {
						response, _ := Post(url, newBody1, nil, nil)
						activitySeq := response["data"].(string)
						EnableCustomStyles(environment, activitySeq) //后置操作，开启自定义样式
					}

				}

			case dataType == "specifyProduct":
				/*
				   指定商品需要前置新增商品
				*/
				{

					spu, sku := AddProduct(environment)

					// 替换spu和sku
					newBody1, _ := sjson.Set(newBody, "effectScopeValueList.0.productId", spu)
					newBody2, _ := sjson.Set(newBody1, "effectScopeValueList.0.spuSeq", spu)
					newBody3, _ := sjson.Set(newBody2, "effectScopeValueList.0.skuList.0.skuSeq", sku)

					if method == "POST" {
						response, _ := Post(url, newBody3, nil, nil)

						activitySeq := response["data"].(string)
						EnableCustomStyles(environment, activitySeq) //后置操作，开启自定义样式
					}

				}

			case dataType == "specifyProductAndCustomer":
				/*
				   指定商品需要前置新增商品
				*/
				{
					memberId := AddCustomer(environment)
					spu, sku := AddProduct(environment)

					//替换客户ID
					newBody1, _ := sjson.Set(newBody, "targetUserInfoList.0.memberId", memberId)

					// 替换spu和sku

					newBody2, _ := sjson.Set(newBody1, "effectScopeValueList.0.productId", spu)
					newBody3, _ := sjson.Set(newBody2, "effectScopeValueList.0.spuSeq", spu)
					newBody4, _ := sjson.Set(newBody3, "effectScopeValueList.0.skuList.0.skuSeq", sku)

					if method == "POST" {
						response, _ := Post(url, newBody4, nil, nil)
						activitySeq := response["data"].(string)
						EnableCustomStyles(environment, activitySeq) //后置操作，开启自定义样式
					}

				}

			case dataType == "buy X get Y free":
				/*
				   指定商品需要前置新增商品
				*/
				{
					spu, sku := AddProduct(environment)
					// 替换spu和sku
					newBody1, _ := sjson.Set(newBody, "benefitConditions.0.benefit.benefitProducts.0.productId", spu)
					newBody2, _ := sjson.Set(newBody1, "benefitConditions.0.benefit.benefitProducts.0.spuSeq", spu)
					newBody3, _ := sjson.Set(newBody2, "benefitConditions.0.benefit.benefitProducts.0.skuList.0.skuSeq", sku)

					newBody4, _ := sjson.Set(newBody3, "effectScopeValueList.0.spuSeq", spu)
					newBody5, _ := sjson.Set(newBody4, "effectScopeValueList.0.productId", spu)
					newBody6, _ := sjson.Set(newBody5, "effectScopeValueList.0.skuList.0.skuSeq", sku)

					if method == "POST" {
						response, _ := Post(url, newBody6, nil, nil)
						activitySeq := response["data"].(string)
						EnableCustomStyles(environment, activitySeq) //后置操作，开启自定义样式
					}

				}

			}

		} else if activityType == "automaticDiscount" {
			switch {
			case dataType == "specifyCustomer":
				/*
				   指定客户需要前置新增客户
				*/
				{
					memberId := AddCustomer(environment)

					//替换客户ID
					newBody1, _ := sjson.Set(string(jsonData), "targetUserInfoList.0.memberId", memberId)

					if method == "POST" {
						response, _ := Post(url, newBody1, nil, nil)
						activitySeq := response["data"].(map[string]interface{})["activitySeq"].(string)
						EnableCustomStyles(environment, activitySeq) //后置操作，开启自定义样式
					}

				}

			case dataType == "specifyProduct":
				/*
				   指定商品需要前置新增商品
				*/
				{

					spu, sku := AddProduct(environment)

					// 替换spu和sku
					newBody1, _ := sjson.Set(string(jsonData), "effectScopeValueList.0.productId", spu)
					newBody2, _ := sjson.Set(newBody1, "effectScopeValueList.0.spuSeq", spu)
					newBody3, _ := sjson.Set(newBody2, "effectScopeValueList.0.skuList.0.skuSeq", sku)

					if method == "POST" {
						response, _ := Post(url, newBody3, nil, nil)
						activitySeq := response["data"].(map[string]interface{})["activitySeq"].(string)
						EnableCustomStyles(environment, activitySeq) //后置操作，开启自定义样式
					}

				}

			case dataType == "specifyProductAndCustomer":
				/*
				   指定商品需要前置新增商品
				*/
				{
					memberId := AddCustomer(environment)
					spu, sku := AddProduct(environment)

					//替换客户ID
					newBody1, _ := sjson.Set(string(jsonData), "targetUserInfoList.0.memberId", memberId)

					// 替换spu和sku

					newBody2, _ := sjson.Set(newBody1, "effectScopeValueList.0.productId", spu)
					newBody3, _ := sjson.Set(newBody2, "effectScopeValueList.0.spuSeq", spu)
					newBody4, _ := sjson.Set(newBody3, "effectScopeValueList.0.skuList.0.skuSeq", sku)

					if method == "POST" {
						response, _ := Post(url, newBody4, nil, nil)
						activitySeq := response["data"].(map[string]interface{})["activitySeq"].(string)
						EnableCustomStyles(environment, activitySeq) //后置操作，开启自定义样式
					}

				}

			case dataType == "buy X get Y free":
				/*
				   指定商品需要前置新增商品
				*/
				{
					spu, sku := AddProduct(environment)
					// 替换spu和sku
					newBody1, _ := sjson.Set(string(jsonData), "benefitConditions.0.benefit.benefitProducts.0.productId", spu)
					newBody2, _ := sjson.Set(newBody1, "benefitConditions.0.benefit.benefitProducts.0.spuSeq", spu)
					newBody3, _ := sjson.Set(newBody2, "benefitConditions.0.benefit.benefitProducts.0.skuList.0.skuSeq", sku)

					newBody4, _ := sjson.Set(newBody3, "effectScopeValueList.0.spuSeq", spu)
					newBody5, _ := sjson.Set(newBody4, "effectScopeValueList.0.productId", spu)
					newBody6, _ := sjson.Set(newBody5, "effectScopeValueList.0.skuList.0.skuSeq", sku)

					if method == "POST" {
						response, _ := Post(url, newBody6, nil, nil)
						activitySeq := response["data"].(map[string]interface{})["activitySeq"].(string)
						EnableCustomStyles(environment, activitySeq) //后置操作，开启自定义样式
					}

				}

			}

		}
	}

}
