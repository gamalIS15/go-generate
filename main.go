package main

import (
	"fmt"
	dbs "generateSertifikat/utils"
)

type auth struct {
	Id   int
	Nip  string
	Nama string
}

func sqlQuery() {
	db, _ := dbs.Connect()
	defer db.Close()

	var rows, _ = db.Query(`SELECT id, nip,nama FROM v2_authorized`)

	var result []auth
	for rows.Next() {
		each := auth{}
		err := rows.Scan(&each.Id, &each.Nama, &each.Nip)
		if err != nil {
			fmt.Println("Error")
			panic(err)
		}

		result = append(result, each)
	}

	for _, each := range result {
		fmt.Println(each)
	}
}
func main() {

	//pdfd, err := pdf.NewPDFGenerator()
	//if err != nil {
	//	panic(err)
	//}
	//
	//page := pdf.NewPage("https://www.google.com/")
	//page.FooterRight.Set("[page]")
	//page.FooterFontSize.Set(10)
	//
	//pdfd.AddPage(page)
	//err = pdfd.Create()
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = pdfd.WriteFile("./output.pdf")
	//if err != nil {
	//	panic(err)
	//}
	//sqlQuery()
	var filepath = "./data/p2.json"

	dataa, _ := dbs.ReadJson(filepath)
	for _, item := range dataa.Peserta {
		if item.Type == "EKSTERNAL" {
			fmt.Println(item)
		}
	}

	fmt.Println(dataa.Peserta[0])
}
