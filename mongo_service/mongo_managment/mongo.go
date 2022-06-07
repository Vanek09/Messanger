package mongo_managment

import (
	"fmt"
	p "mongo_service/util"
	"time"

	// "reflect"
	// "fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const props_name = "config/api.properties"

func Setup() {
	props := p.ReadProperties(props_name)
	host := props.GetString("mongo.host", "localhost")
	port := props.GetString("mongo.port", "27017")
	connectionURI := "mongodb://" + host + ":" + port
	cOptions := options.Client().ApplyURI(connectionURI)
	client, err := mongo.Connect(Ctx, cOptions)
	
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(Ctx)
	err = client.Ping(Ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("messanger")
	userCol = db.Collection("user")
	messageBoxCol = db.Collection("messageBox")
	// messagesCol = db.Collection("messages")
	messageCol = db.Collection("message")
}

func CreateUser(u User) (string, error) {
	props := p.ReadProperties(props_name)
	host := props.GetString("mongo.host", "localhost")
	port := props.GetString("mongo.port", "27017")
	connectionURI := "mongodb://" + host + ":" + port
	cOptions := options.Client().ApplyURI(connectionURI)
	client, err := mongo.Connect(Ctx, cOptions)

	db := client.Database("messanger")
	userCol = db.Collection("user")
	messageBoxCol = db.Collection("messageBox")
	// messagesCol = db.Collection("messages")
	messageCol = db.Collection("message")
	
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(Ctx)
	//To create user we need to init MessageBox for him
	mBoxId := insert(MessageBox{Messages: make([]Message, 0, 1000000), Id: u.Nickname}, messageBoxCol)
	if mBoxId == "" {
		return "", nil
	}
	u.MessageBox = mBoxId.(string)
	// u.Id = u.Nickname
	userId := insert(u, userCol)
	
	return userId.(string), nil
}

func GetUser(id string) (User, error) {
	props := p.ReadProperties(props_name)
	host := props.GetString("mongo.host", "localhost")
	port := props.GetString("mongo.port", "27017")
	connectionURI := "mongodb://" + host + ":" + port
	cOptions := options.Client().ApplyURI(connectionURI)
	client, err := mongo.Connect(Ctx, cOptions)

	db := client.Database("messanger")
	userCol = db.Collection("user")
	messageBoxCol = db.Collection("messageBox")
	// messagesCol = db.Collection("messages")
	messageCol = db.Collection("message")
	
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(Ctx)

	var u User
	encoded, err := findOne(id, userCol)
	encoded.Decode(&u)
	if err != nil {
		return u, err
	}
	return u, nil
}

func GetUsers() ([]User, error) {
	props := p.ReadProperties(props_name)
	host := props.GetString("mongo.host", "localhost")
	port := props.GetString("mongo.port", "27017")
	connectionURI := "mongodb://" + host + ":" + port
	cOptions := options.Client().ApplyURI(connectionURI)
	client, err := mongo.Connect(Ctx, cOptions)

	db := client.Database("messanger")
	userCol = db.Collection("user")
	messageBoxCol = db.Collection("messageBox")
	// messagesCol = db.Collection("messages")
	messageCol = db.Collection("message")
	
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(Ctx)

	var users []User
	var user User
	cursor, err := userCol.Find(Ctx, bson.D{})
	if err != nil {
		return users, nil
	}
	defer cursor.Close(Ctx)
	
	for cursor.Next(Ctx) {
		err := cursor.Decode(&user)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

func GetUserMessages(from, to string) []Message {
	props := p.ReadProperties(props_name)
	host := props.GetString("mongo.host", "localhost")
	port := props.GetString("mongo.port", "27017")
	connectionURI := "mongodb://" + host + ":" + port
	cOptions := options.Client().ApplyURI(connectionURI)
	client, err := mongo.Connect(Ctx, cOptions)

	db := client.Database("messanger")
	userCol = db.Collection("user")
	messageBoxCol = db.Collection("messageBox")
	// messagesCol = db.Collection("messages")
	messageCol = db.Collection("message")
	
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(Ctx)
	log.Printf("Getting messages from: %s to %s", from, to)
	mBox := getMessageBox(User{Id:from})
	return getMessageList(mBox, User{Id:to})
}

func SaveMessage(msg Message) {

	props := p.ReadProperties(props_name)
	host := props.GetString("mongo.host", "localhost")
	port := props.GetString("mongo.port", "27017")
	connectionURI := "mongodb://" + host + ":" + port
	cOptions := options.Client().ApplyURI(connectionURI)
	client, err := mongo.Connect(Ctx, cOptions)
	
	db := client.Database("messanger")
	userCol = db.Collection("user")
	messageBoxCol = db.Collection("messageBox")
	// messagesCol = db.Collection("messages")
	messageCol = db.Collection("message")
	if err != nil {
		log.Fatal(err)
	}
	
	t := time.Now()
	msg.Time = fmt.Sprintf("%d:%d:%d", t.Hour(), t.Minute(), t.Second())
	msg.Date = fmt.Sprintf("%d.%d.%d", t.Year(), t.Month(), t.Day())
	insert(msg, messageCol)
	client.Disconnect(Ctx)
	// log.Printf("Updated MessageBox %s\n", msg.Id)
}

func getMessageBox(u User) MessageBox {
	mBox := MessageBox{Messages: make([]Message, 0, 1000000)}
	encoded, err := findOne(u.Id, messageBoxCol)
	if err != nil {
		log.Println(err)
	}	
	err = encoded.Decode(&mBox)
	if err != nil {
		log.Println(err)
	}
	cur, err := messageBoxCol.Find(Ctx, bson.D{{"_id", u.Id}})
	cur.All(Ctx, &mBox.Messages)
	if err != nil {
		log.Println(err)
	}

	return mBox
}

func getMessageList(mBox MessageBox, to User) []Message {
	var msg Message
	var msgs []Message

	cursor, err := messageCol.Find(Ctx, bson.D{})
	if err != nil {
		defer cursor.Close(Ctx)
		return msgs
	}

	for cursor.Next(Ctx) {
		err := cursor.Decode(&msg)
		if err != nil {
			return msgs
		}
		if (msg.From == mBox.Id && msg.To == to.Id) || (msg.From == to.Id && msg.To == mBox.Id) {
			msgs = append(msgs, msg)
		}
		
	}

	return msgs
}