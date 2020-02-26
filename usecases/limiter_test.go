package usecases_test

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"kohotakehome.com/m/adapters"
	"kohotakehome.com/m/domain"
	"kohotakehome.com/m/usecases"
)

func loadExpectedOutputFile() []domain.OutputLoadEvent {
	file, err := os.Open("./../output.txt")
	if err != nil {
		log.Fatal("unable to open output file")
	}
	defer file.Close()

	var outputRes []domain.OutputLoadEvent
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		output := domain.OutputLoadEvent{}
		json.Unmarshal(scanner.Bytes(), &output)
		outputRes = append(outputRes, output)
	}

	return outputRes
}

func Test_Limiter(t *testing.T) {
	inputFile, err := os.Open("./../input.txt")
	if err != nil {
		log.Fatal("unable to open input file")
	}
	defer inputFile.Close()

	fileDataLoader := adapters.NewFileDataLoader()

	customers := fileDataLoader.ReadInput(inputFile)

	t.Run("usecase output matches expected output file", func(t *testing.T) {
		expectedOutput := loadExpectedOutputFile()

		generateOutputUsecase := usecases.NewLimiterUsecase(usecases.LimiterUsecaseConfig{
			TimeBalanceLedger: adapters.NewBalanceLedger(),
		})

		actualOutput := generateOutputUsecase.GenerateOutput(customers)

		assert.Equal(t, expectedOutput, actualOutput)
	})
}
