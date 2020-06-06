package main

import (
	"github.com/adrianosela/harvest"
	"github.com/adrianosela/harvest/store"
)

// Controller manages the game server
type Controller struct {
	// TODO: authenticator
	// TODO: rooms manager

	store harvest.Store
}

// NewController is the Controller constructor
func NewController(conf Config) (*Controller, error) {
	db, err := store.NewMongo(conf.MongoDBConnStr, conf.MongoDBName)
	if err != nil {
		return nil, err
	}
	return &Controller{
		store: db,
	}, nil
}
