package zabbix

type (
	// ReportPeriodType period for which the report is prepared
	// see "period" in: https://www.zabbix.com/documentation/7.0/en/manual/api/reference/report/object
	ReportPeriodType int

	// ReportCycleType schedule repeating cycle
	// see "cycle" in: https://www.zabbix.com/documentation/7.0/en/manual/api/reference/report/object
	ReportCycleType int

	// ReportStatusType indicates whether the report is enabled
	// see "status" in: https://www.zabbix.com/documentation/7.0/en/manual/api/reference/report/object
	ReportStatusType int

	// ReportStateType delivery state of the report
	// see "state" in: https://www.zabbix.com/documentation/7.0/en/manual/api/reference/report/object
	ReportStateType int
)

const (
	// ReportPeriodPreviousDay previous day
	ReportPeriodPreviousDay ReportPeriodType = 0
	// ReportPeriodPreviousWeek previous week
	ReportPeriodPreviousWeek ReportPeriodType = 1
	// ReportPeriodPreviousMonth previous month
	ReportPeriodPreviousMonth ReportPeriodType = 2
	// ReportPeriodPreviousYear previous year
	ReportPeriodPreviousYear ReportPeriodType = 3
)

const (
	// ReportCycleDaily daily schedule
	ReportCycleDaily ReportCycleType = 0
	// ReportCycleWeekly weekly schedule
	ReportCycleWeekly ReportCycleType = 1
	// ReportCycleMonthly monthly schedule
	ReportCycleMonthly ReportCycleType = 2
	// ReportCycleYearly yearly schedule
	ReportCycleYearly ReportCycleType = 3
)

const (
	// ReportDisabled report is disabled
	ReportDisabled ReportStatusType = 0
	// ReportEnabled report is enabled
	ReportEnabled ReportStatusType = 1
)

const (
	// ReportStateNew report has not yet been processed
	ReportStateNew ReportStateType = 0
	// ReportStateDelivered report was generated and sent to all recipients
	ReportStateDelivered ReportStateType = 1
	// ReportStateGenerationFailed report generation failed
	ReportStateGenerationFailed ReportStateType = 2
	// ReportStateSendFailed report generated but sending failed for some/all recipients
	ReportStateSendFailed ReportStateType = 3
)

const (
	// ReportWeekdayMonday bitmask value for Monday
	ReportWeekdayMonday = 1
	// ReportWeekdayTuesday bitmask value for Tuesday
	ReportWeekdayTuesday = 2
	// ReportWeekdayWednesday bitmask value for Wednesday
	ReportWeekdayWednesday = 4
	// ReportWeekdayThursday bitmask value for Thursday
	ReportWeekdayThursday = 8
	// ReportWeekdayFriday bitmask value for Friday
	ReportWeekdayFriday = 16
	// ReportWeekdaySaturday bitmask value for Saturday
	ReportWeekdaySaturday = 32
	// ReportWeekdaySunday bitmask value for Sunday
	ReportWeekdaySunday = 64
)

// ReportUser represents a report recipient user
// https://www.zabbix.com/documentation/7.0/en/manual/api/reference/report/object#users
type ReportUser struct {
	UserID       string `json:"userid"`
	AccessUserID string `json:"access_userid,omitempty"`
	Exclude      int    `json:"exclude,string,omitempty"`
}

// ReportUsers is an array of report recipient users
type ReportUsers []ReportUser

// ReportUserGroup represents a report recipient user group
// https://www.zabbix.com/documentation/7.0/en/manual/api/reference/report/object#user-groups
type ReportUserGroup struct {
	UserGroupID  string `json:"usrgrpid"`
	AccessUserID string `json:"access_userid,omitempty"`
}

// ReportUserGroups is an array of report recipient user groups
type ReportUserGroups []ReportUserGroup

// Report represents a Zabbix scheduled report object
// https://www.zabbix.com/documentation/7.0/en/manual/api/reference/report/object
type Report struct {
	ReportID    string           `json:"reportid,omitempty"`
	UserID      string           `json:"userid"`
	Name        string           `json:"name"`
	DashboardID string           `json:"dashboardid"`
	Period      ReportPeriodType `json:"period,string"`
	Cycle       ReportCycleType  `json:"cycle,string"`
	StartTime   int              `json:"start_time,string,omitempty"`
	Weekdays    int              `json:"weekdays,string,omitempty"`
	ActiveSince string           `json:"active_since,omitempty"`
	ActiveTill  string           `json:"active_till,omitempty"`
	Subject     string           `json:"subject,omitempty"`
	Message     string           `json:"message,omitempty"`
	Status      ReportStatusType `json:"status,string"`
	Description string           `json:"description,omitempty"`
	State       ReportStateType  `json:"state,string,omitempty"`
	LastSent    int64            `json:"lastsent,string,omitempty"`
	Info        string           `json:"info,omitempty"`
	Users       ReportUsers      `json:"users,omitempty"`
	UserGroups  ReportUserGroups `json:"user_groups,omitempty"`
}

// Reports is an array of Report
type Reports []Report

// ReportsGet Wrapper for report.get
// https://www.zabbix.com/documentation/7.0/en/manual/api/reference/report/get
func (api *API) ReportsGet(params Params) (res Reports, err error) {
	if _, present := params["output"]; !present {
		params["output"] = "extend"
	}
	err = api.CallWithErrorParse("report.get", params, &res)
	return
}

// ReportGetByID Gets report by ID only if there is exactly 1 matching report.
func (api *API) ReportGetByID(id string) (res *Report, err error) {
	reports, err := api.ReportsGet(Params{"reportids": id})
	if err != nil {
		return
	}

	if len(reports) == 1 {
		res = &reports[0]
	} else {
		e := ExpectedOneResult(len(reports))
		err = &e
	}
	return
}

// ReportsCreate Wrapper for report.create
// https://www.zabbix.com/documentation/7.0/en/manual/api/reference/report/create
func (api *API) ReportsCreate(reports Reports) (err error) {
	response, err := api.CallWithError("report.create", reports)
	if err != nil {
		return
	}

	result := response.Result.(map[string]interface{})
	reportids := result["reportids"].([]interface{})
	for i, id := range reportids {
		reports[i].ReportID = id.(string)
	}
	return
}

// ReportsUpdate Wrapper for report.update
// https://www.zabbix.com/documentation/7.0/en/manual/api/reference/report/update
func (api *API) ReportsUpdate(reports Reports) (err error) {
	_, err = api.CallWithError("report.update", reports)
	return
}

// ReportsDelete Wrapper for report.delete
// Cleans ReportID in all reports elements if call succeeds.
// https://www.zabbix.com/documentation/7.0/en/manual/api/reference/report/delete
func (api *API) ReportsDelete(reports Reports) (err error) {
	ids := make([]string, len(reports))
	for i, report := range reports {
		ids[i] = report.ReportID
	}

	err = api.ReportsDeleteByIds(ids)
	if err == nil {
		for i := range reports {
			reports[i].ReportID = ""
		}
	}
	return
}

// ReportsDeleteByIds Wrapper for report.delete
// https://www.zabbix.com/documentation/7.0/en/manual/api/reference/report/delete
func (api *API) ReportsDeleteByIds(ids []string) (err error) {
	response, err := api.CallWithError("report.delete", ids)
	if err != nil {
		return
	}

	result := response.Result.(map[string]interface{})
	reportids := result["reportids"].([]interface{})
	if len(ids) != len(reportids) {
		err = &ExpectedMore{len(ids), len(reportids)}
	}
	return
}
