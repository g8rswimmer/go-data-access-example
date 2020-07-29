package user

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/g8rswimmer/go-data-access-example/pkg/api/response"
	"github.com/g8rswimmer/go-data-access-example/pkg/errorx"
	"github.com/g8rswimmer/go-data-access-example/pkg/model"
	"github.com/gorilla/mux"
)

// DAO is the user data access object
type DAO interface {
	Create(ctx context.Context, user *model.User) (*model.UserEntity, error)
	FetchByID(ctx context.Context, id string) (*model.UserEntity, error)
	FetchAll(ctx context.Context) ([]*model.UserEntity, error)
	Update(ctx context.Context, id string, user *model.User) (*model.UserEntity, error)
	Delete(ctx context.Context, id string) error
}

const userID = "id"

type errorMessage struct {
	ID      string `json:"id,omitempty"`
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}

// Handler provides all of the user handlers
type Handler struct {
	UserDAO DAO
}

// create handles the user create request
func (h *Handler) create() http.HandlerFunc {
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

// fetchByID will return an user by its id
func (h *Handler) fetchByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars[userID]
		entity, err := h.UserDAO.FetchByID(r.Context(), id)
		switch {
		case errors.Is(err, errorx.ErrNoUser):
			msg := &errorMessage{
				Message: fmt.Sprintf("user %s does not exist", id),
			}
			response.JSON(w, http.StatusNotFound, msg)
			return
		case errors.Is(err, errorx.ErrDeleteUser):
			msg := &errorMessage{
				Message: fmt.Sprintf("user %s has been deleted", id),
			}
			response.JSON(w, http.StatusGone, msg)
			return
		case err != nil:
			msg := &errorMessage{
				Error:   err.Error(),
				Message: "user datastore error",
			}
			response.JSON(w, http.StatusInternalServerError, msg)
			return
		default:
			response.JSON(w, http.StatusOK, entity)
		}

	}
}

// list will return all of the users
func (h *Handler) list() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		entities, err := h.UserDAO.FetchAll(r.Context())
		switch {
		case errors.Is(err, errorx.ErrNoUser):
			msg := &errorMessage{
				Message: fmt.Sprintf("no users exist"),
			}
			response.JSON(w, http.StatusNotFound, msg)
			return
		case err != nil:
			msg := &errorMessage{
				Error:   err.Error(),
				Message: "user datastore error",
			}
			response.JSON(w, http.StatusInternalServerError, msg)
			return
		default:
			response.JSON(w, http.StatusOK, entities)
		}
	}
}

// update will return the updated user
func (h *Handler) update() http.HandlerFunc {
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
		if len(user.FirstName) == 0 && len(user.LastName) == 0 {
			msg := &errorMessage{
				Message: "user must have fields to update",
			}
			response.JSON(w, http.StatusBadRequest, msg)
			return
		}

		vars := mux.Vars(r)
		id := vars[userID]
		entity, err := h.UserDAO.Update(r.Context(), id, user)
		switch {
		case errors.Is(err, errorx.ErrNoUser):
			msg := &errorMessage{
				Message: fmt.Sprintf("user %s does not exist", id),
			}
			response.JSON(w, http.StatusNotFound, msg)
			return
		case errors.Is(err, errorx.ErrDeleteUser):
			msg := &errorMessage{
				Message: fmt.Sprintf("user %s has been deleted", id),
			}
			response.JSON(w, http.StatusGone, msg)
			return
		case err != nil:
			msg := &errorMessage{
				Error:   err.Error(),
				Message: "user datastore error",
			}
			response.JSON(w, http.StatusInternalServerError, msg)
			return
		default:
			response.JSON(w, http.StatusOK, entity)
		}

	}
}

// delete will remove the user
func (h *Handler) delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars[userID]
		err := h.UserDAO.Delete(r.Context(), id)
		switch {
		case errors.Is(err, errorx.ErrNoUser):
			msg := &errorMessage{
				Message: fmt.Sprintf("user %s does not exist", id),
			}
			response.JSON(w, http.StatusNotFound, msg)
			return
		case errors.Is(err, errorx.ErrDeleteUser):
			msg := &errorMessage{
				Message: fmt.Sprintf("user %s has been deleted", id),
			}
			response.JSON(w, http.StatusGone, msg)
			return
		case err != nil:
			msg := &errorMessage{
				Error:   err.Error(),
				Message: "user datastore error",
			}
			response.JSON(w, http.StatusInternalServerError, msg)
			return
		default:
			response.JSON(w, http.StatusNoContent, nil)
		}
	}

}

// Add will configure the routes for user operations
func (h *Handler) Add(router *mux.Router) {
	router.Methods(http.MethodPost).Path("/user").Handler(h.create()).Name("user-create")
	router.Methods(http.MethodGet).Path(fmt.Sprintf("/users/{%s}", userID)).Handler(h.fetchByID()).Name("user-fetch")
	router.Methods(http.MethodGet).Path("/users").Handler(h.list()).Name("user-fetch-all")
	router.Methods(http.MethodPatch).Path(fmt.Sprintf("/users/{%s}", userID)).Handler(h.update()).Name("user-update")
	router.Methods(http.MethodDelete).Path(fmt.Sprintf("/users/{%s}", userID)).Handler(h.delete()).Name("user-delete")
}
