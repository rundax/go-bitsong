package types

const (
	// ModuleName is the name of the module
	ModuleName = "mpeg21"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// RouterKey to be used for routing msgs
	RouterKey = ModuleName

	// QuerierRoute to be used for querierer msgs
	QuerierRoute = ModuleName

	// TODO: move to params
	MaxContractInfoLength = 2 * 1024
)

// Keys for contract store
// Items are stored with the following key: values
//
// - 0x00<contractID_Bytes>: Contract
var (
	ContractKeyPrefix = []byte{0x00}
)

func GetContractIDBytes(contractID string) []byte {
	return []byte(contractID)
}

func GetContractKey(contractID string) []byte {
	return append(ContractKeyPrefix, GetContractIDBytes(contractID)...)
}
