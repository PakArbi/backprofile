package parkiran

import (
	"context"
	"fmt"
	"os"

	
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// mongodb

func GetConnectionMongo(MONGOSTRING, dbname string) (*mongo.Database, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv(MONGOSTRING)))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}
	return client.Database(dbname), nil
}

func SetConnection(MONGOSTRINGENV, dbname string) (*mongo.Database, error) {
    client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv(MONGOSTRINGENV)))
    if err != nil {
        return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
    }
    return client.Database(dbname), nil
}


func MongoConnect(MONGOSTRINGENV, dbname string) *mongo.Database {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv(MONGOSTRINGENV)))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
		return nil
	}
	return client.Database(dbname)
}

// Parkiran
func CreateNewParkiran(mongoconn *mongo.Database, collection string, parkirandata Parkiran) (*mongo.InsertOneResult, error) {
	coll := mongoconn.Collection(collection)
	result, err := coll.InsertOne(context.TODO(), parkirandata)
	if err != nil {
		return nil, fmt.Errorf("failed to insert document: %v", err)
	}
	return result, nil
}

func InsertParkiranData(db *mongo.Database, collectionName string, parkiranData Parkiran) error {
    collection := db.Collection(collectionName)
    
    _, err := collection.InsertOne(context.Background(), parkiranData)
    if err != nil {
        return err
    }

    return nil
}

func InsertParkiranDataToDB(db *mongo.Database, collectionName string, parkiranData Parkiran) error {
    collection := db.Collection(collectionName)
    _, err := collection.InsertOne(context.Background(), parkiranData)
    return err
}

func DeleteParkiran(mongoconn *mongo.Database, collection string, parkiranID int) (*mongo.DeleteResult, error) {
	coll := mongoconn.Collection(collection)
	filter := bson.M{"parkiranid": parkiranID}
	result, err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, fmt.Errorf("failed to delete document: %v", err)
	}
	return result, nil
}

func UpdateParkiran(mongoconn *mongo.Database, collection string, parkiranID int, parkirandata Parkiran) (*mongo.UpdateResult, error) {
	coll := mongoconn.Collection(collection)
	filter := bson.M{"parkiranid": parkiranID}
	update := bson.M{"$set": parkirandata}
	result, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, fmt.Errorf("failed to update document: %v", err)
	}
	return result, nil
}

func GetAllParkiran(mongoconn *mongo.Database, collection string) ([]Parkiran, error) {
	coll := mongoconn.Collection(collection)
	cursor, err := coll.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve documents: %v", err)
	}
	defer cursor.Close(context.TODO())

	var parkirans []Parkiran
	if err := cursor.All(context.TODO(), &parkirans); err != nil {
		return nil, fmt.Errorf("failed to decode documents: %v", err)
	}
	return parkirans, nil
}

func GetParkiranByID(mongoconn *mongo.Database, collection string, parkiranID int) (*Parkiran, error) {
	coll := mongoconn.Collection(collection)
	filter := bson.M{"parkiranid": parkiranID}
	var parkiran Parkiran
	if err := coll.FindOne(context.TODO(), filter).Decode(&parkiran); err != nil {
		return nil, fmt.Errorf("failed to retrieve document: %v", err)
	}
	return &parkiran, nil
}

func CreateResponse(status bool, message string, data interface{}) Response {
	return Response{
		Status: status,
        Message: message,
        Data:    data,
    }
}

// SaveQRCodeToMongoDB simulates the function to save the generated QR code to MongoDB
func SaveQRCodeToMongoDB(database *mongo.Database, collectionName string, qrCode []byte) error {
    collection := database.Collection(collectionName)
    ctx := context.Background()
    
    // Simpan kode QR ke dalam database MongoDB
    _, err := collection.InsertOne(ctx, bson.M{"qrcode": qrCode})
    if err != nil {
        return fmt.Errorf("failed to save QR code to MongoDB: %v", err)
    }

    return nil
}
