package usecases

import "kohotakehome.com/m/domain"

// LimiterUsecase handles limiting loading a customer's balance ledger
type LimiterUsecase interface {
	GenerateOutputFile(customerEvents []domain.CustomerLoadEvent) []domain.LoadEventOutput
}
