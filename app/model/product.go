package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Product ...
type Product struct {
	ID        primitive.ObjectID `json:"ID" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name,omitempty"`
	Desc      string             `json:"desc" bson:"desc,omitempty"`
	CreatedAt time.Time          `json:"createdAt" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updated_at,omitempty"`
}
