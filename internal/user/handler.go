package user

import (
	"encoding/json"
	"net/http"
)

type IUserHandler interface{
	List(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
};

type UserHandler struct {
	service *UserService
}

func NewUserHandler(s *UserService) *UserHandler {
	return &UserHandler{
		service: s,
	}
}

func (h UserHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/users", h.List)
}

func (h UserHandler) List(w http.ResponseWriter, r *http.Request){
	users, err := h.service.List();

	if (err != nil) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return;
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return;
	}
}

