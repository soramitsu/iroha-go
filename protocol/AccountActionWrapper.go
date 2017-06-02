// automatically generated by the FlatBuffers compiler, do not modify

package protocol

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type AccountActionWrapper struct {
	_tab flatbuffers.Table
}

func GetRootAsAccountActionWrapper(buf []byte, offset flatbuffers.UOffsetT) *AccountActionWrapper {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AccountActionWrapper{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *AccountActionWrapper) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AccountActionWrapper) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *AccountActionWrapper) ActionType() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *AccountActionWrapper) MutateActionType(n byte) bool {
	return rcv._tab.MutateByteSlot(4, n)
}

func (rcv *AccountActionWrapper) Action(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

func AccountActionWrapperStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func AccountActionWrapperAddActionType(builder *flatbuffers.Builder, actionType byte) {
	builder.PrependByteSlot(0, actionType, 0)
}
func AccountActionWrapperAddAction(builder *flatbuffers.Builder, action flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(action), 0)
}
func AccountActionWrapperEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}