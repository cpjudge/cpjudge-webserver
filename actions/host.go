package actions

import (
	"errors"
	"fmt"

	"github.com/cpjudge/cpjudge_webserver/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
)

// HostHandler default implementation.
func HostHandler(c buffalo.Context) error {
	fmt.Println("In hosts")
	fmt.Println("GET params were:", c.Request().URL.Query())
	name := c.Request().URL.Query().Get("name")
	if name != "" {
		err := insertHost(c, name)
		if err != nil {
			return c.Render(400, r.JSON(map[string]string{"message": "Host already exists"}))
		}
		return c.Render(200, r.JSON(map[string]string{"message": "Success"}))
	}
	return c.Render(400, r.JSON(map[string]string{"message": "Bad request"}))
}

func insertHost(c buffalo.Context, name string) error {
	host := &models.Host{
		Name: name,
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
