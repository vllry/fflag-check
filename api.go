package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/vllry/fflag-check-api"
	"golang.org/x/net/context"
)

type server struct{}

func (s *server) GetFlag(ctx context.Context, query *fflagcheckapi.FlagQuery) (*fflagcheckapi.FlagResult, error) {
	flagValue, err := getFlag(query.AccountId, query.FlagName)
	if err != nil {
		return nil, err
	}

	return &fflagcheckapi.FlagResult{Found: true, Value: flagValue}, nil
}

func getFlag(accountId string, flagName string) (bool, error) {
	c, err := redis.Dial("tcp", redisAddress+":"+redisPort)
	if err != nil {
		return false, err
	}
	defer c.Close()

	key := accountId + "." + flagName
	val, err := c.Do("GET", key)
	fmt.Println("raw val: ", val, err)
	flagValue := false
	if val.(uint8) == 49 { // Hack - getting around encoding issues.
		flagValue = true
	}

	return flagValue, nil
}
