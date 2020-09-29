package types

const (
	QueryParams    = "params"
	QueryID        = "id"
	QueryContracts = "contracts"
)

// Params for queries
type QueryContractParams struct {
	ID uint64 `json:"id" yaml:"id"`
}

// creates a new instance of QueryContentParams
func NewQueryContractParams(id uint64) QueryContractParams {
	return QueryContractParams{
		ID: id,
	}
}

type QueryContractsParams struct {
	Page  int
	Limit int
}

func DefaultQueryContractsParams(page, limit int) QueryContractsParams {
	return QueryContractsParams{
		Page:  page,
		Limit: limit,
	}
}
