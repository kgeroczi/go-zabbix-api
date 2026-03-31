# Go zabbix api

Note, this is not tested and is adjusted for use of kgeroczi/terraform-provider-zabbix

[![GoDoc](https://godoc.org/github.com/kgeroczi/go-zabbix-api?status.svg)](https://godoc.org/github.com/kgeroczi/go-zabbix-api) [![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

This Go package provides access to Zabbix API.

Requires Zabbix 7.0 or later. Uses Bearer token authentication (Authorization header).

This package supports multiple Zabbix resources from its API: trigger, host group, template group, host, item, template, proxy, user, user group, LLD rule, graph, and macro.

## Install

Install it: `go get github.com/kgeroczi/go-zabbix-api`

## Getting started

```go
package main

import (
	"fmt"

	"github.com/kgeroczi/go-zabbix-api"
)

func main() {
	api, err := zabbix.NewAPI(zabbix.Config{
		Url: "http://localhost/api_jsonrpc.php",
	})
	if err != nil {
		panic(err)
	}

	_, err = api.Login("MyZabbixUsername", "MyZabbixPassword")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Connected to zabbix api v%d\n", api.Config.Version)
}
```

## Security

Debug logging redacts sensitive fields (auth, password, token, tls_psk, macro values) by default. Raw request/response bodies are never logged with secret content exposed.

## Notable API helpers

- `Items.ByKeySafe()` — converts an item slice to a map keyed by item key, returning an error on duplicate keys. Prefer this over the legacy `ByKey()` which panics on duplicates.

## Configuration

`Config` supports the following fields:

- `Url` — Zabbix API endpoint (required)
- `TlsNoVerify` — disable TLS certificate verification (default: false)
- `Serialize` — serialize API calls (default: false)
- `Timeout` — HTTP client timeout (default: 30s if unset)

## Tests

### Considerations

Tests are not expected to be destructive, but you are advised to run them against a non-production instance or at least make a backup.

### Unit tests

Unit tests do not require a Zabbix instance:

```bash
go test -v -short ./...
```

### Acceptance tests

Acceptance tests require a live Zabbix 7.0+ instance:

```bash
export TEST_ZABBIX_URL=http://localhost:8080/zabbix/api_jsonrpc.php
export TEST_ZABBIX_USER=Admin
export TEST_ZABBIX_PASSWORD=zabbix
export TEST_ZABBIX_VERBOSE=1
go test -v ./...
```

`TEST_ZABBIX_URL` may contain HTTP basic auth username and password: `http://username:password@host/api_jsonrpc.php`. Also, in some setups URL should be like `http://host/zabbix/api_jsonrpc.php`.

### CI

A GitHub Actions workflow runs unit tests automatically on push/PR. Acceptance tests run only when `TEST_ZABBIX_URL` is configured in repository variables.

## References

Documentation is available on [godoc.org](https://godoc.org/github.com/kgeroczi/go-zabbix-api).
Also, Rafael Fernandes dos Santos wrote a [great article](http://www.sourcecode.net.br/2014/02/zabbix-api-with-golang.html) about using and extending this package.

License: Simplified BSD License (see [LICENSE](LICENSE)).
