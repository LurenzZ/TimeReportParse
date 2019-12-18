package main

import (
	"TimeReportParser/cmd/parser"
	"os"
)

func main() {
	fileToParse := os.Args[1]
	if fileToParse == ""{
		panic("No file provided to parser. Exit!")
	}

	parser.ParseExcel(fileToParse)
	
}
