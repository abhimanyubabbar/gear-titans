package main

import "testing"

func TestCalculateInventory(t *testing.T) {

	inbound := map[float64]int{
		123: 1,
		456: 2,
	}

	outbound := map[float64]int{
		456: 1,
		789: 1,
	}

	resp := CalculateInventory(inbound, outbound)

	expected := map[float64]int{
		123: 1,
		456: 1,
	}

	for k, v := range expected {

		if _, ok := resp[k]; !ok {
			t.Fail()
		}

		if resp[k] != v {
			t.Fail()
		}

	}
}
