package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterInit(c fiber.Router, dbx *sqlx.DB) {
	repo := NewRepo(dbx)
	svc := NewService(repo)
	handler := NewHandler(svc)

	authApi := c.Group("/auth")
	authApi.Post("/sign-up", handler.CreateAuth)
	authApi.Get("/verify-email/:token", handler.ActivateAuth)
	authApi.Post("/sign-in", handler.SignInAuth)
}
