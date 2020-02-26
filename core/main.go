package main

import (
	"log"
	"os"

	"kohotakehome.com/m/adapters"
	"kohotakehome.com/m/usecases"
)

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("unable to open input file")
	}
	defer inputFile.Close()

	fileDataLoader := adapters.NewFileDataLoader()
	customers := fileDataLoader.ReadInput(inputFile)

	generateOutputUsecase := usecases.NewLimiterUsecase(usecases.LimiterUsecaseConfig{
		TimeBalanceLedger: adapters.NewBalanceLedger(),
	})

	generateOutputUsecase.GenerateOutputFile(customers)
}
