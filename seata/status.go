package seata

import (
	"encoding/json"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"io"
	"net/http"
	"os"
)

type NodeStatusResponse struct {
	BaseResponse
	Data []NodeStatus `json:"data"`
}

type NodeStatus struct {
	Address string `json:"address"`
	Status  string `json:"status"`
	Type    string `json:"type"`
}

func GetStatus() {
	url := HTTPProtocol + GetAuth().GetAddress() + HealthCheckURL
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

	var response NodeStatusResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
	}

	if response.Code != "200" {
		fmt.Println(response.Message)
	}

	t := table.NewWriter()
	header := table.Row{"type", "address", "status"}
	t.AppendHeader(header)
	for _, data := range response.Data {
		row := table.Row{data.Type, data.Address, data.Status}
		t.AppendRow(row)
	}
	fmt.Println(t.Render())
	t.Style()
}
