package db

import "time"

type User struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Goals     []*Goal   `json:"goals" gorm:"foreignKey:UserId;references:ID;OnDelete:CASCADE"`
}

func (db *Database) AddUser() (*User, error) {
	user := User{
		ID: db.GenerateId(),
	}
	result := db.Gorm.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (db *Database) GetUserById(id string) (*User, error) {
	var user User
	result := db.Gorm.First(&user).Where("id = ?", id).Take(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (db *Database) GetUsers() []User {
	var users []User
	db.Gorm.Find(&users)
	return users
}
