package actions

import (
	"fmt"

	"github.com/cpjudge/cpjudge_webserver/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/uuid"
)

// SubmissionHandler : Handles submission
func SubmissionHandler(c buffalo.Context) error {
	userID := c.Request().URL.Query().Get("user_id")
	questionID := c.Request().URL.Query().Get("question_id")
	submissionFile := c.Request().URL.Query().Get("submission_file")
	if userID != "" && questionID != "" && submissionFile != "" {
		err := insertSubmission(userID, questionID, 0, submissionFile)
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
	status int, submissionFile string) error {

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
		Status:         status,
	}
	_, err = models.DB.ValidateAndCreate(submission)
	if err != nil {
		fmt.Println("insert submission error", err.Error())
		return err
	}
	return nil
}
