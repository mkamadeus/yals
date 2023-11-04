package handler

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (h oauthHandler) CheckUser(c *fiber.Ctx) error {
	bearer := c.Get("Authorization")
	parsedBearer := strings.Split(bearer, " ")
	if bearer == "" || len(parsedBearer) != 2 {
		return c.Status(http.StatusUnauthorized).SendString("Unauthorized")
	}
	accessToken := parsedBearer[1]
	fmt.Println(accessToken)

	userEmails, err := h.OAuthService.GetEmails(accessToken)
	fmt.Println(userEmails)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	whitelistedEmails := strings.Split(os.Getenv("WHITELIST"), ",")
	fmt.Println(whitelistedEmails)

	for _, validEmail := range whitelistedEmails {
		for _, userEmail := range userEmails {
			if userEmail == validEmail {
				return c.Next()
			}
		}
	}
	return c.Status(http.StatusUnauthorized).SendString("Unauthorized")
}
