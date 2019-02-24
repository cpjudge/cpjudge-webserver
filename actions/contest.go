package actions

import (
	"fmt"

	"github.com/cpjudge/cpjudge_webserver/models"
	"github.com/gobuffalo/buffalo"
)

// ContestHandler default implementation.
func ContestHandler(c buffalo.Context) error {
	title := c.Request().URL.Query().Get("title")
	description := c.Request().URL.Query().Get("description")
	if title != "" {
		err := insertContest(c, title, description)
		if err != nil {
			return c.Render(400, r.JSON(map[string]interface{}{
				"message": err.Error(),
			}))
		}
		return c.Render(200, r.JSON(map[string]interface{}{
			"message": "Success",
		}))
	}
	return c.Render(400, r.JSON(map[string]interface{}{
		"message": "Bad request",
	}))
}

func insertContest(c buffalo.Context, title string, description string) error {

	contest := &models.Contest{
		Title:       title,
		Description: description,
	}
	verrs, err := models.DB.ValidateAndCreate(contest)
	if err != nil {
		fmt.Println("Error inserting contest", err.Error())
		return err
	}
	fmt.Println(verrs)
	return nil
}
