// Package dbconn contains all the database connection related functionalities.
package dbconn

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/udayangaac/wallet-api/cfgloader"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// postgresFlags environ
var postgresFlags = struct {
	Host     string
	Port     string
	User     string
	Password string
	DB       string
	// 1 : silent log level,
	// 2 : error log level
	// 3 : warn log level
	// 4 : info log level
	LogLevel    string
	AutoMigrate string // Y or y for enable the auto migrate.
}{
	Host:        cfgloader.LookupEnv("POSTGRES_HOST", "localhost"),
	Port:        cfgloader.LookupEnv("POSTGRES_PORT", "5432"),
	User:        cfgloader.LookupEnv("POSTGRES_USER", "xcore"),
	Password:    cfgloader.LookupEnv("POSTGRES_PASS", "corex"),
	DB:          cfgloader.LookupEnv("POSTGRES_DB", "xcore_main"),
	AutoMigrate: cfgloader.LookupEnv("AUTO_MIGRATE", "Y"),
}

type pgDBConnector struct {
	db *gorm.DB
}

// NewPGConnecter create an instance of db connecter.
func NewPGConnecter() (DBConnector, error) {

	// create the data source name (DSN).
	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=UTC",
		postgresFlags.Host,
		postgresFlags.User,
		postgresFlags.Password,
		postgresFlags.DB,
		postgresFlags.Port,
	)

	dialector := postgres.New(postgres.Config{DSN: dsn, PreferSimpleProtocol: true})

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.Logger = logger.Default.LogMode(getLogLevel())
	return &pgDBConnector{db: db}, nil
}

// AddEntities add entities to the database
func (c *pgDBConnector) AddEntities(entities ...interface{}) error {
	if strings.ToLower(postgresFlags.AutoMigrate) == "y" {
		return c.db.AutoMigrate(entities...)
	}
	return nil
}

// GetDB return initialized database.
func (c *pgDBConnector) GetDB() *gorm.DB {
	return c.db
}

func getLogLevel() logger.LogLevel {
	i, err := strconv.Atoi(postgresFlags.LogLevel)
	if err != nil || (i > 0 && i < 5) {
		return logger.Error
	}
	return logger.LogLevel(i)
}
