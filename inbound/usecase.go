package inbound

import (
	"context"
	"github.com/gomarkho/demo-callpicker/model"
)

type UseCase interface {
	InsertDialPlan(ctx context.Context, req *InsertDialPlan) (*model.DialPlan, int, error)
	GetXml(ctx context.Context, req *XMLOptions) (*model.XMLResponse, error)
	DeleteDialPlan(ctx context.Context, req DeleteDialPlan) (int, error)

	GetDialPlan(ctx context.Context, req GetDialPlan) (*model.DialPlan, int, error)

	//DislikeIdea(ctx context.Context, req *DisLikeIdea) (int, error)
	//ListLikeIdeas(ctx context.Context, op GetLikeIdeasRequestOptions) (*[]model.LikeAggr, *model.Page, error)

	//UpdateIdea(ctx context.Context, req *UpdateIdea) (*model.Idea, int, error)
	//UpdateStatus(ctx context.Context, req *UpdateStatus) (*model.Idea, int, error)
	//DeleteIdea(ctx context.Context, ObjectId string) (int, error)
	//ListIdeas(ctx context.Context, op GetIdeaRequestOptions) (*[]model.Idea, *model.Page, error)
	//ListSuggestions(ctx context.Context, op GetSuggestionsRequestOptions) (*[]model.IdeaAggr, *model.Page, error)
}
