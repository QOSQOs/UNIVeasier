package common

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

// NullInt32 is an alias for sql.NullInt32 data type
type NullInt32 struct {
	sql.NullInt32
}

// MarshalJSON for NullInt64
func (ni *NullInt32) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Int32)
}

// UnmarshalJSON implements json.Unmarshaler.
// It supports number and null input.
// 0 will not be considered a null Int.
// It also supports unmarshalling a sql.NullInt32.
func (ni *NullInt32) UnmarshalJSON(data []byte) error {
	var err error
	var val int64
	var v interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch x := v.(type) {
	case float32:
		// Unmarshal again, directly to int32, to avoid intermediate float64
		err = json.Unmarshal(data, &ni.Int32)
	case string:
		str := string(x)
		if len(str) == 0 {
			ni.Valid = false
			return nil
		}
		val, err = strconv.ParseInt(str, 10, 32)
		ni.Int32 = int32(val)
	case map[string]interface{}:
		err = json.Unmarshal(data, &ni.NullInt32)
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
func Int32From(i int32) NullInt32 {
	return NewInt32(i, true)
}

// NewInt creates a new Int
func NewInt32(i int32, valid bool) NullInt32 {
	return NullInt32{
		NullInt32: sql.NullInt32{
			Int32: i,
			Valid: valid,
		},
	}
}

// IsZero returns true for invalid Ints
func (ni NullInt32) IsNull() bool {
	return !ni.Valid
}
