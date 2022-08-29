package packet

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSubmit_Encode(t *testing.T) {
	id := fmt.Sprintf("%08d", 1)
	bytes := []byte("hello")
	submit := Submit{
		ID:      id,
		Payload: bytes,
	}
	encode, err := Encode(&submit)
	if err != nil {
		t.Errorf("want nil, actual %s", err.Error())
	}
	sizeofUint := 1
	//commandID 是uint8
	totalLen := len(id) + len(bytes) + sizeofUint

	if len(encode) != totalLen {
		t.Errorf("want len(encode)  %d, actual   %d", totalLen, len(encode))
	}
	println(encode)
}

func TestSubmit_Decode(t *testing.T) {
	id := fmt.Sprintf("%08d", 1)
	bytes := []byte("hello")
	submit := Submit{
		ID:      id,
		Payload: bytes,
	}
	encode, err := Encode(&submit)
	if err != nil {
		t.Errorf("want nil, actual %s", err.Error())
	}
	packet, err := Decode(encode)
	if err != nil {
		t.Errorf("want nil, actual %s", err.Error())
	}
	submitPacket, ok := packet.(*Submit)
	if !ok {
		t.Error("want submite,but not")
	}
	if submitPacket.ID != submitPacket.ID {
		t.Errorf("ID not equal,want %s ,actual %s", submit.ID, submitPacket.ID)
	}

	if !reflect.DeepEqual(submitPacket.Payload, submit.Payload) {
		t.Error("payload not equal")
	}

}
