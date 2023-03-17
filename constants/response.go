package constants

import (
	"github.com/gofiber/fiber/v2"
)

func ErrorResponse(ctx *fiber.Ctx, err error) error {
	error_map := map[string]interface{}{
		"status":  ServerErrorCode,
		"message": ServerErrorMessage + err.Error(),
	}

	return ctx.Status(ServerErrorCode).JSON(error_map)
}

func ErrorResponseWithMessage(ctx *fiber.Ctx, message string, err error) error {
	error_map := map[string]interface{}{
		"status":  ServerErrorCode,
		"message": message + err.Error(),
	}

	return ctx.Status(ServerErrorCode).JSON(error_map)
}

func SuccessResponse(ctx *fiber.Ctx) error {
	success_map := map[string]interface{}{
		"status":  SuccessCode,
		"message": SuccessMessage,
	}

	return ctx.Status(SuccessCode).JSON(success_map)
}

func SuccessResponseWithData(ctx *fiber.Ctx, data interface{}) error {
	success_map := map[string]interface{}{
		"status":  SuccessCode,
		"message": SuccessMessage,
		"data":    data,
	}

	return ctx.Status(SuccessCode).JSON(success_map)
}

func NotFoundResponse(ctx *fiber.Ctx) error {
	error_map := map[string]interface{}{
		"status":  NotFoundCode,
		"message": NotFoundMessage,
	}

	return ctx.Status(NotFoundCode).JSON(error_map)
}

func ValidationErrorResponse(ctx *fiber.Ctx, err []*ErrorStructValidation) error {
	error_map := map[string]interface{}{
		"status":  ValidationErrorCode,
		"message": err,
	}

	return ctx.Status(ValidationErrorCode).JSON(error_map)
}
