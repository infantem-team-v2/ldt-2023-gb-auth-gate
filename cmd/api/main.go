package main

import (
	"bank_api/internal/pkg/server"
)

// @title CryptoBank API
// @description Core backend specification for B2C app
// @version 1.0.0
// @contact.name Docs developer
// @contact.url https://t.me/sirop_work
// @contact.email bwg.it317@gmail.com

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
	if err := server.NewServer().MapHandlers().Run(); err != nil {
		panic("can't start server")
	}
}
