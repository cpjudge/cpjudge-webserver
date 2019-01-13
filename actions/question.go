package actions

import (
	"errors"
	"fmt"

	"github.com/cpjudge/cpjudge_webserver/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/gofrs/uuid"
)

// QuestionHandler default implementation.
func QuestionHandler(c buffalo.Context) error {
	fmt.Println("In Question")
	fmt.Println("GET params were:", c.Request().URL.Query())
	question := c.Request().URL.Query().Get("question")
	editorial := c.Request().URL.Query().Get("editorial")
	hostID := c.Request().URL.Query().Get("host_id")
	contestID := c.Request().URL.Query().Get("contest_id")
	if question != "" && hostID != "" && contestID != "" {
		err := insertQuestion(c, question, editorial, hostID, contestID)
		if err != nil {
			return c.Render(400, r.JSON(map[string]string{"message": err.Error()}))
		}
		return c.Render(200, r.JSON(map[string]string{"message": "Success"}))
	}
	return c.Render(400, r.JSON(map[string]string{"message": "Bad request"}))
}

func insertQuestion(c buffalo.Context, questionText string, editorial string, hostID string, contestID string) error {
	hostUUID, err := uuid.FromString(hostID)
	if err != nil {
		return err
	}
	contestUUID, err := uuid.FromString(contestID)
	if err != nil {
		return err
	}
	question := &models.Question{
		QuestionText: questionText,
		Editorial:    editorial,
		HostID:       hostUUID,
		ContestID:    contestUUID,
	}
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("Transaction error")
	}
	verrs, err := tx.ValidateAndCreate(question)
	if err != nil {
		fmt.Println("test", err.Error())
		return err
	}
	fmt.Println(verrs)
	return nil
}
