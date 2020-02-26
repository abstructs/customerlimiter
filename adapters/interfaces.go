package adapters

import (
	"os"

	"github.com/shopspring/decimal"
	"kohotakehome.com/m/domain"
)

// DataLoader loads the data from the input into
// the customers slice
type DataLoader interface {
	ReadInput(inputFile *os.File) []domain.CustomerLoadEvent
	WriteOutput(outputFile *os.File, output []domain.OutputLoadEvent) error
}

// TimedBalanceLedger maintains a track of the
// customers's balance on a daily and weekly basis
type TimedBalanceLedger interface {
	LoadDailyAmount(event *domain.CustomerLoadEvent)
	LoadWeeklyAmount(event *domain.CustomerLoadEvent)
	TimesLoadedForDay(event *domain.CustomerLoadEvent) int
	AmountForDay(event *domain.CustomerLoadEvent) decimal.Decimal
	AmountForWeek(event *domain.CustomerLoadEvent) decimal.Decimal
}
