package inbound

import (
	"context"
	model "github.com/gomarkho/demo-callpicker/model"
)

//UserRepository represent the book repository
type Repository interface {
	//GetLikeObject(ctx context.Context, ideaid primitive.ObjectID, userid primitive.ObjectID) (*model.Like, error)
	Store(ctx context.Context, u *model.DialPlan) (*model.DialPlan, error)
	Exists(ctx context.Context, dest string, source string) (bool, *model.DialPlan)
	GetByNumber(ctx context.Context, op XMLOptions) (*model.DialPlan, error)
	GetDialplanByNumber(ctx context.Context, op GetDialPlan) (*model.DialPlan, error)
	DeleteByNumber(ctx context.Context, u DeleteDialPlan) error

	//UpdateTime(ctx context.Context, op DisLikeIdea) (*model.Like,error)
	//
	//DeleteLikeObject(ctx context.Context, op DisLikeIdea) error
	//List(ctx context.Context, op FetchLikeIdeasFilters) (*[]model.LikeAggr, int, error)
	//GetLike(ctx context.Context, idea_id primitive.ObjectID, user_id primitive.ObjectID) bool

	//GetByID(ctx context.Context, id primitive.ObjectID) (*model.Conversation, error)
	//IfConversationExists(ctx context.Context, obj model.Conversation) (bool, error)
	//DeleteByID(ctx context.Context, id primitive.ObjectID) (error)

	//Get(ctx context.Context, ObjectId primitive.ObjectID) (*model.Idea, error)
	//List(ctx context.Context, op FetchIdeaFilters) (*[]model.Idea, int, error)
	//Delete(ctx context.Context, ObjectId primitive.ObjectID) error
	//UpdateStatus(ctx context.Context, ObjectId primitive.ObjectID) error
	//ListSuggestions(ctx context.Context, op FetchSuggestionsFilters) (*[]model.IdeaAggr, int, error)
}
