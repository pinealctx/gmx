package endpoints

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
)

type Token struct {
	Symbol    string         `json:"symbol"`
	Address   common.Address `json:"address"`
	Decimals  int64          `json:"decimals"`
	Synthetic bool           `json:"synthetic,omitempty"`
}

func (e *Endpoint) Tokens(ctx context.Context) ([]*Token, error) {
	resp, err := e.cli.Get(ctx, "/tokens")
	if err != nil {
		return nil, fmt.Errorf("get tokens failed: %w", err)
	}
	var result struct {
		Tokens []*Token `json:"tokens"`
	}
	if err = resp.JSONUnmarshal(&result); err != nil {
		return nil, fmt.Errorf("get tokens failed: %w", err)
	}
	return result.Tokens, nil
}
