package actions

import (
	"fmt"

	"github.com/cpjudge/cpjudge_webserver/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/binding"
	"github.com/gobuffalo/uuid"
)

// SubmissionHandler : Handles submission
func SubmissionHandler(c buffalo.Context) error {
	userID := c.Request().URL.Query().Get("user_id")
	questionID := c.Request().URL.Query().Get("question_id")
	submissionFile, err := c.File("submission_file")
	if err != nil {
		return c.Render(400, r.JSON(map[string]interface{}{
			"message": err.Error(),
		}))
	}
	if userID != "" && questionID != "" {
		err := insertSubmission(userID, questionID, submissionFile)
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
		"message": "Required parameters missing",
	}))
}

func insertSubmission(userID string, questionID string,
	submissionFile binding.File) error {

	userUUID, err := uuid.FromString(userID)
	if err != nil {
		return err
	}
	questionUUID, err := uuid.FromString(questionID)
	if err != nil {
		return err
	}
	submission := &models.Submission{
		UserID:         userUUID,
		QuestionID:     questionUUID,
		SubmissionFile: submissionFile,
	}
	_, err = models.DB.ValidateAndCreate(submission)
	if err != nil {
		fmt.Println("insert submission error", err.Error())
		return err
	}
	return nil
}
