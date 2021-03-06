// automatically generated by the FlatBuffers compiler, do not modify

package iroha

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type AssetQuery struct {
	_tab flatbuffers.Table
}

func GetRootAsAssetQuery(buf []byte, offset flatbuffers.UOffsetT) *AssetQuery {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AssetQuery{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *AssetQuery) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AssetQuery) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *AssetQuery) PubKey() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *AssetQuery) LedgerName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *AssetQuery) DomainName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *AssetQuery) AssetName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *AssetQuery) Uncommitted() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *AssetQuery) MutateUncommitted(n byte) bool {
	return rcv._tab.MutateByteSlot(12, n)
}

func AssetQueryStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func AssetQueryAddPubKey(builder *flatbuffers.Builder, pubKey flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(pubKey), 0)
}
func AssetQueryAddLedgerName(builder *flatbuffers.Builder, ledgerName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(ledgerName), 0)
}
func AssetQueryAddDomainName(builder *flatbuffers.Builder, domainName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(domainName), 0)
}
func AssetQueryAddAssetName(builder *flatbuffers.Builder, assetName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(assetName), 0)
}
func AssetQueryAddUncommitted(builder *flatbuffers.Builder, uncommitted byte) {
	builder.PrependByteSlot(4, uncommitted, 0)
}
func AssetQueryEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
