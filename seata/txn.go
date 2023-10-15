package seata

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

type TxnResponse struct {
	BaseResponse
	Data string
}

func BeginTxn(timeout int) {
	url := HTTPProtocol + GetAuth().GetAddress() + TryBeginTxnURL
	url = url + "?timeout=" + strconv.Itoa(timeout)
	token, err := GetAuth().GetToken()
	if err != nil {
		fmt.Println("Please login again!")
		os.Exit(0)
	}
	request, _ := http.NewRequest("POST", url, nil)
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

	if response.Code != CodeOK {
		fmt.Println(response.Message)
	} else {
		fmt.Printf("Try an example txn successfully, xid=%s\n", response.Data)
	}
}

func CommitTxn(xid string) {
	url := HTTPProtocol + GetAuth().GetAddress() + TryCommitTxnURL
	url = url + "?xid=" + xid
	token, err := GetAuth().GetToken()
	if err != nil {
		fmt.Println("Please login again!")
		os.Exit(0)
	}
	request, _ := http.NewRequest("POST", url, nil)
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

	if response.Code != CodeOK {
		fmt.Println(response.Message)
	} else {
		fmt.Printf("Commit txn successfully, xid=%s\n", response.Data)
	}
}

func RollbackTxn(xid string) {
	url := HTTPProtocol + GetAuth().GetAddress() + TryRollBackTxnURL
	url = url + "?xid=" + xid
	token, err := GetAuth().GetToken()
	if err != nil {
		fmt.Println("Please login again!")
		os.Exit(0)
	}
	request, _ := http.NewRequest("POST", url, nil)
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

	if response.Code != CodeOK {
		fmt.Println(response.Message)
	} else {
		fmt.Printf("Rollback txn successfully, xid=%s\n", response.Data)
	}
}
