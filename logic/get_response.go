package logic

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func GetAccessToken() string {
	tokenBytes, err := ioutil.ReadFile(".env")
	if err != nil {
		log.Fatal("Error occurs when getting access token")
	}
	return strings.TrimSpace(string(tokenBytes))
}

func GetResponse(api string, accessToken string) (*http.Response, error) {
	req, err := http.NewRequest("GET", api, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("token %s", accessToken))
	return http.DefaultClient.Do(req)
}
