package movie

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"github.com/google/uuid"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type Movie struct {
	Id openapi_types.UUID `json:"id" validate:"required"`
	Year int	 `json:"year" validate:"required,min=1800"`
	Title string `json:"title" validate:"required"`
	cert *x509.Certificate 
	privateKey *rsa.PrivateKey
}

type Option func(movie Movie) Movie

func WithTitle(title string) Option {
	return func(movie Movie) Movie {
		m := movie
		m.Title = title
		return m
	}
}

func WithYear(year int) Option {
	return func(movie Movie) Movie {
		m := movie
		m.Year = year
		return m
	}
}

func WithCert(cert *x509.Certificate, privateKey *rsa.PrivateKey) Option {
	return func(movie Movie) Movie {
		m := movie
		m.cert = cert
		m.privateKey = privateKey
		return m
	}
}

func NewWithOptions(opts ...Option) *Movie {
	movie := Movie {
		Id: uuid.New(),
	}

	for _, o := range opts {
		movie = o(movie)
	}

 	return &movie
}

func (m *Movie) GetCert() *x509.Certificate {
	return m.cert
}

func (m *Movie) GetCertString() string {
	return certToPEMString(m.cert)
}

func (m Movie) GetPrivateKey() *rsa.PrivateKey {
	return m.privateKey
}

func certToPEMString(cert *x509.Certificate) string {
    pemBytes := pem.EncodeToMemory(&pem.Block{
        Type:  "CERTIFICATE",
        Bytes: cert.Raw,
    })
    return string(pemBytes)
}
