package users

import (
	infra "github.com/fernandojec/assignment-2/infra/fiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/jmoiron/sqlx"
)

func RouterWebInit(c fiber.Router, dbx *sqlx.DB, sessionStore *session.Store) {
	mid := infra.NewMiddleware(*sessionStore)
	repo := NewRepo(dbx)
	svc := NewService(repo)
	handler := NewHandlerWeb(svc, sessionStore)

	// authApi := c.Group("/web")
	c.Get("/sign-up", handler.SignUp)
	c.Post("/sign-up", handler.SignUpPost)
	c.Get("/verify-email/:token", handler.ActivateAuth)
	c.Post("/send-new-activation-link", handler.SendNewActivationLink)

	c.Get("/sign-in", handler.SignIn)
	c.Post("/sign-in", handler.SignInPost)
	c.Get("/", mid.AuthorizationWeb(), handler.Home)
	c.Get("/logout", mid.AuthorizationWeb(), handler.Logout)
}
