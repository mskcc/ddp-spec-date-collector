package ddp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	timeout = 30 * time.Second
	login   = "https://ddp.mskcc.org/api/v1/authenticate/token"
)

type Authority struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Token   string `json:"auth_token"`
}

func GetToken(userid, pw string) (string, error) {
	body := fmt.Sprintf("{\"username\":\"%s\",\n \"password\":\"%s\"}\n", userid, pw)
	bodyReader := bytes.NewReader([]byte(body))

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, login, bodyReader)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP return code != 200: %d\n", resp.StatusCode)
	}

	rBody, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var auth Authority
	if err := json.Unmarshal(rBody, &auth); err != nil {
		return "", err
	}

	return auth.Token, nil
}
