package tsecure

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"github.com/sarulabs/di"
	"os"
)

type RsaCrypto struct {
	PublicKey  *rsa.PublicKey
	PrivateKey *rsa.PrivateKey
}

func openRsaFile(path string) (key []byte, err error) {
	workDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	pemFile, err := os.Open(fmt.Sprintf("%s%s", workDir, path))
	if err != nil {
		return nil, err
	}
	defer pemFile.Close()
	_, err = pemFile.Read(key)
	return key, err
}

func parseRsaPublicKey(key []byte) (publicKey *rsa.PublicKey, err error) {
	return x509.ParsePKCS1PublicKey(key)
}

func parseRsaPrivateKey(key []byte) (privateKey *rsa.PrivateKey, err error) {
	return x509.ParsePKCS1PrivateKey(key)
}

func BuildRsaCrypto(ctn di.Container) (interface{}, error) {

	return nil, nil
}

func (r *RsaCrypto) Encrypt(message string) (cipher string, err error) {
	cipherRaw, err := rsa.EncryptOAEP(hashesMap[SHA256](), rand.Reader, r.PublicKey, []byte(message), nil)
	return string(cipherRaw), err
}

func (r *RsaCrypto) Decrypt(cipher string) (message string, err error) {
	//TODO implement me
	panic("implement me")
}
