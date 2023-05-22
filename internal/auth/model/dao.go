package model

type AuthServiceDAO struct {
	Id         uint64 `db:"id"`
	Name       string `db:"name"`
	PublicKey  string `db:"public_key"`
	PrivateKey string `db:"private_key"`
	URL        string `db:"url"`
}
