package stores

import (
	"context"
	"database/sql"
	"jwt/models"
	"jwt/settings"

	"github.com/gofrs/uuid"
)

func GetByAuthToken(ctx context.Context, token uuid.UUID) (*models.AuthSession, error) {
	queryStmnt := `
	   SELECT * FROM auth_sessions 
	   WHERE auth_session.token = $1
	`
	row := settings.DBClient.QueryRowContext(ctx, queryStmnt, token)
	obj := &models.AuthSession{}
	if err := row.Scan(
		&obj.ID,
		&obj.UserID,
		&obj.Token,
		&obj.IsValid,
		&obj.ExpiresAt,
		&obj.CreatedAt,
		&obj.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return obj, nil

}

func InsertAuthToken(ctx context.Context, tx *sql.Tx, arg *models.AuthSession) (*models.AuthSession, error) {
	queryStmnt := `
	INSERT INTO
	auth_sessions(
		user_id, 
		token,
		is_valid,
		expires_at
	)
	VALUES($1, $2 ,$3, $4)
	RETURNING *
	`
	row := tx.QueryRowContext(ctx, queryStmnt,
		&arg.UserID,
		&arg.Token,
		&arg.IsValid,
		&arg.ExpiresAt,
	)
	obj := &models.AuthSession{}
	if err := row.Scan(
		&obj.ID,
		&obj.UserID,
		&obj.Token,
		&obj.IsValid,
		&obj.ExpiresAt,
		&obj.CreatedAt,
		&obj.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return obj, nil
}

func UpdateAuthToken(ctx context.Context, tx *sql.Tx, arg *models.AuthSession) error {
	queryStmnt := `
	 UPDATE auth_sessions
	 SET 
	 is_valid=$1
	WHERE id=$2
	`
	_, err := tx.ExecContext(ctx, queryStmnt,
		&arg.IsValid,
		&arg.ID,
	)
	if err != nil {
		return err
	}
	return nil
}
