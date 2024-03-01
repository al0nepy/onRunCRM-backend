package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProductCategoryModel struct {
	ID        primitive.ObjectID  `bson:"_id,omitempty"`
	Name      string              `bson:"name"`
	TenantID  primitive.ObjectID  `bson:"tenantId"`
	CreatedAt primitive.Timestamp `bson:"created_at"`
	UpdatedAt primitive.Timestamp `bson:"updated_at"`
}
