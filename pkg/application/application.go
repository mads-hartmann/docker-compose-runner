package application

import (
	"fmt"

	"github.com/mads-hartmann/docker-compose-runner/pkg/shell"
)

type Application struct {
	GitUrl string
	Watch  bool
}

func (app *Application) Run() error {
	stdout, stderr, err := shell.Command("ls")
	if err != nil {
		fmt.Println(stderr)
		return err
	}
	fmt.Println(stdout)
	return nil
}
