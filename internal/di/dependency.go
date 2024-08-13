package di

import (
	"log"

	"github.com/anandu86130/Microservice_gRPC_admin/config"
	"github.com/anandu86130/Microservice_gRPC_admin/internal/db"
	"github.com/anandu86130/Microservice_gRPC_admin/internal/handlers"
	"github.com/anandu86130/Microservice_gRPC_admin/internal/product"
	"github.com/anandu86130/Microservice_gRPC_admin/internal/repositories"
	"github.com/anandu86130/Microservice_gRPC_admin/internal/server"
	"github.com/anandu86130/Microservice_gRPC_admin/internal/services"
)

func Init() {
	config.LoadConfig()

	db := db.ConnectDB()
	productClient, err := product.CilentDial()
	if err != nil {
		log.Fatalf("something went wrong when dialing product client %v", err)
	}

	adminRepo := repositories.NewAdminRepo(db)
	adminsvc := services.NewAdminService(adminRepo, productClient)
	adminHandler := handlers.NewAdminHandler(adminsvc)
	server.NewGrpcServer(adminHandler)
}
