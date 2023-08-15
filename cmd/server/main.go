package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/Galbeyte1/go-cars-service/api/http"
	"github.com/Galbeyte1/go-cars-service/internal/car"
	"github.com/Galbeyte1/go-cars-service/internal/database"
)

type App struct {}

func (app *App) Run() error {
	fmt.Println("Running Application...")

	db := database.NewDatabase()

	carService := car.NewService(db)

	handler := transportHTTP.NewHandler(carService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}

	return nil
}

func main() {
	fmt.Println("Hello Car Service!")
	app := App{}

	if err := app.Run(); err != nil {
		fmt.Println("Error Starting Application")
	}
}