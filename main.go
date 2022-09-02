package main

import (
	"fmt"
	"log"

	"github.com/lazybark/go-helpers/conv"
	"github.com/lazybark/go-helpers/fsw"
)

func main() {
	csv, converted, err := conv.ConvertCSVFiletoJSON("Namings copy.csv", ";", "Русский", "Английский", "Украинский")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("converted rows:", converted)

	f, err := fsw.MakePathToFile("wiki.json", true)
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.Write(csv)
	if err != nil {
		log.Fatal(err)
	}
}
