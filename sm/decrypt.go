package sm

import (
	"github.com/tjfoc/gmsm/x509"
)

func Decrypt(pri, cipher []byte) ([]byte, error) {
	pem, err := x509.ReadPrivateKeyFromPem(pri, nil)
	if err != nil {
		return nil, err
	}
	msg, err := pem.DecryptAsn1(cipher)
	if err != nil {
		return nil, err
	}
	return msg, nil
}
