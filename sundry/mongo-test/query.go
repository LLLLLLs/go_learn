// Time        : 2019/09/02
// Description :

package mongotest

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	model2 "golearn/sundry/mongo-test/model"
	mongodb "golearn/sundry/mongo-test/mongo-db"
	"golearn/utils"
)

var client *mongo.Client

func init() {
	mongodb.InitClient("mongodb://localhost")
	client = mongodb.GetClient()
}

func queryRole(id string) model2.Role {
	ctx := context.Background()
	collection := client.Database("test").Collection("role")
	cur, err := collection.Find(ctx, bson.D{{"_id", id}})
	utils.OkOrPanic(err)
	if !cur.Next(ctx) {
		panic("no role")
	}
	var role model2.Role
	err = cur.Decode(&role)
	utils.OkOrPanic(err)
	return role
}

func queryStudent(id string) model2.StudentValue {
	ctx := context.Background()
	collection := client.Database("test").Collection("student")
	cur, err := collection.Find(ctx, bson.D{{"name", id}})
	utils.OkOrPanic(err)
	if !cur.Next(ctx) {
		panic("no student")
	}
	var student model2.StudentValue
	err = cur.Decode(&student)
	utils.OkOrPanic(err)
	return student
}

func queryAll() {
	ctx := context.Background()
	dbNames, err := client.Database("test").ListCollectionNames(ctx, bson.D{})
	utils.OkOrPanic(err)
	for i := range dbNames {
		collection := client.Database("test").Collection(dbNames[i])
		cur, err := collection.Find(ctx, bson.D{})
		utils.OkOrPanic(err)
		for cur.Next(ctx) {
			var row bson.M
			err = cur.Decode(&row)
			utils.OkOrPanic(err)
			fmt.Println(row)
		}
	}
}