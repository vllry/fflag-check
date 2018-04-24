#!/usr/bin/env bash

govendor fetch github.com/gomodule/redigo/redis
govendor fetch golang.org/x/net/context
govendor fetch google.golang.org/grpc
govendor fetch google.golang.org/grpc/reflection
govendor fetch github.com/vllry/fflag-check-api