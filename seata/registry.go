package seata

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"io"
	"net/http"
	"net/url"
	"os"
	"sort"
)

type RegistryResponse struct {
	BaseResponse
	Data map[string]string `json:"data"`
}

func GetRegistryConfigurations(key string) {
	v := url.Values{}
	if key != "" {
		v.Set("key", key)
	}

	url := HTTPProtocol + GetAuth().GetAddress() + RegistryURL
	url = fmt.Sprintf("%s?%s", url, v.Encode())

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

	var response RegistryResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
	}

	if response.Code != "200" {
		fmt.Println(response.Message)
	}

	t := table.NewWriter()
	header := table.Row{"key", "value"}
	t.AppendHeader(header)

	// Make output in order
	var keys []string
	for k, _ := range response.Data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		row := table.Row{key, response.Data[key]}
		t.AppendRow(row)
	}

	fmt.Println(t.Render())
	t.Style()
}

func SetRegistryConfiguration(key string, value string) {
	params := make(map[string]string)
	if key != "" {
		params[key] = value
	}
	b, _ := json.Marshal(params)

	url := HTTPProtocol + GetAuth().GetAddress() + RegistryURL
	token, err := GetAuth().GetToken()
	if err != nil {
		fmt.Println("Please login again!")
		os.Exit(0)
	}

	request, _ := http.NewRequest("POST", url, bytes.NewBuffer(b))
	request.Header.Set("authorization", token)
	request.Header.Set("Content-Type", "application/json")
	resp, err := (&http.Client{}).Do(request)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var response RegistryResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
	}

	if response.Code != "200" {
		fmt.Println(response.Message)
	}

	t := table.NewWriter()
	header := table.Row{"key", "value"}
	t.AppendHeader(header)

	// Make output in order
	var keys []string
	for k, _ := range response.Data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		row := table.Row{key, response.Data[key]}
		t.AppendRow(row)
	}
	fmt.Println(t.Render())
	t.Style()
}
