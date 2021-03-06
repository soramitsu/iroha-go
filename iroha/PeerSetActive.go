// automatically generated by the FlatBuffers compiler, do not modify

package iroha

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type PeerSetActive struct {
	_tab flatbuffers.Table
}

func GetRootAsPeerSetActive(buf []byte, offset flatbuffers.UOffsetT) *PeerSetActive {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &PeerSetActive{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *PeerSetActive) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *PeerSetActive) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *PeerSetActive) PeerPubKey() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *PeerSetActive) Active() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *PeerSetActive) MutateActive(n byte) bool {
	return rcv._tab.MutateByteSlot(6, n)
}

func PeerSetActiveStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func PeerSetActiveAddPeerPubKey(builder *flatbuffers.Builder, peerPubKey flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(peerPubKey), 0)
}
func PeerSetActiveAddActive(builder *flatbuffers.Builder, active byte) {
	builder.PrependByteSlot(1, active, 0)
}
func PeerSetActiveEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
