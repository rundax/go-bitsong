<!--
Guiding Principles:

Changelogs are for humans, not machines.
There should be an entry for every single version.
The same types of changes should be grouped.
Versions and sections should be linkable.
The latest version comes first.
The release date of each version is displayed.
Mention whether you follow Semantic Versioning.

Usage:

Change log entries are to be added to the Unreleased section under the
appropriate stanza (see below). Each entry should ideally include a tag and
the Github issue reference in the following format:

* (<tag>) \#<issue-number> message

The issue numbers will later be link-ified during the release process so you do
not have to worry about including a link manually, but you can if you wish.

Types of changes (Stanzas):

"Features" for new features.
"Improvements" for changes in existing functionality.
"Deprecated" for soon-to-be removed features.
"Bug Fixes" for any bug fixes.
"Client Breaking" for breaking CLI commands and REST routes.
"State Machine Breaking" for breaking the AppState

Ref: https://keepachangelog.com/en/1.0.0/
-->

# Changelog

## [Unreleased]

### Features

* (bitsongd) [\#119](https://github.com/bitsongofficial/go-bitsong/pull/119) Add support for the `--inter-block-cache` CLI
flag and configuration.
* (bitsongcli) [\#132](https://github.com/bitsongofficial/go-bitsong/pull/132) Add `tx decode` command to decode
Base64 encoded transactions.
* (modules) [\#190](https://github.com/bitsongofficial/go-bitsong/pull/190) Introduce use of the `x/evidence` module.
* (bitsongd) [\#191](https://github.com/bitsongofficial/go-bitsong/pull/191) Add debug commands to bitsongd:
  * `pubkey`: decode pubkey from base64, hex or bech32
  * `addr`: convert a address between hex and bech32
  * `raw-bytes` convert raw-bytes to hex
* (bitsongcli) [\#191](https://github.com/bitsongofficial/go-bitsong/pull/191) Add cmd `decode-tx`, decodes a tx from hex or base64
* (modules) [\#196](https://github.com/bitsongofficial/go-bitsong/pull/196) Integrate the `x/upgrade` module.

### Client Breaking Changes

* [\#164](https://github.com/bitsongofficial/go-bitsong/pull/164) [\#212](https://github.com/bitsongofficial/go-bitsong/pull/212) The LevelDB-based
keybase implementation has been replaced with a 99 designs Keyring library-backed implementation. Keys created and stored
with previous `gaia` releases need to be migrated through the `bitsongcli keys migrate` command.
* (bitsongcli) [\#326](https://github.com/bitsongofficial/go-bitsong/pull/326) Implement `--offline` flag in all post commands. Some commands
that did not work with `--generate-only` as `bitsongcli staking delegate` now work as long as we don't include the offline flag.

## [v2.0.8] - 2020-04-09

### Improvements

* (sdk) Bump SDK version to [v0.37.9](https://github.com/cosmos/cosmos-sdk/releases/tag/v0.37.9).

## [v2.0.7] - 2020-03-11

### Improvements

* (sdk) Bump SDK version to [v0.37.8](https://github.com/cosmos/cosmos-sdk/releases/tag/v0.37.8).

## [v2.0.6] - 2020-02-10

### Improvements

* (sdk) Bump SDK version to [v0.37.7](https://github.com/cosmos/cosmos-sdk/releases/tag/v0.37.7).

## [v2.0.5] - 2020-01-21

### Improvements

* (sdk) Bump SDK version to [v0.37.6](https://github.com/cosmos/cosmos-sdk/releases/tag/v0.37.6).
* (tendermint) Bump Tendermint version to [v0.32.9](https://github.com/tendermint/tendermint/releases/tag/v0.32.9).

## [v2.0.4] - 2020-01-09

### Improvements

* (sdk) Bump SDK version to [v0.37.5](https://github.com/cosmos/cosmos-sdk/releases/tag/v0.37.5).
* (tendermint) Bump Tendermint version to [v0.32.8](https://github.com/tendermint/tendermint/releases/tag/v0.32.8).

### Bug Fixes

* (cli) Fixed `bitsongcli query txs` to use `events` instead of `tags`. Events take the form of `'{eventType}.{eventAttribute}={value}'`. Please
  see the [events doc](https://github.com/cosmos/cosmos-sdk/blob/master/docs/core/events.md#events-1)
  for further documentation.

## [v2.0.3] - 2019-11-04

### Improvements

* (sdk) Bump SDK version to [v0.37.4](https://github.com/cosmos/cosmos-sdk/releases/tag/v0.37.4).
* (tendermint) Bump Tendermint version to [v0.32.7](https://github.com/tendermint/tendermint/releases/tag/v0.32.7).

## [v2.0.2] - 2019-10-12

### Improvements

* (sdk) Bump SDK version to [v0.37.3](https://github.com/cosmos/cosmos-sdk/releases/tag/v0.37.3).
* (tendermint) Bump Tendermint version to [v0.32.6](https://github.com/tendermint/tendermint/releases/tag/v0.32.6).

## [v2.0.1] - 2019-09-20

### Features

* (bitsongd) [\#119](https://github.com/bitsongofficial/go-bitsong/pull/119) Add support for the `--halt-time` CLI flag and configuration.

### Improvements

* [\#119](https://github.com/bitsongofficial/go-bitsong/pull/119) Refactor and upgrade Circle CI
configuration.
* (sdk) Update SDK version to v0.37.1

## [v2.0.0] - 2019-08-22

### Bug Fixes

* [\#104](https://github.com/bitsongofficial/go-bitsong/issues/104) Fix `ModuleAccountAddrs` to
not rely on the `x/supply` keeper to get module account addresses for blacklisting.

### State Machine Breaking Changes

* (sdk) Update SDK version to v0.37.0

## [v1.0.0] - 2019-08-13

### Bug Fixes

* (bitsongd) [\#4113](https://github.com/cosmos/cosmos-sdk/issues/4113) Fix incorrect `$GOBIN` in `Install Go`
* (bitsongcli) [\#3945](https://github.com/cosmos/cosmos-sdk/issues/3945) There's no check for chain-id in TxBuilder.SignStdTx
* (bitsongcli) [\#4190](https://github.com/cosmos/cosmos-sdk/issues/4190) Fix redelegations-from by using the correct params and query endpoint.
* (bitsongcli) [\#4219](https://github.com/cosmos/cosmos-sdk/issues/4219) Return an error when an empty mnemonic is provided during key recovery.
* (bitsongcli) [\#4345](https://github.com/cosmos/cosmos-sdk/issues/4345) Improved Ledger Nano X detection

### Breaking Changes

* (sdk) Update SDK version to v0.36.0
* (bitsongd) [\#3985](https://github.com/cosmos/cosmos-sdk/issues/3985) ValidatorPowerRank uses potential consensus power
* (bitsongd) [\#4027](https://github.com/cosmos/cosmos-sdk/issues/4027) bitsongd version command does not return the checksum of the go.sum file shipped along with the source release tarball.
  Go modules feature guarantees dependencies reproducibility and as long as binaries are built via the Makefile shipped with the sources, no dependendencies can break such guarantee.
* (bitsongd) [\#4159](https://github.com/cosmos/cosmos-sdk/issues/4159) use module pattern and module manager for initialization
* (bitsongd) [\#4272](https://github.com/cosmos/cosmos-sdk/issues/4272) Merge gaiareplay functionality into bitsongd replay.
  Drop `gaiareplay` in favor of new `bitsongd replay` command.
* (bitsongcli) [\#3715](https://github.com/cosmos/cosmos-sdk/issues/3715) query distr rewards returns per-validator
  rewards along with rewards total amount.
* (bitsongcli) [\#40](https://github.com/cosmos/cosmos-sdk/issues/40) rest-server's --cors option is now gone.
* (bitsongcli) [\#4027](https://github.com/cosmos/cosmos-sdk/issues/4027) bitsongcli version command dooes not return the checksum of the go.sum file anymore.
* (bitsongcli) [\#4142](https://github.com/cosmos/cosmos-sdk/issues/4142) Turn bitsongcli tx send's --from into a required argument.
  New shorter syntax: `bitsongcli tx send FROM TO AMOUNT`
* (bitsongcli) [\#4228](https://github.com/cosmos/cosmos-sdk/issues/4228) Merge gaiakeyutil functionality into bitsongcli keys.
  Drop `gaiakeyutil` in favor of new `bitsongcli keys parse` command. Syntax and semantic are preserved.
* (rest) [\#3715](https://github.com/cosmos/cosmos-sdk/issues/3715) Update /distribution/delegators/{delegatorAddr}/rewards GET endpoint
  as per new specs. For a given delegation, the endpoint now returns the
  comprehensive list of validator-reward tuples along with the grand total.
* (rest) [\#3942](https://github.com/cosmos/cosmos-sdk/issues/3942) Update pagination data in txs query.
* (rest) [\#4049](https://github.com/cosmos/cosmos-sdk/issues/4049) update tag MsgWithdrawValidatorCommission to match type
* (rest) The `/auth/accounts/{address}` now returns a `height` in the response. The
  account is now nested under `account`.

### Features

* (bitsongd) Add `migrate` command to `bitsongd` to provide the ability to migrate exported
  genesis state from one version to another.
* (bitsongd) Update Gaia for community pool spend proposals per Cosmos Hub governance proposal [\#7](https://github.com/cosmos/cosmos-sdk/issues/7) "Activate the Community Pool"

### Improvements

* (bitsongd) [\#4042](https://github.com/cosmos/cosmos-sdk/issues/4042) Update docs and scripts to include the correct `GO111MODULE=on` environment variable.
* (bitsongd) [\#4066](https://github.com/cosmos/cosmos-sdk/issues/4066) Fix 'ExportGenesisFile() incorrectly overwrites genesis'
* (bitsongd) [\#4064](https://github.com/cosmos/cosmos-sdk/issues/4064) Remove `dep` and `vendor` from `doc` and `version`.
* (bitsongd) [\#4080](https://github.com/cosmos/cosmos-sdk/issues/4080) add missing invariants during simulations
* (bitsongd) [\#4343](https://github.com/cosmos/cosmos-sdk/issues/4343) Upgrade toolchain to Go 1.12.5.
* (bitsongcli) [\#4068](https://github.com/cosmos/cosmos-sdk/issues/4068) Remove redundant account check on `bitsongcli`
* (bitsongcli) [\#4227](https://github.com/cosmos/cosmos-sdk/issues/4227) Support for Ledger App v1.5
* (rest) [\#2007](https://github.com/cosmos/cosmos-sdk/issues/2007) Return 200 status code on empty results
* (rest) [\#4123](https://github.com/cosmos/cosmos-sdk/issues/4123) Fix typo, url error and outdated command description of doc clients.
* (rest) [\#4129](https://github.com/cosmos/cosmos-sdk/issues/4129) Translate doc clients to chinese.
* (rest) [\#4141](https://github.com/cosmos/cosmos-sdk/issues/4141) Fix /txs/encode endpoint

<!-- Release links -->

[Unreleased]: https://github.com/bitsongofficial/go-bitsong/compare/v2.0.8...HEAD
[v2.0.8]: https://github.com/bitsongofficial/go-bitsong/releases/tag/v2.0.8
[v2.0.7]: https://github.com/bitsongofficial/go-bitsong/releases/tag/v2.0.7
[v2.0.6]: https://github.com/bitsongofficial/go-bitsong/releases/tag/v2.0.6
[v2.0.5]: https://github.com/bitsongofficial/go-bitsong/releases/tag/v2.0.5
[v2.0.4]: https://github.com/bitsongofficial/go-bitsong/releases/tag/v2.0.4
[v2.0.3]: https://github.com/bitsongofficial/go-bitsong/releases/tag/v2.0.3
[v2.0.2]: https://github.com/bitsongofficial/go-bitsong/releases/tag/v2.0.2
[v2.0.1]: https://github.com/bitsongofficial/go-bitsong/releases/tag/v2.0.1
[v2.0.0]: https://github.com/bitsongofficial/go-bitsong/releases/tag/v2.0.0
[v1.0.0]: https://github.com/bitsongofficial/go-bitsong/releases/tag/v1.0.0
