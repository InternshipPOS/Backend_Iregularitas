package namapackage

import (
	"fmt"
	"testing"
)

func TestInsertRegional(t *testing.T) {
	// Membuat data dummy untuk insert
	id_sistem := "109002024091810201423"
	regAsalP6 := "Asal1"
	kantorAsalP6 := "Kantor Asal1"
	nopendAsalP6 := "NOP001"
	tanggalBeritaAcara := "2024-09-30"
	bulan := "September"
	week := "4"
	tahun := "2024"
	regTujuanP6 := "Tujuan1"
	kantorTujuanP6 := "Kantor Tujuan1"
	nopendTujuanP6 := "NOP002"
	deskripsi := "Deskripsi Test"
	dnln := "DN"
	nomorKiriman := "KRM123"
	uraianBeritaAcara := "Uraian Berita Test"
	deskripsiIregularitas := "Tidak ada"
	rincianRootCause := "Human Error"
	referensiRootCause := "Dokumen A"
	tindakanPencegahan := "Training"
	correctiveAction := "Penggantian Sistem"
	locus := "Area X"
	namaNIKPegawai := "123456"
	nomorEvidence := "EVID001"
	validasiRegional := "Valid"
	validasiPusat := "Valid"

	// Melakukan insert data menggunakan fungsi InsertRegional
	insertedID := InsertRegional(
		id_sistem, regAsalP6, kantorAsalP6, nopendAsalP6, tanggalBeritaAcara,
		bulan, week, tahun, regTujuanP6, kantorTujuanP6, nopendTujuanP6, deskripsi,
		dnln, nomorKiriman, uraianBeritaAcara, deskripsiIregularitas, rincianRootCause,
		referensiRootCause, tindakanPencegahan, correctiveAction, locus, namaNIKPegawai,
		nomorEvidence, validasiRegional, validasiPusat,
	)

	if insertedID == nil {
		t.Errorf("InsertRegional failed, insertedID is nil")
	} else {
		fmt.Printf("Inserted ID: %v\n", insertedID)
	}
}

func TestGetRegionalFromID_Sistem(t *testing.T) {
	id_sistem := "109002024091810201400"
	
	regional, err := GetRegionalFromID_Sistem(id_sistem)
	if err != nil {
		t.Errorf("GetRegionalFromID_Sistem failed: %v", err)
		return
	}

	if regional.ID_Sistem == "" {
		t.Errorf("GetRegionalFromID_Sistem failed, no data found for ID_Sistem: %s", id_sistem)
	} else {
		fmt.Printf("Data Regional: %+v\n", regional)
	}
}
