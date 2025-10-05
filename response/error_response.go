package response

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type CustomError struct {
	StatusCode int ``
	Message    string
	Details    any
}

func (c CustomError) Error() string {
	return "Something went wrong"
}

func Error(c *fiber.Ctx, statusCode int, message string, details interface{}) error {
	var errRes error
	if details != nil {
		errRes = c.Status(statusCode).JSON(errorDetails{
			Errors: details,
		})
	} else {
		errRes = c.Status(statusCode).JSON(common{
			Message: message,
		})
	}

	if errRes != nil {
		fmt.Printf("Failed to send error response : %+v", errRes)
	}

	return errRes
}

type errorDetails struct {
	Errors interface{} `json:"errors"`
}

type common struct {
	Message string `json:"message"`
}
