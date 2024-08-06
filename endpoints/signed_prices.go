package endpoints

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pinealctx/gmx/pkg/bigx"
	"time"
)

type SignedPrice struct {
	ID                        string         `json:"id"`
	MinBlockNumber            int            `json:"minBlockNumber"`
	MinBlockHash              interface{}    `json:"minBlockHash"`
	OracleDecimals            interface{}    `json:"oracleDecimals"`
	TokenSymbol               string         `json:"tokenSymbol"`
	TokenAddress              common.Address `json:"tokenAddress"`
	MinPrice                  interface{}    `json:"minPrice"`
	MaxPrice                  interface{}    `json:"maxPrice"`
	Signer                    interface{}    `json:"signer"`
	Signature                 interface{}    `json:"signature"`
	SignatureWithoutBlockHash interface{}    `json:"signatureWithoutBlockHash"`
	CreatedAt                 time.Time      `json:"createdAt"`
	MinBlockTimestamp         interface{}    `json:"minBlockTimestamp"`
	OracleKeeperKey           string         `json:"oracleKeeperKey"`
	MaxBlockTimestamp         int            `json:"maxBlockTimestamp"`
	MaxBlockNumber            int            `json:"maxBlockNumber"`
	MaxBlockHash              string         `json:"maxBlockHash"`
	MaxPriceFull              *bigx.Int      `json:"maxPriceFull"`
	MinPriceFull              *bigx.Int      `json:"minPriceFull"`
	OracleKeeperRecordID      interface{}    `json:"oracleKeeperRecordId"`
	OracleKeeperFetchType     string         `json:"oracleKeeperFetchType"`
	OracleType                string         `json:"oracleType"`
	Blob                      string         `json:"blob"`
}

func (e *Endpoint) SignedPrices(ctx context.Context) ([]*SignedPrice, error) {
	resp, err := e.cli.Get(ctx, "/signed_prices/latest")
	if err != nil {
		return nil, fmt.Errorf("get signed prices failed: %w", err)
	}
	var result struct {
		SignedPrices []*SignedPrice `json:"signedPrices"`
	}
	if err = resp.JSONUnmarshal(&result); err != nil {
		return nil, fmt.Errorf("get signed prices failed: %w", err)
	}
	return result.SignedPrices, nil
}
