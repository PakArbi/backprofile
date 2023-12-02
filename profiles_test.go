package backprofile

import (
	"fmt"
	"testing"
	"time"
	
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

	// Check whether the profile is successfully created
	// Add your logic here to verify the creation if needed
}

// TestGenerateQRCodeString tests the GenerateQRCodeString function
func TestGenerateQRCodeString(t*testing.T) {
	// String to convert to QR code
	text := "Hello, this is a test QR code string!"

	// Generate QR code as a URL-safe string
	urlEncodedString, err := GenerateQRCodeString(text)
	if err != nil {
		fmt.Printf("Error generating QR code: %v", err)
		return
	}

	// Display the URL-safe string
	fmt.Println("URL-encoded QR code:", urlEncodedString)
}