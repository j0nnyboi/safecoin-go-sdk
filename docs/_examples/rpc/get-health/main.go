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
	ok, err := c.GetHealth(context.TODO())
	if err != nil {
		log.Fatalf("failed to get health, err: %v", err)
	}

	fmt.Println(ok)
}
