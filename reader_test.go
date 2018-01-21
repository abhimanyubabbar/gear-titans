package main

import (
	"testing"
)

var testfilename = "test.xlsx"

func TestReadMasterData(t *testing.T) {

	records, err := GetMasterRecords(testfilename, "Master")

	if err != nil {
		t.Fail()
	}

	if len(records) != 1 {
		t.Fail()
	}

	expected := MasterRecord{
		SKU:         "10319301-7",
		UPC:         float64(4053984348596),
		MRP:         "3999",
		Description: "METEOR SALA",
	}

	for _, v := range records {
		if v != expected {
			t.Fail()
		}
	}
}

func TestInboundRecords(t *testing.T) {

	records, err := GetInboundRecords(testfilename, "Inbound")
	// records, err := GetInboundRecords("Gear Tiatans.xlsx", "inboundinbound  puchase")
	if err != nil {
		t.Fail()
	}

	_, ok := records[0]
	if ok {
		t.Fail()
	}

	resp, ok := records[float64(884497053745)]
	if !ok {
		t.Fail()
	}

	if resp != 1 {
		t.Fail()
	}
}

func TestOutboundRecords(t *testing.T) {

	records, err := GetOutboundRecords(testfilename, "Outbound")
	// records, err := GetInboundRecords("Gear Tiatans.xlsx", "inboundinbound  puchase")
	if err != nil {
		t.Fail()
	}

	_, ok := records[0]
	if ok {
		t.Fail()
	}

	resp, ok := records[float64(91205732315)]

	if !ok {
		t.Fail()
	}

	if resp != 1 {
		t.Fail()
	}
}
