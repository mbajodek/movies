package cert

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"math/big"
	"os"
	"time"
)

var certFileEnv = "CERT_MOVIES"
var privateKeyFileEnv = "PRIV_MOVIES"

type CertGenerator struct {
	cert       *x509.Certificate
	privateKey *rsa.PrivateKey
}

func New() *CertGenerator {
	keyPair, err := tls.LoadX509KeyPair(os.Getenv(certFileEnv), os.Getenv(privateKeyFileEnv))

	if err != nil {
		fmt.Println("LoadX509KeyPair error:", err)
	}

	cert, err := x509.ParseCertificate(keyPair.Certificate[0])

	if err != nil {
		fmt.Println("X509 parse certificate error:", err)
	}

	return &CertGenerator{
		cert:       cert,
		privateKey: keyPair.PrivateKey.(*rsa.PrivateKey),
	}
}

func (cg *CertGenerator) GenerateMovieCert() (*x509.Certificate, *rsa.PrivateKey, error) {
	return cg.generateCert(cg.cert, cg.privateKey)
}

func (cg *CertGenerator) GenerateCharacterCert(cert *x509.Certificate, privateKey *rsa.PrivateKey) (*x509.Certificate, *rsa.PrivateKey, error) {
	return cg.generateCert(cert, privateKey)
}

func (cg *CertGenerator) generateCert(parentCert *x509.Certificate, privateKey *rsa.PrivateKey) (*x509.Certificate, *rsa.PrivateKey, error) {
	template := &x509.Certificate{
		SerialNumber: big.NewInt(24),
		Subject:      pkix.Name{CommonName: "Cert"},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(1, 0, 0),
		KeyUsage:     x509.KeyUsageCertSign,
		IsCA:         true,
	}

	priv, _ := rsa.GenerateKey(rand.Reader, 2048)
	pub := &priv.PublicKey

	certBytes, err := x509.CreateCertificate(rand.Reader, template, parentCert, pub, privateKey)

	if err != nil {
		fmt.Println("Create certificate error:", err)
		return nil, nil, err
	}

	cert, err := x509.ParseCertificate(certBytes)
	if err != nil {
		fmt.Println("Parse certificate error:", err)
		return nil, nil, err
	}

	return cert, priv, nil
}
