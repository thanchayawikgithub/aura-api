package util

import (
	"context"
	"errors"
)

type CtxKey int

const (
	UserID CtxKey = iota + 1
	UserEmail
)

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
