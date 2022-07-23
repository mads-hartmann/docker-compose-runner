package pkg

import "fmt"

type Application struct {
	GitUrl string
	Watch  bool
}

func (app *Application) Run() {
	fmt.Printf("Would have run using %s\n", app.GitUrl)
}
