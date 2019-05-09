// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 43.
//!+

// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch2/lenconv"
	"gopl.io/ch2/tempconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "cf: no arguments given.\nUsage: cf TYPE NUM [NUM...]\n")
		os.Exit(1)
	}

	t := os.Args[1]
	nums := stringToFloats(os.Args[2:])

	switch t {
	case "temp":
		convertTemp(nums)
	case "dist":
		convertDist(nums)
	}
}

func stringToFloats(s []string) []float64 {
	nums := make([]float64, len(s))
	for i, s := range s {
		n, err := strconv.ParseFloat(s, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		nums[i] = n
	}
	return nums
}

func convertDist(nums []float64) {
	for _, n := range nums {
		mi := lenconv.Mile(n)
		km := lenconv.Kilometer(n)
		fmt.Printf("%s = %s, %s = %s\n",
			mi, lenconv.MiToKm(mi), km, lenconv.KmToMi(km))
	}
}

func convertTemp(nums []float64) {
	for _, n := range nums {
		f := tempconv.Fahrenheit(n)
		c := tempconv.Celsius(n)
		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}
