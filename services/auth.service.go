package services

import "github.com/kavindu-prabhashitha/go-crud-gin-mongodb/models"

type AuthService interface {
	SignUpUser(*models.SignUpInput) (*models.DBResponse, error)
	SignInUser(*models.SignInInput) (*models.DBResponse, error)
}
