package actions

import (
	"errors"
	"fmt"

	"github.com/cpjudge/cpjudge_webserver/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"golang.org/x/crypto/bcrypt"
)

// SigninHandler : Handles signin.
func SigninHandler(c buffalo.Context) error {
	fmt.Println("GET params were:", c.Request().URL.Query())
	username := c.Request().URL.Query().Get("username")
	password := c.Request().URL.Query().Get("password")
	if username != "" && password != "" {
		fmt.Println("IN")
		err := verifyPassword(c, username, password)
		if err != nil {
			return c.Render(403, r.JSON(map[string]string{"message": err.Error()}))
		}
		err = GenerateTokenAndSetCookie(c, username)
		if err != nil {
			return c.Render(500, r.JSON(map[string]string{"message": "Internal Server error"}))
		}
		// Sucess.
		return c.Render(200, r.JSON(map[string]string{"message": "Check cookies"}))
	}
	return c.Render(400, r.JSON(map[string]string{"message": "Bad request"}))
}

func verifyPassword(c buffalo.Context, username string, password string) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("Transaction error")
	}
	candidateUser := &[]models.User{}
	fmt.Println("ad", tx.Where("username = ?", username).All(candidateUser))
	if len((*candidateUser)) == 0 {
		return errors.New("User Not Found")
	}
	user := (*candidateUser)[0]
	err := bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err != nil {
		return errors.New("Wrong username/password")
	}
	return nil
}
