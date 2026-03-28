package types

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type DateType struct {
	time.Time
}

var DateFormat = "2006-01-02"

// For GORM
// Implement Scanner interface for GORM
func (dt *DateType) Scan(value any) error {
	if value == nil {
		dt.Time = time.Time{}
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		dt.Time = v
		return nil
	case []byte:
		return dt.UnmarshalText(string(v))
	case string:
		return dt.UnmarshalText(v)
	}

	return fmt.Errorf("cannot scan type %T into DateType", value)
}

// Implement Valuer interface for GORM
func (dt DateType) Value() (driver.Value, error) {
	if dt.Time.IsZero() {
		return nil, nil
	}
	return dt.Time, nil
}

// For JSON
// UnmarshalJSON For JSON deserialization
func (dt *DateType) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "null" || s == "" {
		dt.Time = time.Time{}
		return nil
	}
	t, err := time.Parse(DateFormat, s)
	if err != nil {
		return err
	}
	dt.Time = t
	return nil
}

// MarshalJSON for JSON serialization
func (dt DateType) MarshalJSON() ([]byte, error) {
	if dt.IsZero() {
		return []byte("null"), nil
	}
	return []byte(dt.Format(DateFormat)), nil
	// return fmt.Appendf(nil, `"%s"`, dt.Format(DateFormat)), nil
}

// For FORM
// UnmarshalText For Form fields deserialization
func (dt *DateType) UnmarshalText(text string) error {
	if text == "" {
		dt.Time = time.Time{}
		return nil
	}
	t, err := time.Parse(DateFormat, text)
	if err != nil {
		return err
	}
	dt.Time = t
	return nil
}

// MarshalText for Form fields serialization; and for template rendering (required for HTML forms)
func (dt DateType) MarshalText() ([]byte, error) {
	if dt.IsZero() {
		return []byte(""), nil
	}
	// return dt.Time.Format(DateFormat), nil
	return fmt.Appendf(nil, `"%s"`, dt.Format(DateFormat)), nil
}

func (dt DateType) String() string {
	if dt.IsZero() {
		return ""
	}
	return dt.Format(DateFormat)
}
