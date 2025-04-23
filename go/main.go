package main

import (
	"context"
	"hermes/gen/hermes"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient(
		"localhost:9000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := hermes.NewHermesHandlerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.Handle(ctx, &hermes.HermesRequest{
		Payload: "Hello from Go!",
	})
	if err != nil {
		panic(err)
	}

	log.Printf("Response from PHP: %s\n", res.GetResult())
}
