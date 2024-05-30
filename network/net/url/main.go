package main

import (
	"fmt"
	"net/url"
	"sort"
	"strings"
)

// 测试 url.value{} 的用法
func main() {
	content := Content{
		OutTradeNo:  "1",
		Subject:     "2",
		TotalAmount: "3",
		BuyerOpenId: "4",
		ProductCode: "5",
		OpAppId:     "6",
	}
	v := url.Values{}
	v.Set("out_trade_no", content.OutTradeNo)
	v.Set("subject", content.Subject)
	v.Set("total_amount", "0.01")
	v.Set("buyer_open_id", content.BuyerOpenId)
	v.Set("product_code", "JSAPI_PAY")
	v.Set("op_app_id", content.OpAppId)
	fmt.Println(StringUrl(v))
}

type Content struct {
	OutTradeNo  string `json:"out_trade_no"`
	Subject     string `json:"subject"`
	TotalAmount string `json:"total_amount"`
	BuyerOpenId string `json:"buyer_open_id"`
	ProductCode string `json:"product_code"`
	OpAppId     string `json:"op_app_id"`
}

func StringUrl(param url.Values) string {
	//param就是url参数，privateKey 是私钥
	if param == nil {
		param = make(url.Values)
	}

	var pList = make([]string, 0)
	for key := range param {
		var value = strings.TrimSpace(param.Get(key)) //*去空格 签名必须有的一步
		if len(value) > 0 {
			pList = append(pList, key+"="+value)
		}
	}
	sort.Strings(pList) //*根据ASCII码排序 签名必须有的一步
	return strings.Join(pList, "&")
}
