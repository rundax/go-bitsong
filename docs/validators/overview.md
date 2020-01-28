# Validators Overview

## Introduction

The [BitSong Network](../README.md) is based on [Tendermint](https://github.com/tendermint/tendermint/tree/master/docs/introduction) and [Cosmos-SDK](https://cosmos.network), which relies on a set of validators that are responsible for committing new blocks in the blockchain. These validators participate in the consensus protocol by broadcasting votes which contain cryptographic signatures signed by each validator's private key.

Validator candidates can bond their own Btsgs and have Btsgs ["delegated"](../delegators/delegator-guide-cli.md), or staked, to them by token holders. The BitSong Network will have 32 validators, but over time this will increase to 100 validators according to the governance. The validators are determined by who has the most stake delegated to them — the top 32 validator candidates with the most stake will become BitSong Network validators.

Validators and their delegators will earn Btsgs as block provisions and tokens as transaction fees through execution of the Tendermint consensus protocol. Initially, transaction fees will be paid in Btsgs. Note that validators can set commission on the fees their delegators receive as additional incentive.

If validators double sign, are frequently offline or do not participate in governance, their staked Btsgs (including Btsgs of users that delegated to them) can be slashed. The penalty depends on the severity of the violation.

## Hardware

There currently exists no appropriate cloud solution for validator key management. For this reason, validators must set up a physical operation secured with restricted access. A good starting place, for example, would be co-locating in secure data centers.

Validators should expect to equip their datacenter location with redundant power, connectivity, and storage backups. Expect to have several redundant networking boxes for fiber, firewall and switching and then small servers with redundant hard drive and failover. Hardware can be on the low end of datacenter gear to start out with.

We anticipate that network requirements will be low initially. The current testnet requires minimal resources. Then bandwidth, CPU and memory requirements will rise as the network grows. Large hard drives are recommended for storing years of blockchain history.

## Set Up a Website

Set up a dedicated validator's website and signal your intention to become a validator on our forum [TODO: ADD LINK]. This is important since delegators will want to have information about the entity they are delegating their Btsgs to.

## Seek Legal Advice

Seek legal advice if you intend to run a Validator.

## Community

Discuss the finer details of being a validator on our community chat and forum:

* [Validators Discord Chat](https://discord.gg/qSFUps6)
* [Validator Forum](https://btsg.community/c/validators)
* [btsg.community](https://btsg.community)