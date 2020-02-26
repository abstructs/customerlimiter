package main

import (
	"fmt"
	"log"
	"os"

	"kohotakehome.com/m/adapters"
	"kohotakehome.com/m/usecases"
)

func main() {
	fmt.Println("Opening input file...")
	inputFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("unable to open input file")
	}
	defer inputFile.Close()

	fileDataLoader := adapters.NewFileDataLoader()
	fmt.Println("Reading inputs...")
	customers := fileDataLoader.ReadInput(inputFile)

	generateOutputUsecase := usecases.NewLimiterUsecase(usecases.LimiterUsecaseConfig{
		TimeBalanceLedger: adapters.NewBalanceLedger(),
	})

	fmt.Println("Generating output file...")
	output := generateOutputUsecase.GenerateOutput(customers)

	outputFile, err := os.Create("submission.txt")
	if err != nil {
		log.Fatal("unable to open output file")
	}
	defer outputFile.Close()

	err = fileDataLoader.WriteOutput(outputFile, output)
	if err != nil {
		log.Fatal("failed to write output")
	}
}
