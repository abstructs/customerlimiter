package adapters_test

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"kohotakehome.com/m/adapters"
	"kohotakehome.com/m/domain"
)

func Test_DataLoader(t *testing.T) {
	t.Run("writes the output file", func(t *testing.T) {
		outputFile, err := os.Create("../test_files/test_dataloader_write.txt")
		if err != nil {
			log.Fatal("unable to open output file")
		}
		defer outputFile.Close()

		fileDataLoader := adapters.NewFileDataLoader()

		outputData := []domain.OutputLoadEvent{
			domain.OutputLoadEvent{
				ID:         "1",
				CustomerID: "1",
				Accepted:   true,
			},
		}

		err = fileDataLoader.WriteOutput(outputFile, outputData)
		assert.NoError(t, err)
	})

	t.Run("reads the input file", func(t *testing.T) {
		inputFile, err := os.Open("../test_files/test_dataloader_read.txt")
		if err != nil {
			log.Fatal("unable to open test input file")
		}
		defer inputFile.Close()

		fileDataLoader := adapters.NewFileDataLoader()

		customerLoadEvents := fileDataLoader.ReadInput(inputFile)

		assert.Equal(t, 1, len(customerLoadEvents))
	})
}
