package database

import (
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib" // load pgx driver for PostgreSQL
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func PostgreSQLConnection() (*sqlx.DB, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file, %w", err)
	}

	maxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	maxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	maxLifetimeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))

	db, err := sqlx.Connect("pgx", os.Getenv("DB_SERVER_URL"))
	if err != nil {
		return nil, fmt.Errorf("error, not connect to database, %w", err)
	}

	db.SetMaxOpenConns(maxConn)                           // the default is 0 (unlimited)
	db.SetMaxIdleConns(maxIdleConn)                       // defaultMaxIdleConns = 2
	db.SetConnMaxLifetime(time.Duration(maxLifetimeConn)) // 0, connections are reused forever

	if err := db.Ping(); err != nil {
		defer db.Close()
		return nil, fmt.Errorf("error, not sent ping to database, %w", err)
	}

	return db, nil

}
