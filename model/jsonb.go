package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type JSONB map[string]interface{}

func (j JSONB) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}

func (j *JSONB) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal JSONB value: %v", value)
	}
	return json.Unmarshal(bytes, j)
}

func (j JSONB) MarshalJSON() ([]byte, error) {
	if j == nil {
		return []byte("null"), nil
	}
	return json.Marshal((map[string]interface{})(j))
}

func (j *JSONB) UnmarshalJSON(b []byte) error {
	if j == nil {
		*j = make(JSONB)
	}
	return json.Unmarshal(b, (*map[string]interface{})(j))
}
