package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Shop struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name       string             `bson:"name" json:"name"`
	OpenTime   string             `bson:"opentime,omitempty" json:"opentime"`
	CloseTime  string             `bson:"closetime,omitempty" json:"closetime"`
	DaysOfWeek string             `bson:"daysofweek,omitempty" json:"daysofweek"`
}
