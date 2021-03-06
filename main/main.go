package main

import (
	"fmt"
	"log"
	"strings"
	"github.com/lukesamk/movidesk"
)

func main() {
	api := movidesk.New("API_TOKEN")

	field := []string{"id", "subject", "createdDate"}
	filter := []string{"baseStatus=Stopped", "justification=Versão liberada"}

	errRequest := api.NewRequest(field, filter)
	trataErro(errRequest)

	fmt.Println(api.GetStringRequest())

	errResponse := api.Request.Run()
	trataErro(errResponse)

	groupBy := api.Request.Response.GroupByOrganization()

	for _, organizacao := range groupBy {
		fmt.Printf("%s:\n", organizacao.Nome)
		for _, ticket := range organizacao.Tickets {
			createdDate := strings.Split(ticket.CreatedDate, "T")
			fmt.Printf(
				"%d - %s\n%s | %s\n\n",
				ticket.ID,
				ticket.Subject,
				createdDate[0],
				ticket.Client[0].BusinessName,
			)
		}
	}
}

func trataErro(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
