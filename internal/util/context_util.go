package util

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type CtxKey int

const (
	Tx CtxKey = iota + 1
	UserID
	UserEmail
)

func GetTx(ctx context.Context) (*gorm.DB, error) {
	tx, ok := ctx.Value(Tx).(*gorm.DB)
	if !ok {
		return nil, errors.New("transaction not found in context")
	}
	return tx, nil
}

func GetUserID(ctx context.Context) (uint, error) {
	id, ok := ctx.Value(UserID).(uint)
	if !ok {
		return 0, errors.New("user id not found in context")
	}
	return id, nil
}

func GetUserEmail(ctx context.Context) (string, error) {
	email, ok := ctx.Value(UserEmail).(string)
	if !ok {
		return "", errors.New("user email not found in context")
	}
	return email, nil
}
