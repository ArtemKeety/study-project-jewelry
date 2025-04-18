package handler

import (
	"curs/pkg/service"
	"github.com/gorilla/mux"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *mux.Router {
	router := mux.NewRouter()

	auth := router.PathPrefix("/auth").Subrouter()
	{
		auth.HandleFunc("/sign-in", h.SignIn).Methods("POST")
		auth.HandleFunc("/sign-up", h.SignUp).Methods("POST")
	}

	api := router.PathPrefix("/api").Subrouter()
	api.Use(h.userIdentity)
	{
		product := api.PathPrefix("/product").Subrouter()
		{
			product.HandleFunc("/", h.GetProducts).Methods("GET")
			product.HandleFunc("/{Id:[0-9]+}", h.GetCurProduct).Methods("GET")
		}

		cart := api.PathPrefix("/cart").Subrouter()

		{
			cart.HandleFunc("/", h.GetCart).Methods("GET")
			cart.HandleFunc("/", h.ClearCart).Methods("DELETE")
			cart.HandleFunc("/{Id:[0-9]+}", h.AddInCart).Methods("POST")
			cart.HandleFunc("/{Id:[0-9]+}", h.CheckInCart).Methods("GET")
			cart.HandleFunc("/{Id:[0-9]+}", h.UpdateItemCart).Methods("PUT")
			cart.HandleFunc("/{Id:[0-9]+}", h.RemoveInCart).Methods("DELETE")
		}

	}

	return router
}
