package repositories

import (
	"github.com/anandu86130/Microservice_gRPC_admin/internal/model"
	inter "github.com/anandu86130/Microservice_gRPC_admin/internal/repositories/interfaces"
	"gorm.io/gorm"
)

type AdminRepo struct {
	DB *gorm.DB
}

func NewAdminRepo(db *gorm.DB) inter.AdminRepoInter {
	return &AdminRepo{
		DB: db,
	}
}

func (a *AdminRepo) FindAdmin(user string) (*model.AdminModel, error) {
	var admin model.AdminModel
	err := a.DB.Where("username=?", user).First(&admin).Error
	return &admin, err
}
