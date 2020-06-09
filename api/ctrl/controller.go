package ctrl

import (
	"crypto/rand"
	"crypto/rsa"
	"log"

	"github.com/adrianosela/harvest"
	"github.com/adrianosela/harvest/api/auth"
	"github.com/adrianosela/harvest/api/config"
	"github.com/adrianosela/harvest/api/store"
)

// Controller manages the game server
type Controller struct {
	// TODO: rooms manager

	games harvest.DB
	auth  *auth.Authenticator
}

// NewController is the Controller constructor
func NewController(conf config.Conf) (*Controller, error) {
	db, err := store.NewMongo(conf.MongoDBConnStr, conf.MongoDBName)
	if err != nil {
		return nil, err
	}

	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}

	for i := 0; i < 5; i++ {
		g := harvest.NewGame()
		g.AddPlayer("MOCK_PLAYER_1")
		g.AddPlayer("MOCK_PLAYER_2")
		g.AddPlayer("MOCK_PLAYER_3")
		g.AddPlayer("MOCK_PLAYER_4")
		db.CreateGame(g)
		log.Printf("created mock game %s", g.ID)
	}

	return &Controller{
		games: db,
		auth:  auth.NewAuthenticator(db, priv, "harvest.adrianosela.com", "api"),
	}, nil
}
