// automatically generated by the FlatBuffers compiler, do not modify

package iroha

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ReceiverConfirmation struct {
	_tab flatbuffers.Table
}

func GetRootAsReceiverConfirmation(buf []byte, offset flatbuffers.UOffsetT) *ReceiverConfirmation {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ReceiverConfirmation{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *ReceiverConfirmation) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ReceiverConfirmation) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ReceiverConfirmation) Signature(obj *Signature) *Signature {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Signature)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *ReceiverConfirmation) Hash(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *ReceiverConfirmation) HashLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ReceiverConfirmation) HashBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func ReceiverConfirmationStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func ReceiverConfirmationAddSignature(builder *flatbuffers.Builder, signature flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(signature), 0)
}
func ReceiverConfirmationAddHash(builder *flatbuffers.Builder, hash flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(hash), 0)
}
func ReceiverConfirmationStartHashVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func ReceiverConfirmationEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
