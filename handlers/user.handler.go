package handlers

import (
	"encoding/json"
	"fmt"
	"jwt/models"
	"jwt/services"
	"jwt/stores"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func UserListHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	res, err := services.ListUserService(ctx)
	if err != nil {
		render.JSON(w, r, err)
	}
	render.JSON(w, r, res)
}

func UserGetByIDHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		render.JSON(w, r, err)
	}
	res, err := services.GetBYIDService(ctx, id)
	if err != nil {
		render.JSON(w, r, err)
	}
	render.JSON(w, r, res)
}

func UserInsertHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := models.User{}
	if decodeErr := json.NewDecoder(r.Body).Decode(&req); decodeErr != nil {
		err := fmt.Errorf("invalid Json Request")
		render.JSON(w, r, err.Error())
		return
	}
	defer r.Body.Close()
	tx, err := stores.BeginTx(ctx)
	if err != nil {
		render.JSON(w, r, err)
		return
	}
	defer stores.RollbackTx(ctx, tx)

	res, err := services.InsertService(ctx, tx, req)
	if err != nil {
		render.JSON(w, r, err.Error())
		return
	}
	fmt.Println(req)
	if err := stores.CommitTx(ctx, tx); err != nil {
		render.JSON(w, r, err)
		return
	}
	render.JSON(w, r, res)
}

func UserUpdateHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := models.User{}
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		render.JSON(w, r, err)
		return
	}
	if decodeErr := json.NewDecoder(r.Body).Decode(&req); decodeErr != nil {
		render.JSON(w, r, err)
		return
	}
	defer r.Body.Close()

	tx, err := stores.BeginTx(ctx)
	if err != nil {
		render.JSON(w, r, err)
		return
	}
	defer stores.RollbackTx(ctx, tx)

	res, err := services.UpdateService(ctx, tx, id, req)
	if err != nil {
		fmt.Println(res)
		render.JSON(w, r, err)
		return
	}
	if err := stores.CommitTx(ctx, tx); err != nil {
		render.JSON(w, r, err)
		return
	}
	render.JSON(w, r, res)

}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		render.JSON(w, r, err)
		return
	}
	tx, err := stores.BeginTx(ctx)
	if err != nil {
		render.JSON(w, r, err)
		return
	}
	defer stores.RollbackTx(ctx, tx)

	err = services.DeleteUserService(ctx, tx, id)
	if err != nil {
		render.JSON(w, r, err)
		return
	}
	if err := stores.CommitTx(ctx, tx); err != nil {
		render.JSON(w, r, err)
		return

	}
	msg := "User deleted successfully"
	render.JSON(w, r, msg)

}
