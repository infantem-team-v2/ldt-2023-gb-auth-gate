package repository

import (
	"gb-auth-gate/internal/auth/model"
	"gb-auth-gate/pkg/terrors"
	"github.com/jmoiron/sqlx"
	"github.com/sarulabs/di"
)

type PostgresRepository struct {
	db *sqlx.DB `di:"postgres"`
}

func BuildPostgresRepository(ctn di.Container) (interface{}, error) {
	return &PostgresRepository{
		db: ctn.Get("postgres").(*sqlx.DB),
	}, nil
}

func (p PostgresRepository) FindServiceByPublicKey(publicKey string) (*model.AuthServiceDAO, error) {
	query := `
			 SELECT * FROM service.auth WHERE public_key = $1;
			 `
	var data model.AuthServiceDAO
	err := p.db.Get(&data, query, publicKey)
	if err != nil {
		return nil, terrors.Raise(err, 100002)
	}

	return &data, nil
}

func (p PostgresRepository) FindServiceByName(name string) (*model.AuthServiceDAO, error) {
	query := `
			 SELECT * FROM service.auth WHERE name = $1;
			 `
	var data model.AuthServiceDAO
	err := p.db.Get(&data, query, name)
	if err != nil {
		return nil, terrors.Raise(err, 200002)
	}

	return &data, nil
}
