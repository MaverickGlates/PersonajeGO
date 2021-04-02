package main

import (
	"aprendiendo/controller"
	"aprendiendo/repository"
	"aprendiendo/service"
	"context"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/joho/godotenv"
)

func main() {
	loadEnviromentVariables() // carga las variables de entorno
	firestoreclient := getFirestoreClient()

	personajeRepository := repository.NewPersonajeRepository(firestoreclient)
	personajeService := service.NewPersonajeService(personajeRepository) //inyección de dependencias
	controller.NewPersonajeController(personajeService)                  // registro de ruta

	itemRepository := repository.NewItemRepository(firestoreclient)
	itemService := service.NewItemService(itemRepository)
	controller.NewItemController(itemService)

	port := os.Getenv("PORT")                     // llama al puerto
	log.Fatal(http.ListenAndServe(":"+port, nil)) // levanta servidor HTTP
}

func loadEnviromentVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func getFirestoreClient() *firestore.Client {
	projectId := os.Getenv("PROJECT_ID") //:= inferencia de tipos
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatal("error creating firestore client")
		// TODO: Handle error.
	}
	return client

}
