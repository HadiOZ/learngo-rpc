package main

import (
	"context"
	"grpc/payload"
	"log"
	"net"

	"google.golang.org/grpc"
)

func errorHandel(msg string, err error) {
	if err != nil {
		log.Fatalf("[x] error %s : %v", msg, err)
	}
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	errorHandel("can't listen tcp connection", err)
	grpcServer := grpc.NewServer()

	server := payload.SignInService{}
	server.Auth = func(ctx context.Context, in *payload.SignInRequest) (*payload.SignInResponse, error) {
		log.Printf(`incoming request auth from email: %s password %s`, in.Email, in.Password)
		return &payload.SignInResponse{
			UserId: "123",
			Jwt:    "456",
		}, nil
	}

	payload.RegisterSignInService(grpcServer, &server)

	log.Println("server start at :9000")
	err = grpcServer.Serve(lis)
	errorHandel("can't serve server", err)

}
