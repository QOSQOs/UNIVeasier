package types

import (
	"github.com/QOSQOs/UNIVeasier/pkg/model/errors"

	"database/sql"
	"encoding/json"
	"reflect"
)

type Int32 struct {
	sql.NullInt32
}

func (number Int32) IsNull() bool {
	return !number.Valid
}

func (number *Int32) MarshalJSON() ([]byte, error) {
	if !number.Valid {
		return []byte(`null`), nil
	}
	return json.Marshal(number.Int32)
}

func (number *Int32) UnmarshalJSON(data []byte) error {
	var i interface{}

	err := json.Unmarshal(data, &i)
	if err != nil {
		return err
	}

	switch value := i.(type) {
	case float64:
		err = json.Unmarshal(data, &number.Int32)
		if err != nil {
			return &errors.OverflowError{"Int32"}
		}
		number.Valid = true
	case nil:
		number.Valid = false
	default:
		return &errors.InvalidTypeError{"Int32", reflect.TypeOf(value).Name()}
	}
	return nil
}

func Int32From(value int32) Int32 {
	return newInt32(value, true)
}

func NullInt32() Int32 {
	return newInt32(0, false)
}

func newInt32(value int32, valid bool) Int32 {
	return Int32{
		sql.NullInt32{
			Int32: value,
			Valid: valid,
		},
	}
}
