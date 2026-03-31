package zabbix_test

import (
	"testing"

	zapi "github.com/kgeroczi/go-zabbix-api"
)

func TestProtoItemsGet(t *testing.T) {
	api := getAPI(t)

	items, err := api.ProtoItemsGet(zapi.Params{})
	if maybeSkipRestricted(t, err) {
		return
	}
	if err != nil {
		t.Fatal(err)
	}
	if len(items) == 0 {
		return
	}

	item, err := api.ProtoItemGetByID(items[0].ItemID)
	if err != nil {
		t.Fatal(err)
	}
	if item.ItemID != items[0].ItemID {
		t.Fatalf("unexpected item prototype id: got %s want %s", item.ItemID, items[0].ItemID)
	}
}

func TestProtoTriggersGet(t *testing.T) {
	api := getAPI(t)

	triggers, err := api.ProtoTriggersGet(zapi.Params{})
	if maybeSkipRestricted(t, err) {
		return
	}
	if err != nil {
		t.Fatal(err)
	}
	if len(triggers) == 0 {
		return
	}

	trigger, err := api.ProtoTriggerGetByID(triggers[0].TriggerID)
	if err != nil {
		t.Fatal(err)
	}
	if trigger.TriggerID != triggers[0].TriggerID {
		t.Fatalf("unexpected trigger prototype id: got %s want %s", trigger.TriggerID, triggers[0].TriggerID)
	}
}

func TestGraphProtosGet(t *testing.T) {
	api := getAPI(t)

	graphs, err := api.GraphProtosGet(zapi.Params{})
	if maybeSkipRestricted(t, err) {
		return
	}
	if err != nil {
		t.Fatal(err)
	}
	if len(graphs) == 0 {
		return
	}

	graph, err := api.GraphProtoGetByID(graphs[0].GraphID)
	if err != nil {
		t.Fatal(err)
	}
	if graph.GraphID != graphs[0].GraphID {
		t.Fatalf("unexpected graph prototype id: got %s want %s", graph.GraphID, graphs[0].GraphID)
	}
}
