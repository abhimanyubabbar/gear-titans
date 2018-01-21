package main

import (
	"log"
)

func CalculateInventory(inbound, outbound map[float64]int) map[float64]int {

	inventory := make(map[float64]int)

	for inItem, inCount := range inbound {
		inventory[inItem] = inCount
	}

	for outItem, outCount := range outbound {

		if inCount, ok := inventory[outItem]; ok {
			inventory[outItem] = inCount - outCount
		} else {
			log.Printf("Located item in outbound: %.0f not present in inbound\n", outItem)
		}
	}

	return inventory
}
