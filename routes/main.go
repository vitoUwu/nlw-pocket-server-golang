package routes

import db "nlw/pocket/db"

var (
	Db = db.Db
)

type Goal struct {
	db.Goal
}
