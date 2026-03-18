package crypto

import (
	"crypto/rand"
	"crypto/rsa"
)

const KeySize = 2048

type KeyPair struct {
	Private *rsa.PrivateKey
	Public  *rsa.PublicKey
}

func GenerateKeyPair() (KeyPair, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, KeySize)
	if err != nil {
		return KeyPair{}, err
	}

	return KeyPair{
		privateKey,
		&privateKey.PublicKey,
	}, nil
}
