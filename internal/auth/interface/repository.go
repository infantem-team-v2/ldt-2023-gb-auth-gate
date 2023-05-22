package authInterface

import "gb-auth-gate/internal/auth/model"

type RelationalRepository interface {
	FindServiceByPublicKey(publicKey string) (*model.AuthServiceDAO, error)
	FindServiceByName(publicKey string) (*model.AuthServiceDAO, error)
}

type CacheRepository interface {
}
