package main

import "testing"

func TestWriteInventorySheet(t *testing.T) {

	master := map[float64]MasterRecord{
		124: MasterRecord{
			SKU:         "sku-1",
			MRP:         "1299",
			UPC:         124,
			Description: "desc1",
		},
		345: MasterRecord{
			SKU:         "sku-2",
			MRP:         "2999",
			UPC:         345,
			Description: "desc2",
		},
	}

	inventory := map[float64]int{
		124: 2,
		345: 1,
		789: 0,
	}

	err := WriteInventorySheet(testfilename, "inventory", inventory, master, testfilename)
	if err != nil {
		t.Fail()
	}
}
