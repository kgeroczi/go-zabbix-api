package zabbix

// Condition represents a condition for the action filter
type Condition struct {
	ConditionType string `json:"conditiontype"`
	Operator      string `json:"operator"`
	Value         string `json:"value"`
}

// Filter represents the filter for an action
type Filter struct {
	EvalType   string      `json:"evaltype"`
	Conditions []Condition `json:"conditions"`
}

// OperationGroup represents a group affected by an operation
type OperationGroup struct {
	GroupID string `json:"groupid"`
}

// OperationTemplate represents a template affected by an operation
type OperationTemplate struct {
	TemplateID string `json:"templateid"`
}

// Operation represents an operation in the action
type Operation struct {
	OperationType string              `json:"operationtype"`
	OpGroup       []OperationGroup    `json:"opgroup,omitempty"`
	OpTemplates   []OperationTemplate `json:"optemplate,omitempty"` // Field for template operations
	OpCommand     string              `json:"opcommand,omitempty"`  // New field for command operation
}

// Action represents a Zabbix action object
type Action struct {
	ActionID    string      `json:"actionid,omitempty"`
	Name        string      `json:"name"`
	EventSource string      `json:"eventsource"`
	Filter      Filter      `json:"filter"`
	Operations  []Operation `json:"operations"`
}

// ActionsCreate Wrapper for action.create
func (api *API) ActionsCreate(actions []Action) (err error) {
	response, err := api.CallWithError("action.create", actions)
	if err != nil {
		return
	}

	result := response.Result.(map[string]interface{})
	actionids := result["actionids"].([]interface{})
	for i, id := range actionids {
		actions[i].ActionID = id.(string)
	}
	return
}

// ActionsGet Wrapper for action.get
func (api *API) ActionsGet(params Params) (res []Action, err error) {
	if _, present := params["output"]; !present {
		params["output"] = "extend"
	}
	err = api.CallWithErrorParse("action.get", params, &res)
	return
}

// ActionGetByID Gets action by Id only if there is exactly 1 matching action.
func (api *API) ActionGetByID(id string) (res *Action, err error) {
	actions, err := api.ActionsGet(Params{"actionids": id})
	if err != nil {
		return
	}

	if len(actions) == 1 {
		res = &actions[0]
	} else {
		e := ExpectedOneResult(len(actions))
		err = &e
	}
	return
}

// ActionsUpdate Wrapper for action.update
func (api *API) ActionsUpdate(actions []Action) (err error) {
	_, err = api.CallWithError("action.update", actions)
	return
}

// ActionsDelete Wrapper for action.delete
func (api *API) ActionsDelete(actions []Action) (err error) {
	ids := make([]string, len(actions))
	for i, action := range actions {
		ids[i] = action.ActionID
	}

	err = api.ActionsDeleteByIds(ids)
	if err == nil {
		for i := range actions {
			actions[i].ActionID = ""
		}
	}
	return
}

// ActionsDeleteByIds Wrapper for action.delete
func (api *API) ActionsDeleteByIds(ids []string) (err error) {
	response, err := api.CallWithError("action.delete", ids)
	if err != nil {
		return
	}

	result := response.Result.(map[string]interface{})
	actionids := result["actionids"].([]interface{})
	if len(ids) != len(actionids) {
		err = &ExpectedMore{len(ids), len(actionids)}
	}
	return
}
