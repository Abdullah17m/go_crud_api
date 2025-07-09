package manager

import (
	"crud-api/models"
	"crud-api/mongoDatabase"
	"crud-api/request"
	"time"
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateUser(req request.CreateUserRequest) (*models.User, error) {
	user := &models.User{
		ID:        uuid.New().String(),
		Name:      req.Name,
		Email:     req.Email,
		CreatedAt: time.Now(),
	}
	err := mongoDatabase.CreateUserInDB(user)
	return user, err
}

func GetUser(id string) (*models.User, error) {
	return mongoDatabase.GetUserFromDB(id)
}

func UpdateUser(id string, req request.CreateUserRequest) (*models.User, error) {
	update := bson.M{
		"$set": bson.M{
			"name":  req.Name,
			"email": req.Email,
		},
	}
	if err := mongoDatabase.UpdateUserInDB(id, update); err != nil {
		return nil, err
	}
	return GetUser(id)
}

func DeleteUser(id string) error {
	return mongoDatabase.DeleteUserFromDB(id)
}


func GetAllUsers(page int64, limit int64) ([]models.User, error) {
	return mongoDatabase.GetAllUsersFromDB(page, limit)
}

func IsEmailExists(email string) (bool, error) {
	collection := mongoDatabase.UserCollection()

	filter := bson.M{"email": email}

	count, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}