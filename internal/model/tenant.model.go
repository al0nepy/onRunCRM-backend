package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type TenantModel struct {
	ID        primitive.ObjectID  `bson:"_id,omitempty"`
	Name      string              `bson:"name"`
	CreatedAt primitive.Timestamp `bson:"created_at"`
	UpdatedAt primitive.Timestamp `bson:"updated_at"`
}
