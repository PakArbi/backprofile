package parkiran

import (
	"context"
	"fmt"
	"os"

	"github.com/aiteung/atdb"
	// "github.com/whatsauth/watoken"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// mongodb
func MongoConnect(MongoString, dbname string) *mongo.Database {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv(MongoString)))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func GetConnectionMongo(MongoString, dbname string) *mongo.Database {
	MongoInfo := atdb.DBInfo{
		DBString: os.Getenv(MongoString),
		DBName:   dbname,
	}
	conn := atdb.MongoConnect(MongoInfo)
	return conn
}

func SetConnection(MONGOCONNSTRINGENV, dbname string) *mongo.Database {
	var DBmongoinfo = atdb.DBInfo{
		DBString: os.Getenv(MONGOCONNSTRINGENV),
		DBName:   dbname,
	}
	return atdb.MongoConnect(DBmongoinfo)
}

// Parkiran
func CreateNewParkiran(mongoconn *mongo.Database, collection string, parkirandata Parkiran) interface{} {
	return atdb.InsertOneDoc(mongoconn, collection, parkirandata)
}

// Function Parkiran
func CreateParkiran(mongoconn *mongo.Database, collection string, parkirandata Parkiran) interface{} {
	return atdb.InsertOneDoc(mongoconn, collection, parkirandata)
}

func DeleteParkiran(mongoconn *mongo.Database, collection string, parkirandata Parkiran) interface{} {
	filter := bson.M{"parkiranid": parkirandata.ParkiranId}
	return atdb.DeleteOneDoc(mongoconn, collection, filter)
}

func UpdatedParkiran(mongoconn *mongo.Database, collection string, filter bson.M, parkirandata Parkiran) interface{} {
	filter = bson.M{"parkiranid": parkirandata.ParkiranId}
	return atdb.ReplaceOneDoc(mongoconn, collection, filter, parkirandata)
}

func GetAllParkiran(mongoconn *mongo.Database, collection string) []Parkiran {
	parkiran := atdb.GetAllDoc[[]Parkiran](mongoconn, collection)
	return parkiran
}

func GetAllParkiranID(mongoconn *mongo.Database, collection string, parkirandata Parkiran) Parkiran {
	filter := bson.M{
		"parkiranid":     parkirandata.ParkiranId,
		"nama":           parkirandata.Nama,
		"npm":            parkirandata.NPM,
		"jurusan":        parkirandata.Jurusan,
		"namakendaraan":  parkirandata.NamaKendaraan,
		"nomorkendaraan": parkirandata.NomorKendaraan,
		"jeniskendaraan": parkirandata.JenisKendaraan,
	}
	parkiranID := atdb.GetOneDoc[Parkiran](mongoconn, collection, filter)
	return parkiranID
}

func CreateResponse(status bool, message string, data interface{}) Response {
    return Response{
        Status: status,
        Message: message,
        Data:    data,
    }
}