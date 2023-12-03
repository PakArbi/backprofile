package backprofile

import "go.mongodb.org/mongo-driver/mongo"


type ResponseBack struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    []string `json:"data"`
}

type ResponseProfile struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    Profile  	`json:"data"`
}

type ProfileRepository struct {
    collection *mongo.Collection
}

type CodeQR struct {
	Data string `json:"data"`
}

type Profile struct {
	ID               string                 `json:"id" bson:"_id,omitempty"`
	NamaLengkap      string                 `json:"nama,omitempty" bson:"nama,omitempty"`
	NPM              string                 `json:"npm,omitempty" bson:"npm,omitempty"`
	Prodi            string                 `json:"prodi,omitempty" bson:"prodi,omitempty"`
	NamaKendaraan    string                 `json:"namakendaraan,omitempty" bson:"namakendaraan,omitempty"`
	NomorKendaraan   string                 `json:"nomorkendaraan,omitempty" bson:"nomorkendaraan,omitempty"`
	Time             Time                   `json:"time,omitempty" bson:"time,omitempty"`
}

type Time struct {
	Message     string `json:"message,omitempty" bson:"message,omitempty"`
	WaktuMasuk  string `json:"waktumasuk,omitempty" bson:"waktumasuk,omitempty"`
	WaktuKeluar string `json:"waktukeluar,omitempty" bson:"waktukeluar,omitempty"`
}

type Credential struct {
	Status  bool   `json:"status" bson:"status"`
	Token   string `json:"token,omitempty" bson:"token,omitempty"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}

type Credents struct {
	Status  string `json:"status" bson:"status"`
	Message string `json:"message" bson:"message"`
}

type Response struct {
	Status  bool        `json:"status" bson:"status"`
	Message string      `json:"message" bson:"message"`
	Data    interface{} `json:"data" bson:"data"`
}

type Prodi struct {
	ProdiId   int    `json:"jurusanid" bson:"jurusanid,omitempty"`
	ProdiName string `json:"jurusanname" bson:"jurusanname,omitempty"`
}

type Updated struct {
	Message string `json:"message"`
}

type RequestProfile struct {
	Username string `json:"username"`
	Npm      string `json:"npm"`
	Email    string `json:"email"`
}

