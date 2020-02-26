package adapters

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
	"kohotakehome.com/m/domain"
)

// NewBalanceLedger get a new balance ledger
func NewBalanceLedger() TimedBalanceLedger {
	return &ledger{
		dailyLedgers:  map[domain.CustomerID]map[day]balanceLedgerEntry{},
		weeklyLedgers: map[domain.CustomerID]map[week]balanceLedgerEntry{},
	}
}

type ledger struct {
	dailyLedgers  map[domain.CustomerID]map[day]balanceLedgerEntry
	weeklyLedgers map[domain.CustomerID]map[week]balanceLedgerEntry
}

// balanceLedgerEntry ..
type balanceLedgerEntry struct {
	cumulativeAmount decimal.Decimal
	depositCount     int
}

type day string
type week string

func (l *ledger) getDay(loadTime time.Time) day {
	yearNum := loadTime.Year()
	dayNum := loadTime.YearDay()
	return day(fmt.Sprintf("%d-%d", yearNum, dayNum))
}

func (l *ledger) getWeek(loadTime time.Time) week {
	yearNum, weekNum := loadTime.ISOWeek()
	return week(fmt.Sprintf("%d-%d", yearNum, weekNum))
}

func (l *ledger) LoadDailyAmount(event *domain.CustomerLoadEvent) {
	customerLedger, ok := l.dailyLedgers[event.CustomerID][l.getDay(event.Time)]
	if !ok {
		newDailyLedger := map[day]balanceLedgerEntry{
			l.getDay(event.Time): balanceLedgerEntry{
				cumulativeAmount: event.LoadAmount,
				depositCount:     1,
			},
		}
		l.dailyLedgers[event.CustomerID] = newDailyLedger
		return
	}

	l.dailyLedgers[event.CustomerID][l.getDay(event.Time)] = balanceLedgerEntry{
		cumulativeAmount: customerLedger.cumulativeAmount.Add(event.LoadAmount),
		depositCount:     customerLedger.depositCount + 1,
	}
}

func (l *ledger) LoadWeeklyAmount(event *domain.CustomerLoadEvent) {
	customerLedger, ok := l.weeklyLedgers[event.CustomerID][l.getWeek(event.Time)]
	if !ok {
		newWeeklyLedger := map[week]balanceLedgerEntry{
			l.getWeek(event.Time): balanceLedgerEntry{
				cumulativeAmount: event.LoadAmount,
				depositCount:     1,
			},
		}
		l.weeklyLedgers[event.CustomerID] = newWeeklyLedger
		return
	}

	l.weeklyLedgers[event.CustomerID][l.getWeek(event.Time)] = balanceLedgerEntry{
		cumulativeAmount: customerLedger.cumulativeAmount.Add(event.LoadAmount),
		depositCount:     customerLedger.depositCount + 1,
	}
}

func (l *ledger) TimesLoadedForDay(event *domain.CustomerLoadEvent) int {
	customerLedger, ok := l.dailyLedgers[event.CustomerID][l.getDay(event.Time)]
	if !ok {
		return 0
	}
	return customerLedger.depositCount
}

func (l *ledger) AmountForDay(event *domain.CustomerLoadEvent) decimal.Decimal {
	customerLedger, ok := l.dailyLedgers[event.CustomerID][l.getDay(event.Time)]
	if !ok {
		return decimal.Zero
	}

	return customerLedger.cumulativeAmount
}

func (l *ledger) AmountForWeek(event *domain.CustomerLoadEvent) decimal.Decimal {
	customerLedger, ok := l.weeklyLedgers[event.CustomerID][l.getWeek(event.Time)]
	if !ok {
		return decimal.Zero
	}

	return customerLedger.cumulativeAmount
}
