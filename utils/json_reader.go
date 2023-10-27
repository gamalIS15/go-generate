package utils

import (
	"encoding/json"
	"os"
)

type Sertifikat struct {
	Status  int       `json:"status"`
	Peserta []Peserta `json:"peserta"`
}

type Peserta struct {
	IdSertifikat       string `json:"id_sertifikat"`
	Id                 string `json:"id"`
	IdPelaksanaan      string `json:"id_pelaksanaan"`
	Nip                string `json:"nip"`
	Nama               string `json:"nama"`
	Pretest            string `json:"pretest"`
	Posttest           string `json:"posttest"`
	Type               string `json:"type"`
	Status             string `json:"status"`
	Nosertifikat       string `json:"no_sertifikat"`
	NoSertifikatLan    string `json:"no_sertifikat_lan"`
	Nilai              string `json:"nilai"`
	Sprint             string `json:"sprint"`
	Sertifikat         string `json:"sertifikat"`
	HasilPembelajaran  string `json:"hasil_pembelajaran"`
	Instansi           string `json:"instansi"`
	Unit               string `json:"unit"`
	Jabatan            string `json:"jabatan"`
	IsSave             string `json:"isSave"`
	Jp                 string `json:"jp"`
	Tempkelulusan      string `json:"tempkelulusan"`
	SubPelatihan       string `json:"sub_pelatihan"`
	NamaFileSertifikat string `json:"nama_file_sertifikat"`
	IsSavedPrev        string `json:"isSavedPrev"`
	Email              string `json:"email"`
	Kelulusan          Kelulusan
}

type Kelulusan struct {
	Kelulusan int    `json:"kelulusan"`
	Nama      string `json:"nama"`
}

func ReadJson(path string) (Sertifikat, error) {
	data, _ := os.ReadFile(path)

	retData := Sertifikat{}
	err := json.Unmarshal(data, &retData)
	if err != nil {
		panic(err)
	}

	return retData, nil
}
