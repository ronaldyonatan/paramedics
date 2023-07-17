package users

import (
	"errors"

	"github.com/fernandojec/assignment-2/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type serviceWeb interface {
	createAuthServiceWeb
	activateAuthServiceWeb
	signInServiceWeb
}

type createAuthServiceWeb interface {
	CreateAuth(req authCreateRequest, baseVerifyEmail string) (err error)
}
type activateAuthServiceWeb interface {
	ActivateAuth(token string) (err error)
	SendNewActivationLink(token string, baseVerifyEmail string) (err error)
}
type signInServiceWeb interface {
	SignInAuth(req authSignInRequest) (data authSignInResponse, err error)
}
type handlerWeb struct {
	svcAuth      serviceWeb
	sessionStore *session.Store
}

func NewHandlerWeb(svcauth serviceWeb, sessionStore *session.Store) handlerWeb {
	return handlerWeb{
		svcAuth:      svcauth,
		sessionStore: sessionStore,
	}
}

func (h *handlerWeb) Home(c *fiber.Ctx) error {
	// session, _ := h.sessionStore.Get(c.Context().Request, config.AppConfig.Session.AuthSessionId)
	sess, _ := h.sessionStore.Get(c)
	dataSess := sess.Get(config.AppConfig.Session.AuthSessionId)
	authData := dataSess.(authSignInResponse)
	// authData := authSignInResponse{
	// 	Email: dataSess["Email"],
	// 	Username: dataSess["Email"],
	// 	FirstName: dataSess["Email"],
	// 	LastName: dataSess["Email"],
	// }
	return c.Render("index", fiber.Map{
		"Title":    "Home",
		"AuthData": authData,
	})
}

func (h *handlerWeb) SignUp(c *fiber.Ctx) error {
	return c.Render("sign-up", fiber.Map{
		"Title": "Sign Up",
	})
}

func (h *handlerWeb) SignUpPost(c *fiber.Ctx) error {
	req := new(authCreateRequest)
	req = &authCreateRequest{
		Email:     c.FormValue("Email"),
		Password:  c.FormValue("Password"),
		Username:  c.FormValue("Username"),
		FirstName: c.FormValue("FirstName"),
		LastName:  c.FormValue("LastName"),
	}
	// if err := c.FormValue(req); err != nil {
	// 	return err
	// }

	if err := ValidateStruct(req); err != nil {
		errorMessage := make(map[string]string)
		for _, v := range err {
			errorMessage[v.FailedField] = v.Tag
		}
		return c.Render("sign-up", fiber.Map{
			"Title":        "Sign Up",
			"ErrorType":    "Form",
			"ErrorMessage": errorMessage,
		})
	}

	err := h.svcAuth.CreateAuth(
		*req,
		config.AppConfig.App.BaseUrl+config.AppConfig.App.BaseWebPort+"/web/verify-email/",
	)

	if err != nil {
		return c.Render("sign-up", fiber.Map{
			"Title":        "Sign Up",
			"ErrorType":    "Alert",
			"ErrorMessage": err.Error(),
		})
	}
	sess, _ := h.sessionStore.Get(c)
	sess.Set("success_message_login", "Your Account has been created. Check your email to activate your account.")
	sess.Save()

	return c.Redirect("/web/sign-in")

}
func (h *handlerWeb) SendNewActivationLink(c *fiber.Ctx) error {
	token := c.FormValue("code")
	if token == "" {
		return c.Render("verify-email", map[string]interface{}{
			"Success":      false,
			"ErrorMessage": errors.New("token is required"),
			"Code":         token,
		})
	}

	err := h.svcAuth.SendNewActivationLink(
		token,
		config.AppConfig.App.BaseUrl+config.AppConfig.App.BaseWebPort+"/web/verify-email/",
	)

	if err != nil {
		return c.Render("verify-email", map[string]interface{}{
			"Success":      false,
			"ErrorMessage": err.Error(),
			"Code":         token,
		})
	}

	return c.Render("verify-email", map[string]interface{}{
		"Success":        true,
		"SuccessMessage": "Your activation link has been sended. Please check you email.",
		"ErrorMessage":   "",
		"Code":           token,
	})
}
func (h *handlerWeb) ActivateAuth(c *fiber.Ctx) error {
	token := c.Params("token")
	if token == "" {
		return c.Render("verify-email", map[string]interface{}{
			"Success":      false,
			"ErrorMessage": errors.New("token is required"),
			"Code":         token,
		})
	}

	err := h.svcAuth.ActivateAuth(token)

	if err != nil {
		return c.Render("verify-email", map[string]interface{}{
			"Success":      false,
			"ErrorMessage": err.Error(),
			"Code":         token,
		})
	}

	return c.Render("verify-email", map[string]interface{}{
		"Success":        true,
		"SuccessMessage": "Your account has been activated.",
		"ErrorMessage":   "",
		"Code":           token,
	})
}
func (h *handlerWeb) SignIn(c *fiber.Ctx) error {
	sess, _ := h.sessionStore.Get(c)
	success_message := sess.Get("success_message_login")
	sess.Delete("success_message_login")
	return c.Render("sign-in", fiber.Map{
		"Title":          "Sign In",
		"SuccessMessage": success_message,
	})
}
func (h *handlerWeb) SignInPost(c *fiber.Ctx) error {
	req := authSignInRequest{
		Email:    c.FormValue("Email"),
		Password: c.FormValue("Password"),
	}
	err := ValidateStruct(req)
	if len(err) > 0 {
		errorMessage := make(map[string]string)
		for _, v := range err {
			errorMessage[v.FailedField] = v.Tag
		}
		return c.Render("sign-in", fiber.Map{
			"Title":        "Sign In",
			"ErrorType":    "Form",
			"ErrorMessage": errorMessage,
		})
	}
	resp, errSignIn := h.svcAuth.SignInAuth(req)

	if errSignIn != nil {
		return c.Render("sign-in", fiber.Map{
			"Title":        "Sign In",
			"ErrorType":    "Alert",
			"ErrorMessage": errSignIn.Error(),
		})
	}
	sess, _ := h.sessionStore.Get(c)
	h.sessionStore.RegisterType(resp)
	sess.Set(config.AppConfig.Session.AuthSessionId, resp)
	errSession := sess.Save()
	if errSession != nil {
		return c.Render("sign-in", fiber.Map{
			"Title":        "Sign In",
			"ErrorType":    "Alert",
			"ErrorMessage": errSession.Error(),
		})
	}

	return c.Redirect("/web")
	// return c.Render("sign-in", fiber.Map{
	// 	"Title": "Sign In",
	// 	"data":  resp,
	// })
}
func (h *handlerWeb) Logout(c *fiber.Ctx) error {

	sess, _ := h.sessionStore.Get(c)
	sess.Delete(config.AppConfig.Session.AuthSessionId)
	errSession := sess.Save()
	if errSession != nil {
		return c.Render("sign-in", fiber.Map{
			"Title":        "Sign In",
			"ErrorType":    "Alert",
			"ErrorMessage": errSession.Error(),
		})
	}

	return c.Redirect("/web/sign-in")
	// return c.Render("sign-in", fiber.Map{
	// 	"Title": "Sign In",
	// 	"data":  resp,
	// })
}
