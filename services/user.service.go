package services

import "github.com/kavindu-prabhashitha/go-crud-gin-mongodb/models"

type UserService interface {
	FindUserById(string) (*models.DBResponse, error)
	FindUserByEmail(string) (*models.DBResponse, error)
}
