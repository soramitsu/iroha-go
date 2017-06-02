// automatically generated by the FlatBuffers compiler, do not modify

package protocol

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type AssetSetPermission struct {
	_tab flatbuffers.Table
}

func GetRootAsAssetSetPermission(buf []byte, offset flatbuffers.UOffsetT) *AssetSetPermission {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AssetSetPermission{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *AssetSetPermission) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AssetSetPermission) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *AssetSetPermission) Account() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *AssetSetPermission) AssetUuid() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *AssetSetPermission) Permission(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *AssetSetPermission) PermissionLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *AssetSetPermission) PermissionBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func AssetSetPermissionStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func AssetSetPermissionAddAccount(builder *flatbuffers.Builder, account flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(account), 0)
}
func AssetSetPermissionAddAssetUuid(builder *flatbuffers.Builder, assetUuid flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(assetUuid), 0)
}
func AssetSetPermissionAddPermission(builder *flatbuffers.Builder, permission flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(permission), 0)
}
func AssetSetPermissionStartPermissionVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func AssetSetPermissionEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}