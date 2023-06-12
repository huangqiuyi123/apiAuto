package common

import (
    "bytes"
    "encoding/json"
    "errors"
    "log"
    "net/http"

    "apiAuto/util"
)

func Get(url string, params map[string]string, headers map[string]string) (interface{}, error) {

    // new request
    request, err := http.NewRequest(http.MethodGet, url, nil)
    if err != nil {
        log.Println(err)
        return nil, errors.New("new request is fail ")
    }

    request.Header.Set("cookie", configData.Header.Cookie)

    // add params
    param := request.URL.Query()
    if params != nil {
        for key, value := range params {
            param.Add(key, value)
        }
        request.URL.RawQuery = param.Encode()
    }

    // add headers
    if headers != nil {
        for key, value := range headers {
            request.Header.Add(key, value)
        }
    }

    // 发起请求
    client := &http.Client{}
    log.Printf("请求方式： %s  ｜ 请求URL : %s \n\n", http.MethodGet, request.URL.String())
    res, err := client.Do(request)

    // 格式化返回值
    response, _ := util.ParseResponse(res)

    return response, err

}

func Post(url string, body interface{}, params map[string]string, headers map[string]string) (map[string]any, error) {
    // add post body
    var bodyJson []byte
    var request *http.Request

    // 判断入参类型
    switch body.(type) {
    case string:
        bodyJson = []byte(body.(string))

    default:
        if body != nil {
            var err error
            bodyJson, err = json.Marshal(body) // 转成json格式
            if err != nil {
                log.Println(err)
                return nil, errors.New("http post body to json failed")
            }
        }
    }

    //http.NewRequest() 方法返回一个 *http.Request 指针以及错误对象 err
    request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(bodyJson))
    if err != nil {
        log.Println(err)
        return nil, errors.New("new request is fail: %v \n")
    }
    // set固定数据到请求头
    request.Header.Set("Content-type", "application/json")
    request.Header.Set("cookie", configData.Header.Cookie)

    // add params 查询参数
    param := request.URL.Query()
    if params != nil {
        // 使用 for 循环遍历 map 类型的 params 中的数据，将其添加到 param 对象中，例如： url?param1=value1&param2=value2
        for key, value := range params {
            param.Add(key, value)
        }

        //将查询参数编码为 URL 格式，并将编码后的字符串赋值给 request.URL.RawQuery 属性
        request.URL.RawQuery = param.Encode()
    }

    // add header
    if headers != nil {
        for key, val := range headers {
            request.Header.Add(key, val)
        }
    }

    // 发起请求
    client := &http.Client{}
    log.Printf("请求方式： %s  ｜ 请求URL : %s \n", http.MethodPost, request.URL.String())
    res, err := client.Do(request)
    // 格式化返回值
    response, _ := util.ParseResponse(res)

    jsonData, _ := json.Marshal(response)

    log.Printf("响应数据： %s", string(jsonData))

    return response, err

}

/*
我们首先创建一个map类型的JSON数据。然后使用json.Marshal()函数将其转换为JSON格式的字节数组。
接下来，我们创建一个新的http.Request对象，并将JSON数据作为请求主体。最后，我们使用http.Client发送请求并处理响应。

需要注意的是，在发送HTTPS请求时，无需添加任何特殊的代码。http.Client将自动使用TLS协议进行加密和身份验证。
但是，在生产环境中，我们建议您加载SSL证书以确保通信安全性。在发送请求时，我们需要确保关闭响应主体，以避免资源泄漏。

*/

//func main() {
//    //get请求
//    //url := "https://niki.myshoplinestg.com/admin/api/sale/plugin/common/design_maker/admin/discount/config"
//    //
//    //header := map[string]string{
//    //	"cookie": "_BEAMER_USER_ID_whBStDUm30458=b99b65e5-1186-4c59-933d-f8e9e5f6f984; _BEAMER_FIRST_VISIT_whBStDUm30458=2023-02-01T06:40:35.843Z; _BEAMER_FIRST_VISIT_undefined=2023-02-06T06:09:10.546Z; n_u=24c37fcbcf4d401d2541fbdc24873b94; r_b_ined=1; mafrc=8XYE699L9E17%2Csales_plugin_af; a_ste=m7tH55QQBALTFezrpyp0YE0xb+hZjcgYSHtmcFcx5Iq01UjUVhSQaMCYXlAAsFoNb8aFMObih5h5l/EAByfhmg==; store_id=1665628907335; merchant_id=4211387386; lang.sig=HPZEXM6qRQA3fl9QF0Gl5KM_KZ7FwUtDpVV9UEUrrek; addressLang.sig=fZhLaUxh_564Gt_Ygb8agf56cVb1lYYp6NMpk7wfgaM; userSelectLocale.sig=xaWhkiDLccJKOWtBx98z0KVVx7o_iP0WoEYPBrEqJCw; s_id=432C37B3A7ECD1BD874A64811EFEFE63; s_id.sig=3afec40af11c2c5d2917020221fa9f02; osudb_lang=en; osudb_hdid=2217590e27c9544604c3575811dae79c; f_ds_info.sig=3CKJGA_m714rZP1Sd6rVfiPZrE-e85dCA1u-vNls85c; store_id.sig=jkiGcSihCNQHcIwuJeDT90gbTVd1S9PYIlgxgvQOJtM; merchant_id.sig=JdG9OdjYaypCshQ26ZRHpUNyoOuPv3UW7kt9Ba4A7VU; tsao=c79f003d-f1bd-4e08-b0a1-fe5c1198b515#ad50675722ed913e; tsui=03eb85fa-4176-4e1e-be02-463406306734#ca17543de9e9389f; r_b_in=1; _tracking_consent=%7B%22con%22%3A%7B%22GDPR%22%3A%22%22%7D%2C%22v%22%3A%221.0%22%2C%22lim%22%3A%5B%5D%2C%22reg%22%3A%22%22%7D; lang=en; addressLang=en; userSelectLocale=en; multi_lang_user_select_lang=en; googtrans=/auto/en; googtrans=/auto/en; t_cart=a7c6adb7110847ffa2de77b2fb527bcb; t_cart.sig=cba6fd8121e83dcb7d10dd69b3f22470; client_abandoned_order_seq=2405911021363771480096%2334eba825db785c3e; a_lang=zh-hans-cn; f_ds_info=vIkVV50Dkzv2zmQ98TKTgtk3b1cl9Ma+OJEOEXxG5CSnNgeHzfIbqulQrZ39Pa3cYEZKj5kOO+OxU5eNBLt8HQ==; country_code=CN; currency_code=CNY; currency_code.sig=BPe38eX0xU_6vUcN8hBCl2wo-vS61R5Qc9m2Ao_uba0; localization=CN; currency_code_userSetting=CNY; currency_code_userSetting.sig=N6p784I4PXJLkvcoM99TDGRh2ydbDLU4t_lmbHvHf5I; history_browse_products=16056550630997722466600637,16057240157463233196040637,16057240158879062455150637,16057240160056990786230637,16059098486729306369010637,16057240155040267655000637,16057240159038278226880637,16057240158655422162910637,16057240156118707106640637,16058402539549580405810637; history_browse_products.sig=x2KDNfYK1g3NYnh1n2Om1iGzYgbRq5P-FjKXFtjSbK0; n_sess={\"session_id\":\"ffcd0a76-0682-49c4-a34a-09f471b258c7\",\"created_at\":1682676917631,\"last_session_id\":\"f10e8c77-bf2c-4545-8686-bee7b14d25db\",\"session_create_type\":102}; a_osudb_appid=1163336839; a_osudb_subappid=1; a_osudb_csrf=#C01#SID0000073BAYkYifkDdmIpM2sbSU3ijaHLz2dj2Z7ylY45OjkwATsBgruf/cZdJxkWxmW0uSbnN+0crQ1fHdzxduuzOSjzQBU+BHcxinfg5tLW8LwLRk7YMiI6oJW04izllZeUoTw9cJcERlxWYYULc/4wjef7hXeUelYI6D2DGj6GEyuFMUQ; a_cstgc=1163336839_1_4211387386_1682945618344TOoBqp; a_osudb_uid=4211387386; a_osudb_oar=#01#SID0000073BFsmZjmOwKuBBGUYID9e9ufc1Y8SdKvic2pxWMMYfoIUdmEm0NQlRV9KXZDjh+GC30QbYvFUDNUYVBh1YTMu8s135O8D6fVBOkTeB6irq3gUu6XxqeNx/NgLUey8tyoZpoDVIQnrA8iXRMct72u1bg==; a_lce=1683550418318; a_dhch=niki; JSESSIONID=39CE8C5958B674F900DD5AE71AD2FDD2; _BEAMER_FILTER_BY_URL_whBStDUm30458=false",
//    //}
//    //Get(url, nil, header)
//
//    // post请求
//    //url := "https://niki.myshoplinestg.com/admin/api/sale/plugin/common/design_maker/admin/discount/config"
//    //body := map[string]interface{}{
//    //	"productDetailConfigEnable": false,
//    //	"settlementConfigEnable":    true,
//    //}
//    //
//    //header := map[string]string{
//    //	"cookie": "_BEAMER_USER_ID_whBStDUm30458=b99b65e5-1186-4c59-933d-f8e9e5f6f984; _BEAMER_FIRST_VISIT_whBStDUm30458=2023-02-01T06:40:35.843Z; _BEAMER_FIRST_VISIT_undefined=2023-02-06T06:09:10.546Z; n_u=24c37fcbcf4d401d2541fbdc24873b94; r_b_ined=1; mafrc=8XYE699L9E17%2Csales_plugin_af; a_ste=m7tH55QQBALTFezrpyp0YE0xb+hZjcgYSHtmcFcx5Iq01UjUVhSQaMCYXlAAsFoNb8aFMObih5h5l/EAByfhmg==; store_id=1665628907335; merchant_id=4211387386; lang.sig=HPZEXM6qRQA3fl9QF0Gl5KM_KZ7FwUtDpVV9UEUrrek; addressLang.sig=fZhLaUxh_564Gt_Ygb8agf56cVb1lYYp6NMpk7wfgaM; userSelectLocale.sig=xaWhkiDLccJKOWtBx98z0KVVx7o_iP0WoEYPBrEqJCw; s_id=432C37B3A7ECD1BD874A64811EFEFE63; s_id.sig=3afec40af11c2c5d2917020221fa9f02; osudb_lang=en; osudb_hdid=2217590e27c9544604c3575811dae79c; f_ds_info.sig=3CKJGA_m714rZP1Sd6rVfiPZrE-e85dCA1u-vNls85c; store_id.sig=jkiGcSihCNQHcIwuJeDT90gbTVd1S9PYIlgxgvQOJtM; merchant_id.sig=JdG9OdjYaypCshQ26ZRHpUNyoOuPv3UW7kt9Ba4A7VU; tsao=c79f003d-f1bd-4e08-b0a1-fe5c1198b515#ad50675722ed913e; tsui=03eb85fa-4176-4e1e-be02-463406306734#ca17543de9e9389f; r_b_in=1; _tracking_consent=%7B%22con%22%3A%7B%22GDPR%22%3A%22%22%7D%2C%22v%22%3A%221.0%22%2C%22lim%22%3A%5B%5D%2C%22reg%22%3A%22%22%7D; lang=en; addressLang=en; userSelectLocale=en; multi_lang_user_select_lang=en; googtrans=/auto/en; googtrans=/auto/en; t_cart=a7c6adb7110847ffa2de77b2fb527bcb; t_cart.sig=cba6fd8121e83dcb7d10dd69b3f22470; client_abandoned_order_seq=2405911021363771480096%2334eba825db785c3e; a_lang=zh-hans-cn; f_ds_info=vIkVV50Dkzv2zmQ98TKTgtk3b1cl9Ma+OJEOEXxG5CSnNgeHzfIbqulQrZ39Pa3cYEZKj5kOO+OxU5eNBLt8HQ==; country_code=CN; currency_code=CNY; currency_code.sig=BPe38eX0xU_6vUcN8hBCl2wo-vS61R5Qc9m2Ao_uba0; localization=CN; currency_code_userSetting=CNY; currency_code_userSetting.sig=N6p784I4PXJLkvcoM99TDGRh2ydbDLU4t_lmbHvHf5I; history_browse_products=16056550630997722466600637,16057240157463233196040637,16057240158879062455150637,16057240160056990786230637,16059098486729306369010637,16057240155040267655000637,16057240159038278226880637,16057240158655422162910637,16057240156118707106640637,16058402539549580405810637; history_browse_products.sig=x2KDNfYK1g3NYnh1n2Om1iGzYgbRq5P-FjKXFtjSbK0; n_sess={\"session_id\":\"ffcd0a76-0682-49c4-a34a-09f471b258c7\",\"created_at\":1682676917631,\"last_session_id\":\"f10e8c77-bf2c-4545-8686-bee7b14d25db\",\"session_create_type\":102}; a_osudb_appid=1163336839; a_osudb_subappid=1; a_osudb_csrf=#C01#SID0000073BAYkYifkDdmIpM2sbSU3ijaHLz2dj2Z7ylY45OjkwATsBgruf/cZdJxkWxmW0uSbnN+0crQ1fHdzxduuzOSjzQBU+BHcxinfg5tLW8LwLRk7YMiI6oJW04izllZeUoTw9cJcERlxWYYULc/4wjef7hXeUelYI6D2DGj6GEyuFMUQ; a_cstgc=1163336839_1_4211387386_1682945618344TOoBqp; a_osudb_uid=4211387386; a_osudb_oar=#01#SID0000073BFsmZjmOwKuBBGUYID9e9ufc1Y8SdKvic2pxWMMYfoIUdmEm0NQlRV9KXZDjh+GC30QbYvFUDNUYVBh1YTMu8s135O8D6fVBOkTeB6irq3gUu6XxqeNx/NgLUey8tyoZpoDVIQnrA8iXRMct72u1bg==; a_lce=1683550418318; a_dhch=niki; JSESSIONID=39CE8C5958B674F900DD5AE71AD2FDD2; _BEAMER_FILTER_BY_URL_whBStDUm30458=false",
//    //}
//    //
//    //Post(url, body, nil, header)
//
//
//}
