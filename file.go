package namapackage

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mendapatkan connection string MongoDB dari environment variables
var MongoString string = os.Getenv("MONGOSTRING")

// Fungsi untuk menghubungkan ke MongoDB
func MongoConnect(dbname string) *mongo.Database {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		log.Fatalf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

// Fungsi untuk insert satu dokumen ke MongoDB
func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		log.Fatalf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

// Fungsi untuk insert data Regional
func InsertRegional(
	id_sistem string, regAsalP6 string, kantorAsalP6 string, nopendAsalP6 string, 
	tanggalBeritaAcara string, bulan string, week string, tahun string, regTujuanP6 string, 
	kantorTujuanP6 string, nopendTujuanP6 string, deskripsi string, dnln string, 
	nomorKiriman string, uraianBeritaAcara string, deskripsiIregularitas string, 
	rincianRootCause string, referensiRootCause string, tindakanPencegahan string, 
	correctiveAction string, locus string, namaNIKPegawai string, nomorEvidence string, 
	validasiRegional string, validasiPusat string) (InsertedID interface{}) {
	
	var regional Regional
	regional.ID_Sistem = id_sistem
	regional.RegAsalP6 = regAsalP6
	regional.KantorAsalP6 = kantorAsalP6
	regional.NopendAsalP6 = nopendAsalP6
	regional.TanggalBeritaAcara = tanggalBeritaAcara
	regional.Bulan = bulan
	regional.Week = week
	regional.Tahun = tahun
	regional.RegTujuanP6 = regTujuanP6
	regional.KantorTujuanP6 = kantorTujuanP6
	regional.NopendTujuanP6 = nopendTujuanP6
	regional.Deskripsi = deskripsi
	regional.DNLN = dnln
	regional.NomorKiriman = nomorKiriman
	regional.UraianBeritaAcara = uraianBeritaAcara
	regional.DeskripsiIregularitas = deskripsiIregularitas
	regional.RincianRootCause = rincianRootCause
	regional.ReferensiRootCause = referensiRootCause
	regional.TindakanPencegahan = tindakanPencegahan
	regional.CorrectiveAction = correctiveAction
	regional.Locus = locus
	regional.NamaNIKPegawai = namaNIKPegawai
	regional.NomorEvidence = nomorEvidence
	regional.ValidasiRegional = validasiRegional
	regional.ValidasiPusat = validasiPusat

	return InsertOneDoc("iregularitas", "regional", regional)
}

// Fungsi untuk mengambil data Regional berdasarkan ID Sistem
func GetRegionalFromID_Sistem(id_sistem string) (regional Regional, err error) {
	collection := MongoConnect("iregularitas").Collection("regional")
	filter := bson.M{"id_sistem": id_sistem}
	err = collection.FindOne(context.TODO(), filter).Decode(&regional)
	if err != nil {
		fmt.Printf("GetRegionalFromID_Sistem: %v\n", err)
		return regional, err
	}
	return regional, nil
}