package test

import (
	"generateSertifikat/utils"
	"testing"
)

func TestGenerateCode(t *testing.T) {
	var filepath = "../data/p2.json"
	data, _ := utils.ReadJson(filepath)

	for _, item := range data.Peserta {
		if item.Type == "EKSTERNAL" {
			utils.GenerateQr(item.Nosertifikat+"-"+item.Nama, "../asset/qrcode/"+item.IdSertifikat+".png", 256)
		}
	}
}
