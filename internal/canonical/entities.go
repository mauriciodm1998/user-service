package canonical

import "time"

type Customer struct {
	Id       string
	UserID   string
	Document string
	Name     string
	Email    string
}

type User struct {
	Id            string
	Login         string
	Password      string
	AccessLevelID int
	CreatedAt     time.Time
}
