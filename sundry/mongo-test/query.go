// Time        : 2019/09/02
// Description :

package mongotest

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"golearn/sundry/mongo-test/model"
	mongodb "golearn/sundry/mongo-test/mongo-db"
	"golearn/util"
)

var client *mongo.Client

func init() {
	mongodb.InitClient("mongodb://localhost")
	client = mongodb.GetClient()
}

func queryByRegex(regex string) model.Role {
	ctx := context.Background()
	collection := client.Database("test").Collection("role")
	cur, err := collection.Find(ctx, bsonx.Doc{{Key: "_id", Value: bsonx.Regex(regex, "")}})
	util.MustNil(err)
	if !cur.Next(ctx) {
		panic("no role")
	}
	var role model.Role
	err = cur.Decode(&role)
	util.MustNil(err)
	return role
}

func queryRole(id string, stu interface{}) model.Role {
	ctx := context.Background()
	collection := client.Database("test").Collection("role")
	cur, err := collection.Find(ctx, bson.D{{"_id", id}})
	util.MustNil(err)
	if !cur.Next(ctx) {
		panic("no role")
	}
	var role model.Role
	err = cur.Decode(&role)
	util.MustNil(err)
	fmt.Printf("%+v", role.Students)
	role.MarshalStudents(stu)
	return role
}

func queryStudent(id string) model.StudentValue {
	ctx := context.Background()
	collection := client.Database("test").Collection("student")
	cur, err := collection.Find(ctx, bson.D{{"name", id}})
	util.MustNil(err)
	if !cur.Next(ctx) {
		panic("no student")
	}
	var student model.StudentValue
	err = cur.Decode(&student)
	util.MustNil(err)
	return student
}

func queryAll() {
	ctx := context.Background()
	dbNames, err := client.Database("test").ListCollectionNames(ctx, bson.D{})
	util.MustNil(err)
	for i := range dbNames {
		collection := client.Database("test").Collection(dbNames[i])
		cur, err := collection.Find(ctx, bson.D{})
		util.MustNil(err)
		for cur.Next(ctx) {
			var row bson.M
			err = cur.Decode(&row)
			util.MustNil(err)
			fmt.Println(row)
		}
	}
}

func queryMap() {
	ctx := context.Background()
	collection := client.Database("test").Collection("test")
	cur, err := collection.Find(ctx, bson.D{{"b", 123}})
	util.MustNil(err)
	if !cur.Next(ctx) {
		panic("no value")
	}
	var mm ModelWithMap
	err = cur.Decode(&mm)
	util.MustNil(err)
	fmt.Println(mm)
}

func queryMax() {
	ctx := context.Background()
	collection := client.Database("test").Collection("student")
	pipeline := []bson.M{
		{
			"$match": bson.M{
				"power": bson.M{
					"$gt": 550,
					"$lt": 700,
				},
			},
		},
		{
			"$group": bson.M{
				"_id": "$power",
				"doc": bson.M{
					"$push": "$$ROOT",
				},
			},
		},
		{
			"$sort": bson.M{
				"power": -1,
			},
		},
		{
			"$limit": 1,
		},
	}
	cur, err := collection.Aggregate(ctx, pipeline)
	util.MustNil(err)
	for cur.Next(ctx) {
		var result struct {
			Doc []model.StudentValue
		}
		err = cur.Decode(&result)
		fmt.Printf("%+v\n", result.Doc)
	}
}
