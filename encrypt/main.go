package main

import (
	"flag"
	"log"

	pdf "github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

func main() {
	var filepath = flag.String("file", "", "file path")
	var password = flag.String("password", "", "password")
	flag.Parse()
	log.Println("file ", *filepath)
	log.Println("password ", *password)
	conf := pdfcpu.NewDefaultConfiguration()
	conf.OwnerPW = *password
	conf.UserPW = *password
	conf.EncryptUsingAES = true
	conf.EncryptKeyLength = 256
	conf.Cmd = pdfcpu.ENCRYPT
	err := pdf.OptimizeFile(*filepath, "./output.pdf", conf)
	if err != nil {
		log.Fatal(err)
	}
}
