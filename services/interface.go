package services

type FileLocker interface {
	Init(password string) error
	Lock() error
	Unlock(password string) error
}

type Crypt interface {
	Encrypt(creds []byte) []byte
	Decrypt(cred string) ([]byte, error)
}
