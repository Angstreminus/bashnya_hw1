package main

import (
	"context"
	"log"
	"net"

	"my_service/server/pb"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedSayServer
}

func (s *Server) Hello(ctx context.Context, mess *pb.Mess) (*pb.Mess, error) {
	return &pb.Mess{
		Str: mess.Str,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterSayServer(s, &Server{})
	log.Print("Grpc server started")
	if err := s.Serve(listener); err != nil {
		log.Fatal("Failed to run server")
	}
}
