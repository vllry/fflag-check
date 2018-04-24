package main

import (
	"log"
	"net"

	"github.com/vllry/fflag-check-api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	loadConfigGlobals()

	listener, err := net.Listen("tcp", ":"+listenPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	fflagcheckapi.RegisterFeatureFlagServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
