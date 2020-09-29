package types

type Party struct {
	ID string `json:"@id" yaml:"id"`
	//Label string `json:"label" yaml:"label"`
	Type []string `json:"@type" yaml:"type"`
}

func (p Party) Validate() error {
	return nil
}
