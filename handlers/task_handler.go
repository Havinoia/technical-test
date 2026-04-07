package handlers

import (
	"strconv"
	"strings"

	"stability-test-task-api/models"
	"stability-test-task-api/store"

	"github.com/gofiber/fiber/v2"
)

func GetTasks(c *fiber.Ctx) error {
	tasks := store.GetAllTasks()
	return c.JSON(fiber.Map{
		"data": tasks,
	})
}

func GetTask(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid task id format",
		})
	}

	task, err := store.GetTaskByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "task not found",
		})
	}

	return c.JSON(fiber.Map{
		"data": task,
	})
}

func CreateTask(c *fiber.Ctx) error {
	var task models.Task

	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "failed to parse request body",
		})
	}

	// Simple validation: title should not be empty
	if strings.TrimSpace(task.Title) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "task title is required",
		})
	}

	id := store.AddTask(&task)
	task.ID = id

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "task created successfully",
		"data":    task,
	})
}

func DeleteTask(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid task id format",
		})
	}

	if err := store.DeleteTask(id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "task not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "task deleted successfully",
	})
}
