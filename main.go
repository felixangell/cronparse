package main

import (
	"fmt"
	"github.com/felixangell/cronparse/internal"
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

	stdOutWriter := func(out string) {
		fmt.Println(out)
	}
	internal.PrettyPrintCron(*result, stdOutWriter)
}