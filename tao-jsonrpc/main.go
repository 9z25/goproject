package main

import (

	"github.com/btcsuite/btcd/rpcclient"

	"log"
)

func main() {
	// create new client instance
	client, err := rpcclient.New(&rpcclient.ConnConfig{
		HTTPPostMode: true,
		DisableTLS:   true,
		Host:         "34.221.217.31:8332",
		User:         "rsenejr",
		Pass:         "seventy7",
	}, nil)
	if err != nil {
		log.Fatalf("error creating new btc client: %v", err)
	}

	// list accounts
	accounts, err := client.GetInfo()
	if err != nil {
		log.Fatalf("error listing accounts: %v", err)
	}


	log.Println(err, accounts)

}