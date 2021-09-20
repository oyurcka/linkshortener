package main

import (
	"log"
	"net"

	"github.com/yurioganesyan/LinkShortener/pkg/api"
	"github.com/yurioganesyan/LinkShortener/pkg/shortener"
	"google.golang.org/grpc"
)

func main() {
	grpcserver := grpc.NewServer()
	server := &shortener.GRPCServer{}
	api.RegisterShortenerServer(grpcserver, server)

	listener, error := net.Listen("tcp", ":9080")
	if error != nil {
		log.Fatalln(error)
	}

	if error := grpcserver.Serve(listener); error != nil {
		log.Fatalln(error)
	}
}
