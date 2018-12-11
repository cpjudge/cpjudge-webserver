package actions

import "github.com/gobuffalo/buffalo"

// SignupHandler : Handles signup.
func SignupHandler(c buffalo.Context) error {
	err := GenerateTokenAndSetCookie(c)
	if err != nil {
		return c.Render(400, r.JSON(map[string]string{"message": "Bad request"}))
	}
	return c.Render(200, r.JSON(map[string]string{"message": "Check cookies"}))
}
