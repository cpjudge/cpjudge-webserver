package actions

import (
	"errors"
	"fmt"

	"github.com/cpjudge/cpjudge_webserver/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/binding"
	"github.com/gobuffalo/uuid"
)

// QuestionHandler default implementation.
func QuestionHandler(c buffalo.Context) error {
	question := c.Request().URL.Query().Get("question")
	editorial := c.Request().URL.Query().Get("editorial")
	contestID := c.Request().URL.Query().Get("contest_id")
	testCaseInputZip, err1 := c.File("test_cases_input")
	testCaseOutputZip, err2 := c.File("test_cases_output")
	if err1 != nil && err2 != nil {
		var err error
		if err1 != nil {
			err = err1
		} else {
			err = err2
		}
		return c.Render(400, r.JSON(map[string]interface{}{
			"message": err.Error(),
		}))
	}
	if question != "" && contestID != "" {
		question, err := insertQuestion(c, question, editorial, contestID,
			testCaseInputZip, testCaseOutputZip)
		if err != nil {
			return c.Render(400, r.JSON(map[string]interface{}{
				"message": err.Error(),
			}))
		}
		return c.Render(200, r.JSON(question))
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

// GetQuestionsHandler : get all questions
func GetQuestionsHandler(c buffalo.Context) error {
	questions, err := getQuestions()
	if err != nil {
		return c.Render(403, r.JSON(map[string]interface{}{
			"message": err.Error(),
		}))
	}
	return c.Render(200, r.JSON(questions))
}

func getQuestions() ([]models.Question, error) {
	questions := &[]models.Question{}
	err := models.DB.All(questions)

	if err != nil {
		fmt.Println("getQuestions error", err)
		return *questions, errors.New("Questions doesn't exist")
	}

	// fmt.Println((*questions))
	return (*questions), nil
}

func insertQuestion(c buffalo.Context, questionText string,
	editorial string, contestID string,
	testCaseInputZip binding.File, testCaseOutputZip binding.File) (*models.Question, error) {

	contestUUID, err := uuid.FromString(contestID)
	if err != nil {
		return nil, err
	}
	question := &models.Question{
		QuestionText:      questionText,
		Editorial:         editorial,
		ContestID:         contestUUID,
		TestCaseInputZip:  testCaseInputZip,
		TestCaseOutputZip: testCaseOutputZip,
	}
	verrs, err := models.DB.ValidateAndCreate(question)
	if err != nil {
		fmt.Println("Error inserting question", err.Error())
		return nil, err
	}
	fmt.Println(verrs)
	return question, nil
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
