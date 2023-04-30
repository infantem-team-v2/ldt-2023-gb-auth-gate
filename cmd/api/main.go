package main

import (
	"bank_api/internal/pkg/server"
	"fmt"
)

// @title GO Fiber Structure Template
// @description Simple structure template for web-application w/ Fiber, DI, Postgres, Redis etc.
// @version 1.0.0
// @contact.name Docs developer
// @contact.url https://t.me/KlenoviySirop
// @contact.email KlenoviySir@yandex.ru

// @host bank.onecrypto.pro
// @schemes https

// @securityDefinitions ApiPublic
// @in header
// @name ApiPublic
// @description Api public key

// @securityDefinitions Signature
// @in header
// @name Signature
// @description Signature calculated by the private key

// @securityDefinitions TimeStamp
// @in header
// @name TimeStamp
// @description TimeStamp
func main() {
	if err := server.
		NewServer().
		MapHandlers().
		Run(); err != nil {
		panic(fmt.Sprintf("can't start server %+v", err))
	}
}
