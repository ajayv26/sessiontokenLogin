package stores

import (
	"context"
	"database/sql"
	"jwt/models"
	"jwt/settings"
)

func GetBySessionToken(ctx context.Context, token string) (*models.OTPSession, error) {
	queryStmnt := `
	   SELECT * FROM otp_sessions 
	   WHERE otp_sessions.token = $1
	`
	row := settings.DBClient.QueryRowContext(ctx, queryStmnt, token)
	obj := &models.OTPSession{}
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

func InsertOTPSessionToken(ctx context.Context, tx *sql.Tx, arg *models.OTPSession) (*models.OTPSession, error) {
	queryStmnt := `
	INSERT INTO
	otp_sessions(
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
	obj := &models.OTPSession{}
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

func UpdateSessionToken(ctx context.Context, tx *sql.Tx, arg *models.OTPSession) error {
	queryStmnt := `
	 UPDATE otp_sessions
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
