package handler

import (
	"curs/pkg/service"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *mux.Router {
	router := mux.NewRouter()

	router.Use(h.getLog)

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
			product.HandleFunc("/by_category_id/{Id:[0-9]+}", h.GetFilterProduct).Methods("GET")
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
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	return router
}

func GetId(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	result, err := strconv.Atoi(vars["Id"])
	if err != nil {
		return -1, err
	}

	if result < 1 {
		return -1, errors.New("invalid id")
	}

	return result, nil
}
