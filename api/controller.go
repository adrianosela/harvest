package main

import (
	"github.com/adrianosela/harvest"
	"github.com/adrianosela/harvest/store"
)

// Controller manages the game server
type Controller struct {
	store harvest.Store
}

// NewController is the Controller constructor
func NewController() *Controller {
	return &Controller{
		store: store.NewMock(),
	}
}

// TODO
