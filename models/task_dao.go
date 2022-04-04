package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoList struct {
	ID         primitive.ObjectID `bson:"_id"`
	TodoItem   string             `json:"todo_item"`
	Created_at time.Time          `json:"created_at"`
	Done       bool               `json:"is_done"`
}
