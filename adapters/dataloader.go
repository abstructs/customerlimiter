package adapters

import (
	"bufio"
	"encoding/json"
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
func (fdl *fileDataLoader) ReadInput(file *os.File) []domain.CustomerLoadEvent {
	var customerRes []domain.CustomerLoadEvent
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		customer := domain.CustomerLoadEvent{}
		json.Unmarshal(scanner.Bytes(), &customer)
		customerRes = append(customerRes, customer)
	}

	return customerRes
}
