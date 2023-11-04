package domain

import "github.com/gofiber/fiber/v2"

type YALSHandler interface {
	RedirectFromAlias(c *fiber.Ctx)
	PutHandler(c *fiber.Ctx)
}

type YALSService interface {
	GetURLFromAlias(alias string) (string, error)
	SetURLToAlias(alias, url string, force bool) error
}

type YALSRepository interface {
	FetchURLFromAlias(alias string) (string, error)
	IsAliasUsed(alias string) (bool, error)
	SetURLToAlias(alias, url string) error
}
