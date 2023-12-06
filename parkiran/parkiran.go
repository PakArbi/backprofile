package parkiran

import (
	"encoding/json"
	"net/http"
	// "os"
	"fmt"

	qrcode "github.com/skip2/go-qrcode"
	// "github.com/whatsauth/watoken"
	// "go.mongodb.org/mongo-driver/bson"
)



func GCFCreateParkiran(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn, err := SetConnection(MONGOCONNSTRINGENV, dbname)
	if err != nil {
		return err.Error()
	}

	var dataparkiran Parkiran
	err = json.NewDecoder(r.Body).Decode(&dataparkiran)
	if err != nil {
		return err.Error()
	}

	result, err := CreateNewParkiran(mconn, collectionname, dataparkiran)
	if err != nil {
		return GCFReturnStruct(CreateResponse(true, fmt.Sprintf("Failed Create Parkiran: %v", err), dataparkiran))
	}

	return GCFReturnStruct(CreateResponse(false, fmt.Sprintf("Success Create Parkiran: %v", result.InsertedID), dataparkiran))
}

// Delete Parkiran
func GCFDeleteParkiran(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn, err := SetConnection(MONGOCONNSTRINGENV, dbname)
	if err != nil {
		return err.Error()
	}

	var dataparkiran struct {
		ParkiranID int `json:"parkiranid"`
	}
	err = json.NewDecoder(r.Body).Decode(&dataparkiran)
	if err != nil {
		return err.Error()
	}

	result, err := DeleteParkiran(mconn, collectionname, dataparkiran.ParkiranID)
	if err != nil {
		return GCFReturnStruct(CreateResponse(true, fmt.Sprintf("Failed Delete Parkiran: %v", err), dataparkiran))
	}

	if result.DeletedCount == 0 {
		return GCFReturnStruct(CreateResponse(false, "No matching document found to delete", dataparkiran))
	}

	return GCFReturnStruct(CreateResponse(false, "Success Delete Parkiran", dataparkiran))
}

// Update Parkiran
func GCFUpdateParkiran(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn, err := SetConnection(MONGOCONNSTRINGENV, dbname)
	if err != nil {
		return err.Error()
	}

	var dataparkiran Parkiran
	err = json.NewDecoder(r.Body).Decode(&dataparkiran)
	if err != nil {
		return err.Error()
	}

	result, err := UpdateParkiran(mconn, collectionname, dataparkiran.ParkiranId, dataparkiran)
	if err != nil {
		return GCFReturnStruct(CreateResponse(true, fmt.Sprintf("Failed Update Parkiran: %v", err), dataparkiran))
	}

	if result.ModifiedCount == 0 {
		return GCFReturnStruct(CreateResponse(false, "No matching document found to update", dataparkiran))
	}

	return GCFReturnStruct(CreateResponse(false, "Success Update Parkiran", dataparkiran))
}


// Get All Parkiran
func GCFGetAllParkiran(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn, err := SetConnection(MONGOCONNSTRINGENV, dbname)
	if err != nil {
		return err.Error()
	}

	dataparkiran, err := GetAllParkiran(mconn, collectionname)
	if err != nil {
		return GCFReturnStruct(CreateResponse(false, fmt.Sprintf("Failed Get All Parkiran: %v", err), dataparkiran))
	}

	return GCFReturnStruct(CreateResponse(true, "Success Get All Parkiran", dataparkiran))
}

func generateCodeQR(parkiran Parkiran) ([]byte, error) {
    // Convert data to JSON
    jsonData, err := json.Marshal(parkiran)
    if err != nil {
        return nil, fmt.Errorf("failed to convert data to JSON: %v", err)
    }

    // Generate QR code
    qrCode, err := qrcode.Encode(string(jsonData), qrcode.Medium, 256)
    if err != nil {
        return nil, fmt.Errorf("failed to generate QR code: %v", err)
    }

    return qrCode, nil
}

// GCFPostParkiran is an example of an HTTP request handler function
func GCFPostParkiran(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
    mconn, err := SetConnection(MONGOCONNSTRINGENV, dbname)
    if err != nil {
        return err.Error()
    }

    var parkiranData Parkiran // Gantilah "Parkiran" dengan struktur data yang sesuai

    // Mendekode data dari body request menjadi variabel parkiranData
    err = json.NewDecoder(r.Body).Decode(&parkiranData)
    if err != nil {
        return GCFReturnStruct(CreateResponse(false, fmt.Sprintf("Failed to parse parkiran data: %v", err), nil))
    }

    // Memasukkan data parkiran ke dalam database
    err = InsertParkiranData(mconn, collectionname, parkiranData)
    if err != nil {
        return GCFReturnStruct(CreateResponse(false, fmt.Sprintf("Failed to insert parkiran data: %v", err), nil))
    }

    // Generate QR code
    qrCode, err := generateCodeQR(parkiranData)
    if err != nil {
        return GCFReturnStruct(CreateResponse(false, fmt.Sprintf("Failed to generate QR code: %v", err), nil))
    }

    // Simpan QR code ke MongoDB
    err = SaveQRCodeToMongoDB(mconn, "PakArbi", qrCode) // Ganti dengan fungsi yang sesuai
    if err != nil {
        return GCFReturnStruct(CreateResponse(false, fmt.Sprintf("Failed to save QR code to MongoDB: %v", err), nil))
    }

    // Create notification based on Parkiran data
    notification := Notifikasi{
        Status:  200,
        Message: "QR code generated successfully and saved to MongoDB",
        Data:    parkiranData,
    }

    return GCFReturnStruct(CreateResponse(true, "Success inserting parkiran data", notification))
}





// Get All Parkiran By Id
func GCFGetAllParkiranID(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn, err := SetConnection(MONGOCONNSTRINGENV, dbname)
	if err != nil {
		return err.Error()
	}

	var dataparkiran Parkiran
	err = json.NewDecoder(r.Body).Decode(&dataparkiran)
	if err != nil {
		return err.Error()
	}

	parkiran, err := GetParkiranByID(mconn, collectionname, dataparkiran.ParkiranId)
	if err != nil {
		return GCFReturnStruct(CreateResponse(false, fmt.Sprintf("Failed to Get ID Parkiran: %v", err), dataparkiran))
	}

	if parkiran != nil {
		return GCFReturnStruct(CreateResponse(true, "Success: Get ID Parkiran", parkiran))
	}

	return GCFReturnStruct(CreateResponse(false, "Failed to Get ID Parkiran", dataparkiran))
}


func GCFReturnStruct(DataStuct any) string {
	jsondata, _ := json.Marshal(DataStuct)
	return string(jsondata)
}

func ReturnStringStruct(Data any) string {
	jsonee, _ := json.Marshal(Data)
	return string(jsonee)
}