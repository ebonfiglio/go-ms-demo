package domain

import (
	"encoding/json"
	"fmt"
)

type Job struct {
	ID             int64  `json:"id" db:"id"`
	Name           string `json:"name" db:"name"`
	OrganizationID int64  `json:"organization_id" db:"organization_id"`
}

func JsonToJob(input string) (*Job, error) {
	var j Job
	err := json.Unmarshal([]byte(input), &j)
	if err != nil {
		return nil, fmt.Errorf("invalid JSON: %w", err)
	}
	return &j, err
}
