package service

import (
	"Book-Service/database"
	"Book-Service/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var mongoConnection = database.NewMongoCollection("books-ega")

type BookService struct {
	collection       *mongo.Collection
	ratingCollection *mongo.Collection
	ctx              context.Context
}

func NewBookService() *BookService {
	return &BookService{
		collection:       mongoConnection.GetCollection(),
		ratingCollection: RatingConnection.GetCollection(),
		ctx:              context.Background(),
	}
}

func (se *BookService) Create(book models.Books, user models.User) (models.Books, error) {
	book.Author = user.Email
	book.ID = primitive.NewObjectID().Hex()
	book.Rating_id = []string{}
	book.Create_at = time.Now().Local().Unix()
	book.Update_at = time.Now().Local().Unix()
	_, err := se.collection.InsertOne(se.ctx, book)
	return book, err
}

func (se *BookService) FindAll(book interface{}) (interface{}, error) {
	cursor, err := se.collection.Find(se.ctx, bson.M{})
	cursor.All(se.ctx, book)
	return book, err
}

func (se *BookService) FindByID(id string) (models.Books, error) {
	var book models.Books
	err := se.collection.FindOne(se.ctx, bson.M{"_id": id}).Decode(&book)
	return book, err
}

func (se *BookService) Update(book models.Books, id string) (models.Books, error) {
	oldbook, _ := se.FindByID(id)

	book.Author = oldbook.Author
	book.ID = oldbook.ID
	book.Update_at = time.Now().Local().Unix()
	book.Create_at = oldbook.Create_at
	book.Rating_id = oldbook.Rating_id

	_, err := se.collection.UpdateOne(se.ctx, bson.M{"_id": id}, bson.M{"$set": book})
	return book, err
}

func (se *BookService) Delete(id string, rating_ids []string) error {
	_, err := se.collection.DeleteOne(se.ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	_, err = se.ratingCollection.DeleteMany(se.ctx, bson.M{"_id": bson.M{"$in": rating_ids}})
	if err != nil {
		return err
	}
	return nil
}
