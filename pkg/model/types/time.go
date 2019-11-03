package types

import (
	"github.com/QOSQOs/UNIVeasier/pkg/model/errors"

	"database/sql"
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

const timeFormat = time.RFC3339

type Time struct {
	sql.NullTime
}

func (timestamp Time) IsNull() bool {
	return !timestamp.Valid
}

func (timestamp *Time) MarshalJSON() ([]byte, error) {
	if !timestamp.Valid {
		return []byte(`null`), nil
	}
	value := fmt.Sprintf("%q", timestamp.Time.Format(timeFormat))
	return []byte(value), nil
}

func (timestamp *Time) UnmarshalJSON(data []byte) error {
	var i interface{}

	err := json.Unmarshal(data, &i)
	if err != nil {
		return err
	}

	switch value := i.(type) {
	case string:
		timestamp.Time, err = time.Parse(timeFormat, value)
		if err != nil {
			return &errors.OverflowError{"Time"}
		}
		timestamp.Valid = true
	case nil:
		timestamp.Valid = false
	default:
		return &errors.InvalidTypeError{"Time", reflect.TypeOf(value).Name()}
	}
	return nil
}

func TimeFrom(value time.Time) Time {
	return newTime(value, true)
}

func NullTime() Time {
	return newTime(time.Time{}, false)
}

func newTime(value time.Time, valid bool) Time {
	return Time{
		sql.NullTime{
			Time:  value,
			Valid: valid,
		},
	}
}
