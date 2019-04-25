package main

import (
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch2/tempconv"
)

func main() {
	fmt.Printf("Hello World\n")
	c := tempconv.Celsius(3)

	fmt.Printf("Hello, it is cold today: %s\n", c)

	for i, arg := range os.Args[1:] {
		num, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "brendan_main: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Input #%d:\n", i+1)

		c = tempconv.Celsius(num)
		fmt.Printf("%s is %s is %s\n", c, tempconv.CToF(c), tempconv.CToK(c))

		f := tempconv.Fahrenheit(num)
		fmt.Printf("%s is %s is %s\n", f, tempconv.FToC(f), tempconv.FToK(f))
	}
}
