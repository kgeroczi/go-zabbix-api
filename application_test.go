package zabbix_test

import (
	"testing"

	zapi "github.com/kgeroczi/go-zabbix-api"
)

func TestApplicationsDeprecated(t *testing.T) {
	api := getAPI(t)

	_, err := api.ApplicationsGet(zapi.Params{})
	if err == nil {
		t.Fatal("Expected error for deprecated applications API")
	}
}
