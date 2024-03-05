package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProductModel struct {
	ID          primitive.ObjectID  `bson:"_id,omitempty"`
	TenantID    primitive.ObjectID  `bson:"tenantId"`
	Name        string              `bson:"name"`
	Description string              `bson:"description"`
	Price       float64             `bson:"price"`
	Quantity    int64               `bson:"quantity"`
	Category    primitive.ObjectID  `bson:"category"`
	Attributes  map[string]string   `bson:"attributes"`
	CreatedAt   primitive.Timestamp `bson:"created_at"`
	UpdatedAt   primitive.Timestamp `bson:"updated_at"`
}
