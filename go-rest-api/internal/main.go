package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Aryan1302/Go-Tutorials/go-rest-api/pkg/swagger/server/restapi"
	"github.com/go-openapi/runtime"


	"github.com/Aryan1302/Go-Tutorials/go-rest-api/pkg/swagger/server/restapi/operations"
)

func main() {

	// Initialize Swagger
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatal(err)
	}

	api := operations.NewHelloAPIAPI(swaggerSpec)
	server := restapi.NewServer(api)

	defer func() {
		if err := server.Shutdown(); err != nil {
			// error handle
			log.Fatal(err)
		}
	}()

	server.Port = 8080
	api.CheckHealthHandler = operations.CheckHealthHandlerFunc(Health)
	api.GetHelloUserHandler = operations.GetHelloUserHandlerFunc(GetHelloUser)
	api.GetGopherNameHandler = operations.GetGopherNameHandlerFunc(GetGopherByName)

	// Start server which listens
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

// Health route returns OK
func Health(operations.CheckHealthParams) middleware.Responder {
	return operations.NewCheckHealthOK().WithPayload("OK")
}

// GetHelloUser returns Hello + your name
func GetHelloUser(user operations.GetHelloUserParams) middleware.Responder {
	return operations.NewGetHelloUserOK().WithPayload("Hello " + user.User + "!")
}

// GetGopherByName returns a goher in png
func GetGopherByName(gopher operations.GetGopherNameParams) middleware.Responder {
	var URL string
	if gopher.Name != "" {
		URL = "https://github.com/Aryan1302/gophers/raw/main/" + gopher.Name + ".png"
	} else {
		//by default we return dr-who gopher
		URL = "https://github.com/Aryan1302/gophers/raw/main/dr-who.png"
	}

	response, err := http.Get(URL)
	if err != nil {
		fmt.Println("error")
	}

	return operations.NewGetGopherNameOK().WithPayload(response.Body)
}
