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

	tokenAccount, err := c.GetTokenAccount(context.Background(), "81Ck4pb8sZVYacLVHh4KbyiYHX8qnR4JvuZcyPiJN5cc")
	if err != nil {
		log.Fatalf("failed to get token account, err: %v", err)
	}

	fmt.Printf("%+v\n", tokenAccount)
}
