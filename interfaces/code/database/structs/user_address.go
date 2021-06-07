package structs

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

var (
	ErrCannotCastUserAddress = fmt.Errorf("cannot cast UserAddress")
)

type UserAddress struct {
	Valid  bool   `json:"-"`
	City   string `json:"city"`
	Street string `json:"street"`
	Home   int64  `json:"home"`
	Flat   int64  `json:"flat"`
}

// Scan implements the sql.Scanner interface.
func (a *UserAddress) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	switch data := value.(type) {
	case string:
		if data == "null" || data == "{}" {
			return nil
		}

		if err := json.Unmarshal([]byte(data), a); err != nil {
			return fmt.Errorf("cannot unmarshal UserAddress: %w", err)
		}
	case []byte:
		if string(data) == "null" || string(data) == "{}" {
			return nil
		}

		if err := json.Unmarshal(data, a); err != nil {
			return fmt.Errorf("cannot unmarshal UserAddress: %w", err)
		}
	default:
		return ErrCannotCastUserAddress
	}

	a.Valid = true

	return nil
}

// Value implements the driver driver.Valuer interface.
func (a UserAddress) Value() (driver.Value, error) {
	if !a.Valid {
		return nil, nil
	}

	value, err := json.Marshal(&a)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal UserAddress: %w", err)
	}

	return value, nil
}

var (
	_ driver.Valuer = (*UserAddress)(nil)
	_ sql.Scanner   = (*UserAddress)(nil)
)
