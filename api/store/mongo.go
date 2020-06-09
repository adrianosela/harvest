package store

import (
	"context"

	"github.com/adrianosela/harvest"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mongoCollectionsPrimaryKey = "_id"
	mongoGamesCollectionName   = "games"
)

// Mongo implements the Store
// interface in mongodb
type Mongo struct {
	games *mongo.Collection
}

// NewMongo is the constructor for a Mongo type Store
//
// the format of the mongo connection string is:
// mongodb://<user>:<pass>@<url>:<port>/<dbname>
func NewMongo(connStr, db string) (*Mongo, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connStr))
	if err != nil {
		return nil, err
	}

	// liveliness check
	if err = client.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}

	return &Mongo{
		games: client.Database(db).Collection(mongoGamesCollectionName),
	}, nil
}

// Create writes a new game to the Store
func (m *Mongo) Create(game *harvest.Game) error {
	_, err := m.games.InsertOne(context.TODO(), game)
	return err
}

// Read reads a game from the Store
func (m *Mongo) Read(gameID string) (*harvest.Game, error) {
	var game harvest.Game
	err := m.games.FindOne(context.TODO(), querySingle(gameID)).Decode(&game)
	return &game, err
}

// Watch gets a change stream for a particular game
func (m *Mongo) Watch(gameID string) (*mongo.ChangeStream, error) {
	return m.games.Watch(context.TODO(), pipelineWatchGame(gameID))
}

// Update updates a game in the Store
func (m *Mongo) Update(game *harvest.Game) error {
	update := bson.M{
		"$set": bson.M{
			"players": game.Players,
			"stack":   game.Stack,
			"rejects": game.Rejects,
			"ongoing": game.Ongoing,
			"turn":    game.Turn,
			"round":   game.Round,
		},
	}
	_, err := m.games.UpdateOne(context.TODO(), querySingle(game.ID), update)
	return err
}

// Delete deletes a game from the Store
func (m *Mongo) Delete(gameID string) error {
	_, err := m.games.DeleteOne(context.TODO(), querySingle(gameID))
	return err
}

func querySingle(id string) bson.D {
	return bson.D{{Key: mongoCollectionsPrimaryKey, Value: id}}
}

func pipelineWatchGame(gameID string) []bson.M {
	return []bson.M{{"$match": bson.M{"_id": gameID}}}
}