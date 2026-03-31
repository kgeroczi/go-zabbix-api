package zabbix_test

import (
	"fmt"
	"math/rand"
	"testing"

	zapi "github.com/kgeroczi/go-zabbix-api"
)

func TestReportsGet(t *testing.T) {
	api := getAPI(t)

	reports, err := api.ReportsGet(zapi.Params{})
	if err != nil {
		t.Fatal(err)
	}

	if len(reports) == 0 {
		return
	}

	report, err := api.ReportGetByID(reports[0].ReportID)
	if err != nil {
		t.Fatal(err)
	}
	if report.ReportID != reports[0].ReportID {
		t.Fatalf("unexpected report id: got %s want %s", report.ReportID, reports[0].ReportID)
	}
}

func TestServicesGet(t *testing.T) {
	api := getAPI(t)

	services, err := api.ServicesGet(zapi.Params{})
	if err != nil {
		t.Fatal(err)
	}

	if len(services) == 0 {
		return
	}

	service, err := api.ServiceGetByID(services[0].ServiceID)
	if err != nil {
		t.Fatal(err)
	}
	if service.ServiceID != services[0].ServiceID {
		t.Fatalf("unexpected service id: got %s want %s", service.ServiceID, services[0].ServiceID)
	}
}

func TestSLAsGet(t *testing.T) {
	api := getAPI(t)

	slas, err := api.SLAsGet(zapi.Params{})
	if err != nil {
		t.Fatal(err)
	}

	if len(slas) == 0 {
		return
	}

	sla, err := api.SLAGetByID(slas[0].SLAID)
	if err != nil {
		t.Fatal(err)
	}
	if sla.SLAID != slas[0].SLAID {
		t.Fatalf("unexpected sla id: got %s want %s", sla.SLAID, slas[0].SLAID)
	}
}

func TestReportsCRUD(t *testing.T) {
	api := getAPI(t)

	reports, err := api.ReportsGet(zapi.Params{"output": "extend", "selectUsers": "extend", "selectUserGroups": "extend"})
	if err != nil {
		t.Fatal(err)
	}
	if len(reports) == 0 {
		t.Skip("no existing reports to clone")
	}

	seed := reports[0]
	if len(seed.Users) == 0 && len(seed.UserGroups) == 0 {
		t.Skip("seed report has no recipients and report.create requires users or user_groups")
	}

	newReport := zapi.Report{
		UserID:      seed.UserID,
		Name:        fmt.Sprintf("go-zabbix-report-%d", rand.Int()),
		DashboardID: seed.DashboardID,
		Period:      seed.Period,
		Cycle:       seed.Cycle,
		StartTime:   seed.StartTime,
		Weekdays:    seed.Weekdays,
		ActiveSince: seed.ActiveSince,
		ActiveTill:  seed.ActiveTill,
		Subject:     "Go Zabbix API test report",
		Message:     "report create/update/delete integration test",
		Status:      zapi.ReportEnabled,
		Description: "created by go-zabbix-api tests",
		Users:       seed.Users,
		UserGroups:  seed.UserGroups,
	}

	reportList := zapi.Reports{newReport}
	err = api.ReportsCreate(reportList)
	if maybeSkipRestricted(t, err) {
		return
	}
	if err != nil {
		t.Fatal(err)
	}
	if reportList[0].ReportID == "" {
		t.Fatal("report id is empty after create")
	}

	reportList[0].Description = "updated by go-zabbix-api tests"
	err = api.ReportsUpdate(reportList)
	if err != nil {
		t.Fatal(err)
	}

	err = api.ReportsDelete(reportList)
	if err != nil {
		t.Fatal(err)
	}
	if reportList[0].ReportID != "" {
		t.Fatal("report id was not cleared after delete")
	}
}

func TestServicesCRUD(t *testing.T) {
	api := getAPI(t)

	services, err := api.ServicesGet(zapi.Params{})
	if err != nil {
		t.Fatal(err)
	}

	newService := zapi.Service{
		Name:      fmt.Sprintf("go-zabbix-service-%d", rand.Int()),
		Algorithm: zapi.ServiceAlgorithmNone,
		Sortorder: 0,
		Status:    0,
	}
	if len(services) > 0 {
		newService.Algorithm = services[0].Algorithm
		newService.Sortorder = services[0].Sortorder
		newService.Status = services[0].Status
	}

	serviceList := zapi.Services{newService}
	err = api.ServicesCreate(serviceList)
	if maybeSkipRestricted(t, err) {
		return
	}
	if err != nil {
		t.Fatal(err)
	}
	if serviceList[0].ServiceID == "" {
		t.Fatal("service id is empty after create")
	}

	serviceList[0].Description = "updated by go-zabbix-api tests"
	err = api.ServicesUpdate(serviceList)
	if err != nil {
		t.Fatal(err)
	}

	err = api.ServicesDelete(serviceList)
	if err != nil {
		t.Fatal(err)
	}
	if serviceList[0].ServiceID != "" {
		t.Fatal("service id was not cleared after delete")
	}
}

func TestSLAsCRUD(t *testing.T) {
	api := getAPI(t)

	slas, err := api.SLAsGet(zapi.Params{})
	if err != nil {
		t.Fatal(err)
	}
	if len(slas) == 0 {
		t.Skip("no existing slas to clone")
	}

	seed := slas[0]
	slo := seed.SLO
	if slo <= 0 {
		slo = 99.9
	}

	newSLA := zapi.SLA{
		Name:              fmt.Sprintf("go-zabbix-sla-%d", rand.Int()),
		Period:            seed.Period,
		SLO:               slo,
		EffectiveDate:     seed.EffectiveDate,
		Timezone:          seed.Timezone,
		Status:            seed.Status,
		Description:       "created by go-zabbix-api tests",
		ServiceTags:       seed.ServiceTags,
		Schedule:          seed.Schedule,
		ExcludedDowntimes: seed.ExcludedDowntimes,
	}

	slaList := zapi.SLAs{newSLA}
	err = api.SLAsCreate(slaList)
	if maybeSkipRestricted(t, err) {
		return
	}
	if err != nil {
		t.Fatal(err)
	}
	if slaList[0].SLAID == "" {
		t.Fatal("sla id is empty after create")
	}

	slaList[0].Description = "updated by go-zabbix-api tests"
	err = api.SLAsUpdate(slaList)
	if err != nil {
		t.Fatal(err)
	}

	err = api.SLAsDelete(slaList)
	if err != nil {
		t.Fatal(err)
	}
	if slaList[0].SLAID != "" {
		t.Fatal("sla id was not cleared after delete")
	}
}
