package backprofile

import (
	"context"
	"fmt"
	"bytes"
	// "time"
	"encoding/json"
	"encoding/base64"
	"image"
	"image/png"
	"strconv"
	"net/http"
	"os"

	// "github.com/nfnt/resize"
	"github.com/disintegration/imaging"
	qrcode "github.com/skip2/go-qrcode"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"

)

func GCHandlerFunc(Mongostring, dbname, colname string) (string, error) {
    koneksyen, err := GetConnectionMongo(Mongostring, dbname)
    if err != nil {
        return "", fmt.Errorf("failed to connect to MongoDB: %v", err)
    }

    dataprofile, err := GetAllDataProfile(koneksyen, colname)
    if err != nil {
        return "", fmt.Errorf("failed to get data from MongoDB: %v", err)
    }

    jsoncihuy, err := json.Marshal(dataprofile)
    if err != nil {
        return "", fmt.Errorf("failed to marshal data to JSON: %v", err)
    }

    return string(jsoncihuy), nil
}


func GCFPostDataProf(Mongostring, dbname, colname string, r *http.Request) string {
	req := new(Credents)
	conn, err := GetConnectionMongo(Mongostring, dbname)
	if err != nil {
		req.Status = strconv.Itoa(http.StatusInternalServerError)
		req.Message = "error connecting to MongoDB: " + err.Error()
		return ReturnStringStruct(req)
	}

	resp := new(Profile)
	err = json.NewDecoder(r.Body).Decode(&resp)
	if err != nil {
		req.Status = strconv.Itoa(http.StatusNotFound)
		req.Message = "error parsing application/json: " + err.Error()
		return ReturnStringStruct(req)
	}

	req.Status = strconv.Itoa(http.StatusOK)
	insResult, insErr := InsertDataProfile(conn, colname, nil /* Untuk koordinat ([]float64), berikan nilai yang sesuai */, resp.ID, resp.NamaLengkap, resp.NPM, resp.Prodi, resp.NamaKendaraan, resp.NomorKendaraan, resp.Time.WaktuMasuk)
	if insErr != nil {
		req.Message = "Failed to insert data: " + insErr.Error()
		return ReturnStringStruct(req)
	}

	req.Message = fmt.Sprintf("Berhasil Input data. ID yang disisipkan: %v", insResult)
	return ReturnStringStruct(req)
}

func GCFUpdateProfile(Mongostring, dbname, colname string, r *http.Request) string {
    req := new(Credents)
    resp := new(Profile)
    conn, err := GetConnectionMongo(Mongostring, dbname)
    if err != nil {
        req.Status = strconv.Itoa(http.StatusInternalServerError)
        req.Message = "error connecting to MongoDB: " + err.Error()
        return ReturnStringStruct(req)
    }

    err = json.NewDecoder(r.Body).Decode(&resp)
    if err != nil {
        req.Status = strconv.Itoa(http.StatusNotFound)
        req.Message = "error parsing application/json: " + err.Error()
        return ReturnStringStruct(req)
    }

    req.Status = strconv.Itoa(http.StatusOK)
    err = UpdateDataProfile(conn, colname, resp.ID, resp.NamaLengkap, resp.NPM, resp.Prodi, resp.NamaKendaraan, resp.NomorKendaraan, resp.Time.WaktuMasuk)
    if err != nil {
        req.Message = "Failed to update data: " + err.Error()
        return ReturnStringStruct(req)
    }

    req.Message = "Successfully updated data"
    return ReturnStringStruct(req)
}


func GCFDeleteDataProfile(Mongostring, dbname, colname string, r *http.Request) string {
    req := new(Credents)
    resp := new(Profile)
    conn, errConn := GetConnectionMongo(Mongostring, dbname)
    if errConn != nil {
        req.Status = strconv.Itoa(http.StatusInternalServerError)
        req.Message = "error connecting to MongoDB: " + errConn.Error()
        return ReturnStringStruct(req)
    }

    errDecode := json.NewDecoder(r.Body).Decode(&resp)
    if errDecode != nil {
        req.Status = strconv.Itoa(http.StatusNotFound)
        req.Message = "error parsing application/json: " + errDecode.Error()
        return ReturnStringStruct(req)
    }

    req.Status = strconv.Itoa(http.StatusOK)
    delResult, delErr := DeleteDataProfile(conn, colname, resp.ID)
    if delErr != nil {
        req.Status = strconv.Itoa(http.StatusInternalServerError)
        req.Message = "error deleting data: " + delErr.Error()
    } else {
        req.Message = fmt.Sprintf("Berhasil menghapus data. Jumlah data terhapus: %v", delResult.DeletedCount)
    }
    return ReturnStringStruct(req)
}



func CreateProfile(db *mongo.Database, profile Profile) error {
	collection := db.Collection("profiles")

	_, err := collection.InsertOne(context.Background(), profile)
	if err != nil {
		return err
	}

	return nil
}

func UpdateProfile(db *mongo.Database, id string, profile Profile) error {
	collection := db.Collection("profiles")

	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "nama", Value: profile.NamaLengkap},
			{Key: "npm", Value: profile.NPM},
			// Update other fields accordingly
		}},
	}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func DeleteProfile(db *mongo.Database, id string) error {
	collection := db.Collection("profiles")

	filter := bson.D{{Key: "_id", Value: id}}

	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}

func GetProfile(db *mongo.Database, id string) (Profile, error) {
	collection := db.Collection("profiles")

	filter := bson.D{{Key: "_id", Value: id}}

	var profile Profile
	err := collection.FindOne(context.Background(), filter).Decode(&profile)
	if err != nil {
		return profile, err
	}

	return profile, nil
}

func GetAllProfiles(db *mongo.Database) ([]Profile, error) {
	collection := db.Collection("profiles")

	filter := bson.D{{}}

	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	var profiles []Profile
	for cur.Next(context.Background()) {
		var profile Profile
		err := cur.Decode(&profile)
		if err != nil {
			return nil, err
		}

		profiles = append(profiles, profile)
	}

	return profiles, nil
}

// GenerateCodeQr menghasilkan kode QR dari data JSON dan menyimpannya di MongoDB
func GenerateCodeQr(dataJSON string, dbName, collectionName string, client *mongo.Client) error {
	code := &CodeQR{
		Data: dataJSON,
	}

	// Mengonversi data menjadi JSON
	jsonData, err := json.Marshal(code)
	if err != nil {
		return fmt.Errorf("gagal mengonversi data ke JSON: %v", err)
	}

	// Membuat kode QR dari data JSON
	qrCode, err := qrcode.Encode(string(jsonData), qrcode.Medium, 256)
	if err != nil {
		return fmt.Errorf("gagal membuat kode QR: %v", err)
	}

	// Menyimpan kode QR ke dalam MongoDB
	collection := client.Database(dbName).Collection(collectionName)
	ctx := context.Background()
	_, err = collection.InsertOne(ctx, bson.M{"qrcode": qrCode})
	if err != nil {
		return fmt.Errorf("gagal menyimpan kode QR ke MongoDB: %v", err)
	}

	return nil
}
	
//generateqrLogo
// GenerateCodeQRLogo generates a QR code with a logo and saves it to MongoDB
func GenerateCodeQRLogo(dataJSON string, dbName, collectionName string, client *mongo.Client) error {
    code := &CodeQR{
        Data: dataJSON,
    }

    // Convert data to JSON
    jsonData, err := json.Marshal(code)
    if err != nil {
        return fmt.Errorf("failed to convert data to JSON: %v", err)
    }

    // Generate QR code
    qrCode, err := qrcode.Encode(string(jsonData), qrcode.Medium, 256)
    if err != nil {
        return fmt.Errorf("failed to generate QR code: %v", err)
    }

    // Open ULBI logo image
    imagePath := "./img/logo_ulbi.png" // Replace with the correct path to the image
    logoFile, err := os.Open(imagePath)
    if err != nil {
        return fmt.Errorf("failed to open image file: %v", err)
    }
    defer logoFile.Close()

    // Decode logo image
    logoImg, _, err := image.Decode(logoFile)
    if err != nil {
        return fmt.Errorf("failed to decode logo image: %v", err)
    }

    // Resize logo to fit the QR code
    resizedLogo := imaging.Resize(logoImg, 80, 80, imaging.Lanczos)

    // Convert QR code to image format
    qrImage, err := qrcodeToImage(qrCode)
    if err != nil {
        return fmt.Errorf("failed to convert QR code to image: %v", err)
    }

    // Overlay logo on QR code
    qrWithLogo := imaging.OverlayCenter(qrImage, resizedLogo, 1.0)

    // Encode QR code with logo to PNG
    var buf bytes.Buffer
    err = png.Encode(&buf, qrWithLogo)
    if err != nil {
        return fmt.Errorf("failed to encode QR code with logo to PNG: %v", err)
    }

    // Encode PNG data to base64
    base64QR := base64.StdEncoding.EncodeToString(buf.Bytes())

    // Save QR code with logo (base64 encoded) to MongoDB
    collection := client.Database(dbName).Collection(collectionName)
    ctx := context.Background()
    _, err = collection.InsertOne(ctx, bson.M{"qrcode_logo": base64QR})
    if err != nil {
        return fmt.Errorf("failed to save QR code with logo to MongoDB: %v", err)
    }

    return nil
}

// GenerateQRCodeString generates a QR code from a string and encodes it as a URL-safe string
func GenerateQRCodeString(text string) (string, error) {
	// Generate QR code from the text
	qr, err := qrcode.Encode(text, qrcode.Medium, 256)
	if err != nil {
		return "", fmt.Errorf("failed to generate QR code: %v", err)
	}

	// Encode the QR code as a URL-safe string
	urlEncodedString := base64.URLEncoding.EncodeToString(qr)
	return urlEncodedString, nil
}

// qrcodeToImage converts QR code bytes to an image
func qrcodeToImage(qrCode []byte) (image.Image, error) {
    reader := bytes.NewReader(qrCode)
    img, _, err := image.Decode(reader)
    if err != nil {
        return nil, fmt.Errorf("failed to convert QR code to image: %v", err)
    }
    return img, nil
}


// Fungsi untuk mengonversi data JSON menjadi gambar QR code
func JSONToQRCodeImage(jsonData []byte) (image.Image, error) {
    // Buat QR code dari data JSON
    qrCode, err := qrcode.Encode(string(jsonData), qrcode.Medium, 256)
    if err != nil {
        return nil, err
    }

    // Konversi byte array QR code ke dalam gambar
    img, err := qrcodeToImage(qrCode)
    if err != nil {
        return nil, err
    }

    return img, nil
}







