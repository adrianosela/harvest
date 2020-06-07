package auth

import (
	"crypto/rsa"

	"github.com/adrianosela/harvest/api/store"
)

// Authenticator is the module in charge of authentication
type Authenticator struct {
	db     store.Store
	signer *rsa.PrivateKey
	iss    string
	aud    string
}

// NewAuthenticator is the Authenticator constructor
func NewAuthenticator(db store.Store, key *rsa.PrivateKey, aud, iss string) *Authenticator {
	a := &Authenticator{
		db:     db,
		signer: key,
		aud:    aud,
		iss:    iss,
	}
	return a
}
