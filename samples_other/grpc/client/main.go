package main

import (
	"log"

	"grpc_example/mygrpc"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	c := mygrpc.NewSampleServiceClient(conn)

	response, err := c.GetData(context.Background(), &mygrpc.Message{Body: "送信データ"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Print(response.Body)

	defer conn.Close()
}
