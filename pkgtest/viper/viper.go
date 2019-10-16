// Time        : 2019/09/03
// Description :

package viper

import (
	"context"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongodb "golearn/sundry/mongo-test/mongo-db"
	"golearn/util"
	"golearn/util/convert"
	"os"
	"path/filepath"
)

func initConfigFromMongo() map[string]interface{} {
	ctx := context.Background()
	mongodb.InitClient("mongodb://localhost")
	client := mongodb.GetClient()
	dbNames, err := client.Database("test").ListCollectionNames(ctx, bson.D{})
	util.OkOrPanic(err)
	allConf := make(map[string]interface{})
	for i := range dbNames {
		collection := client.Database("test").Collection(dbNames[i])
		cur, err := collection.Find(ctx, bson.D{})
		util.OkOrPanic(err)
		conf := make(map[string]interface{})
		for cur.Next(ctx) {
			var row bson.M
			err = cur.Decode(&row)
			util.OkOrPanic(err)
			var key string
			switch id := row["_id"].(type) {
			case primitive.ObjectID:
				key = id.Hex()
			case string:
				key = id
			case int8, int16, int32, int, int64:
				key = convert.ToString(id)
			}
			conf[key] = row
		}
		allConf[dbNames[i]] = conf
	}
	return allConf
}

func InitConfigFromFile(path string) {
	execPath, err := os.Getwd()
	util.OkOrPanic(err)
	absPath, _ := filepath.Abs(filepath.Dir(execPath) + path)
	viper.SetConfigFile(absPath)
	err = viper.ReadInConfig()
	util.OkOrPanic(err)
}
