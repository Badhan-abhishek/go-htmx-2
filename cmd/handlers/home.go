package handlers

import (
	"net/http"
	"tasks/cmd/models"

	"github.com/gofiber/fiber/v2"
)

type TaskTableHeaders struct {
	Name string
}

type TaskTable struct {
	Headers []TaskTableHeaders
	Data    []models.Task
}

func Home(c *fiber.Ctx) error {
	var tasks []models.Task
	result := models.Db.Preload("Project").Find(&tasks)
	if result.Error != nil {
		// Do nothing
	}
	header := []TaskTableHeaders{
		{
			Name: "Title",
		},
		{
			Name: "Content",
		},
		{
			Name: "Start Time",
		},
		{
			Name: "End Time",
		},
		{
			Name: "Project",
		},
	}
	table := GetProjects()
	return c.Render("index", fiber.Map{
		"Title": "New page!",
		"Tasks": &TaskTable{
			Headers: header,
			Data:    tasks,
		},
		"Projects": table.Data,
	})
}

type Task struct {
	Project     string  `form:"project"`
	Title       string  `form:"title"`
	TaskDetails string  `form:"taskDetails"`
	StartTime   *string `form:"startTime"`
	EndTime     *string `form:"endTime"`
}

func AddTask(c *fiber.Ctx) error {
	t := new(Task)
	if err := c.BodyParser(t); err != nil {
		return c.Status(http.StatusBadRequest).Render("errors", fiber.Map{
			"Error": err.Error(),
		})
	}

	var project models.Project
	result := models.Db.Where("slug = ?", t.Project).First(&project)
	if result.Error != nil {
		return c.Status(http.StatusBadRequest).Render("errors", fiber.Map{
			"Error": result.Error.Error(),
		})
	}

	task := models.Task{
		Title:     t.Title,
		Content:   t.TaskDetails,
		ProjectID: project.ID,
	}
	if t.StartTime != nil {
		task.StartTime = t.StartTime
	}

	if t.EndTime != nil {
		task.EndTime = t.EndTime
	}

	result = models.Db.Save(&task)
	if result.Error != nil {
		return c.Status(http.StatusBadRequest).Render("errors", fiber.Map{
			"Error": result.Error.Error(),
		})
	}

	return c.Render("message", fiber.Map{
		"Message": "Task Added!",
		"Title":   "New page!",
	})
}
