package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	v1 "zkzj_resource_project/pkg/question/api"
	question "zkzj_resource_project/pkg/question/service"
)

const (
	port = ":9090"
)

type QuestionServer struct {
	v1.UnimplementedQuestionServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	server := question.NewServer()
	v1.RegisterQuestionServer(s, server)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}



