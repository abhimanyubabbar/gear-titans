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

	config, err := GetConfig()
	if err != nil {
		log.Fatalf("Unable to get the configuration")
	}

	masterRecords, err := GetMasterRecords(config.InputFile, config.MasterSheet)
	if err != nil {
		log.Fatalf("Unable to create a list of master records")
	}

	inbound, err := GetInboundRecords(config.InputFile, config.InboundSheet)
	if err != nil {
		log.Fatalf("Unable to create a list of inbound records")
	}

	outbound, err := GetOutboundRecords(config.InputFile, config.OutboundSheet)
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
		config.InputFile, config.InventorySheet,
		inventory, masterRecords,
		config.OutputFile)

	if err != nil {
		log.Fatalf("Unable to complete the inventory write process")
	}

	log.Printf("Finished the process of inventory calculation")
}
