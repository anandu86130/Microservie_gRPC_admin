package services

import (
	"errors"

	pb "github.com/anandu86130/Microservice_gRPC_admin/internal/pb"
	productpb "github.com/anandu86130/Microservice_gRPC_admin/internal/product/pb"
	inter "github.com/anandu86130/Microservice_gRPC_admin/internal/repositories/interfaces"
	"github.com/anandu86130/Microservice_gRPC_admin/internal/services/interfaces"
	"github.com/anandu86130/Microservice_gRPC_admin/internal/utils"
)

type AdminService struct{
	AdminRepo inter.AdminRepoInter
	client productpb.ProductServiceClient
}

func (a *AdminService) AdminLoginService(p *pb.AdminRequest) (*pb.AdminResponse, error){
	admin, err := a.AdminRepo.FindAdmin(p.Username)
	if err != nil{
		return nil, err
	}

	if admin.Password != p.Password{
		return nil, errors.New("incorrect password")
	}

	token, err := utils.GenerateToken(p.Username)
	if err != nil{
		return nil, err
	}

	response := pb.AdminResponse{
		Status: "Success",
		Error: "",
		Message: "Admin Logged in successfully, Token"+token,
	}
	return &response,nil
} 

func NewAdminService(repo inter.AdminRepoInter, client productpb.ProductServiceClient) interfaces.AdminServiceInter{
	return &AdminService{
		AdminRepo: repo,
		client: client,
	}
}