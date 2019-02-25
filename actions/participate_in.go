package actions

import (
	"fmt"

	"github.com/cpjudge/cpjudge_webserver/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gofrs/uuid"
)

// ParticipateInHandler default implementation.
func ParticipateInHandler(c buffalo.Context) error {
	userID := c.Request().URL.Query().Get("user_id")
	contestID := c.Request().URL.Query().Get("contest_id")
	if userID != "" && contestID != "" {
		err := insertParticiapteIn(c, userID, contestID)
		if err != nil {
			return c.Render(400, r.JSON(map[string]string{"message": err.Error()}))
		}
		return c.Render(200, r.JSON(map[string]string{"message": "Success"}))
	}
	return c.Render(400, r.JSON(map[string]string{"message": "Bad request"}))
}

// GetParticipatesInHandler : get all participate_ins
func GetParticipatesInHandler(c buffalo.Context) error {
	userID := c.Param("user_id")
	contests, err := getContestsWithUserID(userID)
	if err != nil {
		return c.Render(403, r.JSON(map[string]interface{}{
			"message": err.Error(),
		}))
	}
	return c.Render(200, r.JSON(contests))
}

func getContestsWithUserID(userID string) ([]models.Contest, error) {
	participateIns := &[]models.ParticipateIn{}
	contests := []models.Contest{}
	userUUID, err := uuid.FromString(userID)
	if err != nil {
		return contests, err
	}
	err = models.DB.Where("user_id=?", userUUID).All(participateIns)
	for _, v := range *participateIns {
		contest := &models.Contest{}
		models.DB.Find(contest, v.ContestID)
		contests = append(contests, *contest)
	}
	if err != nil {
		return contests, err
	}
	return contests, nil
}

func insertParticiapteIn(c buffalo.Context, userID string, contestID string) error {
	userUUID, err := uuid.FromString(userID)
	if err != nil {
		return err
	}
	contestUUID, err := uuid.FromString(contestID)
	if err != nil {
		return err
	}
	participateIn := &models.ParticipateIn{
		UserID:    userUUID,
		ContestID: contestUUID,
	}
	_, err = models.DB.ValidateAndCreate(participateIn)
	if err != nil {
		fmt.Println("insert participate_in", err.Error())
		return err
	}
	return nil
}
