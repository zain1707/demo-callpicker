package inbound

import (
	"github.com/gomarkho/demo-callpicker/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InsertDialPlan struct {
	Source      string         `json:"source_number"`
	Destination string         `json:"destination_number"`
	Say         []model.Say    `json:"say"`
	Play        []model.Play   `json:"play"`
	Record      []model.Record `json:"record"`
	Dial        []model.Dial   `json:"dial"`
}

type DeleteDialPlan struct {
	Source      string `json:"source_number"`
	Destination string `json:"destination_number"`
}

type GetDialPlan struct {
	Source      string `json:"source_number"`
	Destination string `json:"destination_number"`
}

type XMLOptions struct {
	Source      string
	Destination string
	CallSid     string
}

type DisLikeIdea struct {
	UserID primitive.ObjectID `json:"user_id"`
	IdeaID primitive.ObjectID `json:"idea_id"`
}

type FetchLikeIdeasFilters struct {
	SortBy     string
	Order      int
	Skip       int64
	Limit      int64
	Filter     map[string]interface{}
	Search     string
	Visibility []string
}

type GetLikeIdeasRequestOptions struct {
	UserID primitive.ObjectID
	Search string
	Page   string
	Limit  string
}
