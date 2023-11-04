package handler

import (
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)

func (h oauthHandler) Authorize(c *fiber.Ctx) error {
	clientID := os.Getenv("GITHUB_CLIENT_ID")
	if clientID == "" {
		return c.Status(http.StatusInternalServerError).SendString("GITHUB_CLIENT_ID is not set")
	}

	authorizationURL := h.OAuthService.GetAuthorizationURL(clientID)
	return c.Status(http.StatusTemporaryRedirect).Redirect(authorizationURL)
}
