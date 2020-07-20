// tokucore
//
// Copyright 2019 by KeyFuse Labs
// BSD License

package xbase

import (
	"fmt"
)

// WitnessDecode -- decodes the segwit address to hrp, version and pubkeyscript.
func WitnessDecode(addr string) (string, byte, []byte, error) {
	hrp, data, err := Bech32Decode(addr)
	if err != nil {
		return "", 0, nil, err
	}
	res, err := convertBits(data[1:], 5, 8, false)
	if err != nil {
		return "", 0, nil, err
	}
	return hrp, data[0], res, nil
}

// WitnessEncode -- encodes to segwit address.
func WitnessEncode(hrp string, version byte, program []byte) (string, error) {
	data, err := convertBits(program, 8, 5, true)
	if err != nil {
		return "", err
	}
	return Bech32Encode(hrp, append([]byte{version}, data...))
}

func convertBits(data []byte, frombits, tobits uint, pad bool) ([]byte, error) {
	acc := 0
	bits := uint(0)
	ret := []byte{}
	maxv := (1 << tobits) - 1
	maxAcc := (1 << (frombits + tobits - 1)) - 1

	for _, value := range data {
		acc = (acc << frombits) | int(value)&maxAcc
		bits += frombits
		for bits >= tobits {
			bits -= tobits
			ret = append(ret, byte(acc>>bits&maxv))
		}
	}

	if pad {
		if bits > 0 {
			ret = append(ret, byte(acc<<(tobits-bits)&maxv))
		}
	} else {
		if bits >= frombits {
			return nil, fmt.Errorf("illegal zero padding")
		}
		if ((acc << (tobits - bits)) & maxv) != 0 {
			return nil, fmt.Errorf("non-zero padding")
		}
	}
	return ret, nil
}
