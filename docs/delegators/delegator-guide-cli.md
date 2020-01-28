# Delegator Guide (CLI)

This document contains all the necessary information for delegators to interact with the BitSong Network through the Command-Line Interface (CLI).

It also contains instructions on how to manage accounts, restore accounts from the fundraiser and use a ledger nano device.

::: danger
**Very Important**: Please assure that you follow the steps described hereinafter
carefully, as negligence in this significant process could lead to an indefinite
loss of your Btsgs. Therefore, read through the following instructions in their 
entirety prior to proceeding and reach out to us in case you need support.

Please also note that you are about to interact with the BitSong Network, a
blockchain technology containing highly experimental software. While the
blockchain has been developed in accordance to the state of the art and audited
with utmost care, we can nevertheless expect to have issues, updates and bugs.
Furthermore, interaction with blockchain technology requires
advanced technical skills and always entails risks that are outside our control.
By using the software, you confirm that you understand the inherent risks
associated with cryptographic software. The BitSong Team may not be held liable for potential damages arising out of the use of the
software. Any use of this open source software released under the Apache 2.0 license is
done at your own risk and on a "AS IS" basis, without warranties or conditions
of any kind.
:::

Please exercise extreme caution!

## Table of Contents

- [Installing `bitsongcli`](#installing-bitsongcli)
- [BitSong Accounts](#bitsong-accounts)
    + [Creating an Account](#creating-an-account)
- [Accessing the BitSong Network](#accessing-the-bitsong-network)
    + [Running Your Own Full-Node](#running-your-own-full-node)
    + [Connecting to a Remote Full-Node](#connecting-to-a-remote-full-node)
- [Setting Up `bitsongcli`](#setting-up-bitsongcli)
- [Querying the State](#querying-the-state)
- [Sending Transactions](#sending-transactions)
    + [A Note on Gas and Fees](#a-note-on-gas-and-fees)
    + [Bonding Btsgs and Withdrawing Rewards](#bonding-btsgs-and-withdrawing-rewards)
    + [Participating in Governance](#participating-in-governance)
    + [Signing Transactions from an Offline Computer](#signing-transactions-from-an-offline-computer)

## Installing `bitsongcli` 

`bitsongcli`: This is the command-line interface to interact with a `bitsongd` full-node. 

::: warning
**Please check that you download the latest stable release of `bitsongcli` that is available**
:::

[**Download the binaries**]
Not available yet.

[**Install from source**](../guide/installation.md)

::: tip
`bitsongcli` is used from a terminal. To open the terminal, follow these steps:
- **Windows**: `Start` > `All Programs` > `Accessories` > `Command Prompt`
- **MacOS**: `Finder` > `Applications` > `Utilities` > `Terminal`
- **Linux**: `Ctrl` + `Alt` + `T`
:::

## BitSong Accounts

At the core of every BitSong account, there is a seed, which takes the form of a 12 or 24-words mnemonic. From this mnemonic, it is possible to create any number of BitSong accounts, i.e. pairs of private key/public key. This is called an HD wallet (see [BIP32](https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki) for more information on the HD wallet specification).

```
     Account 0                         Account 1                         Account 2

+------------------+              +------------------+               +------------------+
|                  |              |                  |               |                  |
|    Address 0     |              |    Address 1     |               |    Address 2     |
|        ^         |              |        ^         |               |        ^         |
|        |         |              |        |         |               |        |         |
|        |         |              |        |         |               |        |         |
|        |         |              |        |         |               |        |         |
|        +         |              |        +         |               |        +         |
|  Public key 0    |              |  Public key 1    |               |  Public key 2    |
|        ^         |              |        ^         |               |        ^         |
|        |         |              |        |         |               |        |         |
|        |         |              |        |         |               |        |         |
|        |         |              |        |         |               |        |         |
|        +         |              |        +         |               |        +         |
|  Private key 0   |              |  Private key 1   |               |  Private key 2   |
|        ^         |              |        ^         |               |        ^         |
+------------------+              +------------------+               +------------------+
         |                                 |                                  |
         |                                 |                                  |
         |                                 |                                  |
         +--------------------------------------------------------------------+
                                           |
                                           |
                                 +---------+---------+
                                 |                   |
                                 |  Mnemonic (Seed)  |
                                 |                   |
                                 +-------------------+
```

The funds stored in an account are controlled by the private key. This private key is generated using a one-way function from the mnemonic. If you lose the private key, you can retrieve it using the mnemonic. However, if you lose the mnemonic, you will lose access to all the derived private keys. Likewise, if someone gains access to your mnemonic, they gain access to all the associated accounts. 

::: danger
**Do not lose or share your 12 words with anyone. To prevent theft or loss of funds, it is best to ensure that you keep multiple copies of your mnemonic, and store it in a safe, secure place and that only you know how to access. If someone is able to gain access to your mnemonic, they will be able to gain access to your private keys and control the accounts associated with them.**
:::

The address is a public string with a human-readable prefix (e.g. `bitsong1wuz32uyxja0wpmz7dhxnfswza8yu48jrqf7z50`) that identifies your account. When someone wants to send you funds, they send it to your address. It is computationally infeasible to find the private key associated with a given address. 

### Creating an Account

To create an account, you just need to have `bitsongcli` installed. Before creating it, you need to know where you intend to store and interact with your private keys. The best options are to store them in an offline dedicated computer or a ledger device. Storing them on your regular online computer involves more risk, since anyone who infiltrates your computer through the internet could exfiltrate your private keys and steal your funds.

#### Using a Ledger Device

::: warning
**Only use Ledger devices that you bought factory new or trust fully**
:::

When you initialize your ledger, a 24-word mnemonic is generated and stored in the device. This mnemonic is compatible with Cosmos and BitSong accounts can be derived from it. Therefore, all you have to do is make your ledger compatible with `bitsongcli`. To do so, you need to go through the following steps:

1. Download the Ledger Live app [here](https://www.ledger.com/pages/ledger-live). 
2. Connect your ledger via USB and update to the latest firmware
3. Go to the ledger live app store, and download the "Cosmos" application (this can take a while). **Note: You may have to enable `Dev Mode` in the `Settings` of Ledger Live to be able to download the "Cosmos" application**. 
4. Navigate to the Cosmos app on your ledger device

Then, to create an account, use the following command:

```bash
bitsongcli keys add <yourAccountName> --ledger 
```

::: warning
**This command will only work while the Ledger is plugged in and unlocked**
:::

- `<yourKeyName>` is the name of the account. It is a reference to the account number used to derive the key pair from the mnemonic. You will use this name to identify your account when you want to send a transaction.
- You can add the optional `--account` flag to specify the path (`0`, `1`, `2`, ...) you want to use to generate your account. By default, account `0` is generated. 

#### Using a Computer 

::: warning
**NOTE: It is more secure to perform this action on an offline computer**
:::

To generate an account, just use the following command:

```bash
bitsongcli keys add <yourKeyName>
```

The command will generate a 24-words mnemonic and save the private and public keys for account `0`
at the same time.
Each time you want to send a transaction, you will need to unlock your system's credentials store.
If you lose access to your credentials storage, you can always recover the private key with the
mnemonic.

::: tip
**You may not be prompted for password each time you send a transaction since most operating systems
unlock user's credentials store upon login by default. If you want to change your credentials
store security policies please refer to your operating system manual.**
:::

::: danger
**Do not lose or share your 12 words with anyone. To prevent theft or loss of funds, it is best to ensure that you keep multiple copies of your mnemonic, and store it in a safe, secure place and that only you know how to access. If someone is able to gain access to your mnemonic, they will be able to gain access to your private keys and control the accounts associated with them.**
::: 

::: warning
After you have secured your mnemonic (triple check!), you can delete bash history to ensure no one can retrieve it:

```bash
history -c
rm ~/.bash_history
```
:::

- `<yourKeyName>` is the name of the account. It is a reference to the account number used to derive the key pair from the mnemonic. You will use this name to identify your account when you want to send a transaction.
- You can add the optional `--account` flag to specify the path (`0`, `1`, `2`, ...) you want to use to generate your account. By default, account `0` is generated. 


You can generate more accounts from the same mnemonic using the following command:

```bash
bitsongcli keys add <yourKeyName> --recover --account 1
```

This command will prompt you to input a passphrase as well as your mnemonic. Change the account number to generate a different account. 


## Accessing the BitSong Network

In order to query the state and send transactions, you need a way to access the network. To do so, you can either run your own full-node, or connect to someone else's.

::: danger
**NOTE: Do not share your mnemonic (12 or 24 words) with anyone. The only person who should ever need to know it is you. This is especially important if you are ever approached via email or direct message by someone requesting that you share your mnemonic for any kind of blockchain services or support. No one from BitSong will ever send an email that asks for you to share any kind of account credentials or your mnemonic."**.
::: 

### Running Your Own Full-Node

This is the most secure option, but comes with relatively high resource requirements. In order to run your own full-node, you need good bandwidth and at least 1TB of disk space. 

You will find the tutorial on how to install `bitsongd` [here](../guide/installation.md), and the guide to run a full-node [here](../guide/join-testnet.html).

### Connecting to a Remote Full-Node

If you do not want or cannot run your own node, you can connect to someone else's full-node. You should pick an operator you trust, because a malicious operator could return incorrect query results or censor your transactions. However, they will never be able to steal your funds, as your private keys are stored locally on your computer or ledger device. Possible options of full-node operators include validators, wallet providers or exchanges. 

In order to connect to the full-node, you will need an address of the following form: `https://0.0.0.0:26657` (*Note: This is a placeholder*). This address has to be communicated by the full-node operator you choose to trust. You will use this address in the [following section](#setting-up-bitsongcli).

## Setting Up `bitsongcli`

::: tip
**Before setting up `bitsongcli`, make sure you have set up a way to [access the BitSong Network](#accessing-the-bitsong-network)**
:::

::: warning
**Please check that you are always using the latest stable release of `bitsongcli`**
:::

`bitsongcli` is the tool that enables you to interact with the node that runs on the BitSong Network, whether you run it yourself or not. Let us set it up properly.

In order to set up `bitsongcli`, use the following command:

```bash
bitsongcli config <flag> <value>
```

It allows you to set a default value for each given flag. 

First, set up the address of the full-node you want to connect to:

```bash
bitsongcli config node <host>:<port

// example: bitsongcli config node https://0.0.0.0:26657
```

If you run your own full-node, just use `tcp://localhost:26657` as the address. 

Then, let us set the default value of the `--trust-node` flag:

```bash
bitsongcli config trust-node false

// Set to true if you run a light-client node, false otherwise
```

Finally, let us set the `chain-id` of the blockchain we want to interact with:

```bash
bitsongcli config chain-id bitsong-testnet-3
```

## Querying the State

::: tip
**Before you can bond btsgs and withdraw rewards, you need to [set up `bitsongcli`](#setting-up-bitsongcli)**
:::

`bitsongcli` lets you query all relevant information from the blockchain, like account balances, amount of bonded tokens, outstanding rewards, governance proposals and more. Next is a list of the most useful commands for delegator. 

```bash
// query account balances and other account-related information
bitsongcli query account <yourAddress>

// query the list of validators
bitsongcli query staking validators

// query the information of a validator given their address (e.g. bitsongvaloper14dtkqvr604k3vw3c37ldn8pa9elfeatkacrjtv)
bitsongcli query staking validator <validatorAddress>

// query all delegations made from a delegator given their address (e.g. bitsong14dtkqvr604k3vw3c37ldn8pa9elfeatkuulmm3)
bitsongcli query staking delegations <delegatorAddress>

// query a specific delegation made from a delegator (e.g. bitsong14dtkqvr604k3vw3c37ldn8pa9elfeatkuulmm3) to a validator (e.g. bitsongvaloper14dtkqvr604k3vw3c37ldn8pa9elfeatkacrjtv) given their addresses
bitsongcli query staking delegation <delegatorAddress> <validatorAddress>

// query the rewards of a delegator given a delegator address (e.g. bitsong14dtkqvr604k3vw3c37ldn8pa9elfeatkuulmm3)
bitsongcli query distribution rewards <delegatorAddress> 

// query all proposals currently open for depositing
bitsongcli query gov proposals --status deposit_period

// query all proposals currently open for voting
bitsongcli query gov proposals --status voting_period

// query a proposal given its proposalID
bitsongcli query gov proposal <proposalID>
```

For more commands, just type:

```bash
bitsongcli query
```

For each command, you can use the `-h` or `--help` flag to get more information.

## Sending Transactions

::: warning
On BitSong testnet, the accepted denom is `ubtsg`, where `1btsg = 1,000,000ubtsg`
:::

### A Note on Gas and Fees

Transactions on the BitSong Network need to include a transaction fee in order to be processed. This fee pays for the gas required to run the transaction. The formula is the following:

```
fees = ceil(gas * gasPrices)
```

The `gas` is dependent on the transaction. Different transaction require different amount of `gas`. The `gas` amount for a transaction is calculated as it is being processed, but there is a way to estimate it beforehand by using the `auto` value for the `gas` flag. Of course, this only gives an estimate. You can adjust this estimate with the flag `--gas-adjustment` (default `1.0`) if you want to be sure you provide enough `gas` for the transaction. For the remainder of this tutorial, we will use a `--gas-adjustment` of `1.5`.

The `gasPrice` is the price of each unit of `gas`. Each validator sets a `min-gas-price` value, and will only include transactions that have a `gasPrice` greater than their `min-gas-price`. 

The transaction `fees` are the product of `gas` and `gasPrice`. As a user, you have to input 2 out of 3. The higher the `gasPrice`/`fees`, the higher the chance that your transaction will get included in a block. 

::: tip
For mainnet, the recommended `gas-prices` is `0.025ubtsg`. 
::: 

### Sending Tokens

::: tip
**Before you can bond btsgs and withdraw rewards, you need to [set up `bitsongcli`](#setting-up-bitsongcli) and [create an account](#creating-an-account)**
:::

::: warning
**Note: These commands need to be run on an online computer. It is more secure to perform them commands using a Ledger Nano S device. For the offline procedure, click [here](#signing-transactions-from-an-offline-computer).**
::: 

```bash
// Send a certain amount of tokens to an address
// Ex value for parameters (do not actually use these values in your tx!!): <to_address>=bitsong14dtkqvr604k3vw3c37ldn8pa9elfeatkuulmm3 <amount>=1000000ubtsg 
// Ex value for flags: <gasPrice>=0.025ubtsg

bitsongcli tx send <to_address> <amount> --from <yourKeyName> --gas auto --gas-adjustment 1.5 --gas-prices <gasPrice>
```

### Bonding Btsgs and Withdrawing Rewards

::: tip
**Before you can bond btsgs and withdraw rewards, you need to [set up `bitsongcli`](#setting-up-bitsongcli) and [create an account](#creating-an-account)**
:::

::: warning
**Before bonding Btsgs, please read the [delegator faq](./delegator-faq.md) to understand the risk and responsibilities involved with delegating**
:::

::: warning
**Note: These commands need to be run on an online computer. It is more secure to perform them commands using a ledger device. For the offline procedure, click [here](#signing-transactions-from-an-offline-computer).**
::: 

```bash
// Bond a certain amount of Btsgs to a given validator
// ex value for flags: <validatorAddress>=bitsongvaloper14dtkqvr604k3vw3c37ldn8pa9elfeatkacrjtv, <amountToBound>=10000000ubtsg, <gasPrice>=0.025ubtsg

bitsongcli tx staking delegate <validatorAddress> <amountToBond> --from <delegatorKeyName> --gas auto --gas-adjustment 1.5 --gas-prices <gasPrice>


// Redelegate a certain amount of Btsgs from a validator to another
// Can only be used if already bonded to a validator
// Redelegation takes effect immediately, there is no waiting period to redelegate
// After a redelegation, no other redelegation can be made from the account for the next 3 weeks
// ex value for flags: <stcValidatorAddress>=bitsongvaloper14dtkqvr604k3vw3c37ldn8pa9elfeatkacrjtv, <amountToRedelegate>=100000000ubtsg, <gasPrice>=0.025ubtsg

bitsongcli tx staking redelegate <srcValidatorAddress> <destValidatorAddress> <amountToRedelegate> --from <delegatorKeyName> --gas auto --gas-adjustment 1.5 --gas-prices <gasPrice>

// Withdraw all rewards
// ex value for flag: <gasPrice>=0.025ubtsg

bitsongcli tx distribution withdraw-all-rewards --from <delegatorKeyName> --gas auto --gas-adjustment 1.5 --gas-prices <gasPrice>


// Unbond a certain amount of Btsgs from a given validator 
// You will have to wait 3 weeks before your Btsgs are fully unbonded and transferrable 
// ex value for flags: <validatorAddress>=bitsongvaloper14dtkqvr604k3vw3c37ldn8pa9elfeatkacrjtv, <amountToUnbound>=10000000ubtsg, <gasPrice>=0.025ubtsg

bitsongcli tx staking unbond <validatorAddress> <amountToUnbond> --from <delegatorKeyName> --gas auto --gas-adjustment 1.5 --gas-prices <gasPrice>
```

::: warning
**If you use a connected Ledger, you will be asked to confirm the transaction on the device before it is signed and broadcast to the network. Note that the command will only work while the Ledger is plugged in and unlocked.**
::: 

To confirm that your transaction went through, you can use the following queries:

```bash
// your balance should change after you bond Btsgs or withdraw rewards
bitsongcli query account

// you should have delegations after you bond Btsg
bitsongcli query staking delegations <delegatorAddress>

// this returns your tx if it has been included
// use the tx hash that was displayed when you created the tx
bitsongcli query tx <txHash>

```

Double check with a block explorer if you interact with the network through a trusted full-node. 

## Participating in Governance

#### Primer on Governance

The BitSong Network has a built-in governance system that lets bonded Btsg holders vote on proposals. There are three types of proposal:

- `Text Proposals`: These are the most basic type of proposals. They can be used to get the opinion of the network on a given topic. 
- `Parameter Proposals`: These are used to update the value of an existing parameter.
- `Software Upgrade Proposal`: These are used to propose an upgrade of the Network's software.

Any Btsg holder can submit a proposal. In order for the proposal to be open for voting, it needs to come with a `deposit` that is greater than a parameter called `minDeposit`. The `deposit` need not be provided in its entirety by the submitter. If the initial proposer's `deposit` is not sufficient, the proposal enters the `deposit_period` status. Then, any Btsg holder can increase the deposit by sending a `depositTx`. 

Once the `deposit` reaches `minDeposit`, the proposal enters the `voting_period`, which lasts 2 weeks. Any **bonded** Btsg holder can then cast a vote on this proposal. The options are `Yes`, `No`, `NoWithVeto` and `Abstain`. The weight of the vote is based on the amount of bonded Btsgs of the sender. If they don't vote, delegator inherit the vote of their validator. However, delegators can override their validator's vote by sending a vote themselves. 

At the end of the voting period, the proposal is accepted if there are more than 50% `Yes` votes (excluding `Abstain ` votes) and less than 33.33% of `NoWithVeto` votes (excluding `Abstain` votes).

#### In Practice

::: tip
**Before you can bond btsgs and withdraw rewards, you need to [bond Btsgs](#bonding-btsgs-and-withdrawing-rewards)**
:::

::: warning
**Note: These commands need to be run on an online computer. It is more secure to perform them commands using a ledger device. For the offline procedure, click [here](#signing-transactions-from-an-offline-computer).**
::: 

```bash
// Submit a Proposal
// <type>=text/parameter_change/software_upgrade
// ex value for flag: <gasPrice>=0.025ubtsg

bitsongcli tx gov submit-proposal --title "Test Proposal" --description "My awesome proposal" --type <type> --deposit=10000000ubtsg --gas auto --gas-adjustment 1.5 --gas-prices <gasPrice> --from <delegatorKeyName>

// Increase deposit of a proposal
// Retrieve proposalID from $bitsongcli query gov proposals --status deposit_period
// ex value for parameter: <deposit>=10000000ubtsg

bitsongcli tx gov deposit <proposalID> <deposit> --gas auto --gas-adjustment 1.5 --gas-prices <gasPrice> --from <delegatorKeyName>

// Vote on a proposal
// Retrieve proposalID from $bitsongcli query gov proposals --status voting_period 
// <option>=yes/no/no_with_veto/abstain

bitsongcli tx gov vote <proposalID> <option> --gas auto --gas-adjustment 1.5 --gas-prices <gasPrice> --from <delegatorKeyName>
```

### Signing Transactions From an Offline Computer

If you do not have a ledger device and want to interact with your private key on an offline computer, you can use the following procedure. First, generate an unsigned transaction on an **online computer** with the following command (example with a bonding transaction):

```bash
// Bond Btsgs 
// ex value for flags: <amountToBound>=10000000ubtsg, <bech32AddressOfValidator>=bitsongvaloper14dtkqvr604k3vw3c37ldn8pa9elfeatkacrjtv, <gasPrice>=0.025ubtsg, <delegatorAddress>=bitsong14dtkqvr604k3vw3c37ldn8pa9elfeatkuulmm3

bitsongcli tx staking delegate <validatorAddress> <amountToBond> --from <delegatorAddress> --gas auto --gas-adjustment 1.5 --gas-prices <gasPrice> --generate-only > unsignedTX.json
```

In order to sign, you will also need the `chain-id`, `account-number` and `sequence`. The `chain-id` is a unique identifier for the blockchain on which you are submitting the transaction. The `account-number` is an identifier generated when your account first receives funds. The `sequence` number is used to keep track of the number of transactions you have sent and prevent replay attacks.

Get the chain-id from the genesis file (`bitsong-testnet-3`), and the two other fields using the account query:

```bash
bitsongcli query account <yourAddress> --chain-id bitsong-testnet-3
```

Then, copy `unsignedTx.json` and transfer it (e.g. via USB) to the offline computer. If it is not done already, [create an account on the offline computer](#using-a-computer). For additional security, you can double check the parameters of your transaction before signing it using the following command:

```bash
cat unsignedTx.json
```

Now, sign the transaction using the following command. You will need the `chain-id`, `sequence` and `account-number` obtained earlier:

```bash
bitsongcli tx sign unsignedTx.json --from <delegatorKeyName> --offline --chain-id bitsong-testnet-3 --sequence <sequence> --account-number <account-number> > signedTx.json
```

Copy `signedTx.json` and transfer it back to the online computer. Finally, use the following command to broadcast the transaction:

```bash
bitsongcli tx broadcast signedTx.json
```
