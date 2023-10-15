package seata

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

var auth Auth

type Auth struct {
	ServerIp   string
	ServerPort int
	Username   string
	Password   string
	token      string
}

type Response struct {
	Code    string
	Message string
	Data    string
	Success bool
}

func (auth *Auth) GetToken() (string, error) {
	if auth.token == "" {
		return auth.token, errors.New("login failed")
	}
	return auth.token, nil
}

func (auth *Auth) GetAddress() string {
	return auth.ServerIp + ":" + strconv.Itoa(auth.ServerPort)
}

func GetAuth() *Auth {
	return &auth
}

func (auth *Auth) Login() error {
	url := HTTPProtocol + auth.GetAddress() + LoginURL
	jsonStr := []byte(fmt.Sprintf(`{"username":"%s","password":"%s"}`, auth.Username, auth.Password))
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var jsonResp Response
	err = json.Unmarshal(body, &jsonResp)
	if err != nil {
		return err
	}
	auth.token = jsonResp.Data
	return nil
}
