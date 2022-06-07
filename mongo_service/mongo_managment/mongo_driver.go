package mongo_managment

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func insert(m interface{}, col *mongo.Collection) interface{} {
	result, err := col.InsertOne(Ctx, m)
	if err != nil {
		log.Println(err)
		return ""
	}

	return result.InsertedID
}

// func insertMany(m []interface{}, col *mongo.Collection) string {
// 	result, err := col.InsertMany(Ctx, m)
// 	if err != nil {
// 		log.Println(err)
// 		return ""
// 	}

// 	return fmt.Sprintf("%v", result)
// }

func findOne(id string, col *mongo.Collection) (*mongo.SingleResult, error) {
	filter := bson.D{{"_id", id}}
	return userCol.FindOne(Ctx, filter), nil
}

// func findAll(col *mongo.Collection, item interface{}, items ...interface{}) ([]interface{}, error) {

// }

// func updateItem(id string, col *mongo.Collection, messages []Message) error {
// 	filter := bson.D{{"_id", id}}
// 	update := bson.D{{"$set", bson.D{{"messages", messages}}}}
// 	_, err := col.UpdateOne(
// 		Ctx,
// 		filter,
// 		update,
// 	)
// 	return err
// }


// func agregateMessages(from string, to string) ([]Message, error) {
// 	matchStage := bson.D{{"$match", bson.D{{"adress", from}}}}

// 	lookupStage := bson.D{{"$lookup",
// 		bson.D{{"from", "messages"},
// 			{"localField", "_id"},
// 			{"foreignField", "adress"},
// 			{"as", "messages"}}}}

// 	showLoadedCursor, err := messageBoxCol.Aggregate(Ctx,
// 		mongo.Pipeline{matchStage, lookupStage})
// 	if err != nil {
// 		return nil, err
// 	}

// 	var a []MessageBox
// 	if err = showLoadedCursor.All(Ctx, &a); err != nil {
// 		return nil, err

// 	}
// 	log.Println(a)
// 	return a[0].Messages, err
// }