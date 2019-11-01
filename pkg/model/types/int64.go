package types

import (
	"github.com/QOSQOs/UNIVeasier/pkg/model/errors"

	"database/sql"
	"encoding/json"
	"reflect"
)

type Int64 struct {
	sql.NullInt64
}

func (number Int64) IsNull() bool {
	return !number.Valid
}

func (number *Int64) MarshalJSON() ([]byte, error) {
	if !number.Valid {
		return []byte(`null`), nil
	}
	return json.Marshal(number.Int64)
}

func (number *Int64) UnmarshalJSON(data []byte) error {
	var i interface{}

	err := json.Unmarshal(data, &i)
	if err != nil {
		return err
	}

	switch value := i.(type) {
	case float64:
		err = json.Unmarshal(data, &number.Int64)
		if err != nil {
			return &errors.OverflowError{"Int64"}
		}
		number.Valid = true
	case nil:
		number.Valid = false
	default:
		return &errors.InvalidTypeError{"int64", reflect.TypeOf(value).Name()}
	}
	return nil
}

func Int64From(value int64) Int64 {
	return newInt64(value, true)
}

func NullInt64() Int64 {
	return newInt64(0, false)
}

func newInt64(value int64, valid bool) Int64 {
	return Int64{
		sql.NullInt64{
			Int64: value,
			Valid: valid,
		},
	}
}
