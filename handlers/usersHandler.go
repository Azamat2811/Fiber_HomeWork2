package handlers

import (
	"fiber_hw_2/models"
	"fiber_hw_2/responses"
	"fiber_hw_2/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {

	users, err := services.GetAllUsers()
	if err != nil {
		return responses.Error(c, 500, err.Error())
	}

	limitStr := c.Query("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 || limit > len(users) {
		limit = len(users)
	}

	return responses.Success(c, users[:limit])
}

func GetUsersByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return responses.Error(c, 400, "invalid id")
	}

	track, err := services.GetUsersByID(id)
	if err != nil {
		return responses.Error(c, 404, err.Error())
	}

	return responses.Success(c, track)
}

func CreateUsers(c *fiber.Ctx) error {
	var user models.Users

	if err := c.BodyParser(&user); err != nil {
		return responses.Error(c, 400, "invalid JSON")
	}

	createdUser, err := services.CreateUsers(user)
	if err != nil {
		return responses.Error(c, 500, err.Error())
	}

	return responses.Success(c, createdUser)
}

func UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return responses.Error(c, 400, "invalid id")
	}

	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		return responses.Error(c, 400, "invalid JSON")
	}

	user, err := services.UpdateUser(id, data)
	if err != nil {
		return responses.Error(c, 404, err.Error())
	}

	return responses.Success(c, user)
}

func FullUpdate(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var updated models.Users

	if err := c.BodyParser(&updated); err != nil {
		return responses.Error(c, 400, "invalid JSON")
	}

	updatedUser, err := services.FullUpdateUsers(id, updated)
	if err != nil {
		return responses.Error(c, 404, err.Error())
	}

	return responses.Success(c, updatedUser)
}
func DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return responses.Error(c, 400, "invalid id")
	}

	err = services.DeleteUsers(id)
	if err != nil {
		return responses.Error(c, 404, err.Error())
	}

	return c.SendStatus(204)
}

func GetUsersPagination(c *fiber.Ctx) error {
    page := c.QueryInt("page", 1)
    limit := c.QueryInt("limit", 10)

    users, err := services.GetUsersPagination(page, limit)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }

    return c.JSON(users)
}


func GetUsersByName(c *fiber.Ctx) error {
   
    name := c.Query("name", "")

    users, err := services.GetUsersByName(name)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Не удалось получить список пользователей",
        })
    }

  
    return c.Status(fiber.StatusOK).JSON(users)
}

func GetUserStats(c *fiber.Ctx) error {
    stats, err := services.GetUserStats()
    if err != nil {
        return c.Status(500).SendString(err.Error())
    }

    return c.JSON(stats)
}
