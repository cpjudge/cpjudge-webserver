package actions

import (
	"fmt"
	"log"

	"github.com/cpjudge/cpjudge_webserver/models"
	evaluatorClient "github.com/cpjudge/cpjudge_webserver/proto/evaluator"
	spb "github.com/cpjudge/cpjudge_webserver/proto/submission"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/uuid"
)

type submissionJSON struct {
	UserID         string `json:"user_id"`
	QuestionID     string `json:"question_id"`
	SubmissionFile string `json:"submission_file"`
	Language       string `json:"language"`
}

// SubmissionHandler : Handles submission
func SubmissionHandler(c buffalo.Context) error {
	submissionJSON := &submissionJSON{}
	if err := c.Bind(submissionJSON); err != nil {
		return err
	}
	log.Println("JSON", submissionJSON)
	userID := submissionJSON.UserID
	questionID := submissionJSON.QuestionID
	submissionFile := submissionJSON.SubmissionFile
	language := submissionJSON.Language
	if userID != "" && questionID != "" &&
		submissionFile != "" && language != "" {
		log.Println(userID, questionID, submissionFile, language)
		s, err := insertSubmission(userID, questionID, 1, language, submissionFile)
		if err != nil {
			return c.Render(400, r.JSON(map[string]interface{}{
				"message": err.Error(),
			}))
		}
		// Path which will be mounted.
		testcasesPath := "/media/vaibhav/Coding/go/src/github.com/cpjudge/" +
			"cpjudge_webserver/questions/testcases/" +
			s.QuestionID.String() +
			"/input/"

		// Path which will be mounted
		submissionPath := "/media/vaibhav/Coding/go/src/github.com/cpjudge/" +
			"cpjudge_webserver/submissions/" +
			s.ID.String() + "/"

		codeStatus := evaluatorClient.EvaluateCode(&spb.Submission{
			Language:       s.Language,
			QuestionId:     s.QuestionID.String(),
			SubmissionId:   s.ID.String(),
			SubmissionPath: submissionPath,
			TestcasesPath:  testcasesPath,
			UserId:         s.UserID.String(),
		})
		store := int(evaluatorClient.EvaluationStatus_value[codeStatus.CodeStatus.String()])
		log.Println(store)
		err = updateCodeStatus(s, store)
		if err != nil {
			return c.Render(500, r.JSON(map[string]interface{}{
				"message": err.Error(),
			}))
		}
		question, err := getQuestion(s.QuestionID.String())
		if err != nil {
			return c.Render(500, r.JSON(map[string]interface{}{
				"message": err.Error(),
			}))
		}
		TriggerLeaderboards(question.ContestID.String())
		return c.Render(200, r.JSON(map[string]interface{}{
			"code_status": codeStatus.CodeStatus.String(),
		}))
	}
	return c.Render(400, r.JSON(map[string]interface{}{
		"message": "Required parameters missing",
	}))
}

func updateCodeStatus(submission *models.Submission, codeStatus int) error {
	submission.Status = codeStatus
	log.Println("updateCodeStatus", codeStatus)
	_, err := models.DB.ValidateAndSave(submission)
	if err != nil {
		log.Println("Error while saving", err.Error())
		return err
	}
	return nil
}

func insertSubmission(userID string, questionID string,
	status int, language string, submissionFile string) (*models.Submission, error) {

	userUUID, err := uuid.FromString(userID)
	if err != nil {
		return nil, err
	}
	questionUUID, err := uuid.FromString(questionID)
	if err != nil {
		return nil, err
	}
	submission := &models.Submission{
		UserID:         userUUID,
		QuestionID:     questionUUID,
		SubmissionFile: submissionFile,
		Language:       language,
		Status:         status,
	}
	_, err = models.DB.ValidateAndCreate(submission)
	if err != nil {
		fmt.Println("insert submission error", err.Error())
		return nil, err
	}
	return submission, nil
}
