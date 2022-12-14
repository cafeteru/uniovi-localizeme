package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"uniovi-localizeme/internal/core/domain"
	"uniovi-localizeme/internal/core/domain/dto"
	"uniovi-localizeme/internal/core/domain/xmlDto"
)

type BaseStringService interface {
	Create(request domain.BaseString, user *domain.User) (domain.BaseString, error)
	Delete(id primitive.ObjectID, user *domain.User) (bool, error)
	Disable(id primitive.ObjectID, user *domain.User) (*domain.BaseString, error)
	FindAll() (*[]domain.BaseString, error)
	FindByGroup(id primitive.ObjectID, user *domain.User) (*[]domain.BaseString, error)
	FindByIdentifier(identifier string, user *domain.User) (*domain.BaseString, error)
	FindByIdentifierAndLanguageAndState(identifier string, isoCode string, stageName string, user *domain.User) (*string, error)
	FindByLanguage(id primitive.ObjectID, user *domain.User) (*[]domain.BaseString, error)
	FindByPermissions(id primitive.ObjectID) (*[]domain.BaseString, error)
	Read(xliff xmlDto.Xliff, user *domain.User, stageId primitive.ObjectID, groupId primitive.ObjectID) (*[]domain.BaseString, error)
	Update(baseString domain.BaseString, user *domain.User) (*domain.BaseString, error)
	Write(xliff dto.XliffDto) (*xmlDto.Xliff, error)
}
