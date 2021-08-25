package sm

import (
	"github.com/tjfoc/gmsm/x509"
)

func Verify(pub, sign, msg []byte) (bool, error) {
	pem, err := x509.ReadPublicKeyFromPem(pub)
	if err != nil {
		return false, err
	}
	verify := pem.Verify(msg, sign)
	return verify, nil
}
