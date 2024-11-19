package storage

// import (
// 	"aura/internal/model"
// 	"context"
// 	"database/sql"
// 	"testing"

// 	"github.com/DATA-DOG/go-sqlmock"
// 	"github.com/stretchr/testify/suite"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// type StorageTestSuite struct {
// 	suite.Suite
// 	db      *gorm.DB
// 	mock    sqlmock.Sqlmock
// 	storage *Storage
// 	ctx     context.Context
// }

// func (s *StorageTestSuite) SetupSuite() {
// 	// Create SQL mock
// 	var (
// 		db  *sql.DB
// 		err error
// 	)
// 	db, s.mock, err = sqlmock.New()
// 	s.Require().NoError(err)

// 	// Create GORM DB with mock
// 	dialector := postgres.New(postgres.Config{
// 		Conn: db,
// 	})
// 	s.db, err = gorm.Open(dialector, &gorm.Config{})
// 	s.Require().NoError(err)

// 	s.storage = &Storage{db: s.db}
// 	s.ctx = context.Background()
// }

// func TestStorageTestSuite(t *testing.T) {
// 	t.Parallel()
// 	suite.Run(t, new(StorageTestSuite))
// }

// func (s *StorageTestSuite) TestSave() {
// 	testCases := []struct {
// 		name string
// 		user *model.User
// 	}{
// 		{
// 			name: "success",
// 			user: &model.User{
// 				Username: "testuser",
// 				Email:    "test@example.com",
// 			},
// 		},
// 	}

// 	for _, tc := range testCases {
// 		s.Run(tc.name, func() {
// 			savedUser, err := s..Save(s.ctx, tc.user)
// 			s.Require().NoError(err)
// 			s.Equal(tc.user.Username, savedUser.Username)
// 		})
// 	}
// }

// // func (s *StorageTestSuite) TestUserCRUD() {
// // 	userStorage := &AbstractStorage[*model.User]{
// // 		db:        s.db,
// // 		tableName: "users",
// // 	}

// // 	// Mock Save
// // 	s.mock.ExpectBegin()
// // 	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users" ("username","email","id") VALUES ($1,$2,$3) RETURNING "id"`)).
// // 		WithArgs("testuser", "test@example.com", sqlmock.AnyArg()).
// // 		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
// // 	s.mock.ExpectCommit()

// // 	testUser := &model.User{
// // 		Username: "testuser",
// // 		Email:    "test@example.com",
// // 	}
// // 	savedUser, err := userStorage.Save(s.ctx, testUser)
// // 	s.NoError(err)
// // 	s.Equal("testuser", savedUser.Username)

// // 	// Mock FindByID
// // 	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE id = $1`)).
// // 		WithArgs(1).
// // 		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "email"}).
// // 			AddRow(1, "testuser", "test@example.com"))

// // 	foundUser, err := userStorage.FindByID(s.ctx, 1)
// // 	s.NoError(err)
// // 	s.Equal(uint(1), foundUser.ID)
// // 	s.Equal("testuser", foundUser.Username)

// // 	// Mock Update
// // 	s.mock.ExpectBegin()
// // 	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "users" SET "username"=$1,"email"=$2 WHERE "id" = $3`)).
// // 		WithArgs("updated_user", "test@example.com", 1).
// // 		WillReturnResult(sqlmock.NewResult(1, 1))
// // 	s.mock.ExpectCommit()

// // 	foundUser.Username = "updated_user"
// // 	updatedUser, err := userStorage.Update(s.ctx, foundUser.ID, foundUser)
// // 	s.NoError(err)
// // 	s.Equal("updated_user", updatedUser.Username)

// // 	// Mock FindAll
// // 	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users"`)).
// // 		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "email"}).
// // 			AddRow(1, "updated_user", "test@example.com"))

// // 	users, err := userStorage.FindAll(s.ctx)
// // 	s.NoError(err)
// // 	s.Len(users, 1)

// // 	// Ensure all expectations were met
// // 	s.NoError(s.mock.ExpectationsWereMet())
// // }
