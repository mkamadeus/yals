package service

import "net/http"

type oauthService struct {
	HTTPClient http.Client
}

func NewOAuthService() oauthService {
	return oauthService{
		HTTPClient: http.Client{},
	}
}
