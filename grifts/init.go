package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/petrnikolasprokop/kick/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
