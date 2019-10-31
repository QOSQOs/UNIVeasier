package utils

import "fmt"

func FailedTest(number int) string {
	return fmt.Sprintf("Test #%d failed!", number)
}
