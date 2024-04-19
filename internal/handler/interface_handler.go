package handler

import (
	"github.com/alefwhite/api-users-go/internal/service/userservice"
	"net/http"
)

type Handler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	GetUserByID(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	FindManyUsers(w http.ResponseWriter, r *http.Request)
	UpdateUserPassword(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	service userservice.UserService
}

func NewHandler(service userservice.UserService) Handler {
	return &handler{
		service,
	}
}
