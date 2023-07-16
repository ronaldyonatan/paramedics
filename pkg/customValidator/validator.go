package customvalidator

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func HttpErrorHandler(c *fiber.Ctx, err error) error {
	report, ok := err.(*fiber.Error)
	// fmt.Println("error handling")
	if !ok {
		report = fiber.NewError(http.StatusInternalServerError, err.Error())
	}
	var reportMessage interface{} = report.Message
	// fmt.Println("error validation")
	if castedObject, ok := err.(validator.ValidationErrors); ok {
		var message []string
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				message = append(message, fmt.Sprintf("%s is required",
					err.Field()))
			case "email":
				message = append(message, fmt.Sprintf("%s is not valid email",
					err.Field()))
			case "gte":
				message = append(message, fmt.Sprintf("%s value must be greater than %s",
					err.Field(), err.Param()))
			case "lte":
				message = append(message, fmt.Sprintf("%s value must be lower than %s",
					err.Field(), err.Param()))
			}
		}
		// report.Message = message
		reportMessage = message
	}

	// c.Logger().Error(report)
	// c.JSON(report.Code, report)
	return c.Status(report.Code).JSON(map[string]interface{}{
		"status":  report.Code,
		"message": reportMessage,
	})
}
