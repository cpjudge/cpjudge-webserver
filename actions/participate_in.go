package actions

import (
	"errors"
	"fmt"

	"github.com/cpjudge/cpjudge_webserver/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/gofrs/uuid"
)

// ParticipateInHandler default implementation.
func ParticipateInHandler(c buffalo.Context) error {
	fmt.Println("In ParticipateIn")
	fmt.Println("GET params were:", c.Request().URL.Query())
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
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("Transaction error")
	}
	verrs, err := tx.ValidateAndCreate(participateIn)
	if err != nil {
		fmt.Println("test", err.Error())
		return err
	}
	fmt.Println(verrs)
	return nil
}
