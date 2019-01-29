package actions

import (
	"errors"
	"fmt"

	"github.com/cpjudge/cpjudge_webserver/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/gofrs/uuid"
)

// TestCaseHandler default implementation.
func TestCaseHandler(c buffalo.Context) error {
	fmt.Println("In TestCase")
	fmt.Println("GET params were:", c.Request().URL.Query())
	testCase := c.Request().URL.Query().Get("test_case")
	questionID := c.Request().URL.Query().Get("question_id")
	if testCase != "" && questionID != "" {
		err := insertTestCase(c, testCase, questionID)
		if err != nil {
			return c.Render(400, r.JSON(map[string]string{"message": err.Error()}))
		}
		return c.Render(200, r.JSON(map[string]string{"message": "Success"}))
	}
	return c.Render(400, r.JSON(map[string]string{"message": "Bad request"}))
}

func insertTestCase(c buffalo.Context, testCaseLink string, questionID string) error {
	questionUUID, err := uuid.FromString(questionID)
	if err != nil {
		return err
	}
	testCase := &models.TestCase{
		TestCaseLink: testCaseLink,
		QuestionID:   questionUUID,
	}
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("Transaction error")
	}
	verrs, err := tx.ValidateAndCreate(testCase)
	if err != nil {
		fmt.Println("test", err.Error())
		return err
	}
	fmt.Println(verrs)
	return nil
}
