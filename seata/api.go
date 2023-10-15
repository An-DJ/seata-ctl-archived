package seata

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/jedib0t/go-pretty/v6/table"
	"net/http"
	"sort"
)

func BuildPostRequestWithArrayData(urlStr string, data []string) (*http.Request, error) {
	token, err := GetAuth().GetToken()
	if err != nil {
		return nil, errors.New("please login")
	}

	body, _ := json.Marshal(data)

	request, _ := http.NewRequest("POST", urlStr, bytes.NewBuffer(body))
	request.Header.Set("authorization", token)
	request.Header.Set("Content-Type", "application/json")

	return request, nil
}

func BuildPostRequestWithMapData(urlStr string, data map[string]string) (*http.Request, error) {
	token, err := GetAuth().GetToken()
	if err != nil {
		return nil, errors.New("please login")
	}

	body, _ := json.Marshal(data)

	request, _ := http.NewRequest("POST", urlStr, bytes.NewBuffer(body))
	request.Header.Set("authorization", token)
	request.Header.Set("Content-Type", "application/json")

	return request, nil
}

func FormatKVResponse(kv map[string]string) string {
	t := table.NewWriter()
	header := table.Row{"key", "value"}
	t.AppendHeader(header)

	// Make output in order
	var keys []string
	for k, _ := range kv {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		row := table.Row{key, kv[key]}
		t.AppendRow(row)
	}
	return t.Render()
}

func FormatDiffResponse(kv map[string][]string) string {
	t := table.NewWriter()
	header := table.Row{"key", "from", "to"}
	t.AppendHeader(header)

	// Make output in order
	var keys []string
	for k, _ := range kv {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		row := table.Row{key, kv[key][0], kv[key][1]}
		t.AppendRow(row)
	}
	return t.Render()
}
