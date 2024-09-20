package main

import (
	"fmt"
	"math-skills/utils"
	"os"
)

func main() {
	args := os.Args
	if len(args) > 2 {
		fmt.Fprintln(os.Stderr, "Too many arguments")
		return
	} else if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "Invalid Arguments")
	}

	inputFile, err := os.ReadFile(args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	formatedTxt := utils.Format(inputFile)
	list, err := utils.Str2Int(formatedTxt)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Invalid List")
		return
	}
	average := utils.Average(list)
	median := utils.Median(list)
	variance := utils.Variance(list)
	stdDeviation := utils.StandardDeviation(list)
	fmt.Println("Average:", average, "\nMedian:", median, "\nVariance:", variance, "\nStandard Deviation:", stdDeviation)

}
