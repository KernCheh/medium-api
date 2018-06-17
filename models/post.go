package post

import (
	"time"
)

type Post struct {
	Id        int
	Creator   string `sql:"type:varchar(255)"`
	Content   string
	CreatedAt time.Time `sql:"default:now()"`
	UpdatedAt time.Time
}
