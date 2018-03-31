package main

import (
	"flag"
	"fmt"
	"github.com/wy3148/dc/date"
)

func main() {

	start := flag.String("start", "", "start time: DD/MM/YYYY")
	end := flag.String("end", "", "end time: DD/MM/YYYY")
	flag.Parse()

	if len(*start) == 0 || len(*end) == 0 {
		fmt.Println("Usage: ./dc -start=DD/MM/YYYY -end=DD/MM/YYYY")
		fmt.Println("Example: ./dc -start=01/02/2018 -end=02/02/2018")
		return
	}

	elapsed, err := date.Elapsed(*start, *end)

	if err != nil {
		fmt.Printf("failed to caculate the ecplased days:%s", err.Error())
		return
	}

	fmt.Printf("%s - %s: %d days\n", *start, *end, elapsed)
}
