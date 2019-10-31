package types

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"reflect"
)

// NullString is an alias for sql.NullString data type
type String struct {
	sql.NullString
}

// MarshalJSON for NullString
func (ns *String) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

// UnmarshalJSON implements json.Unmarshaler.
// It supports string and null input. Blank string input does not produce a null String.
// It also supports unmarshalling a sql.NullString.
func (ns *String) UnmarshalJSON(data []byte) error {
	var err error
	var v interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch x := v.(type) {
	case string:
		ns.String = x
	case map[string]interface{}:
		err = json.Unmarshal(data, &ns.NullString)
	case nil:
		ns.Valid = false
		return nil
	default:
		err = fmt.Errorf("json: cannot unmarshal %v into Go value of type null.String", reflect.TypeOf(v).Name())
	}
	ns.Valid = err == nil
	return err
}

// IntFrom creates a new Int that will always be valid.
func StringFrom(s string) String {
	return NewString(s, true)
}

// NewInt creates a new Int
func NewString(s string, valid bool) String {
	return String{
		NullString: sql.NullString{
			String: s,
			Valid:  valid,
		},
	}
}

// IsZero returns true for invalid Ints
func (ns String) IsNull() bool {
	return !ns.Valid
}
