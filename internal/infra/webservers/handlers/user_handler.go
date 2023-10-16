package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/danielzinhors/APIS-go/internal/dto"
	"github.com/danielzinhors/APIS-go/internal/entity"
	"github.com/danielzinhors/APIS-go/internal/infra/database"
)

type UserHandler struct {
	UserDB database.UserInterface
}

func NewUserHanbler(userDb database.UserInterface) *UserHandler {
	return &UserHandler{UserDB: userDb}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.UserDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
