package service

import (
	"Book-Service/database"
	"Book-Service/helper"
	"Book-Service/models"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var NewUserConnection = database.NewMongoCollection("users-ega")

type UserService struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewUserService() *UserService {
	return &UserService{
		collection: NewUserConnection.GetCollection(),
		ctx:        context.Background(),
	}
}

func (us *UserService) countUser(email string) (int64, error) {
	count, err := us.collection.CountDocuments(us.ctx, bson.M{"email": email})
	return count, err
}

func (us *UserService) SingIn(user models.User) string {
	count, err := us.countUser(user.Email)
	if err != nil || count > 0 {
		return "Failed Insert, User Already Exist"
	}

	password := helper.HashPassword(user.Password)
	user.Role = "user"
	user.Disable = false
	user.ID = primitive.NewObjectID().Hex()
	user.Created_at = time.Now().Local().Unix()
	user.Updated_at = time.Now().Local().Unix()
	user.Password = password

	_, err = us.collection.InsertOne(us.ctx, user)
	if err != nil {
		return "Failed Insert"
	}

	return ""
}

func (us *UserService) FindUser(email string) (models.User, error) {
	var User models.User
	err := us.collection.FindOne(us.ctx, bson.M{"email": email}).Decode(&User)
	return User, err
}

func (us *UserService) Token(email string) (models.Token, error) {
	var tokens models.Token
	token, refreshToken, err := helper.GenereteToken(email)
	tokens.Token = token
	tokens.RefreshToken = refreshToken

	return tokens, err
}

func (us *UserService) Update(user models.User) error {
	user.Updated_at = time.Now().Local().Unix()
	_, err := us.collection.UpdateOne(us.ctx, bson.M{"email": user.Email}, bson.M{"$set": user})
	if err != nil {
		return err
	}
	return nil
}

func (us *UserService) UpdateProfile(user models.User, email string) (models.User, error) {
	oldData, _ := us.FindUser(email)

	fmt.Println(oldData)
	user.Password = oldData.Password
	user.Created_at = oldData.Created_at
	user.Disable = oldData.Disable
	user.ID = oldData.ID
	user.Email = oldData.Email
	user.Updated_at = time.Now().Local().Unix()
	user.Role = oldData.Role
	fmt.Println(user)
	_, err := us.collection.UpdateOne(us.ctx, bson.M{"email": user.Email}, bson.M{"$set": user})
	return user, err
}
