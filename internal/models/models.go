package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	Img      Content
	Desc     string
	Comments []primitive.ObjectID
	Author   string
	Id       primitive.ObjectID `bson:"_id" json:"id"`
}

type Content struct {
	Type string
	Path string
}

type Comment struct {
	Id     primitive.ObjectID
	Text   string
	Author primitive.ObjectID
}
