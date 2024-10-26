package main

import (
	"context"
	"log"

	"my_service/server/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// fix this
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	client := pb.NewSayClient(conn)

	msgText := "Hello world"

	message, err := client.Hello(context.Background(), &pb.Mess{Str: msgText})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Message from grpc server accepted: %s", message.GetStr())
}
