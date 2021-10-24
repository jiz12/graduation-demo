package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
	question "zkzj_resource_project/pkg/question/api"
)

const (
	address  = "47.119.140.240:9018"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := question.NewQuestionClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetQuestionNum(ctx, &question.QuestionNumRequest{QuestionId: 1040})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %v", r.Count)
	fmt.Println(r.Count)

}
