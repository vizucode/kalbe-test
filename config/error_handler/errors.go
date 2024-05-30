package errorhandler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// ErrorHandler handles errors and returns an appropriate response.
func ErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	// Check if the error is a validation error
	if validationErr, ok := err.(validator.ValidationErrors); ok {
		// Create a map to hold the validation errors
		errsMap := make(map[string]interface{})
		// Populate the map with validation errors
		for _, val := range validationErr {
			switch strings.ToLower(val.Tag()) {
			case "required":
				errsMap[strings.ToLower(val.Field())] = "required field " + val.Field()
			}
		}

		// Return a response with bad request status and the validation error map
		return ctx.Status(http.StatusBadRequest).JSON(map[string]interface{}{
			"data":    nil,
			"message": "validation error",
			"errors":  errsMap,
			"success": false,
		})
	}

	// Check if the error is of type fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	// If there is an error, return a response with the error message and status code
	if err != nil {
		return ctx.Status(code).JSON(map[string]interface{}{
			"data":    nil,
			"message": err.Error(),
			"success": false,
		})
	}

	// If there is no error, return nil
	return nil
}
