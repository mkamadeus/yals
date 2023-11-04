package domain

type OAuthUser struct {
	Email string
}

type OAuthService interface {
	GetAuthorizationURL(clientID string) string
	GetAccessToken(clientID, clientSecret, code string) (string, error)
	GetEmails(accessToken string) ([]string, error)
}
