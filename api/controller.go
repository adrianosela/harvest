package main

import (
	"crypto/rand"
	"crypto/rsa"

	"github.com/adrianosela/harvest/api/auth"
	"github.com/adrianosela/harvest/api/store"
)

// Controller manages the game server
type Controller struct {
	// TODO: authenticator
	// TODO: rooms manager

	store store.Store
	auth  *auth.Authenticator
}

// NewController is the Controller constructor
func NewController(conf Config) (*Controller, error) {
	db, err := store.NewMongo(conf.MongoDBConnStr, conf.MongoDBName)
	if err != nil {
		return nil, err
	}

	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}

	return &Controller{
		store: db,
		auth:  auth.NewAuthenticator(db, priv, "harvest.adrianosela.com", "api"),
	}, nil
}
