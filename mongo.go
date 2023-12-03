package backprofile

import (
	"context"
	"time"
	"fmt"
	"bytes"
	"os"
	"image"
	"image/png"

	
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	
)

func GetConnectionMongo(MongoString, dbname string) (*mongo.Database, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv(MongoString)))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}
	db := client.Database(dbname)
	return db, nil
}

func GetAllDataProfile(MongoConnect *mongo.Database, colname string) ([]Profile, error) {
	cur, err := MongoConnect.Collection(colname).Find(context.Background(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to get data from MongoDB: %v", err)
	}
	defer cur.Close(context.Background())

	var data []Profile
	for cur.Next(context.Background()) {
		var profile Profile
		if err := cur.Decode(&profile); err != nil {
			return nil, fmt.Errorf("failed to decode data from MongoDB: %v", err)
		}
		data = append(data, profile)
	}
	if err := cur.Err(); err != nil {
		return nil, fmt.Errorf("cursor error while fetching data from MongoDB: %v", err)
	}
	return data, nil
}

func InsertDataProfile(MongoConn *mongo.Database, colname string, coordinate []float64, id, namaLengkap, npm, prodi, namakendaraan, nomorkendaraan, timeString string) (interface{}, error) {
	req := Profile{
		ID:               id,
		NamaLengkap:      namaLengkap,
		NPM:              npm,
		Prodi:            prodi,
		NamaKendaraan:    namakendaraan,
		NomorKendaraan:   nomorkendaraan,
		Time:             Time{WaktuMasuk: timeString},
	}

	insResult, err := MongoConn.Collection(colname).InsertOne(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("failed to insert data into MongoDB: %v", err)
	}
	return insResult.InsertedID, nil
}


func UpdateDataProfile(MongoConn *mongo.Database, colname, id, namaLengkap, npm, prodi, namakendaraan, nomorkendaraan, timeString string) error {
    filter := bson.M{"id": id}

    update := bson.M{
        "$set": bson.M{
            "nama":            namaLengkap,
            "npm":             npm,
            "prodi":           prodi,
            "namakendaraan":   namakendaraan,
            "nomorkendaraan":  nomorkendaraan,
            "time.waktumasuk": timeString,
            "time.waktukeluar": timeString,
        },
    }

    _, err := MongoConn.Collection(colname).UpdateOne(context.TODO(), filter, update)
    if err != nil {
        return err
    }

    return nil
}

func DeleteDataProfile(MongoConn *mongo.Database, colname string, ID string) (*mongo.DeleteResult, error) {
    filter := bson.M{"id": ID}
    del, err := MongoConn.Collection(colname).DeleteOne(context.TODO(), filter)
    if err != nil {
        return nil, err
    }
    return del, nil
}

func ConnectDB() (*mongo.Database, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://faisalTampan:9byL9bOl3rhqbSrO@soren.uwshwr6.mongodb.net/test"))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	db := client.Database("PakArbi") // menyimpan data di mongoDB
	return db, nil
}


// Fungsi untuk menyimpan gambar ke dalam MongoDB menggunakan GridFS
func saveImageToMongoDB(img image.Image, bucket *gridfs.Bucket, filename string) error {
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return err
	}

	uploadStream, err := bucket.OpenUploadStream(filename)
	if err != nil {
		return err
	}
	defer uploadStream.Close()

	_, err = uploadStream.Write(buf.Bytes())
	if err != nil {
		return err
	}

	return nil
}
