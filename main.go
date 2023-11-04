package main

import (
	"fmt"
	"log"
	"os"

	_ "embed"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	oauthHandler "github.com/mkamadeus/yals/oauth/handler"
	oauthService "github.com/mkamadeus/yals/oauth/service"
	"github.com/mkamadeus/yals/yals/db"
	yalsHandler "github.com/mkamadeus/yals/yals/handler"
	yalsService "github.com/mkamadeus/yals/yals/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	bolt "go.etcd.io/bbolt"
)

// link shortener
func main() {
	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// create fiber app
	app := fiber.New()
	defer app.Shutdown()
	app.Use(limiter.New(limiter.ConfigDefault))

	// init db connection
	boltDB, err := bolt.Open("yals.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer boltDB.Close()

	// create bucket if not exists
	err = boltDB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("urls"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	// init services and handlers
	yalsDB := db.NewBoltDatabaseConnection(boltDB, "urls")
	yalsService := yalsService.NewYALSService(yalsDB)
	yalsHandler := yalsHandler.NewYALSHandler(yalsService)

	oauthService := oauthService.NewOAuthService()
	oauthHandler := oauthHandler.NewOAuthHandler(oauthService)

	// define routes
	app.Get("/:alias", yalsHandler.RedirectFromAlias)
	app.Get("/api/yals/:alias", yalsHandler.RedirectFromAlias)
	app.Put("/api/yals/:alias", oauthHandler.CheckUser, yalsHandler.AssignAliasToURL)

	app.Get("/api/oauth/authorize", oauthHandler.Authorize)
	app.Get("/api/oauth/callback", oauthHandler.Callback)

	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
