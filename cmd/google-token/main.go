package main

import "github.com/lazybark/go-helpers/gapi"

func main() {
	gapi.GetTokenSheetsRead([]string{"https://www.googleapis.com/auth/spreadsheets"})
}
