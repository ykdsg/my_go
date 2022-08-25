package frame

import (
	"encoding/binary"
	"errors"
	"io"
)

type FramePayload []byte
type StreamFrameCodec interface {
	Encode(writer io.Writer, payload FramePayload) error // data->frame，并写入Write
	Decode(reader io.Reader) (FramePayload, error)       //从io.Reader中提取frame payload，并返回给上层
}

var ErrShortWrite = errors.New("short wrirte")
var ErrShortRead = errors.New("short read")

type myFrameCodec struct {
}

func (p *myFrameCodec) Encode(write io.Writer, framePayload FramePayload) error {
	var f = framePayload
	var totalLen int32 = int32(len(framePayload)) + 4

	//binary.Read 或 Write 会根据参数的宽度，读取或写入对应的字节个数的字节，这里 totalLen 使用 int32，
	//那么 Read 或 Write 只会操作数据流中的 4 个字节；
	err := binary.Write(write, binary.BigEndian, &totalLen)
	if err != nil {
		return err
	}

	n, err := write.Write([]byte(f)) // write the frame payload to outbound stream
	if n != len(framePayload) {
		return ErrShortWrite
	}
	return nil
}

func (p *myFrameCodec) Decode(reader io.Reader) (FramePayload, error) {
	var totalLen int32
	err := binary.Read(reader, binary.BigEndian, totalLen)
	if err != nil {
		return nil, err
	}
	bytes := make([]byte, 0, totalLen-4)
	length, err := reader.Read(bytes)
	if err != nil {
		return nil, err
	}
	if length != int(totalLen-4) {
		return nil, ErrShortRead
	}
	return bytes, nil

}

func NewMyFrameCodec() StreamFrameCodec {
	return &myFrameCodec{}
}
