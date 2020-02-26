package domain

// OutputLoadEvent is the expected output
type OutputLoadEvent struct {
	ID         string `json:"id"`
	CustomerID string `json:"customer_id"`
	Accepted   bool   `json:"accepted"`
}
