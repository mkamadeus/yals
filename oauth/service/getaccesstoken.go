package service

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type callbackRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
	Accept       string `json:"accept"`
}

type callbackResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

func (s oauthService) GetAccessToken(clientID, clientSecret, code string) (string, error) {
	callbackBody := callbackRequest{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Code:         code,
		Accept:       "json",
	}
	body, err := json.Marshal(callbackBody)
	if err != nil {
		return "", err
	}
	reader := bytes.NewReader(body)
	request, err := http.NewRequest(
		"POST",
		"https://github.com/login/oauth/access_token",
		reader,
	)
	if err != nil {
		return "", err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := s.HTTPClient.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	var ghResponse callbackResponse
	json.Unmarshal(responseBody, &ghResponse)

	return ghResponse.AccessToken, nil
}
