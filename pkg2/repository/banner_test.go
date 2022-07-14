package repository

import (
	"database/sql"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"

	"github.com/DATA-DOG/go-sqlmock"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

var (
	fakerTime = time.Now()
)

type Suite struct {
	suite.Suite
	DB         *gorm.DB
	mock       sqlmock.Sqlmock
	repository RepositoryBanner
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	s.DB, err = gorm.Open(dialector, &gorm.Config{
		SkipDefaultTransaction: true,
	})
	require.NoError(s.T(), err)
	// require.NotNil(s.T(), s.DB)

	s.repository = NewRepositoryBanner(s.DB)
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestSaveBanner() {
	var (
		id = uuid.NewV4()
	)

	s.mock.ExpectExec(regexp.QuoteMeta(
		`INSERT INTO "banners"`)).
		// WithArgs(1).
		WithArgs(id, 1, 1, "123-gambar.png", 1, "https://google.com", 1234, "indra", fakerTime, 1234, "indra", fakerTime, 1234, "indra", fakerTime).
		WillReturnResult(
			sqlmock.NewResult(1, 1))

	err := s.repository.SaveBanner2(id, fakerTime)

	fmt.Println("ERROR::", err)
	require.NoError(s.T(), err)
}
