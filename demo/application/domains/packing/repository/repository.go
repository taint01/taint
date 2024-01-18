package repository

import (
	"context"
	"demo/application/model"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"

	"demo/application/domains/packing/entity"
	"demo/application/domains/packing/models"
)

type repo struct {
	lib *model.Lib
}

func InitRepo(lib *model.Lib) IRepository {
	return &repo{
		lib: lib,
	}
}

func (r *repo) Read(ctx context.Context, params *models.Params) ([]*entity.User, error) {
	var results []*entity.User

	collection := r.lib.MongoDb.Collection("localcollection")
	paramBytes, err := bson.Marshal(params)
	filter := bson.M{}
	lastFilter := bson.M{}
	err = bson.Unmarshal(paramBytes, &filter)
	for key, val := range filter {
		if val != nil && val != int64(0) && val != "" {
			lastFilter[key] = val
		}
	}
	if err != nil {
		return nil, err
	}
	cur, err := collection.Find(context.TODO(), lastFilter)

	if err != nil {
		panic(err)
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var user *entity.User
		err = cur.Decode(&user)
		if err != nil {
			panic(err)
		}
		results = append(results, user)
	}

	if err = cur.Err(); err != nil {
		panic(err)
	}

	fmt.Println("Found users:", results)
	return results, nil
}

func (r *repo) Insert(ctx context.Context, params *entity.User) error {
	collection := r.lib.MongoDb.Collection("localcollection")
	result, err := collection.InsertOne(context.TODO(), params)

	if err != nil {
		panic(err)
	}

	fmt.Println("Inserted document ID:", result.InsertedID)
	return nil
}

func (r *repo) Update(ctx context.Context, params *entity.User) error {
	collection := r.lib.MongoDb.Collection("localcollection")
	filter := bson.D{{"id", params.Id}}
	paramsBytes, err := bson.Marshal(params)
	if err != nil {
		// Handle error
	}

	save := bson.M{}
	lastSave := bson.M{}
	err = bson.Unmarshal(paramsBytes, &save)
	if err != nil {
		// Handle error
	}
	for key, val := range save {
		if val != nil && val != int64(0) && val != "" {
			lastSave[key] = val
		}
	}
	if err != nil {
		return err
	}

	update := bson.M{"$set": lastSave}
	result, err := collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		panic(err)
	}

	fmt.Println("Updated documents:", result.ModifiedCount)
	return nil
}
func (r *repo) Delete(ctx context.Context, params *entity.User) error {
	collection := r.lib.MongoDb.Collection("localcollection")
	paramBytes, err := bson.Marshal(params)
	filter := bson.M{}
	err = bson.Unmarshal(paramBytes, &filter)
	if err != nil {
		return err
	}

	lastFilter := bson.M{}
	for key, val := range filter {
		if val != nil && val != int64(0) && val != "" {
			lastFilter[key] = val
		}
	}
	result, err := collection.DeleteOne(context.TODO(), lastFilter)

	if err != nil {
		panic(err)
	}

	fmt.Println("Deleted documents:", result.DeletedCount)
	return nil
}
