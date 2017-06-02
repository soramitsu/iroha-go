// automatically generated by the FlatBuffers compiler, do not modify

package protocol

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type DomainDelete struct {
	_tab flatbuffers.Table
}

func GetRootAsDomainDelete(buf []byte, offset flatbuffers.UOffsetT) *DomainDelete {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DomainDelete{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *DomainDelete) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DomainDelete) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *DomainDelete) TargetDomain() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func DomainDeleteStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func DomainDeleteAddTargetDomain(builder *flatbuffers.Builder, targetDomain flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(targetDomain), 0)
}
func DomainDeleteEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}