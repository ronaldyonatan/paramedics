package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterWebInit(c fiber.Router, dbx *sqlx.DB) {
	repo := NewRepo(dbx)
	svc := NewService(repo)
	handler := NewHandlerWeb(svc)

	// authApi := c.Group("/web")
	c.Get("/sign-up", handler.SignUp)
	// authApi.Get("/verify-email/:token", handler.ActivateAuth)
	c.Get("/sign-in", handler.SignIn)
	c.Get("/", handler.Home)
}
