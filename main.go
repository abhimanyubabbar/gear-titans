package main

import (
	"log"
)

var (
	filename    = "Gear Tiatans.xlsx"
	newfilename = "Gear Tiatans new.xlsx"

	masterSheet    = "master till date live products"
	inboundSheet   = "inbound  puchase"
	outboundSheet  = "OUT BOUND"
	inventorySheet = "inventory"
)

func main() {
	log.Printf("Starting the inventory creation operation")

	masterRecords, err := GetMasterRecords(filename, masterSheet)
	if err != nil {
		log.Fatalf("Unable to create a list of master records")
	}

	inbound, err := GetInboundRecords(filename, inboundSheet)
	if err != nil {
		log.Fatalf("Unable to create a list of inbound records")
	}

	outbound, err := GetOutboundRecords(filename, outboundSheet)
	if err != nil {
		log.Fatalf("Unable to create a list of outbound records")
	}

	for upc, _ := range inbound {

		if _, ok := masterRecords[upc]; !ok {
			log.Printf("Unable to locate the inbound record: %.0f in the master records\n", upc)
		}
	}

	for upc, _ := range outbound {

		if _, ok := masterRecords[upc]; !ok {
			log.Printf("Unable to locate the outbound record: %.0f in the master records\n", upc)
		}
	}

	inventory := CalculateInventory(inbound, outbound)
	// fmt.Printf("%v\n", inventory)

	err = WriteInventorySheet(
		filename, inventorySheet,
		inventory, masterRecords,
		newfilename)

	if err != nil {
		log.Fatalf("Unable to complete the inventory write process")
	}

	log.Printf("Finished the process of inventory calculation")
}
