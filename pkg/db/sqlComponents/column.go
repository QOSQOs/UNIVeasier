package sqlComponents

import (
	"github.com/QOSQOs/UNIVeasier/internal/common"

	"fmt"
	"time"
)

type SQLColumn struct {
	name   string
	value  interface{}
	enable bool
}

func (col *SQLColumn) IsUsed() bool {
	return col.enable
}

func (col *SQLColumn) ToString() string {
	return col.name
}

func (col *SQLColumn) GetValue() string {
	switch value := col.value.(type) {
	case string:
		return fmt.Sprintf("%q", value)
	case bool:
		boolValue := fmt.Sprintf("%v", value)
		if boolValue == "true" {
			return "1"
		}
		return "0"
	case time.Time:
		return fmt.Sprintf("%q", value.Format(common.TimeFormat))
	default:
		//TODO() -> use ()
		return fmt.Sprintf("%v", value)
	}
}
