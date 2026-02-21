# Changelog

All notable changes to this project will be documented in this file.

## [5.0.0] - 2026-02-21

### Upgrade Notes (4.x -> 5.x)

- Terraform resource schemas and state formats remain compatible with 4.x.
- No manual state migration is required when upgrading from 4.x.
- Existing configurations should continue to plan/apply without schema refactors.

### Changed

- Generated API client/types using `models/headscale.28.0.json`.
- Updated API key lifecycle handling to use stable key metadata and id-aware expiration behavior.
- Updated pre-auth key expiration flow to align with regenerated request model semantics.
- Updated node tag handling to use the unified node `tags` field from newer generated models.
- Normalized device subnet route state ordering and added convergence reads to avoid transient apply inconsistencies.
- Expanded API key acceptance coverage with creation metadata assertions.
- Added replacement checks when `time_to_expire` changes.
