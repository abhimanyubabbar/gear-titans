package main

import (
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func GetMasterRecords(filename, sheetname string) (map[float64]MasterRecord, error) {

	xlsx, err := excelize.OpenFile(filename)

	if err != nil {
		return nil, err
	}

	rows := xlsx.GetRows(sheetname)
	records := make(map[float64]MasterRecord)

	for _, row := range rows {

		if len(row) == 0 {
			continue
		}

		upc, err := strconv.ParseFloat(row[2], 64)
		if err != nil {
			continue
		}

		records[upc] = MasterRecord{
			SKU:         row[0],
			MRP:         row[1],
			UPC:         upc,
			Description: row[4],
		}
	}

	return records, nil
}

func GetInboundRecords(filename, sheetname string) (map[float64]int, error) {
	return getInventory(filename, sheetname, 0)
}

func GetOutboundRecords(filename, sheetname string) (map[float64]int, error) {
	return getInventory(filename, sheetname, 4)
}

func getInventory(filename, sheetname string, pkloc int) (map[float64]int, error) {

	xlsx, err := excelize.OpenFile(filename)

	if err != nil {
		return nil, err
	}

	records := make(map[float64]int)

	rows := xlsx.GetRows(sheetname)
	for _, row := range rows {

		val := row[pkloc]

		upc, err := strconv.ParseFloat(val, 64)
		if err != nil {
			continue
		}

		// fmt.Println("%v\n", row)
		// fmt.Println(upsc)

		if _, ok := records[upc]; !ok {
			records[upc] = 1
		} else {
			records[upc] = records[upc] + 1
		}

	}

	return records, nil
}
