package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/lucassamuel/movidesk"
	"log"
	"os"
)

func main() {
	errEnv := godotenv.Load("./.env")
	trataErro(errEnv)

	api := movidesk.API{
		URL:   os.Getenv("URL_MOVIDESK"),
		Token: os.Getenv("TOKEN_MOVIDESK"),
	}

	field := []string{"id", "subject"}
	filter := []string{"baseStatus=Stopped", "justification=Versão liberada"}

	errRequest := api.NewRequest(field, filter)
	trataErro(errRequest)

	fmt.Println(api.GetStringRequest())

	api.Request.Run()
	ticket := api.Request.GetAll()

	errResponse := api.Request.Run()
	trataErro(errResponse)

	ticket, errGet := api.GetAll()
	trataErro(errGet)

	for i := 0; i < len(ticket); i++ {
		fmt.Printf("ID: %d, Assunto: %s, Cliente: %s, Organização: %s\n", ticket[i].ID, ticket[i].Subject, ticket[i].Client[0].BusinessName, ticket[i].Client[0].Organization.BusinessName)
	}

	fmt.Println(ticket)
}

func trataErro(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
