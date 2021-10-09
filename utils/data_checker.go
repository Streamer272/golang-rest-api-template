package utils

import (
	"github.com/gofiber/fiber/v2"
	"password-manager-backend/errors"
	"password-manager-backend/logger"
)

func CheckData(c *fiber.Ctx, requiredData ...string) (map[string]interface{}, error) {
	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		logger.LogError(err)

		c.Status(fiber.StatusUnprocessableEntity)
		err := c.JSON(errors.ErrorMessage{
			Error:   "UnprocessableEntity",
			Message: "Couldn't parse JSON",
		})
		if err != nil {
			return nil, err
		}

		return nil, fiber.ErrUnprocessableEntity
	}

	for _, required := range requiredData {
		if data[required] == nil {
			c.Status(fiber.StatusUnprocessableEntity)
			err := c.JSON(errors.ErrorMessage{
				Error:   "UnprocessableEntity",
				Message: "Missing " + required,
			})
			if err != nil {
				return nil, err
			}

			return nil, fiber.ErrUnprocessableEntity
		}
	}

	return data, nil
}
