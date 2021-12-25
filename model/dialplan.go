package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type DialPlan struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Source      string             `json:"source_number" bson:"source_number"`
	Destination string             `json:"destination_number" bson:"destination_number"`
	Say         []Say              `json:"say"  bson:"say"  `
	Play        []Play             `json:"play" bson:"play"`
	Record      []Record           `json:"record" bson:"record"`
	Dial        []Dial             `json:"dial" bson:"dial"`
	CreateDate  time.Time          `json:"create_date" bson:"create_date"`
	UpdateDate  time.Time          `json:"update_date" bson:"update_date"`
}
