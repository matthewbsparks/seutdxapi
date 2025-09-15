package main

import (
	"fmt"
	"log"
)

func main() {
	apiKey, err := GetAPIKey()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("API Key: ", apiKey)

	client := NewClient(apiKey)

	ticketID := 347848
	ticket, err := client.GetTicket(ticketID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ticket)
}
