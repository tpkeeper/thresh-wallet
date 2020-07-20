// tokucore
//
// Copyright 2019 by KeyFuse Labs
// BSD License

package xcrypto

import (
	"testing"

	"encoding/hex"
	"github.com/stretchr/testify/assert"
)

type pubKeyTest struct {
	name    string
	key     []byte
	format  byte
	isValid bool
}

var pubKeyTests = []pubKeyTest{
	// pubkey from bitcoin blockchain tx
	// 0437cd7f8525ceed2324359c2d0ba26006d92d85
	{
		name: "uncompressed ok",
		key: []byte{0x04, 0x11, 0xdb, 0x93, 0xe1, 0xdc, 0xdb, 0x8a,
			0x01, 0x6b, 0x49, 0x84, 0x0f, 0x8c, 0x53, 0xbc, 0x1e,
			0xb6, 0x8a, 0x38, 0x2e, 0x97, 0xb1, 0x48, 0x2e, 0xca,
			0xd7, 0xb1, 0x48, 0xa6, 0x90, 0x9a, 0x5c, 0xb2, 0xe0,
			0xea, 0xdd, 0xfb, 0x84, 0xcc, 0xf9, 0x74, 0x44, 0x64,
			0xf8, 0x2e, 0x16, 0x0b, 0xfa, 0x9b, 0x8b, 0x64, 0xf9,
			0xd4, 0xc0, 0x3f, 0x99, 0x9b, 0x86, 0x43, 0xf6, 0x56,
			0xb4, 0x12, 0xa3,
		},
		isValid: true,
		format:  pubkeyUncompressed,
	},
	{
		name: "uncompressed x changed",
		key: []byte{0x04, 0x15, 0xdb, 0x93, 0xe1, 0xdc, 0xdb, 0x8a,
			0x01, 0x6b, 0x49, 0x84, 0x0f, 0x8c, 0x53, 0xbc, 0x1e,
			0xb6, 0x8a, 0x38, 0x2e, 0x97, 0xb1, 0x48, 0x2e, 0xca,
			0xd7, 0xb1, 0x48, 0xa6, 0x90, 0x9a, 0x5c, 0xb2, 0xe0,
			0xea, 0xdd, 0xfb, 0x84, 0xcc, 0xf9, 0x74, 0x44, 0x64,
			0xf8, 0x2e, 0x16, 0x0b, 0xfa, 0x9b, 0x8b, 0x64, 0xf9,
			0xd4, 0xc0, 0x3f, 0x99, 0x9b, 0x86, 0x43, 0xf6, 0x56,
			0xb4, 0x12, 0xa3,
		},
		isValid: false,
	},
	{
		name: "uncompressed y changed",
		key: []byte{0x04, 0x11, 0xdb, 0x93, 0xe1, 0xdc, 0xdb, 0x8a,
			0x01, 0x6b, 0x49, 0x84, 0x0f, 0x8c, 0x53, 0xbc, 0x1e,
			0xb6, 0x8a, 0x38, 0x2e, 0x97, 0xb1, 0x48, 0x2e, 0xca,
			0xd7, 0xb1, 0x48, 0xa6, 0x90, 0x9a, 0x5c, 0xb2, 0xe0,
			0xea, 0xdd, 0xfb, 0x84, 0xcc, 0xf9, 0x74, 0x44, 0x64,
			0xf8, 0x2e, 0x16, 0x0b, 0xfa, 0x9b, 0x8b, 0x64, 0xf9,
			0xd4, 0xc0, 0x3f, 0x99, 0x9b, 0x86, 0x43, 0xf6, 0x56,
			0xb4, 0x12, 0xa4,
		},
		isValid: false,
	},
	{
		name: "uncompressed claims compressed",
		key: []byte{0x03, 0x11, 0xdb, 0x93, 0xe1, 0xdc, 0xdb, 0x8a,
			0x01, 0x6b, 0x49, 0x84, 0x0f, 0x8c, 0x53, 0xbc, 0x1e,
			0xb6, 0x8a, 0x38, 0x2e, 0x97, 0xb1, 0x48, 0x2e, 0xca,
			0xd7, 0xb1, 0x48, 0xa6, 0x90, 0x9a, 0x5c, 0xb2, 0xe0,
			0xea, 0xdd, 0xfb, 0x84, 0xcc, 0xf9, 0x74, 0x44, 0x64,
			0xf8, 0x2e, 0x16, 0x0b, 0xfa, 0x9b, 0x8b, 0x64, 0xf9,
			0xd4, 0xc0, 0x3f, 0x99, 0x9b, 0x86, 0x43, 0xf6, 0x56,
			0xb4, 0x12, 0xa3,
		},
		isValid: false,
	},
	{
		name: "uncompressed as hybrid ok",
		key: []byte{0x07, 0x11, 0xdb, 0x93, 0xe1, 0xdc, 0xdb, 0x8a,
			0x01, 0x6b, 0x49, 0x84, 0x0f, 0x8c, 0x53, 0xbc, 0x1e,
			0xb6, 0x8a, 0x38, 0x2e, 0x97, 0xb1, 0x48, 0x2e, 0xca,
			0xd7, 0xb1, 0x48, 0xa6, 0x90, 0x9a, 0x5c, 0xb2, 0xe0,
			0xea, 0xdd, 0xfb, 0x84, 0xcc, 0xf9, 0x74, 0x44, 0x64,
			0xf8, 0x2e, 0x16, 0x0b, 0xfa, 0x9b, 0x8b, 0x64, 0xf9,
			0xd4, 0xc0, 0x3f, 0x99, 0x9b, 0x86, 0x43, 0xf6, 0x56,
			0xb4, 0x12, 0xa3,
		},
		isValid: true,
		format:  pubkeyHybrid,
	},
	{
		name: "uncompressed as hybrid wrong",
		key: []byte{0x06, 0x11, 0xdb, 0x93, 0xe1, 0xdc, 0xdb, 0x8a,
			0x01, 0x6b, 0x49, 0x84, 0x0f, 0x8c, 0x53, 0xbc, 0x1e,
			0xb6, 0x8a, 0x38, 0x2e, 0x97, 0xb1, 0x48, 0x2e, 0xca,
			0xd7, 0xb1, 0x48, 0xa6, 0x90, 0x9a, 0x5c, 0xb2, 0xe0,
			0xea, 0xdd, 0xfb, 0x84, 0xcc, 0xf9, 0x74, 0x44, 0x64,
			0xf8, 0x2e, 0x16, 0x0b, 0xfa, 0x9b, 0x8b, 0x64, 0xf9,
			0xd4, 0xc0, 0x3f, 0x99, 0x9b, 0x86, 0x43, 0xf6, 0x56,
			0xb4, 0x12, 0xa3,
		},
		isValid: false,
	},
	// from tx 0b09c51c51ff762f00fb26217269d2a18e77a4fa87d69b3c363ab4df16543f20
	{
		name: "compressed ok (ybit = 0)",
		key: []byte{0x02, 0xce, 0x0b, 0x14, 0xfb, 0x84, 0x2b, 0x1b,
			0xa5, 0x49, 0xfd, 0xd6, 0x75, 0xc9, 0x80, 0x75, 0xf1,
			0x2e, 0x9c, 0x51, 0x0f, 0x8e, 0xf5, 0x2b, 0xd0, 0x21,
			0xa9, 0xa1, 0xf4, 0x80, 0x9d, 0x3b, 0x4d,
		},
		isValid: true,
		format:  pubkeyCompressed,
	},
	// from tx fdeb8e72524e8dab0da507ddbaf5f88fe4a933eb10a66bc4745bb0aa11ea393c
	{
		name: "compressed ok (ybit = 1)",
		key: []byte{0x03, 0x26, 0x89, 0xc7, 0xc2, 0xda, 0xb1, 0x33,
			0x09, 0xfb, 0x14, 0x3e, 0x0e, 0x8f, 0xe3, 0x96, 0x34,
			0x25, 0x21, 0x88, 0x7e, 0x97, 0x66, 0x90, 0xb6, 0xb4,
			0x7f, 0x5b, 0x2a, 0x4b, 0x7d, 0x44, 0x8e,
		},
		isValid: true,
		format:  pubkeyCompressed,
	},
	{
		name: "compressed claims uncompressed (ybit = 0)",
		key: []byte{0x04, 0xce, 0x0b, 0x14, 0xfb, 0x84, 0x2b, 0x1b,
			0xa5, 0x49, 0xfd, 0xd6, 0x75, 0xc9, 0x80, 0x75, 0xf1,
			0x2e, 0x9c, 0x51, 0x0f, 0x8e, 0xf5, 0x2b, 0xd0, 0x21,
			0xa9, 0xa1, 0xf4, 0x80, 0x9d, 0x3b, 0x4d,
		},
		isValid: false,
	},
	{
		name: "compressed claims uncompressed (ybit = 1)",
		key: []byte{0x05, 0x26, 0x89, 0xc7, 0xc2, 0xda, 0xb1, 0x33,
			0x09, 0xfb, 0x14, 0x3e, 0x0e, 0x8f, 0xe3, 0x96, 0x34,
			0x25, 0x21, 0x88, 0x7e, 0x97, 0x66, 0x90, 0xb6, 0xb4,
			0x7f, 0x5b, 0x2a, 0x4b, 0x7d, 0x44, 0x8e,
		},
		isValid: false,
	},
	{
		name:    "wrong length)",
		key:     []byte{0x05},
		isValid: false,
	},
	{
		name: "X == P",
		key: []byte{0x04, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
			0xFF, 0xFF, 0xFE, 0xFF, 0xFF, 0xFC, 0x2F, 0xb2, 0xe0,
			0xea, 0xdd, 0xfb, 0x84, 0xcc, 0xf9, 0x74, 0x44, 0x64,
			0xf8, 0x2e, 0x16, 0x0b, 0xfa, 0x9b, 0x8b, 0x64, 0xf9,
			0xd4, 0xc0, 0x3f, 0x99, 0x9b, 0x86, 0x43, 0xf6, 0x56,
			0xb4, 0x12, 0xa3,
		},
		isValid: false,
	},
	{
		name: "X > P",
		key: []byte{0x04, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
			0xFF, 0xFF, 0xFE, 0xFF, 0xFF, 0xFD, 0x2F, 0xb2, 0xe0,
			0xea, 0xdd, 0xfb, 0x84, 0xcc, 0xf9, 0x74, 0x44, 0x64,
			0xf8, 0x2e, 0x16, 0x0b, 0xfa, 0x9b, 0x8b, 0x64, 0xf9,
			0xd4, 0xc0, 0x3f, 0x99, 0x9b, 0x86, 0x43, 0xf6, 0x56,
			0xb4, 0x12, 0xa3,
		},
		isValid: false,
	},
	{
		name: "Y == P",
		key: []byte{0x04, 0x11, 0xdb, 0x93, 0xe1, 0xdc, 0xdb, 0x8a,
			0x01, 0x6b, 0x49, 0x84, 0x0f, 0x8c, 0x53, 0xbc, 0x1e,
			0xb6, 0x8a, 0x38, 0x2e, 0x97, 0xb1, 0x48, 0x2e, 0xca,
			0xd7, 0xb1, 0x48, 0xa6, 0x90, 0x9a, 0x5c, 0xFF, 0xFF,
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFE, 0xFF,
			0xFF, 0xFC, 0x2F,
		},
		isValid: false,
	},
	{
		name: "Y > P",
		key: []byte{0x04, 0x11, 0xdb, 0x93, 0xe1, 0xdc, 0xdb, 0x8a,
			0x01, 0x6b, 0x49, 0x84, 0x0f, 0x8c, 0x53, 0xbc, 0x1e,
			0xb6, 0x8a, 0x38, 0x2e, 0x97, 0xb1, 0x48, 0x2e, 0xca,
			0xd7, 0xb1, 0x48, 0xa6, 0x90, 0x9a, 0x5c, 0xFF, 0xFF,
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFE, 0xFF,
			0xFF, 0xFD, 0x2F,
		},
		isValid: false,
	},
	{
		name: "hybrid",
		key: []byte{0x06, 0x79, 0xbe, 0x66, 0x7e, 0xf9, 0xdc, 0xbb,
			0xac, 0x55, 0xa0, 0x62, 0x95, 0xce, 0x87, 0x0b, 0x07,
			0x02, 0x9b, 0xfc, 0xdb, 0x2d, 0xce, 0x28, 0xd9, 0x59,
			0xf2, 0x81, 0x5b, 0x16, 0xf8, 0x17, 0x98, 0x48, 0x3a,
			0xda, 0x77, 0x26, 0xa3, 0xc4, 0x65, 0x5d, 0xa4, 0xfb,
			0xfc, 0x0e, 0x11, 0x08, 0xa8, 0xfd, 0x17, 0xb4, 0x48,
			0xa6, 0x85, 0x54, 0x19, 0x9c, 0x47, 0xd0, 0x8f, 0xfb,
			0x10, 0xd4, 0xb8,
		},
		format:  pubkeyHybrid,
		isValid: true,
	},
}

func TestPubKeys(t *testing.T) {
	for _, test := range pubKeyTests {
		pk, err := PubKeyFromBytes(test.key)
		if err != nil {
			if test.isValid {
				t.Errorf("%s pubkey failed when shouldn't %v",
					test.name, err)
			}
			continue
		}
		if !test.isValid {
			t.Errorf("%s counted as valid when it should fail",
				test.name)
			continue
		}

		var pkStr []byte
		switch test.format {
		case pubkeyUncompressed:
			pkStr = (*PubKey)(pk).SerializeUncompressed()
		case pubkeyCompressed:
			pkStr = (*PubKey)(pk).SerializeCompressed()
		case pubkeyHybrid:
			pkStr = (*PubKey)(pk).SerializeHybrid()
		}
		assert.Equal(t, test.key, pkStr)
	}
}

func TestPubKey1(t *testing.T) {
	// testnet tx:6678127a46d4cb91ddcee590c81d58f472b5548aa0e7fcb11deb4fe0fdd33a5b
	key, err := hex.DecodeString("02ec74ff01e69625f89202aa8f1748e0e22ac2dcf9d30c5f4b0e6f608970bd8fd0")
	assert.Nil(t, err)

	pubKey, err := PubKeyFromBytes(key)
	assert.Nil(t, err)

	compressed := pubKey.SerializeCompressed()
	t.Logf("compressed: %X, len:%v", compressed, len(compressed))
	t.Logf("compressed.hash: %X", Hash160(compressed))

	uncompressed := pubKey.SerializeUncompressed()
	t.Logf("uncompressed: %X, len:%v", uncompressed, len(uncompressed))
	t.Logf("uncompressed.hash: %X", Hash160(uncompressed))

	hybrid := pubKey.SerializeHybrid()
	t.Logf("hybrid: %x, len:%v", hybrid, len(hybrid))
}

func TestPubKeyAdd(t *testing.T) {
	hex1 := []byte{0x01}
	prvkey1 := PrvKeyFromBytes(hex1)
	pubkey1 := prvkey1.PubKey()

	hex2 := []byte{0x02}
	prvkey2 := PrvKeyFromBytes(hex2)
	pubkey2 := prvkey2.PubKey()
	pubkey2 = pubkey2.Add(pubkey1)

	hex3 := []byte{0x03}
	prvkey3 := PrvKeyFromBytes(hex3)
	pubkey3 := prvkey3.PubKey()
	assert.Equal(t, pubkey2, pubkey3)
}
