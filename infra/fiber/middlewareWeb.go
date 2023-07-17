package infra

import (
	"github.com/fernandojec/assignment-2/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type Middleware struct {
	store session.Store
}

func NewMiddleware(store session.Store) Middleware {
	return Middleware{
		store: store,
	}
}

func (m Middleware) AuthorizationWeb() fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := m.store.Get(c)
		if err != nil {
			return c.Redirect("/web/sign-in", fiber.StatusTemporaryRedirect)
		}
		// if len(sess.Keys()) == 0 {
		// 	return c.Redirect("/web/sign-in", fiber.StatusTemporaryRedirect)
		// }
		dataSess := sess.Get(config.AppConfig.Session.AuthSessionId)
		if dataSess == nil {
			return c.Redirect("/web/sign-in", fiber.StatusTemporaryRedirect)
		}
		return c.Next()
	}
}
