package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// get请求
	url := "https://niki-us.myshopline.com/admin/api/sale/plugin/common/design_maker/admin/discount/config"
	method := "GET"

	//构造请求
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	// 抛出请求异常
	if err != nil {
		fmt.Println(err)
		return
	}

	// 添加请求头信息
	cookie := "a_ste=krobHo7p7fCJvjzWRX1BOtysenebDqVhBM3UUB9n7mIpWryZf4xiHtWPTGkB0atyb8aFMObih5h5l/EAByfhmg==; _BEAMER_USER_ID_vXnSKpuN10685=47d215e2-aefb-4d86-be1e-3c645d117e6a; _BEAMER_FIRST_VISIT_vXnSKpuN10685=2023-04-20T02:58:49.666Z; r_b_ined=1; n_u=d7dbcfcf02198f96efe7649caeb5404b; f_ds_info=xnA90F0gs63umF/ghOwdanrUC1tja3iIUvFz5GfmCnPnSaw+ZHwCHF5VaiHak+9nYEZKj5kOO+OxU5eNBLt8HQ==; store_id=1678861295829; merchant_id=2005188193; currency_code=USD; currency_code.sig=nEGddW1-E-8oJfI_Pm_5XNzC2sMi1n3aVzZ3v01csyY; localization=US; lang=en; lang.sig=HPZEXM6qRQA3fl9QF0Gl5KM_KZ7FwUtDpVV9UEUrrek; addressLang=en; addressLang.sig=fZhLaUxh_564Gt_Ygb8agf56cVb1lYYp6NMpk7wfgaM; userSelectLocale=en; userSelectLocale.sig=xaWhkiDLccJKOWtBx98z0KVVx7o_iP0WoEYPBrEqJCw; currency_code_userSetting=USD; currency_code_userSetting.sig=wreMdGqvcOcZfYXi-Fd1QDxl5OWoQm3s2QLyXkCpvxE; s_id=B8591DD1728567EC40102D889159EDB2; s_id.sig=dd2ba5a52e3783958500ee808d8ba9c4; f_ds_info.sig=-IHUrKyOKPqi3fE7JZ3SJ_oqbA1x0nfO-Jtpka_9bkw; store_id.sig=VEmXRjkNIiFmmixFNFPPl7eJvMBGXEo8kinlasaK31U; merchant_id.sig=jcb6ufq_TwH_ddpzLjU_9jrLftb0Xu1256208WDg258; tsao=8644ad69-a3f3-4189-ba51-95420d05d3fe#6a50d50cc3155218; tsui=33b2eee1-2e8f-408a-a9fb-4db5da103fd4#4ad827447c497c74; osudb_lang=en; osudb_hdid=906d8244f7764f550956fabf71de269b; client_abandoned_order_seq=2405899865880712488388%23c0d84e732d6d08bc; t_cart=57607ae1258b4f1b8b57f4235d913ae2; t_cart.sig=2b54ff2d64c600210a93ba29f0c67e23; r_b_in=1; __stripe_mid=c08e7d40-3458-4617-bc4c-f7d1f8c09de24b3fba; lp_url={%22landingPageHtml%22:%22https://niki-us.myshopline.com/collections%22%2C%22occurredAt%22:1682303277758}; _BEAMER_BOOSTED_ANNOUNCEMENT_DATE_vXnSKpuN10685=2023-04-25T06:49:01.272Z; a_lang=zh-hans-cn; a_osudb_appid=1165600903; a_osudb_subappid=1; a_cstgc=1165600903_1_2005188193_1682577214629fAuRAl; a_osudb_uid=2005188193; a_osudb_oar=#01#SID0000041BALBv4CvNiubnfzHkkwd9u5yHUNVR1zVdNMW9hf8LNOvwBmGAq8IIin9z5DLLMYaTAb7FkZGljIIuSqzQkCWki7T+14UPQS67XFBMxKh2GJHHA1rI80axly+tBqGRSOfEvek51hHNH/00BSnRMj02A==; n_sess={\"session_id\":\"c68fc3db-1b9b-4db0-9cbb-2454ff173b51\",\"created_at\":1682676865167,\"last_session_id\":\"493e6abc-642c-445e-bf69-547f8fbbe17c\",\"session_create_type\":102}; history_browse_products=16058474193355014730353962,16058474190524362850643962,16058474197045834480933962; history_browse_products.sig=4ZI502c_eD-V-dehVf_weobdjt5YD6Ae9TFUK7Ia_uo; JSESSIONID=195AE92691100C441333BC71886ABCE1; _BEAMER_LAST_UPDATE_vXnSKpuN10685=1682842611945; a_dhch=niki-us; intercom-device-id-ryx6jq53=3af03eb0-d3cd-4c25-8255-5ebdb0cb8bab; intercom-session-ryx6jq53=QURHRytlMHB5bU9pNHpyTFBZMkVnblM3NVZSSlRqVTU1cVkyM0xUdlpXQ1dudU9udXN3TFNaaEVCUjJMT1Mwby0tb2xLdlF2WHBnTTdLODV5dHloWDdNZz09--24ddb6fdf3302505770946d7fa76f773e349e90e; _BEAMER_FILTER_BY_URL_vXnSKpuN10685=false"
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("cookie", cookie)

	// 发起请求
	res, err := client.Do(req)
	// 回收请求异常
	if err != nil {
		fmt.Println(err)
		return
	}

	// 关闭请求连接
	defer res.Body.Close()

	// 获取返回体，并回收返回异常
	boby, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// json字符串格式化
	var str bytes.Buffer
	_ = json.Indent(&str, boby, "", "    ")
	fmt.Println("data", str.String())

}
