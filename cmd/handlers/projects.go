package handlers

import (
	"tasks/cmd/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

type ProjectTableHeaders struct {
	Name string
}

type ProjectTableBody struct {
	Name      string
	Sequence  float32
	Slug      string
	CreatedAt string
}

type Table struct {
	Headers []ProjectTableHeaders
	Data    []ProjectTableBody
}

func GetProjects() *Table {
	var projects []models.Project
	result := models.Db.Find(&projects)
	if result.Error != nil {
		// do nothing
	}
	header := []ProjectTableHeaders{
		{
			Name: "Name",
		},
		{
			Name: "Sequence",
		},
		{
			Name: "Slug",
		},
		{
			Name: "CreatedAt",
		},
	}
	var projectsForTable []ProjectTableBody
	for _, p := range projects {
		project := ProjectTableBody{
			Name:      p.Name,
			Sequence:  p.Sequence,
			Slug:      p.Slug,
			CreatedAt: p.CreatedAt.Format(time.DateTime),
		}
		projectsForTable = append(projectsForTable, project)
	}
	return &Table{
		Headers: header,
		Data:    projectsForTable,
	}
}

func Projects(c *fiber.Ctx) error {
	table := GetProjects()
	return c.Render("projects", fiber.Map{
		"Title": "Projects",
		"Table": table,
	})
}

type Project struct {
	Name string `form:"name"`
}

func AddProject(c *fiber.Ctx) error {
	p := new(Project)
	if err := c.BodyParser(p); err != nil {
		return c.Status(402).Render("errors", fiber.Map{
			"Error": err.Error(),
		})
	}
	var count int64
	count = 0
	models.Db.Find(&models.Project{}).Count(&count)
	project := models.Project{
		Name:     p.Name,
		Sequence: float32(count + 1),
		Slug:     p.Name,
	}
	result := models.Db.Save(&project)
	if result.Error != nil {
		return c.Status(402).Render("errors", fiber.Map{
			"Error": result.Error.Error(),
		})
	}
	table := GetProjects()
	return c.Render("message", fiber.Map{
		"Message": "Project added!",
		"Table":   table,
	})
}
