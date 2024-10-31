package parser

import "encoding/json"

type Parser struct {
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Decode(from []byte, to interface{}) error {
	err := json.Unmarshal(from, to)
	if err != nil {
		return err
	}
	return nil
}
