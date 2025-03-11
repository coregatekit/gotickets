package repositoriese

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nrednav/cuid2"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)

	dialector := postgres.New(postgres.Config{
		Conn:       mockDB,
		DriverName: "postgres",
	})
	db, err := gorm.Open(dialector, &gorm.Config{})
	assert.NoError(t, err)

	return db, mock
}

func TestCreateUser(t *testing.T) {
	db, mock := setupMockDB(t)

	defer func() {
		dbSQL, err := db.DB()
		if err != nil {
			fmt.Println(err)
		}

		dbSQL.Close()
	}()

	repo := &UsersRepository{
		dbConnection: db,
	}

	t.Run("should be create user successfully", func(t *testing.T) {
		id := cuid2.Generate()
		rows := sqlmock.NewRows([]string{"id", "name", "username", "email", "password"}).
			AddRow(id, "John Doe", "johndoe", "john@example.com", "password")
		mock.ExpectQuery(`INSERT INTO "auth"."users" (id, name, username, email, password) VALUES ($1, $2, $3, $4, $5)`).
			WithArgs(id, "John Doe", "johndoe", "john@example.com", "password").
			WillReturnRows(rows)

		user, err := repo.CreateUser("John Doe", "johndoe", "john@example.com", "password")

		assert.NoError(t, err)
		// assert.Equal(t, id, user.ID)
		assert.Equal(t, "John Doe", user.Name)
		assert.Equal(t, "johndoe", user.Username)
		assert.Equal(t, "password", user.Password)
	})
}
