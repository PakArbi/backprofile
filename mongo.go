package backprofile

import (
	"context"
	"time"
	"bytes"
	"os"
	"image"
	"image/png"

	"github.com/aiteung/atdb"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	
)

func GetConnectionMongo(MongoString, dbname string) *mongo.Database {
	MongoInfo := atdb.DBInfo{
		DBString: os.Getenv(MongoString),
		DBName:   dbname,
	}
	conn := atdb.MongoConnect(MongoInfo)
	return conn
}

func GetAllData(MongoConnect *mongo.Database, colname string) []Profile {
	data := atdb.GetAllDoc[[]Profile](MongoConnect, colname)
	return data
}

func InsertDataProfile(MongoConn *mongo.Database, colname string, coordinate []float64, id, namaLengkap, npm, prodi, namakendaraan, nomorkendaraan, timeString string) (InsertedID interface{}) {
	req := new(Profile)
	req.ID = id
	req.NamaLengkap = namaLengkap
	req.NPM = npm
	req.NamaKendaraan = namakendaraan
	req.NomorKendaraan = nomorkendaraan
	req.Time = Time{Message: "Message", WaktuMasuk: timeString}
	req.Time = Time{Message: "Message", WaktuKeluar: timeString}

	ins := atdb.InsertOneDoc(MongoConn, colname, req)
	return ins
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
