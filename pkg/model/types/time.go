package types

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

// NullTime is an alias for mysql.NullTime data type
type Time struct {
	sql.NullTime
}

// MarshalJSON for NullTime
func (nt *Time) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return []byte("null"), nil
	}
	val := fmt.Sprintf("\"%s\"", nt.Time.Format(time.RFC3339))
	return []byte(val), nil
}

// UnmarshalJSON implements json.Unmarshaler.
// It supports string, object (e.g. pq.NullTime and friends)
// and null input.
func (nt *Time) UnmarshalJSON(data []byte) error {
	var err error
	var v interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch x := v.(type) {
	case string:
		err = nt.Time.UnmarshalJSON(data)
	case map[string]interface{}:
		ti, tiOK := x["Time"].(string)
		valid, validOK := x["Valid"].(bool)
		if !tiOK || !validOK {
			return fmt.Errorf(`json: unmarshalling object into Go value of type null.Time requires key "Time" to be of type string and key "Valid" to be of type bool; found %T and %T, respectively`, x["Time"], x["Valid"])
		}
		err = nt.Time.UnmarshalText([]byte(ti))
		nt.Valid = valid
		return err
	case nil:
		nt.Valid = false
		return nil
	default:
		err = fmt.Errorf("json: cannot unmarshal %v into Go value of type null.Time", reflect.TypeOf(v).Name())
	}
	nt.Valid = err == nil
	return err
}

// IntFrom creates a new Int that will always be valid.
func TimeFrom(t time.Time) Time {
	return NewTime(t, true)
}

// NewInt creates a new Int
func NewTime(t time.Time, valid bool) Time {
	return Time{
		NullTime: sql.NullTime{
			Time:  t,
			Valid: valid,
		},
	}
}

// IsZero returns true for invalid Ints
func (nt Time) IsNull() bool {
	return !nt.Valid
}
