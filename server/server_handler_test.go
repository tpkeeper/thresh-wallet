// thresh-wallet
//
// Copyright 2019 by KeyFuse Labs
//
// GPLv3 License

package server

import (
	"testing"

	"github.com/tpkeeper/thresh-wallet/proto"

	"github.com/stretchr/testify/assert"
)

func TestServerInfoHandler(t *testing.T) {
	ts, cleanup := MockServer()
	defer cleanup()

	{

		httpRsp, err := proto.NewRequest().Get(ts.URL + "/api/server/info")
		assert.Nil(t, err)
		rsp := &proto.ServerInfoResponse{}
		httpRsp.Json(rsp)
		t.Logf("rsp:%+v", rsp)
		assert.Equal(t, 200, httpRsp.StatusCode())
	}
}
