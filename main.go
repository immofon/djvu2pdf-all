package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func djvu_to_pdf_ext(name string) string {
	return name[:len(name)-5] + ".pdf"
}

func main() {
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	djvu_list := make([]string, 0, 100)

	for _, file := range files {
		name := file.Name()
		if strings.HasSuffix(name, ".djvu") {
			pdf_name := djvu_to_pdf_ext(name)
			_, err := os.Stat(pdf_name)
			if err != nil {
				log.Printf("name: %s\n", name)
				djvu_list = append(djvu_list, name)
			}
		}
	}

	for i, djvu := range djvu_list {
		log.Printf("[%3d/%3d] %s\n", i+1, len(djvu_list), djvu)
		err := exec.Command("djvu2pdf", djvu, djvu_to_pdf_ext(djvu)).Run()
		if err != nil {
			log.Println(err)
			continue
		}
		os.Remove(djvu)
	}
}
