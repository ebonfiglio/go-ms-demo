package domain

import (
	"encoding/json"
	"fmt"
)

type Organization struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func JsonToOrganization(input string) (*Organization, error) {
	var o Organization
	err := json.Unmarshal([]byte(input), &o)
	if err != nil {
		return nil, fmt.Errorf("invalid JSON: %w", err)
	}
	return &o, err
}
