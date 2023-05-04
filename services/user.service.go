package services

import (
	"context"
	"database/sql"
	"fmt"
	"jwt/models"
	"jwt/stores"
)

func ListUserService(ctx context.Context) ([]models.User, error) {
	res, err := stores.ListStores(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetBYIDService(ctx context.Context, id int64) (*models.User, error) {
	return stores.GetByIDStore(ctx, id)
}

func InsertService(ctx context.Context, tx *sql.Tx, req models.User) (*models.User, error) {

	user := models.User{}
	count, err := stores.GetCountStore(ctx)
	if err != nil {
		return nil, err
	}

	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.Email = req.Email
	user.Phone = req.Phone
	user.PasswordHash = req.PasswordHash

	code := fmt.Sprintf("U%05d", count+1)
	req.Code = code

	obj, err := stores.InsertStore(ctx, req)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func UpdateService(ctx context.Context, tx *sql.Tx, id int64, req models.User) (*models.User, error) {
	user, err := stores.GetByIDStore(ctx, id)
	if err != nil {
		return nil, err
	}

	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.Email = req.Email
	user.Phone = req.Phone

	fmt.Println(user)
	if err := stores.UpdateUser(ctx, tx, *user); err != nil {
		return nil, err
	}
	return user, nil
}

func DeleteUserService(ctx context.Context, tx *sql.Tx, id int64) error {
	if err := stores.DeleteUser(ctx, tx, id); err != nil {
		return err
	}
	return nil
}
