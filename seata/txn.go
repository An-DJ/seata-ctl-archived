package seata

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type TxnResponse struct {
	BaseResponse
	Data string
}

func TryTxn() {
	url := HTTPProtocol + GetAuth().GetAddress() + TryTxnURL
	token, err := GetAuth().GetToken()
	if err != nil {
		fmt.Println("Please login again!")
		os.Exit(0)
	}
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("authorization", token)
	resp, err := (&http.Client{}).Do(request)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var response TxnResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
	}

	if response.Code != "200" {
		fmt.Println(response.Message)
	} else {
		fmt.Printf("Try an example txn successfully, xid=%s\n", response.Data)
	}
}
