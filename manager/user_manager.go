package manager

import (
	"crud-api/models"
	"crud-api/mongoDatabase"
	"crud-api/request"
)

func CreateUser(req request.CreateUserRequest) (*models.User, error) {
	return mongoDatabase.CreateUserInDB(req)
}

func GetUser(id string) (*models.User, error) {
	return mongoDatabase.GetUserFromDB(id)
}

// func UpdateUser(id string, req request.CreateUserRequest) (*models.User, error) {
// 	update := bson.M{
// 		"$set": bson.M{
// 			"name":  req.Name,
// 			"email": req.Email,
// 		},
// 	}
// 	if err := mongoDatabase.UpdateUserInDB(id, update); err != nil {
// 		return nil, err
// 	}
// 	return GetUser(id)
// }
func UpdateUser(id string, req request.CreateUserRequest) (*models.User, error) {
	if err := mongoDatabase.UpdateUserInDB(id, req); err != nil {
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

