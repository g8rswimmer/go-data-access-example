package user

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/g8rswimmer/go-data-access-example/pkg/api/response"
	"github.com/g8rswimmer/go-data-access-example/pkg/model"
)

type UserDAO interface {
	Create(ctx context.Context, user *model.User) (*model.UserEntity, error)
	FetchByID(ctx context.Context, id string) (*model.UserEntity, error)
	Update(ctx context.Context, id string, user *model.User) (*model.UserEntity, error)
	Delete(ctx context.Context, id string) error
}

type errorMessage struct {
	ID      string `json:"id,omitempty"`
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}

type Handler struct {
	UserDAO UserDAO
}

func (h *Handler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := &model.User{}
		if err := json.NewDecoder(r.Body).Decode(user); err != nil {
			msg := &errorMessage{
				Error:   err.Error(),
				Message: "user json decode error",
			}
			response.JSON(w, http.StatusBadRequest, msg)
			return
		}

		entity, err := h.UserDAO.Create(r.Context(), user)
		if err != nil {
			msg := &errorMessage{
				Error:   err.Error(),
				Message: "user datastore error",
			}
			response.JSON(w, http.StatusInternalServerError, msg)
			return
		}
		response.JSON(w, http.StatusCreated, entity)
	}
}

func (h *Handler) FetchByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *Handler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *Handler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
