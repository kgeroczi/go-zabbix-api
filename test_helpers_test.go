package zabbix_test

import (
	"strings"
	"testing"
)

func maybeSkipRestricted(t *testing.T, err error) bool {
	if err == nil {
		return false
	}

	msg := strings.ToLower(err.Error())
	if strings.Contains(msg, "permission") ||
		strings.Contains(msg, "access denied") ||
		strings.Contains(msg, "no permissions") ||
		strings.Contains(msg, "cannot be performed") ||
		strings.Contains(msg, "insufficient") ||
		strings.Contains(msg, "not allowed") {
		t.Skipf("skipping because operation is restricted in this environment: %v", err)
		return true
	}

	return false
}
