package utils

import "fmt"

func FailedTest(numberTest int) string {
	return fmt.Sprintf("Test #%d failed!", numberTest)
}

func FailedSQLQuery(procedureName string) string {
	return fmt.Sprintf("The procedure %q was not completed successfully", procedureName)
}
