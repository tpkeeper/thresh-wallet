// tokucore
//
// Copyright 2019 by KeyFuse Labs
// BSD License

package xcore

import (
	"github.com/keyfuse/tokucore/xerror"
)

// Error type.
const (
	ER_KEY_SIGNATURE_VERIFY_FAILED                 int = 2101
	ER_ADDRESS_CHECKSUM_MISMATCH                   int = 3101
	ER_ADDRESS_TYPE_UNKNOWN                        int = 3102
	ER_ADDRESS_FORMAT_MALFORMED                    int = 3103
	ER_ADDRESS_SIZE_MALFORMED                      int = 3104
	ER_ADDRESS_WITNESS_VERSION_UNSUPPORTED         int = 3105
	ER_SCRIPT_TYPE_UNKNOWN                         int = 4100
	ER_SCRIPT_STANDARD_ADDRESS_TYPE_UNSUPPORTED    int = 4105
	ER_SCRIPT_STANDARD_PUBKEYS_LE_NREQUIRED        int = 4106
	ER_SCRIPT_SIGNATURE_TYPE_UNKNOW                int = 4107
	ER_TRANSACTION_SIGN_OUT_INDEX                  int = 5000
	ER_TRANSACTION_SIGN_REDEEM_EMPTY               int = 5001
	ER_TRANSACTION_VERIFY_FAILED                   int = 5002
	ER_TRANSACTION_BUILDER_AMOUNT_NOT_ENOUGH_ERROR int = 5101
	ER_TRANSACTION_BUILDER_FROM_EMPTY              int = 5102
	ER_TRANSACTION_BUILDER_CHANGETO_EMPTY          int = 5103
	ER_TRANSACTION_BUILDER_SENDTO_EMPTY            int = 5104
	ER_TRANSACTION_BUILDER_SIGN_KEY_EMPTY          int = 5105
	ER_TRANSACTION_BUILDER_MIN_FEE_NOT_ENOUGH      int = 5106
	ER_TRANSACTION_BUILDER_FEE_TOO_HIGH            int = 5107
	ER_TRANSACTION_PARTIALLY_MAGIC_MISMATCH        int = 5201
	ER_MICROPAYMENT_LOCKTIME_MISMATCH              int = 5301
	ER_MICROPAYMENT_REFUND_BOND_MISMATCH           int = 5302
)

// Errors -- the jump table of error.
var Errors = map[int]*xerror.Error{
	ER_KEY_SIGNATURE_VERIFY_FAILED:                 {Num: ER_KEY_SIGNATURE_VERIFY_FAILED, State: "TKS00", Message: "key.signature.verify.failed"},
	ER_ADDRESS_CHECKSUM_MISMATCH:                   {Num: ER_ADDRESS_CHECKSUM_MISMATCH, State: "THK00", Message: "address.checksum.mismatch"},
	ER_ADDRESS_TYPE_UNKNOWN:                        {Num: ER_ADDRESS_TYPE_UNKNOWN, State: "TADDR0", Message: "address.unknown.type[%v]"},
	ER_ADDRESS_FORMAT_MALFORMED:                    {Num: ER_ADDRESS_FORMAT_MALFORMED, State: "TADDR0", Message: "address.unknown.format[%v]"},
	ER_ADDRESS_SIZE_MALFORMED:                      {Num: ER_ADDRESS_SIZE_MALFORMED, State: "TADDR0", Message: "address.size[%v].invalid"},
	ER_ADDRESS_WITNESS_VERSION_UNSUPPORTED:         {Num: ER_ADDRESS_WITNESS_VERSION_UNSUPPORTED, State: "TADDR0", Message: "address.witness.address.version[%v].unsupported"},
	ER_SCRIPT_TYPE_UNKNOWN:                         {Num: ER_SCRIPT_TYPE_UNKNOWN, State: "TS000", Message: "script.unknow.type[%v]"},
	ER_SCRIPT_STANDARD_ADDRESS_TYPE_UNSUPPORTED:    {Num: ER_SCRIPT_STANDARD_ADDRESS_TYPE_UNSUPPORTED, State: "TS000", Message: "script.standard.unsupported.address.type[%v]"},
	ER_SCRIPT_STANDARD_PUBKEYS_LE_NREQUIRED:        {Num: ER_SCRIPT_STANDARD_PUBKEYS_LE_NREQUIRED, State: "TS000", Message: "script.standard.pubkeys[%v].less.than.nrequired[%v]"},
	ER_SCRIPT_SIGNATURE_TYPE_UNKNOW:                {Num: ER_SCRIPT_SIGNATURE_TYPE_UNKNOW, State: "TS000", Message: "script.signature.type.unknow[%v]"},
	ER_TRANSACTION_SIGN_OUT_INDEX:                  {Num: ER_TRANSACTION_SIGN_OUT_INDEX, State: "TTX00", Message: "transaction.sign.idx[%v].out.index[%v]"},
	ER_TRANSACTION_SIGN_REDEEM_EMPTY:               {Num: ER_TRANSACTION_SIGN_REDEEM_EMPTY, State: "TTX00", Message: "transaction.sign.idx[%v].redeem.can.not.be.nil.since.keys[%v]>1"},
	ER_TRANSACTION_VERIFY_FAILED:                   {Num: ER_TRANSACTION_VERIFY_FAILED, State: "TTX00", Message: "transaction.verify.for.input[%v].referencing[%v].at[%v].failed"},
	ER_TRANSACTION_BUILDER_AMOUNT_NOT_ENOUGH_ERROR: {Num: ER_TRANSACTION_BUILDER_AMOUNT_NOT_ENOUGH_ERROR, State: "TTB00", Message: "transaction.builder.amount.totalout[%v].more.than.totalin[%v]"},
	ER_TRANSACTION_BUILDER_FROM_EMPTY:              {Num: ER_TRANSACTION_BUILDER_FROM_EMPTY, State: "TTB00", Message: "transaction.builder.from.is.empty"},
	ER_TRANSACTION_BUILDER_CHANGETO_EMPTY:          {Num: ER_TRANSACTION_BUILDER_CHANGETO_EMPTY, State: "TTB00", Message: "transaction.builder.changeto.is.empty"},
	ER_TRANSACTION_BUILDER_SENDTO_EMPTY:            {Num: ER_TRANSACTION_BUILDER_SENDTO_EMPTY, State: "TTB00", Message: "transaction.builder.sendto.is.empty"},
	ER_TRANSACTION_BUILDER_SIGN_KEY_EMPTY:          {Num: ER_TRANSACTION_BUILDER_SIGN_KEY_EMPTY, State: "TTB00", Message: "transaction.builder.sign.but.key.is.empty.at.input.idx[%v]"},
	ER_TRANSACTION_BUILDER_MIN_FEE_NOT_ENOUGH:      {Num: ER_TRANSACTION_BUILDER_MIN_FEE_NOT_ENOUGH, State: "TTB00", Message: "transaction.builder.min.fee[%v].not.enough.from.change.value[%v]"},
	ER_TRANSACTION_BUILDER_FEE_TOO_HIGH:            {Num: ER_TRANSACTION_BUILDER_FEE_TOO_HIGH, State: "TTB00", Message: "transaction.builder.fee[%v].too.high.than.max.fee[%v]"},
	ER_TRANSACTION_PARTIALLY_MAGIC_MISMATCH:        {Num: ER_TRANSACTION_PARTIALLY_MAGIC_MISMATCH, State: "TTP00", Message: "transaction.partially.request.magic.mismatch.want[%x].got[%x]"},
	ER_MICROPAYMENT_LOCKTIME_MISMATCH:              {Num: ER_MICROPAYMENT_LOCKTIME_MISMATCH, State: "TM000", Message: "micropayment.locktime.mismatch.want[%v].got[%v]"},
	ER_MICROPAYMENT_REFUND_BOND_MISMATCH:           {Num: ER_MICROPAYMENT_REFUND_BOND_MISMATCH, State: "TM000", Message: "micropayment.refund.bond.mismatch"},
}
