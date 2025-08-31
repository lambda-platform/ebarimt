package posapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type PosAPI struct {
	baseUrl string
}

func NewPosAPI(baseUrl string) (*PosAPI, error) {
	return &PosAPI{baseUrl}, nil
}

func (api *PosAPI) GetInformation() (InformationResponse, error) {

	url := api.baseUrl + "/rest/info"

	resp, err := http.Get(url)
	if err != nil {
		log.Println("Failed to call API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("Failed to call API, status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Failed to read response body: %v", err)
	}

	var response InformationResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println("Failed to unmarshal JSON: %v", err)
	}

	return response, nil
}

func (api *PosAPI) CreateEBarimt(receiptData ReceiptData) (EBarimtResponse, error) {
	url := api.baseUrl + "/rest/receipt"
	var response EBarimtResponse
	jsonData, err := json.Marshal(receiptData)
	if err != nil {
		return response, fmt.Errorf("failed to marshal receipt data: %w", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return response, fmt.Errorf("failed to call API: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, fmt.Errorf("failed to read response body: %w", err)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, fmt.Errorf(string(body))
	}

	return response, nil
}
func (api *PosAPI) ReturnBill(input BillInput) error {
	url := api.baseUrl + "/rest/receipt"

	jsonData, err := json.Marshal(input)
	if err != nil {
		return fmt.Errorf("failed to marshal receipt data: %w", err)
	}

	fmt.Println(url)
	fmt.Println(string(jsonData))

	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		os.Exit(1)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	return nil

}

func (api *PosAPI) SendData() (map[string]interface{}, error) {
	url := api.baseUrl + "/rest/sendData"

	resp, err := http.Get(url)
	if err != nil {
		log.Println("Failed to call API: %v", err)
	}
	defer resp.Body.Close()

	return map[string]interface{}{}, nil
}
