// Time        : 2019/09/09
// Description :

package loader

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	mongodb "golearn/sundry/mongo-test/mongo-db"
	"golearn/sundry/profile/model"
	"golearn/util"
	"reflect"
)

type MongoLoader struct{}

func (m MongoLoader) Load() (dataMap map[string]map[string]interface{}, dataList map[string]interface{}, err error) {
	ctx := context.Background()
	mongodb.InitClient("mongodb://localhost")
	client := mongodb.GetClient()
	dbNames, err := client.Database("test2").ListCollectionNames(ctx, bson.D{})
	util.OkOrPanic(err)
	dataMap = make(map[string]map[string]interface{})
	dataList = make(map[string]interface{})
	for i := range dbNames {
		collection := client.Database("test2").Collection(dbNames[i])
		cur, err := collection.Find(ctx, bson.D{})
		util.OkOrPanic(err)
		typ := model.GetModelType(dbNames[i])
		data := make(map[string]interface{})
		list := reflect.MakeSlice(reflect.SliceOf(typ), 0, 0)
		for cur.Next(ctx) {
			ptr := reflect.New(typ)
			err = cur.Decode(ptr.Interface())
			key := generateIndexKey(ptr.Elem())
			data[key] = ptr.Elem().Interface()
			list = reflect.Append(list, ptr.Elem())
		}
		dataMap[dbNames[i]] = data
		dataList[dbNames[i]] = list.Interface()
	}
	return
}
