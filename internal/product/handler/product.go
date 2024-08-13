package handler

import (
	"context"
	"log"

	adminpb "github.com/anandu86130/Microservice_gRPC_admin/internal/pb"
	productpb "github.com/anandu86130/Microservice_gRPC_admin/internal/product/pb"
)

func CreateProductHandler(client productpb.ProductServiceClient, p *adminpb.AProductDetails) (*productpb.ProductResponse, error) {
	ctx := context.Background()

	response, err := client.CreateProduct(ctx, &productpb.ProductDetails{
		Category:    p.Category,
		Name:        p.Name,
		Price:       p.Price,
		Imagepath:   p.Imagepath,
		Description: p.Description,
		Size:        p.Size,
		Quantity:    p.Quantity,
	})

	if err != nil {
		log.Fatalf("error while creating productdetails")
		return nil, err
	}
	return response, nil
}

func FetchProductByIDHandler(client productpb.ProductServiceClient, p *adminpb.AProductById) (*productpb.ProductDetails, error) {
	ctx := context.Background()

	response, err := client.FetchByProductID(ctx, &productpb.ProductID{Id: p.Id})
	if err != nil {
		log.Printf("error while loading product by id")
		return nil, err
	}
	return response, nil
}

func FetchProductByNameHandler(client productpb.ProductServiceClient, p *adminpb.AProductByName) (*productpb.ProductDetails, error) {
	ctx := context.Background()

	response, err := client.FetchByName(ctx, &productpb.ProductByName{})
	if err != nil {
		log.Printf("error while loading productlist")
		return nil, err
	}
	return response, nil
}

func FetchAllProductHandler(client productpb.ProductServiceClient, p *adminpb.AdminNoParam) (*productpb.ProductList, error) {
	ctx := context.Background()

	response, err := client.FetchProducts(ctx, &productpb.NoParam{})
	if err != nil {
		log.Printf("error while loading productlist")
		return nil, err
	}
	return response, nil
}
