package main

import services "golang-csv-db/Services"

func main() {

	var filepath = "fixlets.csv"
	services.ReadCSV(filepath)
}
