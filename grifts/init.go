package grifts

import (
	"github.com/cpjudge/cpjudge_webserver/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
