package export

import (
	"aura/internal/storage"
	"context"
	"fmt"

	"github.com/xuri/excelize/v2"
)

const (
	ExportUsersPath = "export/users.xlsx"
)

type (
	IExportUser interface {
		ExportUsers(ctx context.Context) error
	}

	ExportUser struct {
		UserStorage storage.IUserStorage
	}
)

func NewExportUser(userStorage storage.IUserStorage) *ExportUser {
	return &ExportUser{UserStorage: userStorage}
}

func (e *ExportUser) ExportUsers(ctx context.Context) error {
	users, err := e.UserStorage.FindAll(ctx)
	if err != nil {
		return err
	}

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	headers := []string{"ID", "Email", "DisplayName", "CreatedAt"}
	f.SetSheetRow("Sheet1", "A1", &headers)

	for i, user := range users {
		row := []interface{}{user.ID, user.Email, user.DisplayName, user.CreatedAt.Format("2006-01-02 15:04:05")}
		cell, err := excelize.CoordinatesToCellName(1, i+2)
		if err != nil {
			return err
		}
		if err := f.SetSheetRow("Sheet1", cell, &row); err != nil {
			return err
		}
	}
	return f.SaveAs(ExportUsersPath)
}
