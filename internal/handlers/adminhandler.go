package handlers

import (
	"context"
	"log"

	pb "github.com/anandu86130/Microservice_gRPC_admin/internal/pb"
	"github.com/anandu86130/Microservice_gRPC_admin/internal/services/interfaces"
)

type AdminHandler struct {
	AdminService interfaces.AdminServiceInter
	pb.AdminServiceServer
}

func NewAdminHandler(repo interfaces.AdminServiceInter) *AdminHandler{
	return &AdminHandler{
		AdminService: repo,
	}
}

func (a *AdminHandler) AdminLogin(ctx context.Context, p *pb.AdminRequest) (*pb.AdminResponse, error){
	result, err := a.AdminService.AdminLoginService(p)
	if err != nil{
		log.Println("error while logging in")
		return nil, err
	}
	return result, nil
}