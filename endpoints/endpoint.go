package endpoints

import (
	"context"
	"fmt"
	"github.com/pinealctx/gmx"
	"github.com/pinealctx/restgo"
)

const (
	ArbitrumURL  = "https://arbitrum-api.gmxinfra.io"
	AvalancheURL = "https://avalanche-api.gmxinfra.io"
)

func GetDefaultURLByChain(chain gmx.Chain) string {
	switch chain {
	case gmx.ChainArbitrum:
		return ArbitrumURL
	case gmx.ChainAvalanche:
		return AvalancheURL
	default:
		panic(fmt.Sprintf("unsupported chain %d", chain))
	}
}

type Endpoint struct {
	cli *restgo.Client
}

func New(url string) *Endpoint {
	return &Endpoint{cli: restgo.New(restgo.WithBaseURL(url))}
}

func (e *Endpoint) Ping(ctx context.Context) error {
	resp, err := e.cli.Get(ctx, "/ping")
	if err != nil {
		return fmt.Errorf("ping failed: %w", err)
	}

	var result struct {
		Message string `json:"message"`
	}

	if err = resp.JSONUnmarshal(&result); err != nil {
		return fmt.Errorf("ping failed: %w", err)
	}

	if result.Message != "ok" {
		return fmt.Errorf("ping failed: %s", result.Message)
	}

	return nil
}
