package messages

import "testing"

func TestGreet(t *testing.T) {
	gerceklesen := Greet("Mustafa")
	beklenen := "Hello, Mustafa!\n"

	if gerceklesen != beklenen {
		t.Errorf("Beklenen durum olmadı. Gerceklesen : %q, beklenen : %q\n",gerceklesen,beklenen)
	}
}

func TestDepart(t *testing.T) {
	gerceklesen := Depart("Mustafa")
	beklenen := "Goodbye, Mustafa\n"

	if gerceklesen != beklenen {
		t.Errorf("Beklenen durum olmadı. Gerceklesen : %q, beklenen : %q\n",gerceklesen,beklenen)
	}
}

func TestFailureTypes(t *testing.T){
	t.Error("Hatayı tespit eder. Ama diğer testleri durdurmaz")
	t.Fatal("Fatal testin çalışmasını durdurur.")
	t.Error("Bu kod satırı asla çalışmayacak.")
}
