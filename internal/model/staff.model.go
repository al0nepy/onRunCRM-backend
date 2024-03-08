package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type StaffModel struct {
	ID         primitive.ObjectID  `bson:"_id,omitempty"`
	Name       string              `bson:"name"`
	Email      string              `bson:"email"`
	Position   string              `bson:"position"`
	Department string              `bson:"department"`
	TenantID   primitive.ObjectID  `bson:"tenantId"`
	CreatedAt  primitive.Timestamp `bson:"created_at"`
	UpdatedAt  primitive.Timestamp `bson:"updated_at"`
}
