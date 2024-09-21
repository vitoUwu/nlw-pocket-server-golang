package db

import (
	"nlw/pocket/env"
	"time"

	"github.com/lucsky/cuid"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Database struct {
	Goals map[string]Goal
	Users map[string]User
	Gorm  *gorm.DB
}

func NewDatabase() *Database {
	connectionString := env.GetDatabaseURL()

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: false,
		},
	})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{}, &Goal{}, &GoalCompletion{})

	return &Database{
		Goals: map[string]Goal{},
		Users: map[string]User{},
		Gorm:  db,
	}
}

var Db = NewDatabase()

func (db *Database) GenerateId() string {
	return cuid.New()
}

func (db *Database) FirstDayOfWeek() time.Time {
	return time.Now().AddDate(0, 0, -int(time.Now().Weekday())).Truncate(24 * time.Hour)
}

func (db *Database) LastDayOfWeek() time.Time {
	return db.FirstDayOfWeek().AddDate(0, 0, 6).Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)
}
