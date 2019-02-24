package actions

import (
	"fmt"

	"github.com/cpjudge/cpjudge_webserver/models"

	"github.com/gobuffalo/buffalo"
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
			return c.Render(500, r.JSON(map[string]string{
				"message": "Encryption error",
			}))
		}
		err = insertUser(c, firstName, lastName, username, email, hashedPassword, bio)
		if err != nil {
			return c.Render(400, r.JSON(map[string]string{
				"message": "Username already exists",
			}))
		}
		// Success.
		user, err1 := getUser(username)
		if err1 != nil {
			fmt.Println("Error fetching user", err)
			return c.Render(500, r.JSON(map[string]interface{}{
				"message": "Internal Server error",
			}))
		}
		return c.Render(200, r.JSON(map[string]interface{}{
			"username": username,
			"rating":   user.Rating,
		}))
	}
	return c.Render(400, r.JSON(map[string]string{"message": "Bad request"}))
}

// GetUserInfoHandler : Return user info
func GetUserInfoHandler(c buffalo.Context) error {
	username := c.Param("username")
	user, err := getUser(username)
	if err != nil {
		fmt.Println("Error fetching user info", err)
		return c.Render(500, r.JSON(map[string]interface{}{
			"message": err.Error(),
		}))
	}
	return c.Render(200, r.JSON(user))
}
func encrypt(password string) ([]byte, error) {
	hashedPassword, err :=
		bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hashedPassword, nil
}

func insertUser(c buffalo.Context, firstName string, lastName string,
	username string, email string, password []byte, bio string) error {

	user := &models.User{
		FirstName: firstName,
		LastName:  lastName,
		Username:  username,
		Email:     email,
		Password:  password,
		Bio:       bio,
		Rating:    0,
	}
	_, err := models.DB.ValidateAndCreate(user)
	if err != nil {
		fmt.Println("Error inserting user", err.Error())
		return err
	}
	return nil
}
