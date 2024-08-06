package endpoints

import (
	"context"
	"fmt"
	"github.com/pinealctx/restgo"
)

type Period string

const (
	Period1m  Period = "1m"
	Period5m  Period = "5m"
	Period15m Period = "15m"
	Period1h  Period = "1h"
	Period4h  Period = "4h"
	Period1d  Period = "1d"
)

type Candlesticks struct {
	Period  Period      `json:"period"`
	Candles [][]float64 `json:"candles"` // 这里面返回的数据是美元价格
}

// Candlesticks 获取蜡烛图数据: Limit默认为1000
func (e *Endpoint) Candlesticks(ctx context.Context, tokenSymbol string, period Period) (*Candlesticks, error) {
	resp, err := e.cli.Get(ctx, "/prices/candles",
		restgo.NewURLQueryParam("tokenSymbol", tokenSymbol),
		restgo.NewURLQueryParam("period", string(period)))
	if err != nil {
		return nil, fmt.Errorf("get candlesticks failed: %w", err)
	}
	var result = &Candlesticks{}
	if err = resp.JSONUnmarshal(result); err != nil {
		return nil, fmt.Errorf("get candlesticks failed: %w", err)
	}
	return result, nil
}

// CandlesticksWithLimit 获取蜡烛图数据: Limit为指定值
func (e *Endpoint) CandlesticksWithLimit(ctx context.Context, tokenSymbol string, period Period, limit int) (*Candlesticks, error) {
	resp, err := e.cli.Get(ctx, "/prices/candles",
		restgo.NewURLQueryParam("tokenSymbol", tokenSymbol),
		restgo.NewURLQueryParam("period", string(period)),
		restgo.NewURLQueryParam("limit", fmt.Sprintf("%d", limit)),
	)
	if err != nil {
		return nil, fmt.Errorf("get candlesticks failed: %w", err)
	}
	var result = &Candlesticks{}
	if err = resp.JSONUnmarshal(result); err != nil {
		return nil, fmt.Errorf("get candlesticks failed: %w", err)
	}
	return result, nil
}
