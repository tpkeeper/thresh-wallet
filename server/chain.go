// thresh-wallet
//
// Copyright 2019 by KeyFuse Labs
//
// GPLv3 License

package server

import (
	"github.com/tpkeeper/thresh-wallet/xlog"
)

const (
	testnet = "testnet"
	mainnet = "mainnet"
)

// Chain --
type Chain interface {
	GetTxs(address string) ([]Tx, error)
	GetFees() (map[string]float32, error)
	GetUTXO(address string) ([]Unspent, error)
	GetTickers() (map[string]Ticker, error)
	GetTxLink() string
	PushTx(hex string) (string, error)
}

// NewChainProxy -- creates new Chain, default provider is blockstream.info.
func NewChainProxy(log *xlog.Log, conf *Config) Chain {
	return NewBlockstreamChain(log, conf)
}
