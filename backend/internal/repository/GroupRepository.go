package repository

import (
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type GroupRepository interface {
	Create(group domain.Group) (*mongo.InsertOneResult, error)
	FindAll() (*[]domain.Group, error)
	FindById(id primitive.ObjectID) (*domain.Group, error)
	FindByPermissions(email string) (*[]domain.Group, error)
	FindByName(name string) (*domain.Group, error)
	Update(group domain.Group) (*mongo.UpdateResult, error)
}
