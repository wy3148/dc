package date

import (
	"testing"
)

func TestParse(t *testing.T) {

	input := "01/12/1970"

	day := 1
	month := 12
	year := 1970

	d, err := parse(input)

	if err != nil {
		t.Fatal(err)
	}

	if d.day != day || d.month != month ||
		d.year != year {
		t.Fatal("wrong date parsed:", d.day, d.month, d.year)
	}

	input = "29/02/2007"

	d, err = parse(input)

	if err == nil {
		t.Fatalf("date %s is invalid ,shouldn't parse it", input)
	}
}

func TestLeapYear(t *testing.T) {

	//not leap
	year := 2100

	if isLeapYear(year) {
		t.Fatalf("year %d is not leap year\n", year)
	}

	//leap
	year = 2008

	if !isLeapYear(year) {
		t.Fatalf("year %d is  leap year\n", year)
	}

	//not leap
	year = 1901

	if isLeapYear(year) {
		t.Fatalf("year %d is not leap year\n", year)
	}

	//not leap
	year = 2999

	if isLeapYear(year) {
		t.Fatalf("year %d is not leap year\n", year)
	}
}

func TestValidDate(t *testing.T) {

	var d date

	d.day = 1
	d.month = 1
	d.year = 1901

	if res, _ := isValidDate(&d); !res {
		t.Fatalf("date is valid: %d/%d/%d", d.day, d.month, d.year)
	}

	d.day = 31
	d.month = 12
	d.year = 2999

	if res, _ := isValidDate(&d); !res {
		t.Fatalf("date is valid:%d/%d/%d", d.day, d.month, d.year)
	}

	d.day = 29
	d.month = 2
	d.year = 2007

	if res, _ := isValidDate(&d); res {
		t.Fatalf("date is invalid:%d/%d/%d", d.day, d.month, d.year)
	}
}
