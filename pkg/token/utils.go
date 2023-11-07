package token

import (
	"crypto/ed25519"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
)

func GetKeyBytes(key string) []byte {
	data, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil
	}
	return data
}

func GetRSAPrivateKey(data []byte) *rsa.PrivateKey {
	block, _ := pem.Decode(data)
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil
	}
	return key
}

func GetEdDSAPrivateKey(data []byte) ed25519.PrivateKey {
	block, _ := pem.Decode(data)
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil
	}
	return key.(ed25519.PrivateKey)
}
