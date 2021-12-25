package repository

import (
	"context"
	"github.com/gomarkho/demo-callpicker/inbound"
	"github.com/gomarkho/demo-callpicker/model"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const collection = "inbound"

type InboundRepository struct {
	Conn       *mongo.Database
	Collection *mongo.Collection
}

func NewLikeRepository(dbConn *mongo.Database) inbound.Repository {
	return &InboundRepository{dbConn, dbConn.Collection(collection)}
}
func (r *InboundRepository) Store(ctx context.Context, u *model.DialPlan) (*model.DialPlan, error) {
	if u.ID.IsZero() {
		u.ID = primitive.NewObjectID()
		u.CreateDate = time.Now().UTC()
		u.UpdateDate = time.Now().UTC()
		res, err := r.Collection.InsertOne(ctx, u)
		if err != nil {
			logrus.Errorln(err.Error())
			return nil, err
		}
		u.ID = res.InsertedID.(primitive.ObjectID)
	} else {
		u.UpdateDate = time.Now().UTC()
		_, err := r.Collection.UpdateOne(ctx, bson.M{"_id": u.ID}, bson.M{"$set": bson.M{"update_date": u.UpdateDate}})
		if err != nil {
			return nil, err
		}
	}
	return u, nil
}

func (r *InboundRepository) Exists(ctx context.Context, dest string, source string) (bool, *model.DialPlan) {

	var obj *model.DialPlan
	filter := bson.M{
		"destination_number": dest,
		"source_number":      source,
	}
	_ = r.Collection.FindOne(ctx, filter).Decode(&obj)
	if obj == nil {
		return false, nil
	}
	return true, obj
}

func (r *InboundRepository) DeleteByNumber(ctx context.Context, op inbound.DeleteDialPlan) error {
	filter := bson.M{"destination_number": op.Destination, "source_number": op.Source}
	_, err := r.Collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func (r *InboundRepository) GetByNumber(ctx context.Context, req inbound.XMLOptions) (*model.DialPlan, error) {
	var obj model.DialPlan
	filter := bson.M{
		"source_number":      req.Source,
		"destination_number": req.Destination,
	}
	err := r.Collection.FindOne(ctx, filter).Decode(&obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func (r *InboundRepository) GetDialplanByNumber(ctx context.Context, req inbound.GetDialPlan) (*model.DialPlan, error) {
	var obj model.DialPlan
	filter := bson.M{
		"source_number":      req.Source,
		"destination_number": req.Destination,
	}
	err := r.Collection.FindOne(ctx, filter).Decode(&obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}
