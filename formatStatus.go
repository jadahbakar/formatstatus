package formatstatus

import "strings"

func fmtStatus(str, status string) string {
	overallLen := 27
	padStr := "."
	padCountInt := overallLen - (len(str) + len(status))
	retStr := str + strings.Repeat(padStr, padCountInt) + status
	return retStr
}
