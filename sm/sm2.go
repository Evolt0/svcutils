package sm

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/x509"
)

func Sm2KeyPair() (*KeyPair, error) {
	private, err := sm2.GenerateKey(rand.Reader) // 生成密钥对
	if err != nil {
		return nil, err
	}
	public := &private.PublicKey
	privatePem, err := x509.WritePrivateKeyToPem(private, nil)
	if err != nil {
		return nil, err
	}
	publicPem, err := x509.WritePublicKeyToPem(public)
	if err != nil {
		return nil, err
	}
	privateBase := base64.StdEncoding.EncodeToString(privatePem)
	publicBase := base64.StdEncoding.EncodeToString(publicPem)
	return &KeyPair{
		PubKey:  publicBase,
		PrivKey: privateBase,
	}, nil
}

type KeyPair struct {
	// 公钥（base64编码）
	PubKey string `json:"pubKey"`
	// 私钥（base64编码）
	PrivKey string `json:"privKey"`
}