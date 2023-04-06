package main

import (
	"fmt"
	"github.com/lambda-platform/ebarimt/posapi"
)

// This Main used
func main() {
	api, err := posapi.NewPosAPI("/home/khankhulgen/web/ebarimt/sofiles/mmk.so")
	if err != nil {
		panic(err)
	}

	info, _ := api.GetInformation()

	fmt.Println(info.RegisterNo)

	/*input := posapi.PutInput{}
	json.Unmarshal([]byte(`{
		"amount":"1100.00",
			"bankTransactions":[],
	"billIdSuffix":"123458",
	"billType":"1",
	"branchNo":"001",
	"cashAmount":"0.00",
	"cityTax":"0.00",
	"customerNo":"",
	"districtCode":"24",
	"invoiceId":"",
	"nonCashAmount":"1100.00",
	"posNo":"1234",
	"reportMonth":"",
	"returnBillId":"",
	"stocks":[
	{
	"barCode":"1234567890",
	"cityTax":"0.00",
	"code":"1234568",
	"measureUnit":"01",
	"name":"Хоол",
	"qty":"1.00",
	"totalAmount":"1100.00",
	"unitPrice":"1100.00",
	"vat":"100.00"
	}
	],
	"taxType":"",
	"vat":"100.00"
	}`), &input)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	out, err := api.Put(input)

	if err != nil {
		panic(err)
	}

	fmt.Println(out.QRData)
	fmt.Println(out.ErrorCode)
	fmt.Println(out.Message)*/

}
