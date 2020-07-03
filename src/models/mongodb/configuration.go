package mongodb

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ConfigurationObject struct {
	ID        string    `bson:"_id"`
	Value     string    `bson:"value"`
	Key       string    `bson:"key"`
	Type      string    `bson:"type"`
	CreatedAt time.Time `bson:"_created_at"`
	UpdatedAt time.Time `bson:"_updated_at"`
}

// GetConfigurationByKey a function to get value with key as string
func GetConfigurationByKey(key string) (result *ConfigurationObject, err error) {
	mongoURI := os.Getenv("DB_MONGO_URI")
	mongoDBName := os.Getenv("DB_MONGO_DBNAME")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	collection := client.Database(mongoDBName).Collection("Configuration")

	err = collection.FindOne(ctx, bson.M{
		"key": key,
	}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// InsertNewConfiguration a function to insert data into database
func InsertNewConfiguration(key, value, configType string) (result *ConfigurationObject, err error) {
	mongoURI := os.Getenv("DB_MONGO_URI")
	mongoDBName := os.Getenv("DB_MONGO_DBNAME")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	collection := client.Database(mongoDBName).Collection("Configuration")

	doc := bson.M{
		"key":         key,
		"value":       value,
		"type":        configType,
		"enabled":     true,
		"_created_at": time.Now(),
		"_updated_at": time.Now(),
	}
	object, err := collection.InsertOne(ctx, doc)
	if err != nil {
		return nil, err
	}

	err = collection.FindOne(ctx, bson.M{
		"_id": object.InsertedID,
	}).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
