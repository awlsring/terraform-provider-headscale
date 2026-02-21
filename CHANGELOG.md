# Changelog

All notable changes to this project will be documented in this file.

## [5.0.0] - 2026-02-21

### Upgrade Notes (4.x -> 5.x)

- This release is designed to be upgrade-compatible for existing 4.x users.
- No manual state migration is required when upgrading from 4.x.
- Existing configurations should continue to plan/apply without refactoring.
- `headscale_pre_auth_key.acl_tags` changed from list semantics to set semantics.
- Tag ordering is now ignored, so reordering `acl_tags` entries will not trigger replacement.
- On first plan/apply after upgrade, Terraform may normalize stored `acl_tags` ordering/type representation.

### Changed

- Generated API client/types using `models/headscale.28.0.json`.
- Updated API key lifecycle handling to use stable key metadata and id-aware expiration behavior.
- Updated pre-auth key expiration flow to align with regenerated request model semantics.
- Changed `headscale_pre_auth_key.acl_tags` from list to set semantics to avoid order-based drift.
- Updated node tag handling to use the unified node `tags` field from newer generated models.
- Normalized device subnet route state ordering and added convergence reads/fallbacks to avoid transient apply inconsistencies.
- Expanded API key acceptance coverage with creation metadata assertions.
- Added replacement checks when `time_to_expire` changes.
- Expanded pre-auth key acceptance coverage with issue #24 regression checks for ACL tag reordering.
- Added pre-auth key replacement checks when `time_to_expire` changes.
- Added pre-auth key invalid duration and invalid user-id validation coverage.

### Fixed

- Issue #24: ACL tag ordering instability on `headscale_pre_auth_key`.
- Issue #25: `time_to_expire` stability during `create_before_destroy` replacement flows.
- Intermittent acceptance inconsistency when applying dual-stack exit-node routes.
