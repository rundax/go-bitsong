package types

import (
	"encoding/json"
)

type Work struct {
	ID   string   `json:"@id"`
	Type []string `json:"@type"`
}

type PermitsAction struct {
	ID        string   `json:"@id"`
	Type      []string `json:"@type"`
	ActedBy   Party    `json:"actedBy"`
	ActedOver Work     `json:"actedOver"`
}

type ObligatesAction struct {
	ID        string   `json:"@id"`
	Type      []string `json:"@type"`
	ActedBy   Party    `json:"actedBy"`
	ActedOver Work     `json:"actedOver,omitempty"`
	ActedTo   Party    `json:"actedTo,omitempty"`
}

type Issue struct {
	ID              string          `json:"@id"`
	Type            []string        `json:"@type"`
	IssuedBy        Party           `json:"issuedBy"`
	PermitsAction   PermitsAction   `json:"permitsAction,omitempty"`
	ObligatesAction ObligatesAction `json:"obligatesAction,omitempty"`
}

type Contract struct {
	ID      string  `json:"@id"`
	Type    string  `json:"@type"`
	Parties []Party `json:"hasParty"`
	Issues  []Issue `json:"issues"`
}

func NewContract(contractInfo []byte) (*Contract, error) {
	contract := Contract{}
	if err := json.Unmarshal(contractInfo, &contract); err != nil {
		return &contract, err
	}

	return &contract, nil
}

func (m Contract) Validate() error {
	return nil
}
