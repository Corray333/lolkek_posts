package posts

import (
	"context"
	"errors"
	"fmt"
	db "posts_ms/internal/database"
	"posts_ms/internal/database/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func autrorize() bool {
	return true
}

// Create post function
func CreatePost(ctx *fiber.Ctx) error {

	// Post object filling
	newPost := new(models.Post)
	ctx.BodyParser(newPost)
	if !autrorize() {
		return errors.New("User is not authorized!")
	}
	newPost.Id = primitive.NewObjectID()
	newPost.Comments = []primitive.ObjectID{}

	// Push post to db
	_, err := db.MG.Db.Collection("posts").InsertOne(context.TODO(), newPost)
	if err != nil {
		return err
	}

	// Save file on server
	file, err := ctx.FormFile("img")
	if err != nil {
		return err
	}
	fmt.Println(file.Filename)
	ctx.SaveFile(file, "../images/"+file.Filename)
	return nil
}
