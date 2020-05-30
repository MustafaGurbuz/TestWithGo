package main_test

import "testing"

func TestAddition(t *testing.T){
	gerceklesen := 2 + 2
	beklenen := 4
	if gerceklesen != beklenen {
		t.Errorf("Beklenmeyen sonuc. Gerceklesen : '%v', Beklenen : '%v'",gerceklesen,beklenen)
	}
}

func TestDivide(t *testing.T){
	gerceklesen := 100/5
	beklenen := 18
	if gerceklesen != beklenen {
		t.Errorf("Beklenmeyen sonuc. Gerceklesen : '%v', Beklenen : '%v'",gerceklesen,beklenen)
	}
}