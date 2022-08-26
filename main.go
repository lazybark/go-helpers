package main

import (
	"fmt"

	"github.com/lazybark/go-helpers/mock"
)

func main() {
	w := mock.New()

	w.HeaderData["one"] = []string{"one", "one"}
	fmt.Println("Header is: ", w.HeaderData)
	w.Write([]byte{'s', 'o', 'm', 'e', '_', 'd', 'a', 't', 'a'})
	fmt.Println("Data is: ", *w.Data)
	w.WriteHeader(1)
	fmt.Println("Code is: ", *w.StatusCode)
	w.Flush()
	fmt.Println("Data is: ", *w.Data)
}
