package domain

import (
	"encoding/json"
	"time"

	"github.com/shopspring/decimal"
)

// CustomerID ..
type CustomerID string

func (cid *CustomerID) String() string {
	if cid != nil {
		customerID := *cid
		return string(customerID)
	}
	return ""
}

// CustomerLoadEvent is the raw customer from the input file
type CustomerLoadEvent struct {
	ID         string     `json:"id"`
	CustomerID CustomerID `json:"customer_id"`
	LoadAmount decimal.Decimal
	Time       time.Time `json:"time"`
}

// UnmarshalJSON creates an alias to unmarshal the json into, then
// overrides "LoadAmount" with the decimal representation
func (cle *CustomerLoadEvent) UnmarshalJSON(data []byte) error {
	type Alias CustomerLoadEvent
	aux := &struct {
		LoadAmount string `json:"load_amount"`
		*Alias
	}{
		Alias: (*Alias)(cle),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	loadAmount, err := decimal.NewFromString(aux.LoadAmount[1:])
	if err != nil {
		return err
	}
	cle.LoadAmount = loadAmount
	return nil
}
