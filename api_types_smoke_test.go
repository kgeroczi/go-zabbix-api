package zabbix_test

import (
	"fmt"
	"math/rand"
	"testing"

	zapi "github.com/kgeroczi/go-zabbix-api"
)

func TestTemplateGroupsGet(t *testing.T) {
	api := getAPI(t)

	groups, err := api.TemplateGroupsGet(zapi.Params{})
	if err != nil {
		t.Fatal(err)
	}
	if len(groups) == 0 {
		return
	}

	group, err := api.TemplateGroupGetByID(groups[0].GroupID)
	if err != nil {
		t.Fatal(err)
	}
	if group.GroupID != groups[0].GroupID {
		t.Fatalf("unexpected template group id: got %s want %s", group.GroupID, groups[0].GroupID)
	}
}

func TestTemplateGroupsCRUD(t *testing.T) {
	api := getAPI(t)

	groups := zapi.TemplateGroups{{Name: fmt.Sprintf("go-zabbix-tpl-group-%d", rand.Int())}}
	err := api.TemplateGroupsCreate(groups)
	if maybeSkipRestricted(t, err) {
		return
	}
	if err != nil {
		t.Fatal(err)
	}
	if groups[0].GroupID == "" {
		t.Fatal("template group id is empty after create")
	}

	groups[0].Name = fmt.Sprintf("go-zabbix-tpl-group-upd-%d", rand.Int())
	err = api.TemplateGroupsUpdate(groups)
	if err != nil {
		t.Fatal(err)
	}

	err = api.TemplateGroupsDelete(groups)
	if err != nil {
		t.Fatal(err)
	}
	if groups[0].GroupID != "" {
		t.Fatal("template group id was not cleared after delete")
	}
}

func TestGraphsGet(t *testing.T) {
	api := getAPI(t)

	graphs, err := api.GraphsGet(zapi.Params{})
	if err != nil {
		t.Fatal(err)
	}
	if len(graphs) == 0 {
		return
	}

	graph, err := api.GraphGetByID(graphs[0].GraphID)
	if err != nil {
		t.Fatal(err)
	}
	if graph.GraphID != graphs[0].GraphID {
		t.Fatalf("unexpected graph id: got %s want %s", graph.GraphID, graphs[0].GraphID)
	}
}

func TestLLDsGet(t *testing.T) {
	api := getAPI(t)

	llds, err := api.LLDsGet(zapi.Params{})
	if err != nil {
		t.Fatal(err)
	}
	if len(llds) == 0 {
		return
	}

	lld, err := api.LLDGetByID(llds[0].ItemID)
	if err != nil {
		t.Fatal(err)
	}
	if lld.ItemID != llds[0].ItemID {
		t.Fatalf("unexpected lld id: got %s want %s", lld.ItemID, llds[0].ItemID)
	}
}

func TestMacrosGet(t *testing.T) {
	api := getAPI(t)

	macros, err := api.MacrosGet(zapi.Params{})
	if err != nil {
		t.Fatal(err)
	}
	if len(macros) == 0 {
		return
	}

	macro, err := api.MacroGetByID(macros[0].MacroID)
	if err != nil {
		t.Fatal(err)
	}
	if macro.MacroID != macros[0].MacroID {
		t.Fatalf("unexpected macro id: got %s want %s", macro.MacroID, macros[0].MacroID)
	}
}

func TestProxiesGet(t *testing.T) {
	api := getAPI(t)

	proxies, err := api.ProxiesGet(zapi.Params{})
	if err != nil {
		t.Fatal(err)
	}
	if len(proxies) == 0 {
		return
	}

	proxy, err := api.ProxyGetByID(proxies[0].ProxyID)
	if err != nil {
		t.Fatal(err)
	}
	if proxy.ProxyID != proxies[0].ProxyID {
		t.Fatalf("unexpected proxy id: got %s want %s", proxy.ProxyID, proxies[0].ProxyID)
	}
}

func TestProxiesCRUD(t *testing.T) {
	api := getAPI(t)

	proxies := zapi.Proxies{{
		Name:          fmt.Sprintf("go-zabbix-proxy-%d", rand.Int()),
		OperatingMode: 0,
		Description:   "created by go-zabbix-api tests",
	}}
	err := api.ProxiesCreate(proxies)
	if maybeSkipRestricted(t, err) {
		return
	}
	if err != nil {
		t.Fatal(err)
	}
	if proxies[0].ProxyID == "" {
		t.Fatal("proxy id is empty after create")
	}

	proxies[0].Description = "updated by go-zabbix-api tests"
	err = api.ProxiesUpdate(proxies)
	if err != nil {
		t.Fatal(err)
	}

	err = api.ProxiesDelete(proxies)
	if err != nil {
		t.Fatal(err)
	}
	if proxies[0].ProxyID != "" {
		t.Fatal("proxy id was not cleared after delete")
	}
}

func TestUsersGet(t *testing.T) {
	api := getAPI(t)

	users, err := api.UsersGet(zapi.Params{})
	if err != nil {
		t.Fatal(err)
	}
	if len(users) == 0 {
		return
	}

	user, err := api.UserGetByID(users[0].UserID)
	if err != nil {
		t.Fatal(err)
	}
	if user.UserID != users[0].UserID {
		t.Fatalf("unexpected user id: got %s want %s", user.UserID, users[0].UserID)
	}
}

func TestUserGroupsGet(t *testing.T) {
	api := getAPI(t)

	groups, err := api.UserGroupsGet(zapi.Params{})
	if err != nil {
		t.Fatal(err)
	}
	if len(groups) == 0 {
		return
	}

	group, err := api.UserGroupGetByID(groups[0].UserGroupID)
	if err != nil {
		t.Fatal(err)
	}
	if group.UserGroupID != groups[0].UserGroupID {
		t.Fatalf("unexpected user group id: got %s want %s", group.UserGroupID, groups[0].UserGroupID)
	}
}

func TestUserGroupsCRUD(t *testing.T) {
	api := getAPI(t)

	groups := zapi.UserGroups{{Name: fmt.Sprintf("go-zabbix-user-group-%d", rand.Int())}}
	err := api.UserGroupsCreate(groups)
	if maybeSkipRestricted(t, err) {
		return
	}
	if err != nil {
		t.Fatal(err)
	}
	if groups[0].UserGroupID == "" {
		t.Fatal("user group id is empty after create")
	}

	groups[0].Name = fmt.Sprintf("go-zabbix-user-group-upd-%d", rand.Int())
	err = api.UserGroupsUpdate(groups)
	if err != nil {
		t.Fatal(err)
	}

	err = api.UserGroupsDelete(groups)
	if err != nil {
		t.Fatal(err)
	}
	if groups[0].UserGroupID != "" {
		t.Fatal("user group id was not cleared after delete")
	}
}
