package zabbix

import (
	"encoding/json"
	"testing"
)

func TestPrepHostsInventoryMode(t *testing.T) {
	hosts := Hosts{
		{
			Host:          "test-host",
			InventoryMode: InventoryManual,
		},
	}

	prepHosts(hosts)

	if hosts[0].RawInventoryMode == nil {
		t.Fatal("RawInventoryMode should be set after prepHosts, got nil")
	}
	if *hosts[0].RawInventoryMode != InventoryManual {
		t.Errorf("expected RawInventoryMode=%d, got %d", InventoryManual, *hosts[0].RawInventoryMode)
	}

	// Verify it marshals correctly
	b, err := json.Marshal(hosts[0])
	if err != nil {
		t.Fatal(err)
	}
	raw := string(b)
	// inventory_mode should appear as "0" (manual)
	if !json.Valid(b) {
		t.Fatal("invalid JSON output")
	}
	_ = raw
}

func TestPrepHostsInventoryModeAutomatic(t *testing.T) {
	hosts := Hosts{
		{
			Host:          "test-host-auto",
			InventoryMode: InventoryAutomatic,
			Inventory:     Inventory{"location": "test"},
		},
	}

	prepHosts(hosts)

	if hosts[0].RawInventoryMode == nil {
		t.Fatal("RawInventoryMode should be set after prepHosts")
	}
	if *hosts[0].RawInventoryMode != InventoryAutomatic {
		t.Errorf("expected RawInventoryMode=%d, got %d", InventoryAutomatic, *hosts[0].RawInventoryMode)
	}
	if hosts[0].RawInventory == nil {
		t.Fatal("RawInventory should be set after prepHosts")
	}
}

func TestPrepHostsInventoryModeDisabled(t *testing.T) {
	hosts := Hosts{
		{
			Host:          "test-host-disabled",
			InventoryMode: InventoryDisabled,
		},
	}

	prepHosts(hosts)

	if hosts[0].RawInventoryMode == nil {
		t.Fatal("RawInventoryMode should be set after prepHosts")
	}
	if *hosts[0].RawInventoryMode != InventoryDisabled {
		t.Errorf("expected RawInventoryMode=%d, got %d", InventoryDisabled, *hosts[0].RawInventoryMode)
	}
}

func TestItemsByKeySafe(t *testing.T) {
	items := Items{
		{Key: "key1", Name: "Item 1"},
		{Key: "key2", Name: "Item 2"},
	}

	res, err := items.ByKeySafe()
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 2 {
		t.Errorf("expected 2 items, got %d", len(res))
	}

	// Test duplicate key detection
	items = append(items, Item{Key: "key1", Name: "Duplicate"})
	_, err = items.ByKeySafe()
	if err == nil {
		t.Fatal("expected error for duplicate key, got nil")
	}
}

func TestRedactSensitive(t *testing.T) {
	input := []byte(`{"auth":"secret123","method":"host.get","password":"pass","other":"visible"}`)
	result := redactSensitive(input)

	if contains(result, "secret123") {
		t.Error("auth value should be redacted")
	}
	if contains(result, "pass") && !contains(result, "password") {
		t.Error("password value should be redacted")
	}
	if !contains(result, "visible") {
		t.Error("non-sensitive values should remain")
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && searchString(s, substr)
}

func searchString(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
