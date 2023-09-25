package main

import (
	"flag"
)

func main() {

	// var (
	// 	jsonAddr = flag.String("json", ":3000", "listen address of the json transport")
	// 	grpcAddr = flag.String("grpc", ":4000", "listen address of teh grpc trasport")
	// 	svc      = loggingService{&priceService{}}
	// 	ctx      = context.Background()
	// )

	// flag.Parse()

	// grpcClient, err := client.

	// jsonServer := NewJSONAPIServer(*jsonAddr, svc)
	// jsonServer.Run()

	port := flag.String("port", ":3000", "listen address of the service")
	flag.Parse()
	svc := loggingService{&priceService{}}
	server := NewJSONAPIServer(*port, svc)
	server.Run()

}
