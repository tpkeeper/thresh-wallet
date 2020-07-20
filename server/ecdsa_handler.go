// thresh-wallet
//
// Copyright 2019 by KeyFuse Labs
//
// GPLv3 License

package server

import (
	"encoding/json"
	"net/http"

	"github.com/tpkeeper/thresh-wallet/proto"
)

// ecdsaR2 -- the handler of creating R2 of two party.
func (h *Handler) ecdsaR2(w http.ResponseWriter, r *http.Request) {
	log := h.log
	wdb := h.wdb
	resp := newResponse(log, w)

	// UID.
	uid, err := h.userinfo("ecdsaR2", r)
	if err != nil {
		log.Error("api.ecdsa.r2.uid.error:%+v", err)
		resp.writeError(err)
		return
	}

	// Request.
	req := &proto.EcdsaR2Request{}
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		log.Error("api.ecdsa.r2.req.decode.error:%+v", err)
		resp.writeError(err)
	}
	log.Info("api.ecdsa.r2.req:%+v", req)

	// Master Private Key.
	masterPrvKey, err := wdb.MasterPrvKey(uid)
	if err != nil {
		log.Error("api.ecdsa.r2[%v].master.prvkey.error:%+v", uid, err)
		resp.writeError(err)
		return
	}

	// R2.
	r2, shareR, err := createEcdsaR2(req.Pos, masterPrvKey, req.Hash, req.R1)
	if err != nil {
		log.Error("api.ecdsa.r2[%v].create.ecdsar2.error:%+v", uid, err)
		resp.writeError(err)
		return
	}
	rsp := &proto.EcdsaR2Response{
		R2:     r2,
		ShareR: shareR,
	}
	log.Info("api.ecdsa.r2.rsp:%+v", rsp)
	resp.writeJSON(rsp)
}

// ecdsaS2 -- the handler of creates S2 of two party.
func (h *Handler) ecdsaS2(w http.ResponseWriter, r *http.Request) {
	log := h.log
	wdb := h.wdb
	resp := newResponse(log, w)

	// UID.
	uid, err := h.userinfo("ecdsaS2", r)
	if err != nil {
		log.Error("api.ecdsa.s2.uid.error:%+v", err)
		resp.writeError(err)
		return
	}

	// Request.
	req := &proto.EcdsaS2Request{}
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		log.Error("api.ecdsa.s2[%v].decode.body.error:%+v", uid, err)
		resp.writeError(err)
		return
	}
	log.Info("api.ecdsa.s2.req:%+v", req)

	// Master Private Key.
	masterPrvKey, err := wdb.MasterPrvKey(uid)
	if err != nil {
		log.Error("api.ecdsa.s2[%v].master.prvkey.error:%+v", uid, err)
		resp.writeError(err)
		return
	}

	// S2.
	s2, err := createEcdsaS2(req.Pos, masterPrvKey, req.Hash, req.R1, req.ShareR, req.EncPK1, req.EncPub1)
	if err != nil {
		log.Error("api.ecdsa.s2[%v].create.ecdsar2.error:%+v", uid, err)
		resp.writeError(err)
		return
	}
	rsp := &proto.EcdsaS2Response{
		S2: s2,
	}
	log.Info("api.ecdsa.s2.rsp:%+v", rsp)
	resp.writeJSON(rsp)
}
