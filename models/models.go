package models

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/volatiletech/null"
)

type User struct {
	ID           int64     `json:"id"`
	Code         string    `json:"code"`
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	PasswordHash string    `json:"passwordHash"`
	IsAdmin      bool      `json:"isAdmin"`
	IsArchived   bool      `json:"isArchived"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type Auther struct {
	ID      int64     `json:"id"`
	Name    string    `json:"name"`
	IsAdmin bool      `json:"isAdmin"`
	Token   uuid.UUID `json:"token"`
}

type OTPSession struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"userID"`
	Token     string    `json:"token"`
	IsValid   bool      `json:"isValid"`
	ExpiresAt time.Time `json:"expiresAt"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type AuthSession struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"userID"`
	Token     uuid.UUID `json:"token"`
	IsValid   bool      `json:"isValid"`
	ExpiresAt time.Time `json:"expiresAt"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserReq struct {
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	PasswordHash string `json:"passwordHash"`
}

type LoginRequest struct {
	Email        null.String `json:"email"`
	Phone        null.String `json:"phone"`
	PasswordHash string      `json:"passwordHash"`
}
type OTPLoginRequest struct {
	UserName string `json:"userName"`
	OTP      string `json:"otp"`
}

type Token struct {
	Email       string `json:"email"`
	TokenString string `jaon:"tokenString"`
}

type OTPRequest struct {
	Email null.String `json:"email"`
	Phone null.String `json:"phone"`
}

type TokenHeader struct {
	TYP string `json:"typ"`
	ALG string `json:"alg"`
}
