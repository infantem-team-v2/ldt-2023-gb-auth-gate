package relational

import (
	"bank_api/pkg/tstorage/config"
	"fmt"
	"github.com/fiorix/go-redis/redis"
	_ "github.com/jackc/pgx/stdlib" // pgx driver
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"time"
)

//type Comment struct {
//	Id              int64    `db:"id"`
//	UserId          int      `db:"user_id"`
//	CommentId       null.Int `db:"comment_id"`
//	Content         string   `db:"content"`
//	Level           int
//	VoteCount       int     `db:"voteCount"`
//	UpdatedAt       float64 `db:"updated_at"`
//	UpdatedAtNormal string
//}

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
	database.DB.SetConnMaxIdleTime(time.Duration(cfg.Relational.Postgres.ConnMaxIdleTime))
	database.DB.SetMaxOpenConns(cfg.Relational.Postgres.MaxOpenConns)
	return database, nil
}

func InitRedis(cfg *config.TStorageConfig) (*redis.Client, error) {
	url := fmt.Sprintf(
		"%s:%s db=%d passwd=%s",
		cfg.Cache.Redis.Host,
		cfg.Cache.Redis.Port,
		cfg.Cache.Redis.DB,
		cfg.Cache.Redis.Password)
	client, err := redis.NewClient(url)
	if err != nil {
		return nil, err
	}
	return client, nil
}
