package utils

import (
	"fmt"
	"github.com/skip2/go-qrcode"
	"image/png"
	"os"
)

func GenerateQr(content, path string, size int) {
	//fmt.Println(content)
	//fmt.Println(path)

	code, err := qrcode.New(content, qrcode.Medium)
	if err != nil {
		panic(err)
	}

	file, _ := os.Create(path)
	defer file.Close()

	// Encode the QR code as a PNG image and write it to the file
	err = png.Encode(file, code.Image(size))
	if err != nil {
		fmt.Println("Error encoding QR code as PNG:", err)
		return
	}

	fmt.Println("QR code generated and saved to qrcode.png")

}
