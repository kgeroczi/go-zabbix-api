package zabbix_test

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"

	zapi "github.com/kgeroczi/go-zabbix-api"
)

// Helper function to create a new action
func CreateAction(t *testing.T) *zapi.Action {
	actions := zapi.Actions{
		{
			Name:        fmt.Sprintf("zabbix-action-testing-%d", rand.Int()),
			EventSource: "2",
			Filter: zapi.Filter{
				EvalType: "2",
				Conditions: []zapi.Condition{
					{ConditionType: "22", Operator: "2", Value: "SRV"},
					{ConditionType: "24", Operator: "2", Value: "AlmaLinux"},
				},
			},
			Operations: []zapi.Operation{
				{
					OperationType: "4",
					OpGroup: []zapi.OperationGroup{
						{GroupID: "2"},
					},
				},
			},
		},
	}

	err := getAPI(t).ActionsCreate(actions)
	if err != nil {
		t.Fatal(err)
	}
	return &actions[0]
}

// Helper function to delete an action
func DeleteAction(action *zapi.Action, t *testing.T) {
	err := getAPI(t).ActionsDelete(zapi.Actions{*action})
	if err != nil {
		t.Fatal(err)
	}
}

// Test function for actions
func TestActions(t *testing.T) {
	api := getAPI(t)

	// Retrieve existing actions before test
	actions, err := api.ActionsGet(zapi.Params{})
	if err != nil {
		t.Fatal(err)
	}

	// Create a new action
	action := CreateAction(t)
	if action.ActionID == "" || action.Name == "" {
		t.Errorf("Action creation failed: %#v", action)
	}

	// Retrieve the action by its ID
	action2, err := api.ActionGetByID(action.ActionID)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(action, action2) {
		t.Errorf("Error retrieving action.\nOld action: %#v\nNew action: %#v", action, action2)
	}

	// Retrieve actions again to compare the count
	actions2, err := api.ActionsGet(zapi.Params{})
	if err != nil {
		t.Fatal(err)
	}
	if len(actions2) != len(actions)+1 {
		t.Errorf("Error creating action.\nOld actions: %#v\nNew actions: %#v", actions, actions2)
	}

	// Delete the created action
	DeleteAction(action, t)

	// Retrieve actions to confirm deletion
	actions2, err = api.ActionsGet(zapi.Params{})
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(actions, actions2) {
		t.Errorf("Error deleting action.\nOld actions: %#v\nNew actions: %#v", actions, actions2)
	}
}
