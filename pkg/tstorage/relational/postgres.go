package relational

import (
	"bank_api/pkg/tstorage/config"
	"fmt"
	_ "github.com/jackc/pgx/stdlib" // pgx driver
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"time"
)

func InitPsqlDB(cfg *config.TStorageConfig) (*sqlx.DB, error) {
	connectionUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Relational.Postgres.Host,
		cfg.Relational.Postgres.Port,
		cfg.Relational.Postgres.User,
		cfg.Relational.Postgres.Password,
		cfg.Relational.Postgres.DBName,
		cfg.Relational.Postgres.SSLMode)
	database, err := sqlx.Connect(cfg.Relational.Postgres.PgDriver, connectionUrl)
	if err != nil {
		return nil, err
	}
	database.DB.SetConnMaxIdleTime(time.Duration(cfg.Relational.Postgres.ConnMaxIdleTime) * time.Second)
	database.DB.SetMaxOpenConns(cfg.Relational.Postgres.MaxOpenConns)
	return database, nil
}
