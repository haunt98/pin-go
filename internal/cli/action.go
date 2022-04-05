package cli

import (
	"github.com/haunt98/pin-go/internal/pin"
	"github.com/urfave/cli/v2"
)

type action struct {
	handler pin.Handler
}

func (a *action) RunHelp(c *cli.Context) error {
	return cli.ShowAppHelp(c)
}

func (a *action) RunInit(c *cli.Context) error {
	return a.handler.Init(c.Context)
}
