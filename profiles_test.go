package backprofile

import (
	// "image"
	"testing"
	"time"
	// profiles "github.com/PakArbi/backprofile"
) 

func TestCreateProfile(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatalf("Error connecting to the database: %v", err)
	}

	profile := Profile{
		// Buat data profil yang sesuai
		ID:               "1",
		NamaLengkap:      "Muhammad Faisal Ashshidiq",
		NPM:              "1214041",
		Prodi:            "Teknik Informatika",
		NamaKendaraan:    "Motor Yamaha",
		NomorKendaraan:   "D 3316 GXF",
		Time:             Time{Message: "Message", WaktuMasuk: time.Now().Format(time.RFC3339)},
	}

	err = CreateProfile(db, profile)
	if err != nil {
		t.Fatalf("Error creating profile: %v", err)
	}
	
}

// func TestQRCodeToImage(t *testing.T) {
//     // Contoh data QR code
//     qrCodeData := []byte{ /* Masukkan data QR code di sini */ }

//     // Konversi data QR code ke gambar
//     img, err := qrcodeToImage(qrCodeData)
//     if err != nil {
//         t.Fatalf("Failed to convert QR code to image: %v", err)
//     }

//     // Periksa apakah img bukan nil dan merupakan instance dari image.Image
//     if img == nil {
//         t.Fatalf("Expected non-nil image, got nil")
//     }

//     _, isImage := img.(image.Image)
//     if !isImage {
//         t.Fatalf("Expected image.Image type, got different type")
//     }

// }

// // Fungsi test untuk JSONToQRCodeImage
// func TestJSONToQRCodeImage(t *testing.T) {
//     // Data JSON untuk pengujian
//     jsonData := []byte(`{"ID": "1", "NamaLengkap": "Muhammad Faisal Ashshidiq", "NPM": "1214041", "Prodi": "Teknik Informatika","namaKendaraan": "Yamaha Mio Z", "nomorKendaraan":"D 6613 GXF"}`)

//     // Panggil fungsi yang akan diuji
//     img, err := profiles.JSONToQRCodeImage(jsonData)
//     if err != nil {
//         t.Fatalf("Failed to convert JSON to QR code image: %v", err)
//     }

//     // Lakukan pengujian terhadap hasil yang diharapkan
//     if img == nil {
//         t.Fatalf("Expected non-nil image, got nil")
//     }

// }