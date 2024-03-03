package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name string `bson:"name" json:"name"`
	Category string `bson:"category" json:"category"`
	Price int `bson:"price" json:"price"`
	Stock int `bson:"stock" json:"stock"`
}

type CreateProduct struct {
	Name string `bson:"name" json:"name"`
	Category string `bson:"category" json:"category"`
	Price int `bson:"price" json:"price"`
	Stock int `bson:"stock" json:"stock"`
}
