package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net"
	"os"
	"time"
)

var pem string
var key string

func genCert() {
	ca := &x509.Certificate{
		SerialNumber: big.NewInt(1337),
		Subject: pkix.Name{
			Country:            []string{"Neuland"},
			Organization:       []string{"qwertz"},
			OrganizationalUnit: []string{"qwertz"},
		},
		Issuer: pkix.Name{
			Country:            []string{"Neuland"},
			Organization:       []string{"Skynet"},
			OrganizationalUnit: []string{"Computer Emergency Response Team"},
			Locality:           []string{"Neuland"},
			Province:           []string{"Neuland"},
			StreetAddress:      []string{"Mainstreet 23"},
			PostalCode:         []string{"12345"},
			SerialNumber:       "23",
			CommonName:         "23",
		},
		SignatureAlgorithm:    x509.SHA512WithRSA,
		PublicKeyAlgorithm:    x509.ECDSA,
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(0, 0, 10),
		SubjectKeyId:          []byte{1, 2, 3, 4, 5},
		BasicConstraintsValid: true,
		IsCA:        true,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:    x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
	}

	priv, _ := rsa.GenerateKey(rand.Reader, 4096)
	pub := &priv.PublicKey
	ca_b, err := x509.CreateCertificate(rand.Reader, ca, ca, pub, priv)
	if err != nil {
		log.Fatalf("create cert failed %#v", err)
		return
	}
	log.Println("save", pem)
	ioutil.WriteFile(pem, ca_b, 0644)
	log.Println("save", key)
	ioutil.WriteFile(key, x509.MarshalPKCS1PrivateKey(priv), 0644)
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 512)
	log.Print("https: waiting")
	con, err := conn.Read(buf)
	if err != nil {
		if err != nil {
			log.Printf("https: read: %#v", err)
		}
	}

	log.Printf("https: echo %q\n", string(buf[:con]))

	conn.Write([]byte(time.Now().Format(time.RFC3339) + "\r\n\n"))
	con, err = conn.Write(buf[:con])
	log.Printf("https: wrote %d bytes", con)

	if err != nil {
		log.Printf("https: write: %s", err)
	}
	log.Println("https: closed")
}

func main() {
	pem = "cert.pem"
	key = "cert.key"
	if _, err := os.Stat(pem); os.IsNotExist(err) {
		if _, err := os.Stat(key); os.IsNotExist(err) {
			fmt.Println("no certs found, generating new self signed certs.")
			genCert()
		}
	}
	if _, err := os.Stat(key); err == nil {
		ca_b, _ := ioutil.ReadFile(pem)
		ca, _ := x509.ParseCertificate(ca_b)
		priv_b, _ := ioutil.ReadFile(key)
		priv, _ := x509.ParsePKCS1PrivateKey(priv_b)
		pool := x509.NewCertPool()
		pool.AddCert(ca)

		cert := tls.Certificate{
			Certificate: [][]byte{ca_b},
			PrivateKey:  priv,
		}

		config := tls.Config{
			Certificates: []tls.Certificate{cert},
			//MinVersion:   tls.VersionSSL30, //don't use SSLv3, https://www.openssl.org/~bodo/ssl-poodle.pdf
			MinVersion: tls.VersionTLS10,
			//MinVersion:   tls.VersionTLS11,
			//MinVersion:   tls.VersionTLS12,
		}
		config.Rand = rand.Reader
		port := ":4443"
		listener, err := tls.Listen("tcp", port, &config)
		if err != nil {
			log.Fatalf("https: listen: %s", err)
		}
		log.Printf("https: listening on %s", port)

		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Printf("https: accept: %s", err)
				break
			}
			defer conn.Close()
			log.Printf("https: accepted from %s to %s", conn.RemoteAddr(), port)
			go handleClient(conn)
		}
	} else {
		log.Fatalf("https: NO CERT FOUND")
	}
}
