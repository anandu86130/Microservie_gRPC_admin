package interfaces

import "github.com/anandu86130/Microservice_gRPC_admin/internal/model"

type AdminRepoInter interface {
	FindAdmin(username string) (*model.AdminModel,error)
}