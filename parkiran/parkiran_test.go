package parkiran

import (
	"fmt"
	"testing"
)


func TestParkiran(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "PakarbiDB")
	var parkirandata Parkiran
	parkirandata.ParkiranId = 2
	parkirandata.Nama = "Muhammad Faisal Ashshidiq"
	parkirandata.NPM = "1214041"
	parkirandata.Jurusan = "D4 Teknik Informatika"
	parkirandata.NamaKendaraan = "Mio Z"
	parkirandata.NomorKendaraan = "D 3316 GXF"
	parkirandata.JenisKendaraan = "Motor"
	CreateNewParkiran(mconn, "parkiran", parkirandata)
}

func TestAllParkiran(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "PakarbiDB")
	parkiran := GetAllParkiran(mconn, "parkiran")
	fmt.Println(parkiran)
}