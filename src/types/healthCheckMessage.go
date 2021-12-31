package types

type HealthCheck struct {
	Status string `json:"status"`
	Problems []string `json:"problems"`
}