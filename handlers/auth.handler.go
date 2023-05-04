package handlers

import (
	"encoding/json"
	"fmt"
	"jwt/models"
	"jwt/services"
	"jwt/stores"
	"net/http"

	"github.com/go-chi/render"
)

func GetOTPHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := models.OTPLoginRequest{}
	if decodErr := json.NewDecoder(r.Body).Decode(&req); decodErr != nil {
		err := fmt.Errorf("invalid json request")
		render.JSON(w, r, err.Error())
		return
	}
	defer r.Body.Close()

	tx, err := stores.BeginTx(ctx)
	if err != nil {
		render.JSON(w, r, err.Error())
		return
	}
	defer stores.RollbackTx(ctx, tx)

	obj, err := services.GenerateOTP(ctx, tx, req.UserName)
	if err != nil {
		render.JSON(w, r, err.Error())
		return
	}

	if err := stores.CommitTx(ctx, tx); err != nil {
		render.JSON(w, r, err.Error())
		return
	}
	render.JSON(w, r, obj.Token)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := models.OTPLoginRequest{}
	if decodErr := json.NewDecoder(r.Body).Decode(&req); decodErr != nil {
		err := fmt.Errorf("invalid json request")
		render.JSON(w, r, err.Error())
		return
	}
	defer r.Body.Close()

	tx, err := stores.BeginTx(ctx)
	if err != nil {
		render.JSON(w, r, err.Error())
		return
	}
	defer stores.RollbackTx(ctx, tx)

	obj, err := services.Login(ctx, tx, req.UserName, req.OTP)
	if err != nil {
		render.JSON(w, r, err.Error())
		return
	}

	if err := stores.CommitTx(ctx, tx); err != nil {
		render.JSON(w, r, err.Error())
		return
	}
	render.JSON(w, r, obj)

}
