package main

import "fmt"

func main()  {

	var b = []byte{60,120,109,108,62,60,97,112,112,105,100,62,60,33,91,67,68,65,84,65,91,119,120,101,53,101,98,54,100,57,50,48,54,100,50,99,56,56,55,93,93,62,60,47,97,112,112,105,100,62,10,60,98,97,110,107,95,116,121,112,101,62,60,33,91,67,68,65,84,65,91,67,67,66,95,68,69,66,73,84,93,93,62,60,47,98,97,110,107,95,116,121,112,101,62,10,60,99,97,115,104,95,102,101,101,62,60,33,91,67,68,65,84,65,91,51,56,56,48,48,93,93,62,60,47,99,97,115,104,95,102,101,101,62,10,60,102,101,101,95,116,121,112,101,62,60,33,91,67,68,65,84,65,91,67,78,89,93,93,62,60,47,102,101,101,95,116,121,112,101,62,10,60,105,115,95,115,117,98,115,99,114,105,98,101,62,60,33,91,67,68,65,84,65,91,78,93,93,62,60,47,105,115,95,115,117,98,115,99,114,105,98,101,62,10,60,109,99,104,95,105,100,62,60,33,91,67,68,65,84,65,91,49,51,56,56,55,49,57,55,48,50,93,93,62,60,47,109,99,104,95,105,100,62,10,60,110,111,110,99,101,95,115,116,114,62,60,33,91,67,68,65,84,65,91,49,53,50,55,51,51,56,57,57,49,52,50,48,53,55,51,48,53,50,93,93,62,60,47,110,111,110,99,101,95,115,116,114,62,10,60,111,112,101,110,105,100,62,60,33,91,67,68,65,84,65,91,111,70,74,99,50,118,49,117,102,112,118,122,115,49,110,51,79,45,78,69,119,50,97,109,121,52,111,111,93,93,62,60,47,111,112,101,110,105,100,62,10,60,111,117,116,95,116,114,97,100,101,95,110,111,62,60,33,91,67,68,65,84,65,91,86,73,80,50,48,49,56,48,53,50,54,50,48,52,57,53,49,48,48,48,48,48,51,48,56,54,56,54,93,93,62,60,47,111,117,116,95,116,114,97,100,101,95,110,111,62,10,60,114,101,115,117,108,116,95,99,111,100,101,62,60,33,91,67,68,65,84,65,91,83,85,67,67,69,83,83,93,93,62,60,47,114,101,115,117,108,116,95,99,111,100,101,62,10,60,114,101,116,117,114,110,95,99,111,100,101,62,60,33,91,67,68,65,84,65,91,83,85,67,67,69,83,83,93,93,62,60,47,114,101,116,117,114,110,95,99,111,100,101,62,10,60,115,105,103,110,62,60,33,91,67,68,65,84,65,91,67,53,50,67,55,67,51,69,68,56,65,55,70,52,48,68,57,55,66,57,56,52,53,66,54,49,67,68,53,70,57,48,93,93,62,60,47,115,105,103,110,62,10,60,115,117,98,95,97,112,112,105,100,62,60,33,91,67,68,65,84,65,91,119,120,55,99,97,54,56,99,52,101,56,100,99,51,101,99,49,55,93,93,62,60,47,115,117,98,95,97,112,112,105,100,62,10,60,115,117,98,95,105,115,95,115,117,98,115,99,114,105,98,101,62,60,33,91,67,68,65,84,65,91,78,93,93,62,60,47,115,117,98,95,105,115,95,115,117,98,115,99,114,105,98,101,62,10,60,115,117,98,95,109,99,104,95,105,100,62,60,33,91,67,68,65,84,65,91,49,52,50,55,56,55,51,53,48,50,93,93,62,60,47,115,117,98,95,109,99,104,95,105,100,62,10,60,115,117,98,95,111,112,101,110,105,100,62,60,33,91,67,68,65,84,65,91,111,79,117,103,81,116,48,67,71,101,82,50,72,84,84,109,74,102,45,79,105,95,121,97,81,102,88,73,93,93,62,60,47,115,117,98,95,111,112,101,110,105,100,62,10,60,116,105,109,101,95,101,110,100,62,60,33,91,67,68,65,84,65,91,50,48,49,56,48,53,50,54,50,48,52,57,53,57,93,93,62,60,47,116,105,109,101,95,101,110,100,62,10,60,116,111,116,97,108,95,102,101,101,62,51,56,56,48,48,60,47,116,111,116,97,108,95,102,101,101,62,10,60,116,114,97,100,101,95,116,121,112,101,62,60,33,91,67,68,65,84,65,91,65,80,80,93,93,62,60,47,116,114,97,100,101,95,116,121,112,101,62,10,60,116,114,97,110,115,97,99,116,105,111,110,95,105,100,62,60,33,91,67,68,65,84,65,91,52,50,48,48,48,48,48,49,52,50,50,48,49,56,48,53,50,54,48,57,56,54,57,57,48,54,52,55,93,93,62,60,47,116,114,97,110,115,97,99,116,105,111,110,95,105,100,62,10,60,47,120,109,108,62}

	fmt.Println(string(b))
	/*

	<xml>
		<appid><![CDATA[wxe5eb6d9206d2c887]]></appid>
		<bank_type><![CDATA[CCB_DEBIT]]></bank_type>
		<cash_fee><![CDATA[38800]]></cash_fee>
		<fee_type><![CDATA[CNY]]></fee_type>
		<is_subscribe><![CDATA[N]]></is_subscribe>
		<mch_id><![CDATA[1388719702]]></mch_id>
		<nonce_str><![CDATA[1527338991420573052]]></nonce_str>
		<openid><![CDATA[oFJc2v1ufpvzs1n3O-NEw2amy4oo]]></openid>
		<out_trade_no><![CDATA[VIP2018052620495100000308686]]></out_trade_no>
		<result_code><![CDATA[SUCCESS]]></result_code>
		<return_code><![CDATA[SUCCESS]]></return_code>
		<sign><![CDATA[C52C7C3ED8A7F40D97B9845B61CD5F90]]></sign>
		<sub_appid><![CDATA[wx7ca68c4e8dc3ec17]]></sub_appid>
		<sub_is_subscribe><![CDATA[N]]></sub_is_subscribe>
		<sub_mch_id><![CDATA[1427873502]]></sub_mch_id>
		<sub_openid><![CDATA[oOugQt0CGeR2HTTmJf-Oi_yaQfXI]]></sub_openid>
		<time_end><![CDATA[20180526204959]]></time_end>
		<total_fee>38800</total_fee>
		<trade_type><![CDATA[APP]]></trade_type>
		<transaction_id><![CDATA[4200000142201805260986990647]]></transaction_id>
	</xml>

	*/

}