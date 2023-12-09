package main

import (
	"context"
	"fmt"
	"log"

	"github.com/j0nnyboi/safecoin-go-sdk/client"
	"github.com/j0nnyboi/safecoin-go-sdk/rpc"
)

func main() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	tokenAccounts, err := c.GetTokenAccountsByOwner(context.Background(), "9aE476sH92Vz7DMPyq5WLPkrKWivxeuTKEFKd2sZZcde")
	if err != nil {
		log.Fatalf("failed to get token accounts, err: %v", err)
	}

	for pubkey, tokenAccount := range tokenAccounts {
		fmt.Printf("%v => %+v\n", pubkey.ToBase58(), tokenAccount)
	}
}
