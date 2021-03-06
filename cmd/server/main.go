package main

import (
	"log"
	"net"

	"github.com/bytejedi/grpc_example/internal/bank/account"
	"google.golang.org/grpc"

	pb "github.com/bytejedi/grpc_example/pkg/proto/bank/account"
)

func main() {
	log.Println("Server running ...")

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}

	server := grpc.NewServer()

	pb.RegisterDepositServiceServer(server, &account.DepositServer{})

	log.Fatalln(server.Serve(listener))
}
