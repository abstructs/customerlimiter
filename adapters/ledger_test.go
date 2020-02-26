package adapters_test

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"kohotakehome.com/m/adapters"
	"kohotakehome.com/m/domain"
)

func Test_Limiter(t *testing.T) {
	t.Run("adds an amount to a customers daily amount", func(t *testing.T) {
		event := &domain.CustomerLoadEvent{
			CustomerID: "1",
			ID:         "1",
			LoadAmount: decimal.NewFromInt(1000),
		}

		ledger := adapters.NewBalanceLedger()

		ledger.LoadDailyAmount(event)

		assert.Equal(t, event.LoadAmount, ledger.AmountForDay(event))
		assert.Equal(t, 1, ledger.TimesLoadedForDay(event))
	})

	t.Run("adds more than one amount to a customers daily amount", func(t *testing.T) {
		today := time.Now()
		event1 := &domain.CustomerLoadEvent{
			CustomerID: "1",
			ID:         "1",
			LoadAmount: decimal.NewFromInt(1000),
			Time:       today,
		}
		event2 := &domain.CustomerLoadEvent{
			CustomerID: "1",
			ID:         "2",
			LoadAmount: decimal.NewFromInt(3000),
			Time:       today,
		}

		tomorrow := today.AddDate(0, 0, 1)
		event3 := &domain.CustomerLoadEvent{
			CustomerID: "1",
			ID:         "3",
			LoadAmount: decimal.NewFromInt(3000),
			Time:       tomorrow,
		}

		expectedAmount := event1.LoadAmount.Add(event2.LoadAmount)

		ledger := adapters.NewBalanceLedger()

		ledger.LoadDailyAmount(event1)
		ledger.LoadDailyAmount(event2)

		assert.Equal(t, ledger.AmountForDay(event2), expectedAmount)
		assert.Equal(t, 2, ledger.TimesLoadedForDay(event2))

		assert.Equal(t, decimal.Zero, ledger.AmountForDay(event3))
		assert.Equal(t, 0, ledger.TimesLoadedForDay(event3))
	})

	t.Run("adds an amount to a customers weekly amount", func(t *testing.T) {
		event := &domain.CustomerLoadEvent{
			CustomerID: "1",
			ID:         "1",
			LoadAmount: decimal.NewFromInt(1000),
		}

		ledger := adapters.NewBalanceLedger()

		ledger.LoadWeeklyAmount(event)

		assert.Equal(t, event.LoadAmount, ledger.AmountForWeek(event))
	})

	t.Run("adds more than one amount to a customers weekly amount", func(t *testing.T) {
		today := time.Now()
		event1 := &domain.CustomerLoadEvent{
			CustomerID: "1",
			ID:         "1",
			LoadAmount: decimal.NewFromInt(1000),
			Time:       today,
		}
		event2 := &domain.CustomerLoadEvent{
			CustomerID: "1",
			ID:         "2",
			LoadAmount: decimal.NewFromInt(3000),
			Time:       today,
		}

		nextWeek := today.AddDate(0, 1, 0)
		event3 := &domain.CustomerLoadEvent{
			CustomerID: "1",
			ID:         "3",
			LoadAmount: decimal.NewFromInt(3000),
			Time:       nextWeek,
		}

		expectedAmount := event1.LoadAmount.Add(event2.LoadAmount)

		ledger := adapters.NewBalanceLedger()

		ledger.LoadWeeklyAmount(event1)
		ledger.LoadWeeklyAmount(event2)

		assert.Equal(t, ledger.AmountForWeek(event2), expectedAmount)

		assert.Equal(t, decimal.Zero, ledger.AmountForWeek(event3))
	})
}
