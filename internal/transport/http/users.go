package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"github.com/thecaffeinedev/rest-api-ferretdb/internal/users"
)

type UserService interface {
	GetUser(ctx context.Context, email string) (users.User, error)
	CreateUser(ctx context.Context, in users.User) (users.User, error)
	UpdateUser(ctx context.Context, in users.User) (users.User, error)
	DeleteUser(ctx context.Context, email string) error
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")

	if email == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user, err := h.Service.GetUser(r.Context(), email)
	if err != nil {
		if errors.Is(err, users.ErrFetchingUser) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user users.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return
	}
	validate := validator.New()

	err := validate.Struct(user)
	if err != nil {
		log.Info(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err = h.Service.CreateUser(r.Context(), user)
	if err != nil {
		log.Error(err)
		return
	}
	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {

	email := chi.URLParam(r, "email")

	if email == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var user users.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return
	}
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user.Email = email

	user, err = h.Service.UpdateUser(r.Context(), user)
	if err != nil {
		log.Error(err.Error())
		return
	}
	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}

}

// TODO - needs to be fixed
func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")
	if email == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.Service.DeleteUser(r.Context(), email); err != nil {
		if errors.Is(err, users.ErrDeletingUser) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}
	if err := json.NewEncoder(w).Encode(Response{Message: "Successfully Deleted"}); err != nil {
		panic(err)
	}
}
