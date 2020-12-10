package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"example.com/demo-grpc/user_service"
	"google.golang.org/grpc"
)

func main() {
	var id int64
	if len(os.Args) != 2 {
		log.Fatal("Usage: go run main.go <id>")
	}

	id, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial(":8082", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := user_service.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetUser(ctx, &user_service.GetUserRequest{Id: id})
	if err != nil {
		log.Fatalf("GRPC error: %v", err)
	}
	log.Printf("Result: %+v", r.GetUser())
}
