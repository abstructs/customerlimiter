package adapters

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"kohotakehome.com/m/domain"
)

// NewFileDataLoader gets a new data loader
func NewFileDataLoader() DataLoader {
	return &fileDataLoader{}
}

// FileDataLoader is an implementation of DataLoader
type fileDataLoader struct{}

// ReadInput opens the input txt file and parses it into the customer struct
func (fdl *fileDataLoader) ReadInput(inputFile *os.File) []domain.CustomerLoadEvent {
	var customerRes []domain.CustomerLoadEvent
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		customer := domain.CustomerLoadEvent{}
		json.Unmarshal(scanner.Bytes(), &customer)
		customerRes = append(customerRes, customer)
	}

	return customerRes
}

// ReadInput opens the input txt file and parses it into the customer struct
func (fdl *fileDataLoader) WriteOutput(outputFile *os.File, outputEvents []domain.OutputLoadEvent) error {
	w := bufio.NewWriter(outputFile)
	for _, output := range outputEvents {
		jsonOutput, err := json.Marshal(output)
		if err != nil {
			return err
		}
		fmt.Fprintln(w, string(jsonOutput))
	}
	return w.Flush()
}
