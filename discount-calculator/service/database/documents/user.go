package documents

import "time"

type User struct {
	Id          int       `bson:"_id"`
	FirstName   string    `bson:"first_name"`
	LastName    string    `bson:"last_name"`
	DateOfBirth time.Time `bson:"date_of_birth"`
}
