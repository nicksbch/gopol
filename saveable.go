package main

import "crypto/rsa"

type Saveable interface {
	Encrypt(rsaPublicKey *rsa.PublicKey) ([]byte, error)
}
