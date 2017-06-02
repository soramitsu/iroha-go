// automatically generated by the FlatBuffers compiler, do not modify

package protocol

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Asset struct {
	_tab flatbuffers.Table
}

func GetRootAsAsset(buf []byte, offset flatbuffers.UOffsetT) *Asset {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Asset{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Asset) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Asset) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Asset) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Asset) Uuid() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Asset) AssetType() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Asset) MutateAssetType(n byte) bool {
	return rcv._tab.MutateByteSlot(8, n)
}

func (rcv *Asset) Asset(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

func AssetStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func AssetAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func AssetAddUuid(builder *flatbuffers.Builder, uuid flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(uuid), 0)
}
func AssetAddAssetType(builder *flatbuffers.Builder, assetType byte) {
	builder.PrependByteSlot(2, assetType, 0)
}
func AssetAddAsset(builder *flatbuffers.Builder, asset flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(asset), 0)
}
func AssetEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}