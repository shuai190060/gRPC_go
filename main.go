package main

import (
	"context"
	"flag"
)

func main() {

	var (
		jsonAddr = flag.String("json", ":3000", "listen address of the json transport")
		grpcAddr = flag.String("grpc", ":4000", "listen address of teh grpc trasport")
		svc      = loggingService{&priceService{}}
		ctx      = context.Background()
	)

	flag.Parse()

	jsonServer := NewJSONAPIServer(*jsonAddr, svc)
	jsonServer.Run()
}
