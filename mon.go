package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type theme struct {
	IsEnabled    bool
	ThemeID      int
	Title        string
	Description  string
	ImageURI     string
	Genre        int
	Choices      []string
	Keywords     []string
	Formula      string
	SaveInterval int
}

type vote struct {
	ThemeID      int
	Answer       int
	UserProvider string
	UserID       string
	CreatedAt    int
	ExpiredAt    int
}

type result struct {
	Timestamp  int
	Percentage []float64
}

type transition struct {
	ShortTransition []result
	LongTransition  []result
}

func connect(uri string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		return nil, err
	}

	return client, nil
}

func getAllThemes(db *mongo.Database) ([]theme, error) {
	collection := db.Collection("themes")

	var themes []theme

	cur, err := collection.Find(context.TODO(), bson.D{{}})

	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var elem theme
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}

		themes = append(themes, elem)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}
	cur.Close(context.TODO())
	return themes, nil
}

func getAllVotes(db *mongo.Database) ([]vote, error) {
	collection := db.Collection("votes")

	var votes []vote

	cur, err := collection.Find(context.TODO(), bson.D{{}})

	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var elem vote
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}

		votes = append(votes, elem)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}
	cur.Close(context.TODO())
	return votes, nil
}
