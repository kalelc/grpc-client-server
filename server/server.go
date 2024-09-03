package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/kalelc/grpc-example/model"
	pb "github.com/kalelc/grpc-example/proto" // Importa el paquete generado
)

type server struct {
	pb.UnimplementedUserServiceServer
	users [2]model.User
}

func NewServer() *server {
	var users [2]model.User
	users[0] = model.User{Identifier: "123456789-1", Name: "Andres", Email: "andres@gmail.com"}
	users[1] = model.User{Identifier: "987654321-0", Name: "Carlos", Email: "carlos@gmail.com"}

	return &server{users: users}
}

func (s *server) GetUser(ctx context.Context, req *pb.Filter) (*pb.Message, error) {
	var result model.User
	var message string

	for _, u := range s.users {
		if u.Identifier == req.Identifier {
			result = u
			message = "element find it"
			break
		} else {
			message = "element cannot find it"
		}

	}

	return &pb.Message{User: &pb.User{Identifier: result.Identifier, Name: result.Name, Email: result.Email}, Response: message}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, NewServer())
	log.Println("Server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
