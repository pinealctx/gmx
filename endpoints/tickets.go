package endpoints

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pinealctx/gmx/pkg/bigx"
)

type Ticket struct {
	TokenAddress common.Address `json:"tokenAddress"`
	TokenSymbol  string         `json:"tokenSymbol"`
	MinPrice     *bigx.Int      `json:"minPrice"`
	MaxPrice     *bigx.Int      `json:"maxPrice"`
	UpdatedAt    int64          `json:"updatedAt"`
}

func (e *Endpoint) Tickers(ctx context.Context) ([]*Ticket, error) {
	resp, err := e.cli.Get(ctx, "/prices/tickers")
	if err != nil {
		return nil, fmt.Errorf("get tickers failed: %w", err)
	}
	var result []*Ticket
	if err = resp.JSONUnmarshal(&result); err != nil {
		return nil, fmt.Errorf("get tickers failed: %w", err)
	}
	return result, nil
}
