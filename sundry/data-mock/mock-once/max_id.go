//@time:2020/01/20
//@desc:

package mockdata

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"golearn/util"
)

func (m mockManager) maxId(ctx context.Context) int64 {
	pipeline := []bson.M{
		{
			"$sort": bson.M{
				"_id": -1,
			},
		},
		{
			"$limit": 1,
		},
		{
			"$group": bson.M{
				"_id": "$_id",
			},
		},
	}
	cur, err := m.coll.Aggregate(ctx, pipeline)
	util.MustNil(err)
	for cur.Next(ctx) {
		var result map[string]interface{}
		err = cur.Decode(&result)
		util.MustNil(err)
		return result["_id"].(int64)
	}
	return 0
}
