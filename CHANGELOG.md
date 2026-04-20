# Changelog

All notable changes to this project will be documented in this file.

The format is based on Keep a Changelog, and this project follows Semantic Versioning.

## [v0.3.2] - 2026-04-20

### Changed
- Improved JSON compatibility for Zabbix 7.x API responses that return numeric fields as strings.
- Updated service model tags in `service.go` to decode string-backed numeric fields:
  - `algorithm`, `sortorder`, `weight`, `propagation_rule`, `propagation_value`, `status`, `created_at`
  - status rule fields: `type`, `limit_n`, `limit_s`, `new_status`
  - problem tag `operator`
- Updated SLA model tags in `sla.go` to decode string-backed numeric fields:
  - `period`, `slo`, `effective_date`, `status`
  - service tag `operator`
- Updated macro model tag in `macro.go` to decode string-backed `type`.

### Fixed
- Fixed host create/update payload compatibility in `host.go`:
  - `HostGroupIds` now uses `groups` JSON key (instead of `hostgroups`).
- Fixed SLA payload requirements in `sla.go`:
  - `service_tags` and `effective_date` are always serialized when present in the model contract.
- Fixed item payload compatibility in `item.go`:
  - `hosts` is no longer sent on create/update unless explicitly provided.
  - `data_type` and `delta` are omitted when unset to avoid invalid parameters for item types that do not accept them.

## [v0.3.1] - 2026-03-31

### Added
- Added `service` API support in `service.go`:
  - Types and constants for service algorithms, propagation rules, status rules, and tag operators.
  - CRUD wrappers: `ServicesGet`, `ServiceGetByID`, `ServicesCreate`, `ServicesUpdate`, `ServicesDelete`, `ServicesDeleteByIds`.
- Added `sla` API support in `sla.go`:
  - Types for SLA period/status/tag operator, schedule windows, service tags, and excluded downtimes.
  - CRUD wrappers: `SLAsGet`, `SLAGetByID`, `SLAsCreate`, `SLAsUpdate`, `SLAsDelete`, `SLAsDeleteByIds`.
- Added `report` API support in `report.go`:
  - Types/constants for report period, cycle, status, state, weekday bitmask, users and user groups.
  - CRUD wrappers: `ReportsGet`, `ReportGetByID`, `ReportsCreate`, `ReportsUpdate`, `ReportsDelete`, `ReportsDeleteByIds`.
- Added new test suites:
  - `report_test.go` for report/service/sla GET and CRUD coverage.
  - `proto_test.go` for prototype GET coverage (`item`, `trigger`, `graph`).
  - `api_types_smoke_test.go` for additional API resource smoke/CRUD coverage.
  - `test_helpers_test.go` with shared skip helper for restricted environments.

### Changed
- Updated integration test bootstrap in `base_test.go`:
  - API-dependent tests now skip when `TEST_ZABBIX_URL` is not set instead of terminating the test process.
- Updated `README.md`:
  - Resource list now includes `service`, `SLA`, and `report`.
  - Clarified test behavior and environment-variable-driven skip logic.
  - Added test layout section to distinguish unit-focused and integration/API tests.
- Applied `gofmt` normalization to several files.

### Fixed
- Fixed macro create response mapping in `macro.go`:
  - `MacrosCreate` now stores created IDs in `MacroID` (previously assigned to `HostID`).
- Fixed macro delete error handling in `macro.go`:
  - `MacrosDeleteByIDs` now returns early when API call fails, preventing nil response dereference.
- Fixed test compile break in `trigger_test.go` by using current `CreateItem` helper signature.
- Corrected template test type usage in `template_test.go` to use `TemplateGroupID`/`TemplateGroupIDs`.
- Corrected JSON tag typo in `item.go` (`omitEmpty` -> `omitempty`) for `DiscoveryRule`.
- Corrected comment typo in `service.go` (`ServiceStatusRuleTypeType` -> `ServiceStatusRuleType`).
