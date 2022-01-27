package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/nitotang/go-soap/internal/transport/http"
)

// App - the struct which contains things like
// pointers to database connections
type App struct{}

func (app *App) Run() error {
	fmt.Println("Setting Up Our App")

	callSOAPClientSteps()

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}

	return nil
}

func main() {
	fmt.Println("SOAP Call test")

	// ----------------------

	//-----------------------
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error Starting Up")
		fmt.Println(err)
	}
}

func callSOAPClientSteps() {
	/*
		req := populateRequest()

		httpReq, err := generateSOAPRequest(req)
		if err != nil {
			fmt.Println("Some problem occurred in request generation")
		}

		response, err := soapCall(httpReq)
		if err != nil {
			fmt.Println("Problem occurred in making a SOAP call")
		}
	*/
}
