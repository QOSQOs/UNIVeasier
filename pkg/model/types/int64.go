package types

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

// NullInt32 is an alias for sql.NullInt32 data type
type Int64 struct {
	sql.NullInt64
}

// MarshalJSON for NullInt64
func (ni *Int64) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Int64)
}

// UnmarshalJSON implements json.Unmarshaler.
// It supports number and null input.
// 0 will not be considered a null Int.
// It also supports unmarshalling a sql.NullInt64.
func (ni *Int64) UnmarshalJSON(data []byte) error {
	var err error
	var v interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch x := v.(type) {
	case float64:
		// Unmarshal again, directly to int64, to avoid intermediate float64
		err = json.Unmarshal(data, &ni.Int64)
	case string:
		str := string(x)
		if len(str) == 0 {
			ni.Valid = false
			return nil
		}
		ni.Int64, err = strconv.ParseInt(str, 10, 64)
	case map[string]interface{}:
		err = json.Unmarshal(data, &ni.NullInt64)
	case nil:
		ni.Valid = false
		return nil
	default:
		err = fmt.Errorf("json: cannot unmarshal %v into Go value of type null.Int", reflect.TypeOf(v).Name())
	}
	ni.Valid = err == nil
	return err
}

// IntFrom creates a new Int that will always be valid.
func Int64From(i int64) Int64 {
	return NewInt64(i, true)
}

// NewInt creates a new Int
func NewInt64(i int64, valid bool) Int64 {
	return Int64{
		NullInt64: sql.NullInt64{
			Int64: i,
			Valid: valid,
		},
	}
}

// IsZero returns true for invalid Ints
func (i Int64) IsNull() bool {
	return !i.Valid
}
