package main

import (
	"fmt"
	"github.com/lambda-platform/ebarimt/posapi"
)

func main() {
	api, err := posapi.NewPosAPI("/home/khankhulgen/web/ebarimt/sofiles/mmk.so")
	if err != nil {
		panic(err)
	}

	api.SendData()
	api.CheckApi()
	api.SendData()
	api.SendData()

	out, err := api.Put(`{
   "amount":"1100",
   "bankTransactions":[
      
   ],
   "billIdSuffix":"123456",
   "billType":"1",
   "branchNo":"001",
   "cashAmount":"0",
   "cityTax":"0",
   "customerNo":"",
   "districtCode":"24",
   "invoiceId":"",
   "nonCashAmount":"1100",
   "posNo":"0001",
   "reportMonth":"2023-04",
   "returnBillId":"",
   "stocks":[
      {
         "barCode":"1234567890",
         "cityTax":"0",
         "code":"1234567",
         "measureUnit":"01",
         "name":"Хоол",
         "qty":"1",
         "totalAmount":"1100",
         "unitPrice":"1100",
         "vat":"100"
      }
   ],
   "taxType":"",
   "vat":"100"
}`)

	if err != nil {
		panic(err)
	}

	fmt.Println(out.QRData)
	fmt.Println(out.ErrorCode)
	fmt.Println(out.Message)

}
