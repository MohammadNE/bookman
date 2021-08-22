package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mohammadne/bookman/core/failures"
	"github.com/mohammadne/bookman/core/logger"
	"github.com/mohammadne/bookman/user/internal/models"
)

type Database interface {
	CreateUser(user *models.User) failures.Failure
	ReadUserById(id int64) (*models.User, failures.Failure)
	ReadUserByEmailAndPassword(email string, password string) (*models.User, failures.Failure)
	UpdateUser(user *models.User) failures.Failure
	DeleteUser(user *models.User) failures.Failure
}

type mysql struct {
	// injected dependencies
	config *Config
	logger logger.Logger

	// internal dependencies
	connection *sql.DB
}

const (
	driver = "mysql"

	errOpenDatabse = "error in opening mysql database"
	errPingDatabse = "error to ping mysql databse"
)

func NewMysqlDatabase(cfg *Config, log logger.Logger) Database {
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8",
		cfg.Username, cfg.Password, cfg.Host, cfg.Schema,
	)

	client, err := sql.Open(driver, dataSourceName)
	if err != nil {
		log.Fatal(errOpenDatabse, logger.Error(err))
		return nil
	}

	// check client connection (ping it)
	if err = client.Ping(); err != nil {
		log.Fatal(errPingDatabse, logger.Error(err))
		return nil
	}

	return &mysql{config: cfg, logger: log, connection: client}
}
