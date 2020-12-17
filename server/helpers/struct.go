package helpers

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name   string             `bson:"name,omitempty" json:"name,omitempty"`
	Author string             `bson:"author,omitempty" json:"author,omitempty"`
	Pages  int                `bson:"pages,omitempty" json:"pages,omitempty"`
	ImgUrl string             `bson:"imgUrl,omitempty" json:"imgUrl,omitempty"`
}

type Result struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    []Book `json:"data"`
}
