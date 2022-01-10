package main

import (
	"fmt"
	"log"
	"os"

	parse "github.com/felixangell/cronparse/pkg"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("Error: no input args. Example usage: ./cronparse '*/15 0 1,15 * 1-5 /usr/bin/find'")
	}

	input := os.Args[1]
	fmt.Println("cron input:", "'"+input+"'")
	result, err := parse.ParseCronString(input)
	if err != nil {
		log.Fatal(err)
	}
	prettyPrintCron(*result)
}

func prettyPrintCron(result parse.ParseResult) {
	//minute 0 15 30 45
	//hour 0
	//day of month 1 15
	//month 1 2 3 4 5 6 7 8 9 10 11 12
	//day of week 1 2 3 4 5
	//command /usr/bin/find

	dateStrings := []string{
		"minute", "hour", "day of month", "month", "day of week", "command",
	}

	for idx := 0; idx < int(parse.ExpressionIndexCount); idx++ {
		unit, ok := result.ExpressionNode.GetUnit(parse.ExpressionIndex(idx))
		if !ok {
			log.Fatal("Bad unit at index", idx)
		}
		fmt.Println(dateStrings[idx], unit)
	}
}
