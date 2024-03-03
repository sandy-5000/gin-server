package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID primitive.ObjectID 	`bson:"_id,omitempty" json:"id,omitempty"`
	Name string 			`bson:"name" json:"name"`
	Category string 		`bson:"category" json:"category"`
	Price int 				`bson:"price" json:"price"`
	Stock int 				`bson:"stock" json:"stock"`
}

type CreateProduct struct {
	Name string 	`bson:"name" json:"name"`
	Category string `bson:"category" json:"category"`
	Price int 		`bson:"price" json:"price"`
	Stock int 		`bson:"stock" json:"stock"`
}

func (p *Product) CreateIndexes(collection *mongo.Collection) error {
    indexModels := []mongo.IndexModel{
        {
            Keys:    bson.D{{"name", 1}},
            Options: options.Index().SetUnique(true),
        },
        {
            Keys:    bson.D{{"category", 1}},
            Options: nil,
        },
        {
            Keys:    bson.D{{"price", 1}, {"stock", 1}},
            Options: nil,                                 
        },
    }

    _, err := collection.Indexes().CreateMany(context.Background(), indexModels)
    return err
}