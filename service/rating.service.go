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

var RatingConnection = database.NewMongoCollection("rating-ega")

type RatingService struct {
	collection    *mongo.Collection
	bookColection *mongo.Collection
	ctx           context.Context
}

func NewRatingService() *RatingService {
	return &RatingService{
		collection:    RatingConnection.GetCollection(),
		bookColection: mongoConnection.GetCollection(),
		ctx:           context.Background(),
	}
}

func (ra *RatingService) AddRating(book_id string, rating models.Rating, user models.User) (models.Rating, error) {
	rating.Author = user.Email
	rating.Created_at = time.Now().Local().Unix()
	rating.Updated_at = time.Now().Local().Unix()
	rating.ID = primitive.NewObjectID().Hex()
	_, err := ra.collection.InsertOne(ra.ctx, rating)
	_, udateErr := ra.bookColection.UpdateOne(ra.ctx, bson.M{"_id": book_id}, bson.M{"$push": bson.M{"rating_id": rating.ID}})
	if udateErr != nil {
		return models.Rating{}, err
	}
	return rating, err
}

func (ra *RatingService) GetRating(bookId string, rating_ids []string, ratings interface{}) (interface{}, error) {
	cursor, err := ra.collection.Find(ra.ctx, bson.M{"_id": bson.M{"$in": rating_ids}})
	cursor.All(ra.ctx, ratings)
	return ratings, err
}

func (ra *RatingService) Update(id string, rating models.Rating) (models.Rating, error) {
	var OldRating models.Rating
	err := ra.collection.FindOne(ra.ctx, bson.M{"_id": id}).Decode(&OldRating)
	if err != nil {
		return models.Rating{}, err
	}

	rating.Author = OldRating.Author
	rating.Rate = OldRating.Rate
	rating.ID = OldRating.ID
	rating.Created_at = OldRating.Created_at
	rating.Updated_at = time.Now().Local().Unix()

	_, err = ra.collection.UpdateOne(ra.ctx, bson.M{"_id": id}, bson.M{"$set": rating})
	return rating, err
}
