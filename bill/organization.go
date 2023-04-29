package bill

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetOrganizationInfo(regID string) (OrganizationInfo, error) {
	var response OrganizationInfo
	url := "http://info.ebarimt.mn/rest/merchant/info?regno=" + regID
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return response, err
	}
	res, err := client.Do(req)
	if err != nil {
		return response, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

type OrganizationInfo struct {
	VatpayerRegisteredDate string      `json:"vatpayerRegisteredDate"`
	LastReceiptDate        interface{} `json:"lastReceiptDate"`
	ReceiptFound           bool        `json:"receiptFound"`
	Name                   string      `json:"name"`
	FreeProject            bool        `json:"freeProject"`
	Citypayer              bool        `json:"citypayer"`
	Vatpayer               bool        `json:"vatpayer"`
	Found                  bool        `json:"found"`
}
