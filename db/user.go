package db

import "time"

type User struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Goals     []*Goal   `json:"goals" gorm:"foreignKey:UserId;references:ID;OnDelete:CASCADE"`
}

func (db *Database) CreateUser() (*User, error) {
	user := User{
		ID: db.GenerateId(),
	}

	err := db.Gorm.Model(&User{}).Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (db *Database) GetUserById(id string) (*User, error) {
	var user User

	err := db.Gorm.Model(&User{}).Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (db *Database) GetUsers() []User {
	var users []User
	db.Gorm.Model(&User{}).Find(&users)
	return users
}
