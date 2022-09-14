package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/SpaceCadetOG/GoRestAPI/internal/transport/http"
)

// Will Contain Information about app -> Database connections
type App struct{}

// Start Up App
func (app *App) StartUp() error {
	fmt.Println("Rest API Starting Up")
	handler := transportHTTP.NewHandler()
	handler.HandleRoutes()
	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}
	return nil
}

func main() {
	app := App{}
	if err := app.StartUp(); err != nil {
		fmt.Printf("Error Starting Up Rest API...check error: %v", err)
	}

}
