package user

import (
	"github.com/jinzhu/gorm"
)

// type UsersStorage []interface{}

// var users UsersStorage

type Repository interface {
	InsertUser(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) InsertUser(user User) (User, error) {
	// users = append(users, user)

	// fmt.Println(users)
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
