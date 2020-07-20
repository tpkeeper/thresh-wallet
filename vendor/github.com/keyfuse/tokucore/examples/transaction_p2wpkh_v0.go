// tokucore
//
// Copyright 2019 by KeyFuse Labs
// BSD License

package main

import (
	"fmt"

	"github.com/keyfuse/tokucore/xcore"
	"github.com/keyfuse/tokucore/xcore/bip32"
)

func assertNil(err error) {
	if err != nil {
		panic(err)
	}
}

// Demo for sent coin to Native SegWit P2WPKH address and spending from SegWit address to normal address.
func main() {
	seed := []byte("this.is.bohu.seed.")
	bohuHDKey := bip32.NewHDKey(seed)
	bohuPrv := bohuHDKey.PrivateKey()
	bohuPub := bohuHDKey.PublicKey()
	bohu := xcore.NewPayToPubKeyHashAddress(bohuPub.Hash160())

	// Satoshi.
	seed = []byte("this.is.satoshi.seed.")
	satoshiHDKey := bip32.NewHDKey(seed)
	satoshiPrv := satoshiHDKey.PrivateKey()
	satoshiPubKey := satoshiHDKey.PublicKey()
	satoshi := xcore.NewPayToWitnessV0PubKeyHashAddress(satoshiPubKey.Hash160())

	// Funding to SegWit.
	{
		bohuCoin := xcore.NewCoinBuilder().AddOutput(
			"f519a75190312039ddf885231205006b14f2e69f6e5b02314cb0e367b027fa86",
			1,
			127297408,
			"76a9145a927ddadc0ef3ae4501d0d9872b57c9584b9d8888ac",
		).ToCoins()[0]

		tx, err := xcore.NewTransactionBuilder().
			AddCoin(bohuCoin).
			AddKeys(bohuPrv).
			To(satoshi, 666666).
			Then().
			SetChange(bohu).
			SetRelayFeePerKb(20000).
			Then().
			Sign().
			BuildTransaction()
		assertNil(err)

		// Verify.
		err = tx.Verify()
		assertNil(err)

		fmt.Printf("p2wpkh.fund:%v\n", tx.ToString())
		fmt.Printf("p2wpkh.fund.txid:%v\n", tx.ID())
		fmt.Printf("p2wpkh.fund.tx:%x\n", tx.Serialize())
	}

	// Spending From SegWit.
	{
		satoshiCoin := xcore.NewCoinBuilder().AddOutput(
			"c37c3154ae611cfd9a57e684f0c12d51491d09060c643adc292565884e947b2b",
			0,
			666666,
			"00148b7f2212ecc4384abcf1df3fc5783e9c2a24d5a5",
		).ToCoins()[0]

		tx, err := xcore.NewTransactionBuilder().
			AddCoin(satoshiCoin).
			AddKeys(satoshiPrv).
			To(bohu, 66666).
			Then().
			SetChange(satoshi).
			SetRelayFeePerKb(20000).
			Then().
			Sign().
			BuildTransaction()
		assertNil(err)

		// Verify.
		err = tx.Verify()
		assertNil(err)

		fmt.Printf("p2wpkh.spend:%v\n", tx.ToString())
		fmt.Printf("p2wpkh.spend.txid:%v\n", tx.ID())
		fmt.Printf("p2wpkh.spend.tx:%x\n", tx.Serialize())
	}
}
