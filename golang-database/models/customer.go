package models

import "database/sql"

type Customer struct {
	Id        int32
	Name      string
	Email     sql.NullString
	Balance   float64
	Rating    sql.NullFloat64
	BirthDate sql.NullTime
	IsMarried bool
	CreatedAt sql.NullTime
}
