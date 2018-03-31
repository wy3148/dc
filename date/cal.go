package date

import (
	"time"
)

//There is algothrim about this topic, basically it calculate the Julian day number
//for start date and end date, and then get the difference, we assume we do not know this
//formula, we implement in Go in a special way
//https://en.wikipedia.org/wiki/Julian_day
// JDN = (1461 × (Y + 4800 + (M − 14)/12))/4 +(367 × (M − 2 − 12 × ((M − 14)/12)))/12 − (3 × ((Y + 4900 + (M - 14)/12)/100))/4 + D − 32075
// jdStart := toJDN(start.year, start.month, start.day)
// jdEnd := toJDN(end.year, end.month, end.day)
// jdElpased := jdEnd-jdStart - 1

// if jdElpased > 0 {
// 	return jdElpased
// }
// return 0
//convert to Julian day number
// func toJDN(year, month, day int) int {
// 	return day - 32075 +
// 		1461*(year+4800+(month-14)/12)/4 +
// 		367*(month-2-(month-14)/12*12)/12 -
// 		3*((year+4900+(month-14)/12)/100)/4
// }

func caculate(start *date, end *date) int {

	//when we calcuale the days, we know the challenge is acutally from the leap year or nonleap
	//year, the days of month Feburary is different, besides, the start date and end state could be
	//in the same year, or same month or even same day, there are many test cases here,
	//Then I think if it's possible to calculate days without checking the year, the first
	//thing comes to me is UTC timestampe which is widely used in my project, however normally
	//UTC starts from year 1970, this code testing start from 1901, but in fact,
	//in UTC, the timestamp can be negative when the date is before the year 1970, even though
	//we rarely use UTC timestamp in this way(because PC/Unix started from 1970), nowadays 64-bit os system support this, and Golang
	//support it very well, so here is the simple calculation

	//UTC timestamp in seconds
	t1 := time.Date(start.year, time.Month(start.month), start.day, 0, 0, 0, 0, time.UTC).Unix()
	t2 := time.Date(end.year, time.Month(end.month), end.day, 0, 0, 0, 0, time.UTC).Unix()

	//end date is ealier, swap them
	if t1 > t2 {
		t1, t2 = t2, t1
	}

	elapsed := int((t2-t1)/(3600*24) - 1)

	if elapsed > 0 {
		return elapsed
	}

	return 0
}

// Elapsed caculate the eclapsed days
// start and end are within range: 01/01/1901 and 31/12/2999.
func Elapsed(start, end string) (int, error) {

	startD, err := parse(start)

	if err != nil {
		return -1, err
	}

	endD, err := parse(end)

	if err != nil {
		return -1, err
	}

	return caculate(startD, endD), nil
}
