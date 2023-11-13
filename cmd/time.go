// cmd/time.go

package cmd

import (
	"fmt"
	"time"
)

// CustomTime allows parsing of the non-standard time formats used by the API.
type CustomTime struct {
	time.Time
}

// UnmarshalJSON parses the non-standard time format from the API into a CustomTime.
func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := string(b)

	// Remove quotes
	s = s[1 : len(s)-1]

	// You might need to adjust the format to match exactly what your API returns.
	// For example, "2006-01-02T15:04:05.000" is a guess based on the provided error message.
	// If your API returns times in UTC without a Z, you might need to add it manually or handle it here.
	t, err := time.Parse("2006-01-02T15:04:05.000", s)
	if err != nil {
		return err
	}

	ct.Time = t
	return nil
}

// MarshalJSON writes the time in the custom format for the API.
func (ct *CustomTime) MarshalJSON() ([]byte, error) {
	if ct.Time.IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format("2006-01-02T15:04:05.000"))), nil
}

// Then you use CustomTime in place of time.Time in your DeviceDefinition struct.
