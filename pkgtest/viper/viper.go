// Time        : 2019/09/03
// Description :

package viper

import (
	"context"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongodb "golearn/something-sundry/mongo-test/mongo-db"
	"golearn/utils"
	"golearn/utils/typeconvert"
)

func init() {
	viper.AddConfigPath("path")
}

func initConfig() map[string]interface{} {
	ctx := context.Background()
	mongodb.InitClient("mongodb://localhost")
	client := mongodb.GetClient()
	dbNames, err := client.Database("test").ListCollectionNames(ctx, bson.D{})
	utils.OkOrPanic(err)
	allConf := make(map[string]interface{})
	for i := range dbNames {
		collection := client.Database("test").Collection(dbNames[i])
		cur, err := collection.Find(ctx, bson.D{})
		utils.OkOrPanic(err)
		conf := make(map[string]interface{})
		for cur.Next(ctx) {
			var row bson.M
			err = cur.Decode(&row)
			utils.OkOrPanic(err)
			var key string
			switch id := row["_id"].(type) {
			case primitive.ObjectID:
				key = id.String()
			case string:
				key = id
			case int8, int16, int, int64:
				key = typeconvert.NumberToString(id)
			}
			conf[key] = row
		}
		allConf[dbNames[i]] = conf
	}
	return allConf
}