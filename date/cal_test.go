package date

import (
	"testing"
)

func TestEclapsed(t *testing.T) {

	// same year , same month
	expected := 19
	res, err := Elapsed("02/06/1983", "22/06/1983")

	if err != nil {
		t.Fatal(err)
	}

	//same year same month, one day
	expected = 0
	res, err = Elapsed("01/01/1984", "02/01/1984")

	if err != nil {
		t.Fatal(err)
	}

	if res != expected {
		t.Fatalf("expected %d days but return %d", expected, res)
	}

	//exactly same day
	expected = 0
	res, err = Elapsed("01/01/1984", "01/01/1984")

	if err != nil {
		t.Fatal(err)
	}

	if res != expected {
		t.Fatalf("expected %d days but return %d", expected, res)
	}

	//same year, same month, two days
	expected = 1
	res, err = Elapsed("01/01/1984", "03/01/1984")

	if err != nil {
		t.Fatal(err)
	}

	if res != expected {
		t.Fatalf("expected %d days but return %d", expected, res)
	}

	//typical, different year, month, day
	expected = 1979
	res, err = Elapsed("03/01/1989", "03/08/1983")

	if err != nil {
		t.Fatal(err)
	}

	if res != expected {
		t.Fatalf("expected %d days but return %d", expected, res)
	}

	//same year, different month, day
	expected = 173
	res, err = Elapsed("04/07/1984", "25/12/1984")

	if err != nil {
		t.Fatal(err)
	}

	if res != expected {
		t.Fatalf("expected %d days but return %d", expected, res)
	}

	//random
	expected = 761
	res, err = Elapsed("01/01/1901", "02/02/1903")

	if err != nil {
		t.Fatal(err)
	}

	if res != expected {
		t.Fatalf("expected %d days but return %d", expected, res)
	}

	//edge
	expected = 401400
	res, err = Elapsed("01/01/1901", "31/12/2999")

	if err != nil {
		t.Fatal(err)
	}

	if res != expected {
		t.Fatalf("expected %d days but return %d", expected, res)
	}

	//leap year
	expected = 365
	res, err = Elapsed("01/01/2008", "01/01/2009")

	if err != nil {
		t.Fatal(err)
	}

	if res != expected {
		t.Fatalf("expected %d days but return %d", expected, res)
	}
}

func BenchmarkElapsed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Elapsed("01/01/1901", "31/12/2999")
	}
}
