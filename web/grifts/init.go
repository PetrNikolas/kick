package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/petrnikolasprokop/kick/web/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
