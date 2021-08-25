package sm

import (
	"crypto/rand"

	"github.com/tjfoc/gmsm/x509"
)

func Encrypt(pub, msg []byte) ([]byte, error) {
	pem, err := x509.ReadPublicKeyFromPem(pub)
	if err != nil {
		return nil, err
	}
	cipher, err := pem.EncryptAsn1(msg, rand.Reader)
	if err != nil {
		return nil, err
	}
	return cipher, nil
}
