package main

import (
	"fmt"
	"log"
	"net"

	"github.com/gomodule/redigo/redis"
	"github.com/vllry/fflag-check-api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) GetFlag(ctx context.Context, query *fflagcheckapi.FlagQuery) (*fflagcheckapi.FlagResult, error) {
	c, err := redis.Dial("tcp", redisAddress+":"+redisPort)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	key := query.AccountId + "." + query.FlagName
	fmt.Println(key)
	val, err := c.Do("GET", key)
	fmt.Println("raw val: ", val, err)
	valBool := false
	if val.(uint8) == 49 { // Hack - getting around encoding issues.
		valBool = true
	}

	return &fflagcheckapi.FlagResult{Found: true, Value: valBool}, nil
}

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
