package common

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

// NullFloat64 is an alias for sql.NullFloat64 data type
type NullFloat64 struct {
	sql.NullFloat64
}

// MarshalJSON for NullFloat64
func (nf *NullFloat64) MarshalJSON() ([]byte, error) {
	if !nf.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nf.Float64)
}

// UnmarshalJSON implements json.Unmarshaler.
// It supports number and null input.
// 0 will not be considered a null Float.
// It also supports unmarshalling a sql.NullFloat64.
func (nf *NullFloat64) UnmarshalJSON(data []byte) error {
	var err error
	var v interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch x := v.(type) {
	case float64:
		nf.Float64 = float64(x)
	case string:
		str := string(x)
		if len(str) == 0 {
			nf.Valid = false
			return nil
		}
		nf.Float64, err = strconv.ParseFloat(str, 64)
	case map[string]interface{}:
		err = json.Unmarshal(data, &nf.NullFloat64)
	case nil:
		nf.Valid = false
		return nil
	default:
		err = fmt.Errorf("json: cannot unmarshal %v into Go value of type null.Float", reflect.TypeOf(v).Name())
	}
	nf.Valid = err == nil
	return err
}

// IntFrom creates a new Int that will always be valid.
func Float64From(f float64) NullFloat64 {
	return NewFloat64(f, true)
}

// NewInt creates a new Int
func NewFloat64(f float64, valid bool) NullFloat64 {
	return NullFloat64{
		NullFloat64: sql.NullFloat64{
			Float64: f,
			Valid:   valid,
		},
	}
}

// return the value
func (nf NullFloat64) Value() float64 {
	return nf.Float64
}

// IsZero returns true for invalid Ints
func (nf NullFloat64) IsNull() bool {
	return !nf.Valid
}
