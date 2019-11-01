package types

import (
	"database/sql"
	"encoding/json"
	"reflect"

	"github.com/QOSQOs/UNIVeasier/pkg/model/errors"
)

type String struct {
	sql.NullString
}

func (text String) IsNull() bool {
	return !text.Valid
}

func (text *String) MarshalJSON() ([]byte, error) {
	if !text.Valid {
		return []byte(`null`), nil
	}
	return json.Marshal(text.String)
}

func (text *String) UnmarshalJSON(data []byte) error {
	var i interface{}

	err := json.Unmarshal(data, &i)
	if err != nil {
		return err
	}

	switch value := i.(type) {
	case string:
		text.String = value
		text.Valid = true
	case nil:
		text.Valid = false
	default:
		return &errors.InvalidTypeError{"String", reflect.TypeOf(value).Name()}
	}
	return nil
}

func StringFrom(value string) String {
	return newString(value, true)
}

func NullString() String {
	return newString("", false)
}

func newString(value string, valid bool) String {
	return String{
		sql.NullString{
			String: value,
			Valid:  valid,
		},
	}
}
