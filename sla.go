package zabbix

type (
	// SLAPeriodType reporting period of the SLA
	// see "period" in: https://www.zabbix.com/documentation/6.0/en/manual/api/reference/sla/object
	SLAPeriodType int

	// SLAStatusType status of the SLA
	// see "status" in: https://www.zabbix.com/documentation/6.0/en/manual/api/reference/sla/object
	SLAStatusType int

	// SLATagOperatorType operator for service tag matching
	// see "operator" in: https://www.zabbix.com/documentation/6.0/en/manual/api/reference/sla/object
	SLATagOperatorType int
)

const (
	// SLAPeriodDaily daily reporting period
	SLAPeriodDaily SLAPeriodType = 0
	// SLAPeriodWeekly weekly reporting period
	SLAPeriodWeekly SLAPeriodType = 1
	// SLAPeriodMonthly monthly reporting period
	SLAPeriodMonthly SLAPeriodType = 2
	// SLAPeriodQuarterly quarterly reporting period
	SLAPeriodQuarterly SLAPeriodType = 3
	// SLAPeriodAnnually annually reporting period
	SLAPeriodAnnually SLAPeriodType = 4
)

const (
	// SLAEnabled SLA is enabled
	SLAEnabled SLAStatusType = 0
	// SLADisabled SLA is disabled
	SLADisabled SLAStatusType = 1
)

const (
	// SLATagOperatorEquals tag value must equal
	SLATagOperatorEquals SLATagOperatorType = 0
	// SLATagOperatorContains tag value must contain
	SLATagOperatorContains SLATagOperatorType = 1
)

// SLAServiceTag represents a service tag filter for an SLA
// https://www.zabbix.com/documentation/6.0/en/manual/api/reference/sla/object
type SLAServiceTag struct {
	Tag      string             `json:"tag"`
	Value    string             `json:"value,omitempty"`
	Operator SLATagOperatorType `json:"operator,omitempty,string"`
}

// SLASchedule represents a scheduled uptime window for an SLA
// period_from and period_to are seconds from the start of the week (0–604800)
// https://www.zabbix.com/documentation/6.0/en/manual/api/reference/sla/object
type SLASchedule struct {
	PeriodFrom int `json:"period_from"`
	PeriodTo   int `json:"period_to"`
}

// SLAExcludedDowntime represents a named excluded downtime period for an SLA
// period_from and period_to are Unix timestamps
// https://www.zabbix.com/documentation/6.0/en/manual/api/reference/sla/object
type SLAExcludedDowntime struct {
	Name       string `json:"name"`
	PeriodFrom int64  `json:"period_from"`
	PeriodTo   int64  `json:"period_to"`
}

// SLA represent Zabbix SLA object
// https://www.zabbix.com/documentation/6.0/en/manual/api/reference/sla/object
type SLA struct {
	SLAID             string                `json:"slaid,omitempty"`
	Name              string                `json:"name"`
	Period            SLAPeriodType         `json:"period,string"`
	SLO               float64               `json:"slo,string"`
	EffectiveDate     int64                 `json:"effective_date,string"`
	Timezone          string                `json:"timezone,omitempty"`
	Status            SLAStatusType         `json:"status,string"`
	Description       string                `json:"description,omitempty"`
	ServiceTags       []SLAServiceTag       `json:"service_tags"`
	Schedule          []SLASchedule         `json:"schedule,omitempty"`
	ExcludedDowntimes []SLAExcludedDowntime `json:"excluded_downtimes,omitempty"`
}

// SLAs is an array of SLA
type SLAs []SLA

// SLAsGet Wrapper for sla.get
// https://www.zabbix.com/documentation/6.0/en/manual/api/reference/sla/get
func (api *API) SLAsGet(params Params) (res SLAs, err error) {
	if _, present := params["output"]; !present {
		params["output"] = "extend"
	}
	err = api.CallWithErrorParse("sla.get", params, &res)
	return
}

// SLAGetByID Gets SLA by ID only if there is exactly 1 matching SLA.
func (api *API) SLAGetByID(id string) (res *SLA, err error) {
	slas, err := api.SLAsGet(Params{"slaids": id})
	if err != nil {
		return
	}

	if len(slas) == 1 {
		res = &slas[0]
	} else {
		e := ExpectedOneResult(len(slas))
		err = &e
	}
	return
}

// SLAsCreate Wrapper for sla.create
// https://www.zabbix.com/documentation/6.0/en/manual/api/reference/sla/create
func (api *API) SLAsCreate(slas SLAs) (err error) {
	response, err := api.CallWithError("sla.create", slas)
	if err != nil {
		return
	}

	result := response.Result.(map[string]interface{})
	slaids := result["slaids"].([]interface{})
	for i, id := range slaids {
		slas[i].SLAID = id.(string)
	}
	return
}

// SLAsUpdate Wrapper for sla.update
// https://www.zabbix.com/documentation/6.0/en/manual/api/reference/sla/update
func (api *API) SLAsUpdate(slas SLAs) (err error) {
	_, err = api.CallWithError("sla.update", slas)
	return
}

// SLAsDelete Wrapper for sla.delete
// Cleans SLAID in all slas elements if call succeeds.
// https://www.zabbix.com/documentation/6.0/en/manual/api/reference/sla/delete
func (api *API) SLAsDelete(slas SLAs) (err error) {
	ids := make([]string, len(slas))
	for i, sla := range slas {
		ids[i] = sla.SLAID
	}

	err = api.SLAsDeleteByIds(ids)
	if err == nil {
		for i := range slas {
			slas[i].SLAID = ""
		}
	}
	return
}

// SLAsDeleteByIds Wrapper for sla.delete
// https://www.zabbix.com/documentation/6.0/en/manual/api/reference/sla/delete
func (api *API) SLAsDeleteByIds(ids []string) (err error) {
	response, err := api.CallWithError("sla.delete", ids)
	if err != nil {
		return
	}

	result := response.Result.(map[string]interface{})
	slaids := result["slaids"].([]interface{})
	if len(ids) != len(slaids) {
		err = &ExpectedMore{len(ids), len(slaids)}
	}
	return
}
