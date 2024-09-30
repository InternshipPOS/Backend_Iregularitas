package namapackage

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Regional struct {
	ID           		primitive.ObjectID 	`bson:"_id,omitempty" json:"_id,omitempty"`
	ID_Sistem            string             `bson:"id_sistem,omitempty" json:"id_sistem,omitempty"`
	RegAsalP6            string             `bson:"reg_asal_p6,omitempty" json:"reg_asal_p6,omitempty"`
	KantorAsalP6         string             `bson:"kantor_asal_p6,omitempty" json:"kantor_asal_p6,omitempty"`
	NopendAsalP6         string             `bson:"nopend_asal_p6,omitempty" json:"nopend_asal_p6,omitempty"`
	TanggalBeritaAcara   string             `bson:"tanggal_berita_acara,omitempty" json:"tanggal_berita_acara,omitempty"`
	Bulan                string             `bson:"bulan,omitempty" json:"bulan,omitempty"`
	Week                 string             `bson:"week,omitempty" json:"week,omitempty"`
	Tahun                string             `bson:"tahun,omitempty" json:"tahun,omitempty"`
	RegTujuanP6          string             `bson:"reg_tujuan_p6,omitempty" json:"reg_tujuan_p6,omitempty"`
	KantorTujuanP6       string             `bson:"kantor_tujuan_p6,omitempty" json:"kantor_tujuan_p6,omitempty"`
	NopendTujuanP6       string             `bson:"nopend_tujuan_p6,omitempty" json:"nopend_tujuan_p6,omitempty"`
	Deskripsi            string             `bson:"deskripsi,omitempty" json:"deskripsi,omitempty"`
	DNLN                 string             `bson:"dnln,omitempty" json:"dnln,omitempty"`
	NomorKiriman         string             `bson:"nomor_kiriman,omitempty" json:"nomor_kiriman,omitempty"`
	UraianBeritaAcara    string             `bson:"uraian_berita_acara,omitempty" json:"uraian_berita_acara,omitempty"`
	DeskripsiIregularitas string            `bson:"deskripsi_iregularitas,omitempty" json:"deskripsi_iregularitas,omitempty"`
	RincianRootCause     string             `bson:"rincian_root_cause,omitempty" json:"rincian_root_cause,omitempty"`
	ReferensiRootCause   string             `bson:"referensi_root_cause,omitempty" json:"referensi_root_cause,omitempty"`
	TindakanPencegahan   string             `bson:"tindakan_pencegahan,omitempty" json:"tindakan_pencegahan,omitempty"`
	CorrectiveAction     string             `bson:"corrective_action,omitempty" json:"corrective_action,omitempty"`
	Locus                string             `bson:"locus,omitempty" json:"locus,omitempty"`
	NamaNIKPegawai       string             `bson:"nama_nik_pegawai,omitempty" json:"nama_nik_pegawai,omitempty"`
	NomorEvidence        string             `bson:"nomor_evidence,omitempty" json:"nomor_evidence,omitempty"`
	ValidasiRegional     string             `bson:"validasi_regional,omitempty" json:"validasi_regional,omitempty"`
	ValidasiPusat        string             `bson:"validasi_pusat,omitempty" json:"validasi_pusat,omitempty"`
}
