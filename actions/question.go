package actions

import (
	"fmt"

	"github.com/cpjudge/cpjudge_webserver/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/uuid"
)

// QuestionHandler default implementation.
func QuestionHandler(c buffalo.Context) error {
	question := c.Request().URL.Query().Get("question")
	editorial := c.Request().URL.Query().Get("editorial")
	contestID := c.Request().URL.Query().Get("contest_id")
	if question != "" && contestID != "" {
		err := insertQuestion(c, question, editorial, contestID)
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

// GetQuestionHandler : get question from question ID
func GetQuestionHandler(c buffalo.Context) error {
	questionID := c.Param("question_id")
	question, err := getQuestion(questionID)
	if err != nil {
		return c.Render(500, r.JSON(map[string]interface{}{
			"message": err.Error(),
		}))
	}
	return c.Render(200, r.JSON(question))
}

func insertQuestion(c buffalo.Context, questionText string,
	editorial string, contestID string) error {

	contestUUID, err := uuid.FromString(contestID)
	if err != nil {
		return err
	}
	question := &models.Question{
		QuestionText: questionText,
		Editorial:    editorial,
		ContestID:    contestUUID,
	}
	verrs, err := models.DB.ValidateAndCreate(question)
	if err != nil {
		fmt.Println("Error inserting question", err.Error())
		return err
	}
	fmt.Println(verrs)
	return nil
}

func getQuestion(questionID string) (models.Question, error) {
	questionUUID, err := uuid.FromString(questionID)
	question := models.Question{}
	if err != nil {
		return question, err
	}
	err = models.DB.Find(&question, questionUUID)
	if err != nil {
		return question, err
	}
	return question, nil
}
