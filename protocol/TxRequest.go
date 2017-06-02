// automatically generated by the FlatBuffers compiler, do not modify

package protocol

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type TxRequest struct {
	_tab flatbuffers.Table
}

func GetRootAsTxRequest(buf []byte, offset flatbuffers.UOffsetT) *TxRequest {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &TxRequest{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *TxRequest) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *TxRequest) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *TxRequest) Index() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *TxRequest) MutateIndex(n uint64) bool {
	return rcv._tab.MutateUint64Slot(4, n)
}

func (rcv *TxRequest) Sender() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func TxRequestStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func TxRequestAddIndex(builder *flatbuffers.Builder, index uint64) {
	builder.PrependUint64Slot(0, index, 0)
}
func TxRequestAddSender(builder *flatbuffers.Builder, sender flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(sender), 0)
}
func TxRequestEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}