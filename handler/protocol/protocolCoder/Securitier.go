package protocolCoder
type Securitier interface {
	Encrypt(b []byte) ([]byte)
	Decrypt(b []byte) (bool,[]byte)
}