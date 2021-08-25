package sm

import (
	"crypto/rand"

	"github.com/tjfoc/gmsm/x509"
)

func Sign(pri, msg []byte) ([]byte, error) {
	pem, err := x509.ReadPrivateKeyFromPem(pri, nil)
	if err != nil {
		return nil, err
	}
	sign, err := pem.Sign(rand.Reader, msg, nil)
	if err != nil {
		return nil, err
	}
	return sign, nil
}
