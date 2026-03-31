package zabbix

import "fmt"

// Application represent Zabbix application object
// Deprecated: Applications were removed in Zabbix 5.4. Use item tags instead.
type Application struct {
	ApplicationID string `json:"applicationid,omitempty"`
	HostID        string `json:"hostid"`
	Name          string `json:"name"`
	TemplateID    string `json:"templateid,omitempty"`
}

// Applications is an array of Application
type Applications []Application

var errApplicationsRemoved = fmt.Errorf("applications API was removed in Zabbix 5.4, use item tags instead")

// ApplicationsGet Wrapper for application.get
// Deprecated: Applications were removed in Zabbix 5.4.
func (api *API) ApplicationsGet(params Params) (res Applications, err error) {
	err = errApplicationsRemoved
	return
}

// ApplicationGetByID Gets application by Id only if there is exactly 1 matching application.
// Deprecated: Applications were removed in Zabbix 5.4.
func (api *API) ApplicationGetByID(id string) (res *Application, err error) {
	err = errApplicationsRemoved
	return
}

// ApplicationGetByHostIDAndName Gets application by host Id and name only if there is exactly 1 matching application.
// Deprecated: Applications were removed in Zabbix 5.4.
func (api *API) ApplicationGetByHostIDAndName(hostID, name string) (res *Application, err error) {
	err = errApplicationsRemoved
	return
}

// ApplicationsCreate Wrapper for application.create
// Deprecated: Applications were removed in Zabbix 5.4.
func (api *API) ApplicationsCreate(apps Applications) (err error) {
	return errApplicationsRemoved
}

// ApplicationsDelete Wrapper for application.delete
// Deprecated: Applications were removed in Zabbix 5.4.
func (api *API) ApplicationsDelete(apps Applications) (err error) {
	return errApplicationsRemoved
}

// ApplicationsDeleteByIds Wrapper for application.delete
// Deprecated: Applications were removed in Zabbix 5.4.
func (api *API) ApplicationsDeleteByIds(ids []string) (err error) {
	return errApplicationsRemoved
}
