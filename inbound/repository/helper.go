package repository

//func (r *LikeRepository) fetchMany(ctx context.Context, op *like.FetchLikeIdeasFilters) ([]model.LikeAggr, error) {
//	pipeline := r.fetchConversationFinalPipeline(ctx, op)
//	cur, err := r.Collection.Aggregate(ctx, pipeline)
//	if err != nil {
//		return nil, err
//		fmt.Println(err.Error())
//	}
//	result := make([]model.LikeAggr, 0)
//	for cur.Next(ctx) {
//		var obj model.LikeAggr
//		// decode the document
//		if err := cur.Decode(&obj.Like); err != nil {
//			return nil, err
//		}
//		if err := cur.Decode(&obj); err != nil {
//			return nil, err
//		}
//
//		result = append(result, obj)
//	}
//	if err := cur.Err(); err != nil {
//		return nil, err
//	}
//	return result, nil
//}
//
//func (r *LikeRepository) fetchManyCount(ctx context.Context, op *like.FetchLikeIdeasFilters) (int, error) {
//	pipeline := r.fetchConversationBasePipeline()
//
//	pipeline = append(pipeline, bson.M{
//		"$match": op.Filter,
//	})
//	pipeline = append(pipeline, bson.M{
//		"$count": "total_count",
//	})
//	type CountData struct {
//		TotalCount int `json:"total_count" bson:"total_count"`
//	}
//
//	cur, err := r.Collection.Aggregate(ctx, pipeline)
//	if err != nil {
//		return -1, err
//	}
//
//	count := 0
//	// iterate through all documents
//	for cur.Next(ctx) {
//		obj := new(CountData)
//		// decode the document
//		if err := cur.Decode(&obj); err != nil {
//			return -1, err
//		}
//		count = obj.TotalCount
//
//	}
//
//	return count, err
//}
//func (r *LikeRepository) fetchConversationBasePipeline() []bson.M {
//
//	pipeline := []bson.M{
//		{
//			"$lookup": bson.M{
//				"from":         "idea",
//				"localField":   "idea_id",
//				"foreignField": "_id",
//				"as":           "idea_data",
//			},
//		},
//		{"$project": bson.M{
//
//			"_id":         "$_id",
//			"user_id":     "$user_id",
//			"idea_id":     "$idea_id",
//			"create_date": "$create_date",
//			"update_date": "$update_date",
//			"idea_data":   "$idea_data",
//		},
//		},
//	}
//
//	return pipeline
//}
//
//func (r *LikeRepository) fetchConversationFinalPipeline(ctx context.Context, op *like.FetchLikeIdeasFilters) []bson.M {
//
//	pipeline := r.fetchConversationBasePipeline()
//
//	pipeline = append(pipeline, bson.M{
//		"$match": op.Filter,
//	})
//	if op.SortBy != "" {
//		pipeline = append(pipeline, bson.M{
//			"$sort": bson.M{op.SortBy: op.Order},
//		})
//	}
//
//	pipeline = append(pipeline, bson.M{
//		"$skip": op.Skip,
//	})
//
//	if op.Limit > 0 {
//		pipeline = append(pipeline, bson.M{
//			"$limit": op.Limit,
//		})
//	}
//
//	return pipeline
//}
