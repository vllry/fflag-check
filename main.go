package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
	"github.com/vllry/fflag-check-api"
)

const (
	defaultListenPort = "50051"
	defaultRedisAddress = "fflag-redis"
	defaultRedisPort = "6379"
)

type server struct{}

func (s *server) GetFlag(ctx context.Context, query *fflagrpc.FlagQuery) (*fflagrpc.Flag, error) {
	c, err := redis.Dial("tcp", defaultRedisAddress + ":" + defaultRedisPort)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	// 1.beta manually entered showed a 49
	key := query.AccountId + "." + query.FlagName
	fmt.Println(key)
	val, err := c.Do("GET", key)
	fmt.Println("raw val: ", val, err)
	valBool := false
	if val.(uint8) == 49 {
		valBool = true
	}

	return &fflagrpc.Flag{Found:true, Value:valBool}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":"+defaultListenPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	fflagrpc.RegisterFeatureFlagServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	go queryTest()
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func queryTest() {
	time.Sleep(5*time.Second)
	fmt.Println("Running ttest")
	fmt.Println(query("1", "beta"))
	fmt.Println(query("2", "beta"))
	fmt.Println(query("3", "beta"))
	fmt.Println(query("4", "beta"))
	fmt.Println(query("5", "beta"))
}

func query(account string, flag string) bool {

	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := fflagrpc.NewFeatureFlagClient(conn)

	// Contact the server and print out its response.
	r, err := c.GetFlag(context.Background(), &fflagrpc.FlagQuery{AccountId: account, FlagName: flag})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Value)

	return r.Value
}
