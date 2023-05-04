package services

import (
	"context"
	"database/sql"
	"fmt"
	"jwt/helpers"
	"jwt/models"
	"jwt/stores"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gofrs/uuid"
)

func GenerateJWT(email string) (string, error) {
	var mySecretKey = []byte("secretkey")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GenerateOTP(ctx context.Context, tx *sql.Tx, userName string) (*models.OTPSession, error) {

	user, err := stores.GetByUserName(ctx, userName)
	if err != nil {
		return nil, err
	}

	otpArg := models.OTPSession{
		UserID:    user.ID,
		Token:     helpers.GenerateRandomString(5),
		IsValid:   true,
		ExpiresAt: time.Now().Add(time.Minute * time.Duration(5)), // token will expire after 5 minutes
	}

	otp, err := stores.InsertOTPSessionToken(ctx, tx, &otpArg)
	if err != nil {
		return nil, err
	}
	return otp, nil

}

func Login(ctx context.Context, tx *sql.Tx, userName, otpToken string) (*models.AuthSession, error) {

	user, err := stores.GetByUserName(ctx, userName)
	if err != nil {
		return nil, err
	}

	otp, err := stores.GetBySessionToken(ctx, otpToken)
	if err != nil {
		return nil, err
	}
	if !otp.IsValid {
		return nil, fmt.Errorf("otp is invalid")
	}

	if otp.ExpiresAt.UTC().Unix() < time.Now().UTC().Unix() {
		return nil, fmt.Errorf("session is expired")
	}
	if otp.UserID != user.ID {
		return nil, fmt.Errorf("otp is invalid")
	}

	otp.IsValid = false
	if err := stores.UpdateSessionToken(ctx, tx, otp); err != nil {
		return nil, err
	}

	return CreateAuthSession(ctx, tx, user.ID)

}

func CreateAuthSession(ctx context.Context, tx *sql.Tx, userID int64) (*models.AuthSession, error) {
	token, err := uuid.NewV4()
	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}

	arg := &models.AuthSession{
		UserID:    userID,
		Token:     token,
		IsValid:   true,
		ExpiresAt: time.Now().Add(time.Hour * time.Duration(720)), // token will expire after 5 minutes
	}

	return stores.InsertAuthToken(ctx, tx, arg)
}
