package actions

import (
	"errors"
	"fmt"

	"github.com/cpjudge/cpjudge_webserver/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"golang.org/x/crypto/bcrypt"
)

// SignupHandler : Handles signup.
func SignupHandler(c buffalo.Context) error {
	fmt.Println("In signup")
	fmt.Println("GET params were:", c.Request().URL.Query())
	firstName := c.Request().URL.Query().Get("first_name")
	lastName := c.Request().URL.Query().Get("last_name")
	email := c.Request().URL.Query().Get("email")
	username := c.Request().URL.Query().Get("username")
	password := c.Request().URL.Query().Get("password")
	bio := c.Request().URL.Query().Get("bio")
	if firstName != "" && lastName != "" && email != "" && username != "" && password != "" {
		hashedPassword, err := encrypt(password)
		if err != nil {
			return c.Render(500, r.JSON(map[string]string{"message": "Encryption error"}))
		}
		err = insertUser(c, firstName, lastName, username, email, hashedPassword, bio)
		if err != nil {
			return c.Render(400, r.JSON(map[string]string{"message": "Username already exists"}))
		}
		return c.Render(200, r.JSON(map[string]string{"message": "Success"}))
	}
	return c.Render(400, r.JSON(map[string]string{"message": "Bad request"}))
}

func encrypt(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hashedPassword, nil
}

func insertUser(c buffalo.Context, firstName string, lastName string, username string, email string, password []byte, bio string) error {
	user := &models.User{
		FirstName: firstName,
		LastName:  lastName,
		Username:  username,
		Email:     email,
		Password:  password,
		Bio:       bio,
		Rating:    0,
	}
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("Transaction error")
	}
	verrs, err := tx.ValidateAndCreate(user)
	if err != nil {
		fmt.Println("test", err.Error())
		return err
	}
	fmt.Println(verrs)
	return nil
}
