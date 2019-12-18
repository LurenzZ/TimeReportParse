package parser

import (
	"TimeReportParser/cmd/report"
	"errors"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"time"
)

const IN = "1- IN"
const OUT = "2- OUT"

func ParseExcel(filePath string) {
	excel, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get all the rows in the Sheet1.
	sheets := excel.GetSheetMap()
	rows := excel.GetRows(sheets[1])

	var reportRows report.Rows
	for i, row := range rows {
		if i == 0 {
			continue
		}
		rRow, err := validateAndParseRow(row)
		if err == nil {
			reportRows = append(reportRows, *rRow)
		}

	}
	if reportRows != nil {
		reportRows.NewReport()
	}
}

func validateAndParseRow(row []string) (*report.Row, error) {
	if len(row) < 4 {
		panic("Excel must contains [data,esito,operazione,utente]")
	}

	time, err := time.Parse("01/2/06 15:04", row[0])
	if err != nil {
		return nil, errors.New("invalid date")
	}
	operation := row[2]
	if operation != IN && operation != OUT {
		return nil, errors.New("invalid operation")
	}

	user := row[3]
	if user == "" {
		return nil, errors.New("invalid user")
	}

	return &report.Row{time, operation == IN, user}, nil
}
