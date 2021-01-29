package services

import (
	"github.com/halimakibb/bookstore_users-api/domain/users"
	"github.com/halimakibb/bookstore_users-api/utils/errors"
)

// GetUser from database
func GetUser(userID int64) (*users.User, *errors.RestErr) {
	result := users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateUser service
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}
