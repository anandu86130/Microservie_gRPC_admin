package handlers

import (
	"context"
	"log"

	pb "github.com/anandu86130/Microservice_gRPC_admin/internal/pb"
)

func (a *AdminHandler) CreateProduct(ctx context.Context, p *pb.AProductDetails) (*pb.AdminResponse, error) {
	result, err := a.AdminService.CreateProduct(p)
	if err != nil {
		log.Println("error while creating product")
		return nil, err
	}
	return result, nil
}

func (a *AdminHandler) FetchProductByID(ctx context.Context, p *pb.AProductByID) (*pb.AProductDetails, error) {
	result, err := a.AdminService.FetchProductByID(p)
	if err != nil {
		log.Println("error while fetching product by id")
		return nil, err
	}
	return result, nil
}

func (a *AdminHandler) FetchProductByName(ctx context.Context, p *pb.AProductByName) (*pb.AProductDetails, error) {
	result, err := a.AdminService.FetchProductByName(p)
	if err != nil {
		log.Println("error while fetching product by name")
		return nil, err
	}
	return result, nil
}

func (a *AdminHandler) FetchProducts(ctx context.Context, p *pb.AdminNoParam) (*pb.AProductList, error) {
	result, err := a.AdminService.FetchAllProduct(p)
	if err != nil {
		log.Println("error while fetching all product list")
		return nil, err
	}
	return result, nil
}
