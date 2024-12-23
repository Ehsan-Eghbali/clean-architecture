package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"clean-architecture-go/internal/domain"
	"clean-architecture-go/internal/usecase"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler(uuc usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: uuc,
	}
}

// CreateUserHandler هندلر ایجاد کاربر جدید
func (uh *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := uh.userUseCase.CreateUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// GetUserByIDHandler هندلر دریافت کاربر براساس ID
func (uh *UserHandler) GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	user, err := uh.userUseCase.GetUserByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// GetAllUsersHandler هندلر دریافت همه کاربران
func (uh *UserHandler) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := uh.userUseCase.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}
