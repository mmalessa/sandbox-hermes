package grpcmodule

import (
	"context"
	"hermes/gen/hermes"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	target     string
	connection *grpc.ClientConn
	client     hermes.HermesHandlerClient
}

func New(target string) *Client {
	return &Client{
		target: target,
	}
}

func (c *Client) Connect() error {
	var err error
	c.connection, err = grpc.NewClient(
		c.target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return err
	}

	c.client = hermes.NewHermesHandlerClient(c.connection)

	return nil
}

func (c *Client) Close() error {
	return c.connection.Close()
}

func (c *Client) Request(request *hermes.HermesRequest) (*hermes.HermesResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := c.client.Handle(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
