package services

import (
	pb "github.com/anandu86130/Microservice_gRPC_admin/internal/pb"
	productpb "github.com/anandu86130/Microservice_gRPC_admin/internal/product/handler"
)

func (a *AdminService) CreateProduct(p *pb.AProductDetails) (*pb.AdminResponse, error) {
	result, err := productpb.CreateProductHandler(a.client, p)
	if err != nil {
		return nil, err
	}

	return &pb.AdminResponse{
		Status:  "Success",
		Error:   result.Error,
		Message: result.Message,
	}, nil
}

func (a *AdminService) FetchAllProduct(p *pb.AdminNoParam) (*pb.AProductList, error) {
	result, err := productpb.FetchAllProductHandler(a.client, p)
	if err != nil {
		return nil, err
	}
	var product []*pb.AProductDetails
	for _, i := range result.Item {
		product = append(product, &pb.AProductDetails{
			Id:          i.Id,
			Category:    i.Category,
			Name:        i.Name,
			Imagepath:   i.Imagepath,
			Description: i.Description,
			Size:        i.Size,
			Quantity:    i.Quantity,
		})
	}
	return &pb.AProductList{
		Products: product,
	}, nil
}

func (a *AdminService) FetchProductByID(p *pb.AProductByID) (*pb.AProductDetails, error) {
	result, err := productpb.FetchProductByIDHandler(a.client, p)
	if err != nil {
		return nil, err
	}
	return &pb.AProductDetails{
		Id:          result.Id,
		Category:    result.Category,
		Name:        result.Name,
		Imagepath:   result.Imagepath,
		Description: result.Description,
		Size:        result.Size,
		Quantity:    result.Quantity,
	}, nil
}

func (a *AdminService) FetchProductByName(p *pb.AProductByName) (*pb.AProductDetails, error) {
	result, err := productpb.FetchProductByNameHandler(a.client, p)
	if err != nil {
		return nil, err
	}
	return &pb.AProductDetails{
		Id:          result.Id,
		Category:    result.Category,
		Name:        result.Name,
		Imagepath:   result.Imagepath,
		Description: result.Description,
		Size:        result.Size,
		Quantity:    result.Quantity,
	}, nil

}
