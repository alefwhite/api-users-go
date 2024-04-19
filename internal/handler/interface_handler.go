package handler

import (
	"github.com/alefwhite/api-users-go/internal/service/categoryservice"
	"github.com/alefwhite/api-users-go/internal/service/productservice"
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

	CreateCategory(w http.ResponseWriter, r *http.Request)

	CreateProduct(w http.ResponseWriter, r *http.Request)
	UpdateProduct(w http.ResponseWriter, r *http.Request)
	DeleteProduct(w http.ResponseWriter, r *http.Request)
	FindManyProducts(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	userService     userservice.UserService
	categoryService categoryservice.CategoryService
	productService  productservice.ProductService
}

func NewHandler(userService userservice.UserService,
	categoryService categoryservice.CategoryService,
	productService productservice.ProductService) Handler {
	return &handler{
		userService:     userService,
		categoryService: categoryService,
		productService:  productService,
	}
}
