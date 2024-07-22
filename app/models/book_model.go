package models

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID         uuid.UUID `json:"id" bson:"id"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" bson:"updated_at"`
	UserID     uuid.UUID `json:"user_id" bson:"user_id"`
	Title      string    `json:"title" bson:"title"`
	Author     string    `json:"author" bson:"author"`
	BookStatus int       `json:"book_status" bson:"book_status"`
	BookAttrs  BookAttrs `json:"book_attrs" bson:"book_attrs"`
}

type BookAttrs struct {
	Picture     string `json:"picture" bson:"picture"`
	Description string `json:"description" bson:"description"`
	Rating      string `json:"rating" bson:"rating"`
}

//	NOTE: NO NEED FOR MONGODB USAGE
// func (b *BookAttrs) Value() (driver.Value, error) {
// 	return json.Marshal(b)
// }

// func (b *BookAttrs) Scan(value interface{}) error {

// 	j, ok := value.([]byte)

// 	if !ok {
// 		return errors.New("type assertion to []byte failed")
// 	}

// 	return json.Unmarshal(j, &b)
// }
