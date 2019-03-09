package actions

import (
	"errors"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gobuffalo/buffalo"
)

// CpJudgeToken : Authentication Cookie identifier
const CpJudgeToken string = "cp_judge_token"

const cpJudgeSecret string = "Change_during_production"

// AuthenticationMiddleware : Middleware to handle auth part of the application
// Checks whether the cookie has a valid token.
func AuthenticationMiddleware(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		// Verify token
		tokenString, err := c.Cookies().Get(CpJudgeToken)
		if err != nil {
			// No cookie
			return c.Render(401, r.JSON(map[string]string{"message": err.Error()}))
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				// Wrong SigningMethod
				return nil, errors.New("Wrong signing method")
			}
			return []byte(cpJudgeSecret), nil
		})
		if err != nil {
			return c.Render(401, r.JSON(map[string]string{"message": err.Error()}))
		}
		// Process the request only if the token is Valid.
		if token.Valid {
			err := next(c)
			return err
		}
		return nil
	}
}

// GenerateTokenAndSetCookie : Generates a new token from the username
// in header and returns the signed token as a string.
func GenerateTokenAndSetCookie(c buffalo.Context, username string) error {
	// fmt.Println(username)
	if username == "" {
		return errors.New("username not provided")
	}
	claims := jwt.MapClaims{
		"user": username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(cpJudgeSecret))
	if err != nil {
		return err
	}
	c.Cookies().SetWithPath(CpJudgeToken, tokenString, "/") // Expires after a year
	return nil
}
