package handler

import "github.com/mkamadeus/yals/domain"

type oauthHandler struct {
	OAuthService domain.OAuthService
}

func NewOAuthHandler(s domain.OAuthService) oauthHandler {
	return oauthHandler{s}
}
