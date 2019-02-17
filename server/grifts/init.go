package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/petrnikolasprokop/kick/server/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
