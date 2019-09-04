// Time        : 2019/09/02
// Description :

package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golearn/utils"
)

var client *mongo.Client

func GetClient() *mongo.Client {
	if client == nil {
		panic("must init first")
	}
	return client
}

func InitClient(uri string) {
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI(uri))
	utils.OkOrPanic(err)
	err = client.Connect(context.Background())
	utils.OkOrPanic(err)
}
