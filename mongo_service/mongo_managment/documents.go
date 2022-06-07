package mongo_managment

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	Ctx             = context.TODO()
	userCol 		*mongo.Collection
	messageBoxCol   *mongo.Collection
	// messagesCol     *mongo.Collection
	messageCol      *mongo.Collection
)

type User struct {
	Id          string   `bson:"_id"         json:"_id"`
	Nickname    string   `bson:"nickname"    json:"nickname"`
	Hashed_pwd  string   `bson:"hashed_pwd"  json:"hashed_pwd"`
	MessageBox  string   `bson:"messageBox"  json:"messageBox"`
}

type MessageBox struct {
	Id          string   `bson:"_id"        json:"_id"`
	Messages  []Message   `bson:"messages"   json:"messages"`
}

type Message struct {
	Id          primitive.ObjectID   `bson:"_id,omitempty"         json:"_id"`
	From        string   `bson:"adress"      json:"adress"`
	To          string   `bson:"destination" json:"destination"`
	Date        string   `bson:"date"        json:"date"`
	Time        string   `bson:"time"        json:"time"`
	Message     string   `bson:"message"     json:"message"`
}

