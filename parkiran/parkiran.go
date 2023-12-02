package parkiran

import (
	"encoding/json"
	"net/http"
	// "os"

	// "github.com/whatsauth/watoken"
	"go.mongodb.org/mongo-driver/bson"
)



// Post Parkiran
func GCFCreateParkiran(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var dataparkiran Parkiran
	err := json.NewDecoder(r.Body).Decode(&dataparkiran)
	if err != nil {
		return err.Error()
	}
	if err := CreateParkiran(mconn, collectionname, dataparkiran); err != nil {
		return GCFReturnStruct(CreateResponse(true, "Success Create Parkiran", dataparkiran))
	} else {
		return GCFReturnStruct(CreateResponse(false, "Failed Create Parkiran", dataparkiran))
	}
}

// Delete Parkiran
func GCFDeleteParkiran(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)

	var dataparkiran Parkiran
	err := json.NewDecoder(r.Body).Decode(&dataparkiran)
	if err != nil {
		return err.Error()
	}

	if err := DeleteParkiran(mconn, collectionname, dataparkiran); err != nil {
		return GCFReturnStruct(CreateResponse(true, "Success Delete Parkiran", dataparkiran))
	} else {
		return GCFReturnStruct(CreateResponse(false, "Failed Delete Parkiran", dataparkiran))
	}
}

// Update Parkiran
func GCFUpdateParkiran(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)

	var dataparkiran Parkiran
	err := json.NewDecoder(r.Body).Decode(&dataparkiran)
	if err != nil {
		return err.Error()
	}

	if err := UpdatedParkiran(mconn, collectionname, bson.M{"id": dataparkiran.ID}, dataparkiran); err != nil {
		return GCFReturnStruct(CreateResponse(true, "Success Update Parkiran", dataparkiran))
	} else {
		return GCFReturnStruct(CreateResponse(false, "Failed Update Parkiran", dataparkiran))
	}
}

// Get All Parkiran
func GCFGetAllParkiran(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	dataparkiran := GetAllParkiran(mconn, collectionname)
	if dataparkiran != nil {
		return GCFReturnStruct(CreateResponse(true, "success Get All Parkiran", dataparkiran))
	} else {
		return GCFReturnStruct(CreateResponse(false, "Failed Get All Parkiran", dataparkiran))
	}
}

// Get All Parkiran By Id
func GCFGetAllParkiranID(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)

	var dataparkiran Parkiran
	err := json.NewDecoder(r.Body).Decode(&dataparkiran)
	if err != nil {
		return err.Error()
	}

	parkiran := GetAllParkiranID(mconn, collectionname, dataparkiran)
	if parkiran != (Parkiran{}) {
		return GCFReturnStruct(CreateResponse(true, "Success: Get ID Parkiran", dataparkiran))
	} else {
		return GCFReturnStruct(CreateResponse(false, "Failed to Get ID Parkiran", dataparkiran))
	}
}



func GCFReturnStruct(DataStuct any) string {
	jsondata, _ := json.Marshal(DataStuct)
	return string(jsondata)
}

func ReturnStringStruct(Data any) string {
	jsonee, _ := json.Marshal(Data)
	return string(jsonee)
}