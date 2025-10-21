package character

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"movies/internal/entity/movie"

	"github.com/google/uuid"
)

type Character struct {
	Id uuid.UUID 		`json:"id" validate:"required"`
	Name string			`json:"name" validate:"required"`
	Movie movie.Movie	`json:"movie" validate:"required"`
	cert *x509.Certificate 
	privateKey *rsa.PrivateKey
}

type Option func(movie Character) Character

func WithName(name string) Option {
	return func(character Character) Character {
		c := character
		c.Name = name
		return c
	}
}

func WithMovie(movie movie.Movie) Option {
	return func(character Character) Character {
		c := character
		c.Movie = movie
		return c
	}
}

func WithCert(cert *x509.Certificate, privateKey *rsa.PrivateKey) Option {
	return func(character Character) Character {
		c := character
		c.cert = cert
		c.privateKey = privateKey
		return c
	}
}

func NewWithOptions(opts ...Option) *Character {
	character := Character {
		Id: uuid.New(),
	}

	for _, o := range opts {
		character = o(character)
	}

 	return &character
}

func (c *Character) GetCertString() string {
	return certToPEMString(c.cert)
}

func certToPEMString(cert *x509.Certificate) string {
    pemBytes := pem.EncodeToMemory(
		&pem.Block{
			Type:  "CERTIFICATE",
			Bytes: cert.Raw,
		},
	)
    return string(pemBytes)
}
