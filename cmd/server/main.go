package main

import (
	"fmt"
)

type App struct {}

func (app *App) Run() error {
	fmt.Println("Running Application...")

	return nil
}

func main() {
	fmt.Println("Hello Car Service!")
	app := App{}

	if err := app.Run(); err != nil {
		fmt.Println("Error Starting Application")
	}
}