package tsecure

// Encryptor Basic interface for structures which encrypt and decrypt messages
type Encryptor interface {
	Encrypt(message string) (cipher string)
	Decrypt(cipher string) (message string, err error)
}
