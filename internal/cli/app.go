package cli

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/haunt98/pin-go/internal/pin"
	"github.com/make-go-great/color-go"
	"github.com/urfave/cli/v2"
)

const (
	Name  = "pin-go"
	usage = "let the fun begin"

	commandInit = "init"

	usageInit = "init pin, should be run only once"
)

type App struct {
	cliApp *cli.App
}

func NewApp(db *sql.DB) (*App, error) {
	repo, err := pin.NewRepository(context.Background(), db)
	if err != nil {
		return nil, fmt.Errorf("failed to new repository: %w", err)
	}

	service := pin.NewService(repo)
	handler := pin.NewHandler(service)

	a := &action{
		handler: handler,
	}

	cliApp := &cli.App{
		Name:   Name,
		Usage:  usage,
		Action: a.RunHelp,
		Commands: []*cli.Command{
			{
				Name:   commandInit,
				Usage:  usageInit,
				Action: a.RunInit,
			},
		},
	}

	return &App{
		cliApp: cliApp,
	}, nil
}

func (a *App) Run() {
	if err := a.cliApp.Run(os.Args); err != nil {
		color.PrintAppError(Name, err.Error())
	}
}
