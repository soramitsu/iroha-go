// automatically generated by the FlatBuffers compiler, do not modify

package protocol

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type HijiriResponse struct {
	_tab flatbuffers.Table
}

func GetRootAsHijiriResponse(buf []byte, offset flatbuffers.UOffsetT) *HijiriResponse {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &HijiriResponse{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *HijiriResponse) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *HijiriResponse) Table() flatbuffers.Table {
	return rcv._tab
}

func HijiriResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(0)
}
func HijiriResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}