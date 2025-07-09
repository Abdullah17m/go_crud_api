package mongoDatabase

import (
	"context"
	"crud-api/models"
	"crud-api/request"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Create
func CreateUserInDB(req request.CreateUserRequest) (*models.User, error) {
	user := &models.User{
		ID:        uuid.New().String(),
		Name:      req.Name,
		Email:     req.Email,
		CreatedAt: time.Now(),
	}

	_, err := models.UserCollection().InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Read
func GetUserFromDB(id string) (*models.User, error) {
	var user models.User
	err := models.UserCollection().FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)
	return &user, err
}

// Update
// func UpdateUserInDB(id string, update bson.M) error {
// 	_, err := models.UserCollection().UpdateByID(context.TODO(), id, update)
// 	return err
// }

func UpdateUserInDB(id string, req request.CreateUserRequest) error {
	update := bson.M{
		"$set": bson.M{
			"name":  req.Name,
			"email": req.Email,
		},
	}
	_, err := models.UserCollection().UpdateByID(context.TODO(), id, update)
	return err
}

func GetAllUsersFromDB(page int64, limit int64) ([]models.User, error) {
	skip := (page - 1) * limit

	opts := options.Find().
		SetSkip(skip).
		SetLimit(limit)

	cursor, err := models.UserCollection().Find(context.TODO(), bson.M{}, opts)
	if err != nil {
		return nil, err
	}

	var users []models.User
	if err := cursor.All(context.TODO(), &users); err != nil {
		return nil, err
	}

	return users, nil
}
// Delete
func DeleteUserFromDB(id string) error {
	_, err := models.UserCollection().DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}

func IsEmailExists(email string) (bool, error) {
	collection := models.UserCollection()

	filter := bson.M{"email": email}

	count, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}