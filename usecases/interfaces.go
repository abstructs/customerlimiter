package usecases

import "kohotakehome.com/m/domain"

// LimiterUsecase handles limiting loading a customer's balance ledger
type LimiterUsecase interface {
	GenerateOutput(customerEvents []domain.CustomerLoadEvent) []domain.OutputLoadEvent
}
