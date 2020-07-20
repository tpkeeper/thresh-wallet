// tokucore
//
// Copyright 2019 by KeyFuse Labs
// BSD License

package xprotocol

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/keyfuse/tokucore/network"
)

func TestConnection(t *testing.T) {
	network := network.MainNet
	endpoint := "127.0.0.1:3338"

	client, cleanup := newMockNode(network, endpoint)
	defer cleanup()
	// Handshake.
	{
		// Send Version.
		{
			req := NewMsgVersion(network)
			err := client.Send(req)
			assert.Nil(t, err)
		}

		// Send VerAck.
		{
			req := NewMsgVerAck()
			err := client.Send(req)
			assert.Nil(t, err)
		}
	}
	time.Sleep(time.Second)
}
