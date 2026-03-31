package zabbix

type (
	// ServiceAlgorithmType status calculation rule for the service
	// see "algorithm" in: https://www.zabbix.com/documentation/6.0/en/manual/api/reference/service/object
	ServiceAlgorithmType int

	// ServicePropagationRuleType propagation rule for the service status
	// see "propagation_rule" in: https://www.zabbix.com/documentation/6.0/en/manual/api/reference/service/object
	ServicePropagationRuleType int

	// ServiceStatusRuleType type of the service status rule
	// see "type" in: https://www.zabbix.com/documentation/6.0/en/manual/api/reference/service/object#status-rule
	ServiceStatusRuleType int

	// ServiceTagOperatorType operator for service problem tag matching
	// see "operator" in: https://www.zabbix.com/documentation/6.0/en/manual/api/reference/service/object
	ServiceTagOperatorType int
)

const (
	// ServiceAlgorithmNone do not calculate status (default)
	ServiceAlgorithmNone ServiceAlgorithmType = 0
	// ServiceAlgorithmWorstChild most critical of child services
	ServiceAlgorithmWorstChild ServiceAlgorithmType = 1
	// ServiceAlgorithmBestChild least critical of child services
	ServiceAlgorithmBestChild ServiceAlgorithmType = 2
)

const (
	// ServicePropagationRuleAsIs propagate status as-is (default)
	ServicePropagationRuleAsIs ServicePropagationRuleType = 0
	// ServicePropagationRuleIncrease increase the propagated status by propagation_value
	ServicePropagationRuleIncrease ServicePropagationRuleType = 1
	// ServicePropagationRuleDecrease decrease the propagated status by propagation_value
	ServicePropagationRuleDecrease ServicePropagationRuleType = 2
	// ServicePropagationRuleIgnore ignore this service's status in the parent
	ServicePropagationRuleIgnore ServicePropagationRuleType = 3
	// ServicePropagationRuleFixed always use propagation_value as propagated status
	ServicePropagationRuleFixed ServicePropagationRuleType = 4
)

const (
	// ServiceStatusRuleAtLeastNChildrenHaveStatusOrAbove at least N child services have status or above
	ServiceStatusRuleAtLeastNChildrenHaveStatusOrAbove ServiceStatusRuleType = 0
	// ServiceStatusRuleAtLeastNPercentChildrenHaveStatusOrAbove at least N% child services have status or above
	ServiceStatusRuleAtLeastNPercentChildrenHaveStatusOrAbove ServiceStatusRuleType = 1
	// ServiceStatusRuleLessThanNChildrenHaveStatusOrBelow less than N child services have status or below
	ServiceStatusRuleLessThanNChildrenHaveStatusOrBelow ServiceStatusRuleType = 2
	// ServiceStatusRuleLessThanNPercentChildrenHaveStatusOrBelow less than N% child services have status or below
	ServiceStatusRuleLessThanNPercentChildrenHaveStatusOrBelow ServiceStatusRuleType = 3
	// ServiceStatusRuleWeightChildrenWithStatusOrAboveAtLeastW weight of child services with status or above >= W
	ServiceStatusRuleWeightChildrenWithStatusOrAboveAtLeastW ServiceStatusRuleType = 4
	// ServiceStatusRuleWeightChildrenWithStatusOrAboveAtLeastWPercent weight of children with status or above >= W%
	ServiceStatusRuleWeightChildrenWithStatusOrAboveAtLeastWPercent ServiceStatusRuleType = 5
)

const (
	// ServiceTagOperatorEquals problem tag value must equal
	ServiceTagOperatorEquals ServiceTagOperatorType = 0
	// ServiceTagOperatorContains problem tag value must contain
	ServiceTagOperatorContains ServiceTagOperatorType = 1
)

// ServiceTag represents a tag attached to the service
// https://www.zabbix.com/documentation/6.0/en/manual/api/reference/service/object
type ServiceTag struct {
	Tag   string `json:"tag"`
	Value string `json:"value,omitempty"`
}

// ServiceProblemTag represents a problem tag filter for the service
// https://www.zabbix.com/documentation/6.0/en/manual/api/reference/service/object
type ServiceProblemTag struct {
	Tag      string                 `json:"tag"`
	Value    string                 `json:"value,omitempty"`
	Operator ServiceTagOperatorType `json:"operator,omitempty"`
}

// ServiceStatusRule represents a status rule for the service
// https://www.zabbix.com/documentation/6.0/en/manual/api/reference/service/object#status-rule
type ServiceStatusRule struct {
	Type      ServiceStatusRuleType `json:"type"`
	Limit_n   int                   `json:"limit_n"`
	LimitS    int                   `json:"limit_s"`
	NewStatus int                   `json:"new_status"`
}

// ServiceID represents a minimal service reference used in parent/child relations
type ServiceID struct {
	ServiceID string `json:"serviceid"`
}

// ServiceIDs is an array of ServiceID
type ServiceIDs []ServiceID

// Service represents a Zabbix service object
// https://www.zabbix.com/documentation/6.0/en/manual/api/reference/service/object
type Service struct {
	ServiceID        string                     `json:"serviceid,omitempty"`
	Name             string                     `json:"name"`
	Algorithm        ServiceAlgorithmType       `json:"algorithm"`
	Sortorder        int                        `json:"sortorder"`
	Weight           int                        `json:"weight,omitempty"`
	PropagationRule  ServicePropagationRuleType `json:"propagation_rule,omitempty"`
	PropagationValue int                        `json:"propagation_value,omitempty"`
	// Status is read-only, calculated by Zabbix
	Status      int    `json:"status,omitempty"`
	Description string `json:"description,omitempty"`
	UUID        string `json:"uuid,omitempty"`
	// CreatedAt is read-only
	CreatedAt   int64               `json:"created_at,omitempty"`
	Tags        []ServiceTag        `json:"tags,omitempty"`
	ProblemTags []ServiceProblemTag `json:"problem_tags,omitempty"`
	StatusRules []ServiceStatusRule `json:"status_rules,omitempty"`
	Parents     ServiceIDs          `json:"parents,omitempty"`
	Children    ServiceIDs          `json:"children,omitempty"`
}

// Services is an array of Service
type Services []Service

// ServicesGet Wrapper for service.get
// https://www.zabbix.com/documentation/6.0/en/manual/api/reference/service/get
func (api *API) ServicesGet(params Params) (res Services, err error) {
	if _, present := params["output"]; !present {
		params["output"] = "extend"
	}
	err = api.CallWithErrorParse("service.get", params, &res)
	return
}

// ServiceGetByID Gets service by ID only if there is exactly 1 matching service.
func (api *API) ServiceGetByID(id string) (res *Service, err error) {
	services, err := api.ServicesGet(Params{"serviceids": id})
	if err != nil {
		return
	}

	if len(services) == 1 {
		res = &services[0]
	} else {
		e := ExpectedOneResult(len(services))
		err = &e
	}
	return
}

// ServicesCreate Wrapper for service.create
// https://www.zabbix.com/documentation/6.0/en/manual/api/reference/service/create
func (api *API) ServicesCreate(services Services) (err error) {
	response, err := api.CallWithError("service.create", services)
	if err != nil {
		return
	}

	result := response.Result.(map[string]interface{})
	serviceids := result["serviceids"].([]interface{})
	for i, id := range serviceids {
		services[i].ServiceID = id.(string)
	}
	return
}

// ServicesUpdate Wrapper for service.update
// https://www.zabbix.com/documentation/6.0/en/manual/api/reference/service/update
func (api *API) ServicesUpdate(services Services) (err error) {
	_, err = api.CallWithError("service.update", services)
	return
}

// ServicesDelete Wrapper for service.delete
// Cleans ServiceID in all services elements if call succeeds.
// https://www.zabbix.com/documentation/6.0/en/manual/api/reference/service/delete
func (api *API) ServicesDelete(services Services) (err error) {
	ids := make([]string, len(services))
	for i, service := range services {
		ids[i] = service.ServiceID
	}

	err = api.ServicesDeleteByIds(ids)
	if err == nil {
		for i := range services {
			services[i].ServiceID = ""
		}
	}
	return
}

// ServicesDeleteByIds Wrapper for service.delete
// https://www.zabbix.com/documentation/6.0/en/manual/api/reference/service/delete
func (api *API) ServicesDeleteByIds(ids []string) (err error) {
	response, err := api.CallWithError("service.delete", ids)
	if err != nil {
		return
	}

	result := response.Result.(map[string]interface{})
	serviceids := result["serviceids"].([]interface{})
	if len(ids) != len(serviceids) {
		err = &ExpectedMore{len(ids), len(serviceids)}
	}
	return
}
