package main

// func doPostWithAddActivity(AddActivity ){
//
//	common.Post(url, resBody, nil, header)
//
// }
//func main() {
//	var request []common.AddActivity = common.ReadJson()
//
//	header := map[string]string{
//		"cookie": "_BEAMER_USER_ID_whBStDUm30458=b99b65e5-1186-4c59-933d-f8e9e5f6f984; _BEAMER_FIRST_VISIT_whBStDUm30458=2023-02-01T06:40:35.843Z; _BEAMER_FIRST_VISIT_undefined=2023-02-06T06:09:10.546Z; n_u=24c37fcbcf4d401d2541fbdc24873b94; r_b_ined=1; mafrc=8XYE699L9E17%2Csales_plugin_af; a_ste=m7tH55QQBALTFezrpyp0YE0xb+hZjcgYSHtmcFcx5Iq01UjUVhSQaMCYXlAAsFoNb8aFMObih5h5l/EAByfhmg==; store_id=1665628907335; merchant_id=4211387386; lang.sig=HPZEXM6qRQA3fl9QF0Gl5KM_KZ7FwUtDpVV9UEUrrek; addressLang.sig=fZhLaUxh_564Gt_Ygb8agf56cVb1lYYp6NMpk7wfgaM; userSelectLocale.sig=xaWhkiDLccJKOWtBx98z0KVVx7o_iP0WoEYPBrEqJCw; s_id=432C37B3A7ECD1BD874A64811EFEFE63; s_id.sig=3afec40af11c2c5d2917020221fa9f02; osudb_lang=en; osudb_hdid=2217590e27c9544604c3575811dae79c; f_ds_info.sig=3CKJGA_m714rZP1Sd6rVfiPZrE-e85dCA1u-vNls85c; store_id.sig=jkiGcSihCNQHcIwuJeDT90gbTVd1S9PYIlgxgvQOJtM; merchant_id.sig=JdG9OdjYaypCshQ26ZRHpUNyoOuPv3UW7kt9Ba4A7VU; tsao=c79f003d-f1bd-4e08-b0a1-fe5c1198b515#ad50675722ed913e; tsui=03eb85fa-4176-4e1e-be02-463406306734#ca17543de9e9389f; r_b_in=1; _tracking_consent=%7B%22con%22%3A%7B%22GDPR%22%3A%22%22%7D%2C%22v%22%3A%221.0%22%2C%22lim%22%3A%5B%5D%2C%22reg%22%3A%22%22%7D; lang=en; addressLang=en; userSelectLocale=en; multi_lang_user_select_lang=en; googtrans=/auto/en; googtrans=/auto/en; t_cart=a7c6adb7110847ffa2de77b2fb527bcb; t_cart.sig=cba6fd8121e83dcb7d10dd69b3f22470; client_abandoned_order_seq=2405911021363771480096%2334eba825db785c3e; a_lang=zh-hans-cn; f_ds_info=vIkVV50Dkzv2zmQ98TKTgtk3b1cl9Ma+OJEOEXxG5CSnNgeHzfIbqulQrZ39Pa3cYEZKj5kOO+OxU5eNBLt8HQ==; country_code=CN; currency_code=CNY; currency_code.sig=BPe38eX0xU_6vUcN8hBCl2wo-vS61R5Qc9m2Ao_uba0; localization=CN; currency_code_userSetting=CNY; currency_code_userSetting.sig=N6p784I4PXJLkvcoM99TDGRh2ydbDLU4t_lmbHvHf5I; history_browse_products=16056550630997722466600637,16057240157463233196040637,16057240158879062455150637,16057240160056990786230637,16059098486729306369010637,16057240155040267655000637,16057240159038278226880637,16057240158655422162910637,16057240156118707106640637,16058402539549580405810637; history_browse_products.sig=x2KDNfYK1g3NYnh1n2Om1iGzYgbRq5P-FjKXFtjSbK0; n_sess={\"session_id\":\"ffcd0a76-0682-49c4-a34a-09f471b258c7\",\"created_at\":1682676917631,\"last_session_id\":\"f10e8c77-bf2c-4545-8686-bee7b14d25db\",\"session_create_type\":102}; JSESSIONID=39CE8C5958B674F900DD5AE71AD2FDD2; a_osudb_appid=1163336839; a_osudb_csrf=#C01#SID0000073BDzw324C8mN618NqJVRpzb/YMvRVgPsQlubkMBNZuEbBlixnUZ8XtBBktgftpc4UDsB+1Gy275gLe13c1DN0zN7EOWHfJzc3Ce/gsnlVIbjvlVNTDnLvNOzWWS0HPFV6kFegfGpDREsZ3t96ZTUDbKc4UV52eI0bE78LsvmC76uH; a_cstgc=1163336839_1_4211387386_1683107613931LplPEu; a_osudb_uid=4211387386; a_osudb_oar=#01#SID0000073BCyBuG3HqJMALwpB8h9HOUz8cunvWOvhs8fR+CYRdSz2okFTql4GlzmU7JoVebQ/Ib40S2HZak6fMp1yZbFF3qKVVA+V+sL49tw7ZpxkRZEofh34Q4qwg6gsYKV95TZrpOXzHBmUPvnDq9nN1RqR2A==; a_osudb_subappid=1; a_lce=1683712414010; a_dhch=niki; _BEAMER_FILTER_BY_URL_whBStDUm30458=false",
//	}
//
//	for i := 0; i < len(request); i++ {
//		url := request[i].Url
//		resBody := request[i].ResBody
//		name := request[i].Name
//		log.Println(name)
//		common.Post(url, resBody, nil, header)
//	}

//for i := 0; i < len(data); i++ {
//    var tempData = data[i].(interface{})
//fmt.Println(tempData)
//id := tempData["id"]
//name := tempData["name"].(string)
//url := tempData["url"].(string)
//resBody := tempData["resBody"].(map[string]interface{})
//header := map[string]string{
//    "cookie": "_BEAMER_USER_ID_whBStDUm30458=b99b65e5-1186-4c59-933d-f8e9e5f6f984; _BEAMER_FIRST_VISIT_whBStDUm30458=2023-02-01T06:40:35.843Z; _BEAMER_FIRST_VISIT_undefined=2023-02-06T06:09:10.546Z; n_u=24c37fcbcf4d401d2541fbdc24873b94; r_b_ined=1; mafrc=8XYE699L9E17%2Csales_plugin_af; a_ste=m7tH55QQBALTFezrpyp0YE0xb+hZjcgYSHtmcFcx5Iq01UjUVhSQaMCYXlAAsFoNb8aFMObih5h5l/EAByfhmg==; store_id=1665628907335; merchant_id=4211387386; lang.sig=HPZEXM6qRQA3fl9QF0Gl5KM_KZ7FwUtDpVV9UEUrrek; addressLang.sig=fZhLaUxh_564Gt_Ygb8agf56cVb1lYYp6NMpk7wfgaM; userSelectLocale.sig=xaWhkiDLccJKOWtBx98z0KVVx7o_iP0WoEYPBrEqJCw; s_id=432C37B3A7ECD1BD874A64811EFEFE63; s_id.sig=3afec40af11c2c5d2917020221fa9f02; osudb_lang=en; osudb_hdid=2217590e27c9544604c3575811dae79c; f_ds_info.sig=3CKJGA_m714rZP1Sd6rVfiPZrE-e85dCA1u-vNls85c; store_id.sig=jkiGcSihCNQHcIwuJeDT90gbTVd1S9PYIlgxgvQOJtM; merchant_id.sig=JdG9OdjYaypCshQ26ZRHpUNyoOuPv3UW7kt9Ba4A7VU; tsao=c79f003d-f1bd-4e08-b0a1-fe5c1198b515#ad50675722ed913e; tsui=03eb85fa-4176-4e1e-be02-463406306734#ca17543de9e9389f; r_b_in=1; _tracking_consent=%7B%22con%22%3A%7B%22GDPR%22%3A%22%22%7D%2C%22v%22%3A%221.0%22%2C%22lim%22%3A%5B%5D%2C%22reg%22%3A%22%22%7D; lang=en; addressLang=en; userSelectLocale=en; multi_lang_user_select_lang=en; googtrans=/auto/en; googtrans=/auto/en; t_cart=a7c6adb7110847ffa2de77b2fb527bcb; t_cart.sig=cba6fd8121e83dcb7d10dd69b3f22470; client_abandoned_order_seq=2405911021363771480096%2334eba825db785c3e; a_lang=zh-hans-cn; f_ds_info=vIkVV50Dkzv2zmQ98TKTgtk3b1cl9Ma+OJEOEXxG5CSnNgeHzfIbqulQrZ39Pa3cYEZKj5kOO+OxU5eNBLt8HQ==; country_code=CN; currency_code=CNY; currency_code.sig=BPe38eX0xU_6vUcN8hBCl2wo-vS61R5Qc9m2Ao_uba0; localization=CN; currency_code_userSetting=CNY; currency_code_userSetting.sig=N6p784I4PXJLkvcoM99TDGRh2ydbDLU4t_lmbHvHf5I; history_browse_products=16056550630997722466600637,16057240157463233196040637,16057240158879062455150637,16057240160056990786230637,16059098486729306369010637,16057240155040267655000637,16057240159038278226880637,16057240158655422162910637,16057240156118707106640637,16058402539549580405810637; history_browse_products.sig=x2KDNfYK1g3NYnh1n2Om1iGzYgbRq5P-FjKXFtjSbK0; n_sess={\"session_id\":\"ffcd0a76-0682-49c4-a34a-09f471b258c7\",\"created_at\":1682676917631,\"last_session_id\":\"f10e8c77-bf2c-4545-8686-bee7b14d25db\",\"session_create_type\":102}; JSESSIONID=39CE8C5958B674F900DD5AE71AD2FDD2; a_osudb_appid=1163336839; a_osudb_csrf=#C01#SID0000073BDzw324C8mN618NqJVRpzb/YMvRVgPsQlubkMBNZuEbBlixnUZ8XtBBktgftpc4UDsB+1Gy275gLe13c1DN0zN7EOWHfJzc3Ce/gsnlVIbjvlVNTDnLvNOzWWS0HPFV6kFegfGpDREsZ3t96ZTUDbKc4UV52eI0bE78LsvmC76uH; a_cstgc=1163336839_1_4211387386_1683107613931LplPEu; a_osudb_uid=4211387386; a_osudb_oar=#01#SID0000073BCyBuG3HqJMALwpB8h9HOUz8cunvWOvhs8fR+CYRdSz2okFTql4GlzmU7JoVebQ/Ib40S2HZak6fMp1yZbFF3qKVVA+V+sL49tw7ZpxkRZEofh34Q4qwg6gsYKV95TZrpOXzHBmUPvnDq9nN1RqR2A==; a_osudb_subappid=1; a_lce=1683712414010; a_dhch=niki; _BEAMER_FILTER_BY_URL_whBStDUm30458=false",
//}
//
//common.Post(url, resBody, nil, header)
//}

//activity := AddActivity{
//    ID:      data["ID"].(int),
//    Name:    data["Name"].(string),
//    Url:     data["url"].(string),
//    Method:  data["method"].(string),
//    ResBody: data["resBody"].(map[string]interface{}),
//}

//id := data["id"]
//name := data["name"].(string)
//url := data["url"].(string)
////method := data["method"].(string)
//resBody := data["resBody"].(map[string]interface{})
//// 添加请求头信息
//header := map[string]string{
//	"cookie": "_BEAMER_USER_ID_whBStDUm30458=b99b65e5-1186-4c59-933d-f8e9e5f6f984; _BEAMER_FIRST_VISIT_whBStDUm30458=2023-02-01T06:40:35.843Z; _BEAMER_FIRST_VISIT_undefined=2023-02-06T06:09:10.546Z; n_u=24c37fcbcf4d401d2541fbdc24873b94; r_b_ined=1; mafrc=8XYE699L9E17%2Csales_plugin_af; a_ste=m7tH55QQBALTFezrpyp0YE0xb+hZjcgYSHtmcFcx5Iq01UjUVhSQaMCYXlAAsFoNb8aFMObih5h5l/EAByfhmg==; store_id=1665628907335; merchant_id=4211387386; lang.sig=HPZEXM6qRQA3fl9QF0Gl5KM_KZ7FwUtDpVV9UEUrrek; addressLang.sig=fZhLaUxh_564Gt_Ygb8agf56cVb1lYYp6NMpk7wfgaM; userSelectLocale.sig=xaWhkiDLccJKOWtBx98z0KVVx7o_iP0WoEYPBrEqJCw; s_id=432C37B3A7ECD1BD874A64811EFEFE63; s_id.sig=3afec40af11c2c5d2917020221fa9f02; osudb_lang=en; osudb_hdid=2217590e27c9544604c3575811dae79c; f_ds_info.sig=3CKJGA_m714rZP1Sd6rVfiPZrE-e85dCA1u-vNls85c; store_id.sig=jkiGcSihCNQHcIwuJeDT90gbTVd1S9PYIlgxgvQOJtM; merchant_id.sig=JdG9OdjYaypCshQ26ZRHpUNyoOuPv3UW7kt9Ba4A7VU; tsao=c79f003d-f1bd-4e08-b0a1-fe5c1198b515#ad50675722ed913e; tsui=03eb85fa-4176-4e1e-be02-463406306734#ca17543de9e9389f; r_b_in=1; _tracking_consent=%7B%22con%22%3A%7B%22GDPR%22%3A%22%22%7D%2C%22v%22%3A%221.0%22%2C%22lim%22%3A%5B%5D%2C%22reg%22%3A%22%22%7D; lang=en; addressLang=en; userSelectLocale=en; multi_lang_user_select_lang=en; googtrans=/auto/en; googtrans=/auto/en; t_cart=a7c6adb7110847ffa2de77b2fb527bcb; t_cart.sig=cba6fd8121e83dcb7d10dd69b3f22470; client_abandoned_order_seq=2405911021363771480096%2334eba825db785c3e; a_lang=zh-hans-cn; f_ds_info=vIkVV50Dkzv2zmQ98TKTgtk3b1cl9Ma+OJEOEXxG5CSnNgeHzfIbqulQrZ39Pa3cYEZKj5kOO+OxU5eNBLt8HQ==; country_code=CN; currency_code=CNY; currency_code.sig=BPe38eX0xU_6vUcN8hBCl2wo-vS61R5Qc9m2Ao_uba0; localization=CN; currency_code_userSetting=CNY; currency_code_userSetting.sig=N6p784I4PXJLkvcoM99TDGRh2ydbDLU4t_lmbHvHf5I; history_browse_products=16056550630997722466600637,16057240157463233196040637,16057240158879062455150637,16057240160056990786230637,16059098486729306369010637,16057240155040267655000637,16057240159038278226880637,16057240158655422162910637,16057240156118707106640637,16058402539549580405810637; history_browse_products.sig=x2KDNfYK1g3NYnh1n2Om1iGzYgbRq5P-FjKXFtjSbK0; n_sess={\"session_id\":\"ffcd0a76-0682-49c4-a34a-09f471b258c7\",\"created_at\":1682676917631,\"last_session_id\":\"f10e8c77-bf2c-4545-8686-bee7b14d25db\",\"session_create_type\":102}; JSESSIONID=39CE8C5958B674F900DD5AE71AD2FDD2; a_osudb_appid=1163336839; a_osudb_csrf=#C01#SID0000073BDzw324C8mN618NqJVRpzb/YMvRVgPsQlubkMBNZuEbBlixnUZ8XtBBktgftpc4UDsB+1Gy275gLe13c1DN0zN7EOWHfJzc3Ce/gsnlVIbjvlVNTDnLvNOzWWS0HPFV6kFegfGpDREsZ3t96ZTUDbKc4UV52eI0bE78LsvmC76uH; a_cstgc=1163336839_1_4211387386_1683107613931LplPEu; a_osudb_uid=4211387386; a_osudb_oar=#01#SID0000073BCyBuG3HqJMALwpB8h9HOUz8cunvWOvhs8fR+CYRdSz2okFTql4GlzmU7JoVebQ/Ib40S2HZak6fMp1yZbFF3qKVVA+V+sL49tw7ZpxkRZEofh34Q4qwg6gsYKV95TZrpOXzHBmUPvnDq9nN1RqR2A==; a_osudb_subappid=1; a_lce=1683712414010; a_dhch=niki; _BEAMER_FILTER_BY_URL_whBStDUm30458=false",
//}
//
//common.Post(url, resBody, nil, header)
//fmt.Println(id)
//fmt.Printf("%+v \n", name)
//fmt.Printf("%#v \n", url)
//fmt.Printf("%#v \n", method)
//fmt.Printf("%#v \n", resBody)

//}
