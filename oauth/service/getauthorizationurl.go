package service

import "fmt"

func (o oauthService) GetAuthorizationURL(clientID string) string {
	scopes := "user:email"
	return fmt.Sprintf("https://github.com/login/oauth/authorize?scope=%s&client_id=%s", scopes, clientID)
}
