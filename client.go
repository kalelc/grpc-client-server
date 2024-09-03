package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	pb "github.com/kalelc/grpc-example/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserServiceClient(conn)

	resp, err := c.GetUser(context.Background(), &pb.Filter{Identifier: "123456789-1"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Response: %s", resp.Response)
}
