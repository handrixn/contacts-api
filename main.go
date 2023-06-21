package main

import (
	"log"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/handrixn/contacts-api/handler"
	"github.com/handrixn/contacts-api/repository"
	"github.com/handrixn/contacts-api/service"
)

func main() {
	router := mux.NewRouter()

	// Initialize dependencies
	contactRepo := repository.NewFileContactRepository("contacts.json")
	contactService := service.NewContactService(contactRepo)
	contactHandler := handler.NewContactHandler(contactService)

	// Define routes
	router.HandleFunc("/contacts", contactHandler.GetContacts).Methods("GET")
	router.HandleFunc("/contacts/{id}", contactHandler.GetContactByID).Methods("GET")
	router.HandleFunc("/contacts", contactHandler.CreateContact).Methods("POST")
	router.HandleFunc("/contacts/{id}", contactHandler.UpdateContact).Methods("PUT")
	router.HandleFunc("/contacts/{id}", contactHandler.DeleteContact).Methods("DELETE")

	router.Handle("/swagger.yaml", http.FileServer(http.Dir("./docs")))
	opts := middleware.SwaggerUIOpts{SpecURL: "swagger.yaml"}
	sh := middleware.SwaggerUI(opts, nil)
	router.Handle("/docs", sh)

	log.Fatal(http.ListenAndServe(":8080", router))
}
