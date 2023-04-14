package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	ContentType string               `bson:"type" json:"type"`
	ContentPath string               `bson:"path" json:"path"`
	Desc        string               `bson:"desc" json:"desc"`
	Comments    []primitive.ObjectID `bson:"comments" json:"comments"`
	Author      string               `bson:"author" json:"author"`
	Id          primitive.ObjectID   `bson:"_id" json:"id"`
}

// type Content struct {

// }

type Comment struct {
	Id     primitive.ObjectID
	Text   string
	Author primitive.ObjectID
}
