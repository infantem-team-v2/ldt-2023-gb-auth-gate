package tsecure

import "github.com/fernet/fernet-go"

type FernetEncryptor struct {
	EncryptionKey *fernet.Key
}

func NewFernetEncryptor() {

}

func (fe *FernetEncryptor) Encrypt(message string) (cipher string) {
	return ""
}

func (fe *FernetEncryptor) Decrypt(cipher string) (message string, err error) {
	return "", nil
}
