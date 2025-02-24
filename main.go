package main

import (
	_ "fiber_swag_issue/docs"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
)

func ApiAuthMiddleware(user string, pass string) bool {
	if pass != "" {
		log.Printf("PASS IS NOT EMPTY = %q", pass)
		return false
	}
	if user == "" {
		return false
	}
	if user != "user" {
		return false
	}
	return true
}

type SimpleMessage struct {
	Message string `json:"message"`
} // @name Message

// @Summary AUTH Ping-Pong-Title
// @Description AUTH Ping-Pong
// @Tags Ping
// @ID auth-ping
// @Produce  json
// @Success 200 {object} SimpleMessage
// @Router /auth/ping [get]
// @Security BasicAuth
func AuthPingHandler(c *fiber.Ctx) error {
	return c.Status(200).JSON(SimpleMessage{
		Message: "pong",
	})
}

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.

// @host      localhost:8888
// @BasePath  /api
// @securityDefinitions.basic  BasicAuth
func main() {
	app := fiber.New(fiber.Config{})
	app.Use(cors.New())
	api := app.Group("/api")

	authGroup := api.Group("/auth")
	authGroup.Use(basicauth.New(basicauth.Config{Authorizer: ApiAuthMiddleware}))

	authGroup.Get("/ping", AuthPingHandler)

	api.Get("/swagger/*", swagger.HandlerDefault).Name("Swagger")
	if err := app.Listen(":8888"); err != nil {
		panic(err)
	}
}
