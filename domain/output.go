package domain

// LoadEventOutput is the expected output
type LoadEventOutput struct {
	ID         string `json:"id"`
	CustomerID string `json:"customer_id"`
	Accepted   bool   `json:"accepted"`
}
