package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type emailResponse struct {
	Email      string `json:"email"`
	Verified   bool   `json:"verified"`
	Primary    bool   `json:"primary"`
	Visibility string `json:"visibility"`
}

func (s oauthService) GetEmails(accessToken string) ([]string, error) {
	request, err := http.NewRequest(
		"GET",
		"https://api.github.com/user/emails",
		nil,
	)
	if err != nil {
		return nil, err
	}
	authorizationHeaderValue := fmt.Sprintf("token %s", accessToken)
	request.Header.Set("Authorization", authorizationHeaderValue)

	response, err := s.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var emailArray []emailResponse
	err = json.Unmarshal(body, &emailArray)
	if err != nil {
		return nil, err
	}

	emails := make([]string, len(emailArray))
	for i, email := range emailArray {
		emails[i] = email.Email
	}

	return emails, nil
}
