// automatically generated by the FlatBuffers compiler, do not modify

package protocol

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type AccountUndelegate struct {
	_tab flatbuffers.Table
}

func GetRootAsAccountUndelegate(buf []byte, offset flatbuffers.UOffsetT) *AccountUndelegate {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AccountUndelegate{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *AccountUndelegate) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AccountUndelegate) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *AccountUndelegate) AuditorAccount() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func AccountUndelegateStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func AccountUndelegateAddAuditorAccount(builder *flatbuffers.Builder, auditorAccount flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(auditorAccount), 0)
}
func AccountUndelegateEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
