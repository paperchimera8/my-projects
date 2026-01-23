package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Goods struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ShopID      primitive.ObjectID `bson:"shopid" json:"shopid"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description,omitempty" json:"description"`
}
