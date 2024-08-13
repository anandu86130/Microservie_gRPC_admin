package product

import (
	"log"

	productpb "github.com/anandu86130/Microservice_gRPC_admin/internal/product/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CilentDial() (productpb.ProductServiceClient, error){
	grpc, err := grpc.NewClient("localhost:8083", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Printf("error when dialing to gRPC client:8038")
		return nil,err
	}
	log.Printf("Successfully connected to gRPC client:8083")
	return productpb.NewProductServiceClient(grpc), nil
}