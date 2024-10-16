package interfaces

import (
	pb "github.com/anandu86130/Microservice_gRPC_admin/internal/pb"
)

type AdminServiceInter interface{
	AdminLoginService(p *pb.AdminRequest) (*pb.AdminResponse,error)
	CreateProduct(p *pb.AProductDetails) (*pb.AdminResponse,error)
	FetchProductByID(p *pb.AProductByID) (*pb.AProductDetails,error)
	FetchProductByName(p *pb.AProductByName) (*pb.AProductDetails,error)
	FetchAllProduct(p *pb.AdminNoParam) (*pb.AProductList,error)
}