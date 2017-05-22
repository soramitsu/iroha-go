// automatically generated by the FlatBuffers compiler, do not modify

package iroha

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type CheckHashResponse struct {
	_tab flatbuffers.Table
}

func GetRootAsCheckHashResponse(buf []byte, offset flatbuffers.UOffsetT) *CheckHashResponse {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &CheckHashResponse{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *CheckHashResponse) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *CheckHashResponse) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *CheckHashResponse) IsCorrect() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *CheckHashResponse) MutateIsCorrect(n byte) bool {
	return rcv._tab.MutateByteSlot(4, n)
}

func (rcv *CheckHashResponse) IsRoot() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *CheckHashResponse) MutateIsRoot(n byte) bool {
	return rcv._tab.MutateByteSlot(6, n)
}

func (rcv *CheckHashResponse) IsExist() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *CheckHashResponse) MutateIsExist(n byte) bool {
	return rcv._tab.MutateByteSlot(8, n)
}

func CheckHashResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func CheckHashResponseAddIsCorrect(builder *flatbuffers.Builder, isCorrect byte) {
	builder.PrependByteSlot(0, isCorrect, 0)
}
func CheckHashResponseAddIsRoot(builder *flatbuffers.Builder, isRoot byte) {
	builder.PrependByteSlot(1, isRoot, 0)
}
func CheckHashResponseAddIsExist(builder *flatbuffers.Builder, isExist byte) {
	builder.PrependByteSlot(2, isExist, 0)
}
func CheckHashResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
