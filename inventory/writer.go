package main

import (
	"fmt"
	"log"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func WriteInventorySheet(
	filename, sheetname string,
	inventory map[float64]int,
	master map[float64]MasterRecord,
	outputfile string) error {

	log.Println("Started with the writing of the inventory sheet")

	xlsx, err := excelize.OpenFile(filename)
	if err != nil {
		log.Printf("Unable to open file to add the inventory")
		return err
	}

	// Set clean slate
	xlsx.DeleteSheet(sheetname)
	xlsx.NewSheet(sheetname)

	// Push the headings.
	xlsx.SetCellStr(sheetname, "A1", "UPC")
	xlsx.SetCellStr(sheetname, "B1", "SKU")
	xlsx.SetCellStr(sheetname, "C1", "MRP")
	xlsx.SetCellStr(sheetname, "D1", "Description")
	xlsx.SetCellStr(sheetname, "E1", "Inventory Count")

	rowCount := 2
	for upc, itemCount := range inventory {

		if _, ok := master[upc]; !ok {
			log.Printf("Unable to locate the inventory item: %.0f in the master records",
				upc)
			continue
		}

		record := master[upc]

		xlsx.SetCellValue(sheetname, fmt.Sprintf("A%d", rowCount), record.UPC)
		xlsx.SetCellValue(sheetname, fmt.Sprintf("B%d", rowCount), record.SKU)
		xlsx.SetCellValue(sheetname, fmt.Sprintf("C%d", rowCount), record.MRP)
		xlsx.SetCellValue(sheetname, fmt.Sprintf("D%d", rowCount), record.Description)
		xlsx.SetCellValue(sheetname, fmt.Sprintf("E%d", rowCount), itemCount)

		rowCount += 1
	}

	err = xlsx.SaveAs(outputfile)
	if err != nil {
		return err
	}

	log.Printf("Finished with the writing of the inventory sheet in file: %s\n", outputfile)
	return nil
}
