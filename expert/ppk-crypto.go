package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"fmt"
	"io"
	"math/big"
	"os"
)

func genPPKeys(random io.Reader) (private_key_bytes, public_key_bytes []byte) {
	private_key, _ := ecdsa.GenerateKey(elliptic.P224(), random)
	private_key_bytes, _ = x509.MarshalECPrivateKey(private_key)
	public_key_bytes, _ = x509.MarshalPKIXPublicKey(&private_key.PublicKey)
	return private_key_bytes, public_key_bytes
}

func pkSign(hash []byte, private_key_bytes []byte) (r, s *big.Int, err error) {
	zero := big.NewInt(0)
	private_key, err := x509.ParseECPrivateKey(private_key_bytes)
	if err != nil {
		return zero, zero, err
	}

	r, s, err = ecdsa.Sign(rand.Reader, private_key, hash)
	if err != nil {
		return zero, zero, err
	}
	return r, s, nil
}

func pkVerify(hash []byte, public_key_bytes []byte, r *big.Int, s *big.Int) (result bool) {
	public_key, err := x509.ParsePKIXPublicKey(public_key_bytes)
	if err != nil {
		return false
	}

	switch public_key := public_key.(type) {
	case *ecdsa.PublicKey:
		return ecdsa.Verify(public_key, hash, r, s)
	default:
		return false
	}
}

func main() {
	var str []byte
	if len(os.Args) > 1 {
		str = []byte(os.Args[1])
	} else {
		str = []byte("Lorem Ipsum dolor sit Amet")
	}

	private_key, public_key := genPPKeys(rand.Reader)
	fmt.Print("private key: ")
	fmt.Println(private_key)
	fmt.Println()
	fmt.Print("public key: ")
	fmt.Println(public_key)
	fmt.Println()

	r, s, err := pkSign(str, private_key)
	if err != nil {
		fmt.Printf("signing hash error: %s\n", err)
	}

	verify := pkVerify(str, public_key, r, s)
	fmt.Printf("signature verification result: %t\n", verify)

	verify = pkVerify([]byte("some other text"), public_key, r, s)
	fmt.Printf("signature verification result: %t\n", verify)
}
