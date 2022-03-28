package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	// Import Fiber Swagger
	swagger "github.com/arsmn/fiber-swagger/v2"

	// Side Effect import for auto-generated swagger documentation

	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/guionardo/auth-service/golang/data"
	_ "github.com/guionardo/auth-service/golang/docs"
	"github.com/guionardo/auth-service/golang/domain"
	"github.com/guionardo/auth-service/golang/dto"
)

func createServer() *fiber.App {
	repository, err := data.GetRepository()
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	repository.Show()
	// Create new Fiber application
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "/swagger/doc.json",
		DeepLinking: false,
	}))

	// POST /user -> upsert user data
	app.Post("/user", postUser)

	// DELETE /user -> delete user data
	app.Delete("/user/:user_id", deleteUser)
	// Create a route on the default path, only returning some string

	// POST /auth -> Authentication by credentials
	app.Post("/auth", postAuth)

	// GET /auth -> Get user data by token
	app.Get("/auth", getAuth)

	// GET /refresh -> Get refresh token
	app.Get("/refresh", getRefresh)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	return app
}

// postUser godoc
// @Summary Get an item
// @Description Get an item by its ID
// @ID get-item-by-int
// @Accept  json
// @Produce  json
// @Tags Item
// @Param id path int true "Item ID"
// @Success 200 {object} Item
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /api/item/{id} [get]
func postUser(c *fiber.Ctx) error {
	if err := checkFeedAPIKey(c); err != nil {
		return err
	}

	var body dto.PostUserDTO
	if err := c.BodyParser(&body); err != nil {
		return fiber.ErrBadRequest
	}
	repo, _ := data.GetRepository()
	if err := repo.SetUser(&domain.UserData{
		UserID:       body.UserId,
		PasswordHash: body.PasswordHash,
		Payload:      body.Payload,
	}); err != nil {
		return fiber.NewError(fiber.ErrBadGateway.Code, fmt.Sprintf("CANNOT SAVE USER - %v", err))
	}
	c.Status(fiber.StatusAccepted).SendString("ACCEPTED")
	return nil
}

func deleteUser(c *fiber.Ctx) error {
	if err := checkFeedAPIKey(c); err != nil {
		return err
	}
	userId := c.Params("user_id")
	repo, _ := data.GetRepository()
	if err := repo.DeleteUser(userId); err != nil {
		return fiber.NewError(fiber.ErrBadGateway.Code, fmt.Sprintf("CANNOT DELETE USER - %v", err))
	}
	c.Status(fiber.StatusAccepted).SendString("ACCEPTED")
	return nil
}

func getRefresh(c *fiber.Ctx) error {
	return nil
}

func postAuth(c *fiber.Ctx) error {
	return nil
}

func getAuth(c *fiber.Ctx) error {
	return nil
}

type Item struct {
	Id string
}

type HTTPError struct {
	Status  string
	Message string
}
