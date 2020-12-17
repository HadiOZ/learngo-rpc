package main

import (
	"context"
	"grpc/payload"
	"log"
	"time"

	"google.golang.org/grpc"
)

const (
	adress = "localhost:9000"
)

func main() {
	conn, err := grpc.Dial(adress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("can't connect to server")
	}

	defer conn.Close()

	c := payload.NewSignInClient(conn)

	ctx, cencel := context.WithTimeout(context.Background(), time.Second)
	defer cencel()
	r, err := c.Auth(ctx, &payload.SignInRequest{
		Email:    "iqbal@sugoi.com",
		Password: "qwerty",
	})
	if err != nil {
		log.Fatal("can't Auth")
	}

	log.Printf("get message : %s", r)
}
