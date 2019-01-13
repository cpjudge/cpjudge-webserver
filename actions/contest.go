package actions

import (
	"errors"
	"fmt"

	"github.com/cpjudge/cpjudge_webserver/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/gofrs/uuid"
)

// HostHandler default implementation.
func ContestHandler(c buffalo.Context) error {
	fmt.Println("In Contests")
	fmt.Println("GET params were:", c.Request().URL.Query())
	title := c.Request().URL.Query().Get("title")
	description := c.Request().URL.Query().Get("description")
	hostID := c.Request().URL.Query().Get("host_id")
	if title != "" && hostID != "" {
		err := insertContest(c, title, description, hostID)
		if err != nil {
			return c.Render(400, r.JSON(map[string]string{"message": err.Error()}))
		}
		return c.Render(200, r.JSON(map[string]string{"message": "Success"}))
	}
	return c.Render(400, r.JSON(map[string]string{"message": "Bad request"}))
}

func insertContest(c buffalo.Context, title string, description string, hostID string) error {
	hostUuid, err := uuid.FromString(hostID)
	if err != nil {
		return err
	}
	host := &models.Contest{
		Title:       title,
		Description: description,
		HostID:      hostUuid,
	}
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("Transaction error")
	}
	verrs, err := tx.ValidateAndCreate(host)
	if err != nil {
		fmt.Println("test", err.Error())
		return err
	}
	fmt.Println(verrs)
	return nil
}
