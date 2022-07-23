package controllers

import (
	"crud-service-food-recepies/database"
	"crud-service-food-recepies/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func Index(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"title":       "Сайт кулинарных рецептов.",
		"description": "На этом сайте вы найдете разнообразные рецепты.А также статьи о истории появления тех или иных блюд.",
		"data":        "Приятного прочтения.",
	})
}

func About(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"title":       "Сайт кулинарных рецептов.",
		"description": "Истории появления тех или иных блюд",
		"data":        "Приятного прочтения.",
	})
}

func GetAll(c *fiber.Ctx) error {
	recipes, err := database.GetRecipes()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": recipes,
	})
}

func GetRecipe(c *fiber.Ctx) error {
	id := c.Params("id")
	recipe, err := database.GetRecipe(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": recipe,
	})
}

func CreateRecipe(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var recipe models.Recipe
	err := c.BodyParser(&recipe)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "error parsing Json" + err.Error(),
		})
	}

	id := uuid.New()
	recipe.ID = fmt.Sprintf("%v", id)

	err = database.CreateRecipe(recipe)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not create recipe in db: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(recipe)
}

func UpdateRecipe(c *fiber.Ctx) error {
	c.Accepts("application/json")

	id := c.Params("id")

	recipe, err := database.GetRecipe(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could get recipe by id: " + err.Error(),
		})
	}

	if err := c.BodyParser(&recipe); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could parse recipe: " + err.Error(),
		})
	}

	if err := database.UpdateRecipe(recipe); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could update recipe by id: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(recipe)
}

func DeleteRecipe(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := database.DeleteRecipe(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could delete recipe: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "succesfulle delete"})
}
