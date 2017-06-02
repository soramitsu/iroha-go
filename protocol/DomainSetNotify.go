// automatically generated by the FlatBuffers compiler, do not modify

package protocol

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type DomainSetNotify struct {
	_tab flatbuffers.Table
}

func GetRootAsDomainSetNotify(buf []byte, offset flatbuffers.UOffsetT) *DomainSetNotify {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DomainSetNotify{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *DomainSetNotify) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DomainSetNotify) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *DomainSetNotify) TargetDomain() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *DomainSetNotify) IsNotify() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *DomainSetNotify) MutateIsNotify(n byte) bool {
	return rcv._tab.MutateByteSlot(6, n)
}

func DomainSetNotifyStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func DomainSetNotifyAddTargetDomain(builder *flatbuffers.Builder, targetDomain flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(targetDomain), 0)
}
func DomainSetNotifyAddIsNotify(builder *flatbuffers.Builder, isNotify byte) {
	builder.PrependByteSlot(1, isNotify, 0)
}
func DomainSetNotifyEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
