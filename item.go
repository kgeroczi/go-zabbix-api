package zabbix

import (
	"encoding/json"
	"fmt"
)

type (
	// ItemType type of the item
	ItemType int
	// ValueType type of information of the item
	ValueType int
	// DataType data type of the item
	DataType int
	// DeltaType value that will be stored
	DeltaType int
)

const (
	// Different item type, see :
	// - "type" in https://www.zabbix.com/documentation/3.2/manual/api/reference/item/object
	// - https://www.zabbix.com/documentation/3.2/manual/config/items/itemtypes

	// ZabbixAgent type
	ZabbixAgent ItemType = 0
	// SNMPv1Agent type
	SNMPv1Agent ItemType = 1
	// ZabbixTrapper type
	ZabbixTrapper ItemType = 2
	// SimpleCheck type
	SimpleCheck ItemType = 3
	// SNMPv2Agent type
	SNMPv2Agent ItemType = 4
	// ZabbixInternal type
	ZabbixInternal ItemType = 5
	// SNMPv3Agent type
	SNMPv3Agent ItemType = 6
	// ZabbixAgentActive type
	ZabbixAgentActive ItemType = 7
	// WebItem type
	WebItem ItemType = 9
	// ExternalCheck type
	ExternalCheck ItemType = 10
	// DatabaseMonitor type
	DatabaseMonitor ItemType = 11
	//IPMIAgent type
	IPMIAgent ItemType = 12
	// SSHAgent type
	SSHAgent ItemType = 13
	// TELNETAgent type
	TELNETAgent ItemType = 14
	// Calculated type
	Calculated ItemType = 15
	// JMXAgent type
	JMXAgent  ItemType = 16
	SNMPTrap  ItemType = 17
	Dependent ItemType = 18
	HTTPAgent ItemType = 19
	SNMPAgent ItemType = 20
	// Script type (Zabbix 6.4+)
	Script ItemType = 21
	// Browser type (Zabbix 7.0+)
	Browser ItemType = 22
)

const (
	// Type of information of the item
	// see "value_type" in https://www.zabbix.com/documentation/3.2/manual/api/reference/item/object

	// Float value
	Float ValueType = 0
	// Character value
	Character ValueType = 1
	// Log value
	Log ValueType = 2
	// Unsigned value
	Unsigned ValueType = 3
	// Text value
	Text ValueType = 4
)

const (
	// Data type of the item
	// see "data_type" in https://www.zabbix.com/documentation/3.2/manual/api/reference/item/object

	// Decimal data (default)
	Decimal DataType = 0
	// Octal data
	Octal DataType = 1
	// Hexadecimal data
	Hexadecimal DataType = 2
	// Boolean data
	Boolean DataType = 3
)

const (
	// Value that will be stored
	// see "delta" in https://www.zabbix.com/documentation/3.2/manual/api/reference/item/object

	// AsIs as is (default)
	AsIs DeltaType = 0
	// Speed speed per second
	Speed DeltaType = 1
	// Delta simple change
	Delta DeltaType = 2
)

type HttpHeaders map[string]string

// HttpHeaderEntry represents a single header entry in the Zabbix 7.0+ array format
type HttpHeaderEntry struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Item represent Zabbix item object
// https://www.zabbix.com/documentation/7.0/en/manual/api/reference/item/object
type Item struct {
	ItemID       string    `json:"itemid,omitempty"`
	Delay        string    `json:"delay"`
	HostID       string    `json:"hostid"`
	InterfaceID  string    `json:"interfaceid,omitempty"`
	Key          string    `json:"key_"`
	Name         string    `json:"name"`
	Type         ItemType  `json:"type,string"`
	ValueType    ValueType `json:"value_type,string"`
	DataType     DataType  `json:"data_type,string"`
	Delta        DeltaType `json:"delta,string"`
	Description  string    `json:"description"`
	Error        string    `json:"error,omitempty"`
	History      string    `json:"history,omitempty"`
	Trends       string    `json:"trends,omitempty"`
	TrapperHosts string    `json:"trapper_hosts,omitempty"`
	Params       string    `json:"params,omitempty"`

	ItemParent Hosts `json:"hosts"`

	Preprocessors Preprocessors `json:"preprocessing,omitempty"`

	// HTTP Agent Fields
	Url             string          `json:"url,omitempty"`
	RequestMethod   string          `json:"request_method,omitempty"`
	PostType        string          `json:"post_type,omitempty"`
	RetrieveMode    string          `json:"retrieve_mode,omitempty"`
	Posts           string          `json:"posts,omitempty"`
	StatusCodes     string          `json:"status_codes,omitempty"`
	Timeout         string          `json:"timeout,omitempty"`
	VerifyHost      string          `json:"verify_host,omitempty"`
	VerifyPeer      string          `json:"verify_peer,omitempty"`
	AuthType        string          `json:"authtype,omitempty"`
	Username        string          `json:"username,omitempty"`
	Password        string          `json:"password,omitempty"`
	Headers         HttpHeaders     `json:"-"`
	RawHeaders      json.RawMessage `json:"headers,omitempty"`
	Proxy           string          `json:"http_proxy,omitempty"`
	FollowRedirects string          `json:"follow_redirects,omitempty"`

	// SNMP Fields
	SNMPOid              string `json:"snmp_oid,omitempty"`
	SNMPCommunity        string `json:"snmp_community,omitempty"`
	SNMPv3AuthPassphrase string `json:"snmpv3_authpassphrase,omitempty"`
	SNMPv3AuthProtocol   string `json:"snmpv3_authprotocol,omitempty"`
	SNMPv3ContextName    string `json:"snmpv3_contextname,omitempty"`
	SNMPv3PrivPasshrase  string `json:"snmpv3_privpassphrase,omitempty"`
	SNMPv3PrivProtocol   string `json:"snmpv3_privprotocol,omitempty"`
	SNMPv3SecurityLevel  string `json:"snmpv3_securitylevel,omitempty"`
	SNMPv3SecurityName   string `json:"snmpv3_securityname,omitempty"`

	// Dependent Fields
	MasterItemID string `json:"master_itemid,omitempty"`

	// Prototype
	RuleID        string   `json:"ruleid,omitempty"`
	DiscoveryRule *LLDRule `json:"discoveryRule,omitempty"`

	Tags Tags `json:"tags,omitempty"`
}

type Preprocessors []Preprocessor

type Preprocessor struct {
	Type               string `json:"type,omitempty"`
	Params             string `json:"params"`
	ErrorHandler       string `json:"error_handler,omitempty"`
	ErrorHandlerParams string `json:"error_handler_params"`
}

// Items is an array of Item
type Items []Item

// ByKeySafe converts slice to map by key. Returns error on duplicate keys.
func (items Items) ByKeySafe() (res map[string]Item, err error) {
	res = make(map[string]Item, len(items))
	for _, i := range items {
		if _, present := res[i.Key]; present {
			return nil, fmt.Errorf("duplicate key %s", i.Key)
		}
		res[i.Key] = i
	}
	return
}

// ByKey converts slice to map by key.
//
// Deprecated: ByKey panics on duplicate keys. Use ByKeySafe instead, which
// returns an error. This method is kept only for backward compatibility and
// will be removed in a future major version.
func (items Items) ByKey() (res map[string]Item) {
	res, err := items.ByKeySafe()
	if err != nil {
		panic(err)
	}
	return
}

// ItemsGet Wrapper for item.get
// https://www.zabbix.com/documentation/3.2/manual/api/reference/item/get
func (api *API) ItemsGet(params Params) (res Items, err error) {
	if _, present := params["output"]; !present {
		params["output"] = "extend"
	}
	err = api.CallWithErrorParse("item.get", params, &res)
	if err != nil {
		return
	}
	err = api.itemsHeadersUnmarshal(res)
	return
}
func (api *API) ProtoItemsGet(params Params) (res Items, err error) {
	if _, present := params["output"]; !present {
		params["output"] = "extend"
	}
	err = api.CallWithErrorParse("itemprototype.get", params, &res)
	if err != nil {
		return
	}
	err = api.itemsHeadersUnmarshal(res)
	return
}

func (api *API) itemsHeadersUnmarshal(item Items) error {
	for i := 0; i < len(item); i++ {
		h := item[i]

		item[i].Headers = HttpHeaders{}

		if len(h.RawHeaders) == 0 {
			continue
		}

		asStr := string(h.RawHeaders)
		if asStr == "[]" {
			continue
		}

		// Try array-of-objects format first (Zabbix 7.0+)
		var entries []HttpHeaderEntry
		if err := json.Unmarshal(h.RawHeaders, &entries); err == nil {
			for _, e := range entries {
				item[i].Headers[e.Name] = e.Value
			}
			continue
		}

		// Fallback to map format (pre-7.0)
		out := HttpHeaders{}
		if err := json.Unmarshal(h.RawHeaders, &out); err != nil {
			api.printf("got error during unmarshal %s", err)
			return fmt.Errorf("unmarshal item headers: %w", err)
		}
		item[i].Headers = out
	}
	return nil
}

func prepItems(item Items) {
	for i := 0; i < len(item); i++ {
		h := item[i]

		if h.Headers == nil {
			continue
		}
		// Marshal as array-of-objects format (Zabbix 7.0+)
		entries := make([]HttpHeaderEntry, 0, len(h.Headers))
		for k, v := range h.Headers {
			entries = append(entries, HttpHeaderEntry{Name: k, Value: v})
		}
		asB, _ := json.Marshal(entries)
		item[i].RawHeaders = json.RawMessage(asB)
	}
}

// ItemGetByID Gets item by Id only if there is exactly 1 matching host.
func (api *API) ItemGetByID(id string) (res *Item, err error) {
	items, err := api.ItemsGet(Params{"itemids": id})
	if err != nil {
		return
	}

	if len(items) != 1 {
		e := ExpectedOneResult(len(items))
		err = &e
		return
	}
	res = &items[0]
	return
}
func (api *API) ProtoItemGetByID(id string) (res *Item, err error) {
	items, err := api.ProtoItemsGet(Params{"itemids": id})
	if err != nil {
		return
	}

	if len(items) != 1 {
		e := ExpectedOneResult(len(items))
		err = &e
		return
	}
	res = &items[0]
	return
}

// ItemsCreate Wrapper for item.create
// https://www.zabbix.com/documentation/7.0/en/manual/api/reference/item/create
func (api *API) ItemsCreate(items Items) (err error) {
	prepItems(items)
	response, err := api.CallWithError("item.create", items)
	if err != nil {
		return
	}

	result := response.Result.(map[string]interface{})
	itemids := result["itemids"].([]interface{})
	for i, id := range itemids {
		items[i].ItemID = id.(string)
	}
	return
}
func (api *API) ProtoItemsCreate(items Items) (err error) {
	prepItems(items)
	response, err := api.CallWithError("itemprototype.create", items)
	if err != nil {
		return
	}

	result := response.Result.(map[string]interface{})
	itemids := result["itemids"].([]interface{})
	for i, id := range itemids {
		items[i].ItemID = id.(string)
	}
	return
}

// ItemsUpdate Wrapper for item.update
// https://www.zabbix.com/documentation/3.2/manual/api/reference/item/update
func (api *API) ItemsUpdate(items Items) (err error) {
	prepItems(items)
	_, err = api.CallWithError("item.update", items)
	return
}
func (api *API) ProtoItemsUpdate(items Items) (err error) {
	prepItems(items)
	_, err = api.CallWithError("itemprototype.update", items)
	return
}

// ItemsDelete Wrapper for item.delete
// Cleans ItemId in all items elements if call succeed.
// https://www.zabbix.com/documentation/3.2/manual/api/reference/item/delete
func (api *API) ItemsDelete(items Items) (err error) {
	ids := make([]string, len(items))
	for i, item := range items {
		ids[i] = item.ItemID
	}

	err = api.ItemsDeleteByIds(ids)
	if err == nil {
		for i := range items {
			items[i].ItemID = ""
		}
	}
	return
}
func (api *API) ProtoItemsDelete(items Items) (err error) {
	ids := make([]string, len(items))
	for i, item := range items {
		ids[i] = item.ItemID
	}

	err = api.ProtoItemsDeleteByIds(ids)
	if err == nil {
		for i := range items {
			items[i].ItemID = ""
		}
	}
	return
}

// ItemsDeleteByIds Wrapper for item.delete
// https://www.zabbix.com/documentation/3.2/manual/api/reference/item/delete
func (api *API) ItemsDeleteByIds(ids []string) (err error) {
	deleteIds, err := api.ItemsDeleteIDs(ids)
	if err != nil {
		return
	}
	l := len(deleteIds)
	if len(ids) != l {
		err = &ExpectedMore{len(ids), l}
	}
	return
}
func (api *API) ProtoItemsDeleteByIds(ids []string) (err error) {
	deleteIds, err := api.ProtoItemsDeleteIDs(ids)
	if err != nil {
		return
	}
	l := len(deleteIds)
	if len(ids) != l {
		err = &ExpectedMore{len(ids), l}
	}
	return
}

// ItemsDeleteIDs Wrapper for item.delete
// Delete the item and return the id of the deleted item
func (api *API) ItemsDeleteIDs(ids []string) (itemids []interface{}, err error) {
	response, err := api.CallWithError("item.delete", ids)
	if err != nil {
		return
	}

	result := response.Result.(map[string]interface{})
	itemids1, ok := result["itemids"].([]interface{})
	if !ok {
		itemids2 := result["itemids"].(map[string]interface{})
		for _, id := range itemids2 {
			itemids = append(itemids, id)
		}
	} else {
		itemids = itemids1
	}
	return
}
func (api *API) ProtoItemsDeleteIDs(ids []string) (itemids []interface{}, err error) {
	response, err := api.CallWithError("itemprototype.delete", ids)
	if err != nil {
		return
	}

	result := response.Result.(map[string]interface{})
	itemids1, ok := result["prototypeids"].([]interface{})
	if !ok {
		itemids2 := result["prototypeids"].(map[string]interface{})
		for _, id := range itemids2 {
			itemids = append(itemids, id)
		}
	} else {
		itemids = itemids1
	}
	return
}
