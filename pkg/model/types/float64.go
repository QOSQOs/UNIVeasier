package types

import (
	"github.com/QOSQOs/UNIVeasier/pkg/model/errors"

	"database/sql"
	"encoding/json"
	"reflect"
)

type Float64 struct {
	sql.NullFloat64
}

func (number Float64) IsNull() bool {
	return !number.Valid
}

func (number *Float64) MarshalJSON() ([]byte, error) {
	if !number.Valid {
		return []byte(`null`), nil
	}
	return json.Marshal(number.Float64)
}

func (number *Float64) UnmarshalJSON(data []byte) error {
	var i interface{}

	err := json.Unmarshal(data, &i)
	if err != nil {
		return err
	}

	switch value := i.(type) {
	case float64:
		err = json.Unmarshal(data, &number.Float64)
		if err != nil {
			return &errors.OverflowError{"Float64"}
		}
		number.Valid = true
	case nil:
		number.Valid = false
	default:
		return &errors.InvalidTypeError{"Float64", reflect.TypeOf(value).Name()}
	}
	return nil
}

func Float64From(value float64) Float64 {
	return newFloat64(value, true)
}

func NullFloat64() Float64 {
	return newFloat64(0, false)
}

func newFloat64(value float64, valid bool) Float64 {
	return Float64{
		sql.NullFloat64{
			Float64: value,
			Valid:   valid,
		},
	}
}
