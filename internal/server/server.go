package server

import (
	"log"
	"net"

	"github.com/anandu86130/Microservice_gRPC_admin/internal/handlers"
	pb "github.com/anandu86130/Microservice_gRPC_admin/internal/pb"
	"google.golang.org/grpc"
)

func NewGrpcServer(handler *handlers.AdminHandler) {
	log.Println("connecting to gRPC server")
	lis, err := net.Listen("tcp", ":8084")
	if err != nil {
		log.Fatal("error when creating listener on port 8084")
	}

	grp := grpc.NewServer()
	pb.RegisterAdminServiceServer(grp, handler)

	log.Printf("listening on gRPC server on 8084")
	if err := grp.Serve(lis); err != nil {
		log.Fatal("error when connecting to gRPC server")
	}
}
