package actions

import (
	"errors"
	"fmt"

	"github.com/cpjudge/cpjudge_webserver/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gofrs/uuid"
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

// GetContestHandler : get contest from contest ID
func GetContestHandler(c buffalo.Context) error {
	contestID := c.Param("contest_id")
	contest, err := getContest(contestID)
	if err != nil {
		return c.Render(500, r.JSON(map[string]interface{}{
			"message": err.Error(),
		}))
	}
	return c.Render(200, r.JSON(contest))
}

// GetContestsHandler : get all contests
func GetContestsHandler(c buffalo.Context) error {
	contests, err := getContests()
	if err != nil {
		return c.Render(403, r.JSON(map[string]interface{}{
			"message": err.Error(),
		}))
	}
	return c.Render(200, r.JSON(contests))
}

func getContests() ([]models.Contest, error) {
	contests := &[]models.Contest{}
	err := models.DB.All(contests)
	if err != nil {
		fmt.Println("getQuestions error", err)
		return *contests, errors.New("Contests doesn't exist")
	}
	return (*contests), nil
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

func getContest(contestID string) (models.Contest, error) {
	contestUUID, err := uuid.FromString(contestID)
	contest := models.Contest{}
	if err != nil {
		return contest, err
	}
	err = models.DB.Find(&contest, contestUUID)
	if err != nil {
		return contest, err
	}
	return contest, nil
}
