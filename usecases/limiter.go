package usecases

import (
	"kohotakehome.com/m/adapters"
	"kohotakehome.com/m/domain"
)

// LimiterUsecaseConfig contains dependencies for the LimiterUsecase
type LimiterUsecaseConfig struct {
	TimeBalanceLedger adapters.TimedBalanceLedger
}

// NewLimiterUsecase gets a LimiterUsecase
func NewLimiterUsecase(config LimiterUsecaseConfig) LimiterUsecase {
	return &limiter{
		timeBalanceLedger: config.TimeBalanceLedger,
		processed:         map[domain.CustomerID]map[string]bool{},
	}
}

type limiter struct {
	timeBalanceLedger adapters.TimedBalanceLedger
	processed         map[domain.CustomerID]map[string]bool
}

func (l *limiter) GenerateOutput(customerEvents []domain.CustomerLoadEvent) []domain.OutputLoadEvent {
	var res []domain.OutputLoadEvent
	for _, customerEvent := range customerEvents {
		if l.alreadyHandledEvent(&customerEvent) {
			continue
		}

		eventRes := l.handleEvent(&customerEvent)
		l.setHandledEvent(&customerEvent)
		res = append(res, eventRes)
	}

	return res
}

func (l *limiter) alreadyHandledEvent(customerLoadEvent *domain.CustomerLoadEvent) bool {
	eventsProcessed, ok := l.processed[customerLoadEvent.CustomerID]
	if !ok {
		return false
	}

	_, alreadyHandled := eventsProcessed[customerLoadEvent.ID]
	return alreadyHandled
}

func (l *limiter) setHandledEvent(customerLoadEvent *domain.CustomerLoadEvent) {
	_, ok := l.processed[customerLoadEvent.CustomerID]
	if !ok {
		l.processed[customerLoadEvent.CustomerID] = map[string]bool{
			customerLoadEvent.ID: true,
		}
	}

	l.processed[customerLoadEvent.CustomerID][customerLoadEvent.ID] = true

}

func (l *limiter) handleEvent(customerLoadEvent *domain.CustomerLoadEvent) domain.OutputLoadEvent {
	if !l.canFund(customerLoadEvent) {
		return domain.OutputLoadEvent{
			ID:         customerLoadEvent.ID,
			CustomerID: customerLoadEvent.CustomerID.String(),
			Accepted:   false,
		}
	}

	l.timeBalanceLedger.LoadDailyAmount(customerLoadEvent)
	l.timeBalanceLedger.LoadWeeklyAmount(customerLoadEvent)

	return domain.OutputLoadEvent{
		ID:         customerLoadEvent.ID,
		CustomerID: customerLoadEvent.CustomerID.String(),
		Accepted:   true,
	}
}

func (l *limiter) wouldOverdraftDay(customerLoadEvent *domain.CustomerLoadEvent) bool {
	return l.timeBalanceLedger.AmountForDay(customerLoadEvent).
		Add(customerLoadEvent.LoadAmount).
		GreaterThan(domain.DailyAmountLimit)
}

func (l *limiter) wouldOverdraftWeek(customerLoadEvent *domain.CustomerLoadEvent) bool {
	return l.timeBalanceLedger.AmountForWeek(customerLoadEvent).
		Add(customerLoadEvent.LoadAmount).
		GreaterThan(domain.WeeklyAmountLimit)
}

func (l *limiter) wouldExceedDailyDepositCount(customerLoadEvent *domain.CustomerLoadEvent) bool {
	return l.timeBalanceLedger.TimesLoadedForDay(customerLoadEvent)+1 > domain.DailyLoadLimit
}

func (l *limiter) canFund(customerLoadEvent *domain.CustomerLoadEvent) bool {
	return !l.wouldOverdraftDay(customerLoadEvent) &&
		!l.wouldOverdraftWeek(customerLoadEvent) &&
		!l.wouldExceedDailyDepositCount(customerLoadEvent)
}
