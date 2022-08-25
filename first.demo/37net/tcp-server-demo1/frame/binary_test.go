package frame

import (
	"bytes"
	"encoding/binary"
	"testing"
	"unsafe"
)

func TestBinary_write(t *testing.T) {
	buf := make([]byte, 0, 128)
	rw := bytes.NewBuffer(buf)
	var writeNum int32 = 9

	//这里writeNum可以用指针，也可以不用指针，因为实际要写入的数据是9（int32），使用指针的好处是节省了复制的开销。
	//binary.Read 或 Write 会根据参数的宽度，读取或写入对应的字节个数的字节，这里int32 是4字节，那么只会操作数据流中的 4 个字节
	err := binary.Write(rw, binary.BigEndian, &writeNum)
	if err != nil {
		t.Errorf("write want nil, actual %s", err.Error())
	}

	sizeofNum := int(unsafe.Sizeof(writeNum))
	lenBuf := rw.Len()
	if lenBuf != sizeofNum {
		t.Errorf("len buf is %d,actual size writeNum is %d", lenBuf, sizeofNum)
	}
	var readNum int32
	//这里readNum 必须使用指针，不然后面拿不到值
	err = binary.Read(rw, binary.BigEndian, &readNum)
	if err != nil {
		t.Errorf("read want nil, actual %s", err.Error())
	}

	if readNum != writeNum {
		t.Errorf("want %d,actual %d", writeNum, readNum)
	}

}
