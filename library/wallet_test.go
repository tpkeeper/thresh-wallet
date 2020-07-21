// thresh-wallet
//
// Copyright 2019 by KeyFuse Labs
//
// GPLv3 License

package library

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/keyfuse/tokucore/network"
	"github.com/keyfuse/tokucore/xcore/bip32"
	"github.com/keyfuse/tokucore/xcrypto"
	"net/http"
	"testing"

	"github.com/tpkeeper/thresh-wallet/server"

	"github.com/stretchr/testify/assert"
)

func TestWalletCheck(t *testing.T) {
	var token string

	ts, cleanup := server.MockServer()
	defer cleanup()

	// Token.
	{
		body := APIGetToken(ts.URL, mockMobile1, "vcode")
		rsp := &TokenResponse{}
		unmarshal(body, rsp)
		assert.Equal(t, 200, rsp.Code)
		token = rsp.Token
	}

	body := APIWalletCheck(ts.URL, token)
	rsp := &WalletCheckResponse{}
	unmarshal(body, rsp)

	t.Logf("%+v", body)
	assert.Equal(t, 200, rsp.Code)
	assert.False(t, rsp.WalletExists)
	assert.False(t, rsp.BackupExists)
}

func TestWalletCreate(t *testing.T) {
	var token string

	ts, cleanup := server.MockServer()
	defer cleanup()

	// Token.
	{
		body := APIGetToken(ts.URL, mockMobile1, "vcode")
		rsp := &TokenResponse{}
		unmarshal(body, rsp)
		assert.Equal(t, 200, rsp.Code)
		token = rsp.Token
	}

	body := APIWalletCreate(ts.URL, token, mockMasterPrvKey)
	rsp := &WalletCreateResponse{}
	unmarshal(body, rsp)

	t.Logf("%+v", body)
	assert.Equal(t, 200, rsp.Code)
}

func TestWalletPortfolio(t *testing.T) {
	var token string

	ts, cleanup := server.MockServer()
	defer cleanup()

	// Token.
	{
		body := APIGetToken(ts.URL, mockMobile, "vcode")
		rsp := &TokenResponse{}
		unmarshal(body, rsp)
		assert.Equal(t, 200, rsp.Code)
		token = rsp.Token
	}

	body := APIWalletPortfolio(ts.URL, token)
	rsp := &WalletPortfolioResponse{}
	unmarshal(body, rsp)

	t.Logf("%+v", body)
	assert.Equal(t, 200, rsp.Code)
}

func TestWalletBalance(t *testing.T) {
	var token string

	ts, cleanup := server.MockServer()
	defer cleanup()

	// Token.
	{
		body := APIGetToken(ts.URL, mockMobile, "vcode")
		rsp := &TokenResponse{}
		unmarshal(body, rsp)
		assert.Equal(t, 200, rsp.Code)
		token = rsp.Token
	}

	body := APIWalletBalance(ts.URL, token)
	rsp := &WalletBalanceResponse{}
	unmarshal(body, rsp)

	t.Logf("%+v", body)
	assert.Equal(t, 200, rsp.Code)
	assert.Equal(t, uint64(103266), rsp.CoinValue)
}

func TestWalletTxs(t *testing.T) {
	var token string

	ts, cleanup := server.MockServer()
	defer cleanup()

	// Token.
	{
		body := APIGetToken(ts.URL, mockMobile, "vcode")
		rsp := &TokenResponse{}
		unmarshal(body, rsp)
		assert.Equal(t, 200, rsp.Code)
		token = rsp.Token
	}

	body := APIWalletTxs(ts.URL, token, 0, 2)
	rsp := &WalletTxsResponse{}
	unmarshal(body, rsp)

	t.Logf("%+v", body)
	assert.Equal(t, 200, rsp.Code)
	assert.Equal(t, 2, len(rsp.Txs))
}

func TestWalletAddresses(t *testing.T) {
	var token string

	ts, cleanup := server.MockServer()
	defer cleanup()

	// Token.
	{
		body := APIGetToken(ts.URL, mockMobile, "vcode")
		rsp := &TokenResponse{}
		unmarshal(body, rsp)
		assert.Equal(t, 200, rsp.Code)
		token = rsp.Token
	}

	body := APIWalletAddresses(ts.URL, token, 0, 2)
	rsp := &WalletAddressesResponse{}
	unmarshal(body, rsp)

	t.Logf("%+v", body)
	assert.Equal(t, 200, rsp.Code)
	assert.Equal(t, 2, len(rsp.Addresses))
}

func TestAPIEcdsaNewAddress(t *testing.T) {
	var token string

	ts, cleanup := server.MockServer()
	defer cleanup()

	// Token.
	{
		body := APIGetToken(ts.URL, mockMobile, "vcode")
		rsp := &TokenResponse{}
		unmarshal(body, rsp)
		assert.Equal(t, 200, rsp.Code)
		token = rsp.Token
	}

	for i := 0; i < 3; i++ {
		body := APIWalletNewAddress(ts.URL, token)
		rsp := &WalletNewAddressResponse{}
		unmarshal(body, rsp)

		t.Logf("%+v", body)
		assert.Equal(t, 200, rsp.Code)
	}
}

func TestAPIWalletSendFees(t *testing.T) {
	var token string

	ts, cleanup := server.MockServer()
	defer cleanup()

	// Token.
	{
		body := APIGetToken(ts.URL, mockMobile, "vcode")
		rsp := &TokenResponse{}
		unmarshal(body, rsp)
		assert.Equal(t, 200, rsp.Code)
		token = rsp.Token
	}

	{
		body := APIWalletSendFees(ts.URL, token, 100000)
		rsp := &WalletSendFeesResponse{}
		unmarshal(body, rsp)

		t.Logf("%+v", body)
		assert.Equal(t, 200, rsp.Code)
	}
}

func TestAPIWalletSend(t *testing.T) {
	var token string

	ts, cleanup := server.MockServer()
	defer cleanup()

	// Token.
	{
		body := APIGetToken(ts.URL, mockMobile, "vcode")
		rsp := &TokenResponse{}
		unmarshal(body, rsp)
		assert.Equal(t, 200, rsp.Code)
		token = rsp.Token
	}

	{
		body := APIWalletSend(ts.URL, token, "testnet", mockMasterPrvKey, "mmBRSnFG7o1BX5DaK8Da3xKxvjBh6fzNQq", 100000, 1000, "")
		rsp := &WalletSendResponse{}
		unmarshal(body, rsp)

		t.Logf("%+v", body)
		assert.Equal(t, 200, rsp.Code)
	}

	// Suffient value.
	{
		body := APIWalletSend(ts.URL, token, "testnet", mockMasterPrvKey, "mmBRSnFG7o1BX5DaK8Da3xKxvjBh6fzNQq", 1000000, 1000, "")
		rsp := &WalletSendResponse{}
		unmarshal(body, rsp)

		t.Logf("%+v", rsp)
		assert.Equal(t, 500, rsp.Code)
	}
}

func TestSignECDSACommon(t *testing.T) {
	var token string

	ts, cleanup := server.MockServer()
	defer cleanup()

	// Token.
	{
		body := APIGetToken(ts.URL, mockMobile, "vcode")
		rsp := &TokenResponse{}
		unmarshal(body, rsp)
		assert.Equal(t, 200, rsp.Code)
		token = rsp.Token
	}

	var err error

	var masterkey *bip32.HDKey

	rsp := &WalletSendResponse{}
	rsp.Code = http.StatusOK

	// Master pravite key.
	{
		masterkey, err = bip32.NewHDKeyFromString(mockMasterPrvKey)
		if err != nil {
			rsp.Code = http.StatusInternalServerError
			rsp.Message = err.Error()
			t.Fatal(err)
		}
	}

	cliPrvKey, err := masterkey.Derive(0)
	net := network.TestNet
	mockSvrMasterPrvKey := "tprv8ZgxMBicQKsPfNhXDHV93ummM6rEzTmxHf96Mk3FnpgoaoNYPjfSCZyHFnFQnQDLAiMNsvJqEtvjCkvo5P3CPRHQx5GcZxPqRHy31q2oWXD"
	svrpubkey, err := createSvrChildPubKey(0, mockSvrMasterPrvKey, net)
	svrPubKey, err := bip32.NewHDKeyFromString(svrpubkey)
	sigHash := []byte("helloword")
	sigBts, sharePubkey, err := SignECDSACommon(ts.URL, token, 0, sigHash, cliPrvKey, svrPubKey)
	if err != nil {
		t.Fatal(err)
	}

	sig := xcrypto.NewSignatureEcdsa()
	if err := sig.Deserialize(sigBts); err != nil {
		t.Fatal(err)
	}
	//rsBytes := make([]byte, 0)
	//rsBytes = append(rsBytes, sig.R.Bytes()...)
	//rsBytes = append(rsBytes, sig.S.Bytes()...)
	t.Log(hex.EncodeToString(sharePubkey.SerializeUncompressed()))
	//t.Log(hex.EncodeToString(rsBytes))

	ethRsBts, err := TransToEthSig(sig,sharePubkey,sigHash)
	if err != nil {
		t.Fatal(err)
	}

	v := ethRsBts[0] - 27
	copy(ethRsBts, ethRsBts[1:])
	ethRsBts[64] = v
	t.Log(hex.EncodeToString(ethRsBts))

	pubKeyrec, err := crypto.Ecrecover(sigHash, ethRsBts)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(hex.EncodeToString(pubKeyrec))

	if !crypto.VerifySignature(sharePubkey.Serialize(), sigHash, ethRsBts) {
		t.Fatal("verifysignature err")
	}
}


func createSvrChildPubKey(pos uint32, svrMasterPrvKey string, net *network.Network) (string, error) {
	svrmasterkey, err := bip32.NewHDKeyFromString(svrMasterPrvKey)
	if err != nil {
		return "", err
	}
	svrchild, err := svrmasterkey.Derive(pos)
	if err != nil {
		return "", err
	}
	return svrchild.HDPublicKey().ToString(net), nil
}

