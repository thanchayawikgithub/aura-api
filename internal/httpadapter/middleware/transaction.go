package middleware

import (
	"aura/internal/storage"
	"aura/internal/util"
	"context"

	"github.com/labstack/echo/v4"
)

func WithTx(s *storage.Storage) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tx := s.GetDB().Begin()
			defer func() {
				if r := recover(); r != nil {
					tx.Rollback()
				}
			}()

			ctx := context.WithValue(c.Request().Context(), util.Tx, tx)
			c.SetRequest(c.Request().WithContext(ctx))

			if err := next(c); err != nil {
				tx.Rollback()
				return err
			}

			if err := tx.Commit().Error; err != nil {
				tx.Rollback()
				return err
			}
			return nil
		}
	}
}
