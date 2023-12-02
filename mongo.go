package backprofile

import (
	"context"
	"time"
	"bytes"
	"image"
	"image/png"

	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	
)

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
