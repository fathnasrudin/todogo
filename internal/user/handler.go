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
	mux.HandleFunc("GET /api/users", h.List)
	mux.HandleFunc("POST /api/users", h.Create)
	mux.HandleFunc("PATCH /api/users/{id}", h.Update)
	mux.HandleFunc("DELETE /api/users/{id}", h.Delete)
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


func (h UserHandler) Create(w http.ResponseWriter, r *http.Request){
	var uData CreateUserInput

	if err := json.NewDecoder(r.Body).Decode(&uData); err != nil {
		http.Error(w, "Failed to decode body: " + err.Error(), http.StatusBadRequest)
	}

	err := h.service.Create(uData);
	if (err != nil) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return;
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	if err := json.NewEncoder(w).Encode(CreateUserResponse{Message: "Success create user"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return;
	}
}


func (h UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	var updateUserInput UpdateUserInput
	userId := r.PathValue("id");
	
	err := json.NewDecoder(r.Body).Decode(&updateUserInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	
	if err := h.service.Update(userId, updateUserInput); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("id");

	err := h.service.Delete(userId);
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.WriteHeader(http.StatusNoContent)
}