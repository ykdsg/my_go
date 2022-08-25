package main

import (
	"fmt"
	"net"
	"tcp-server-demo1/frame"
	"tcp-server-demo1/packet"
)

func main() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			break
		}
		go handleConn(c)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	frameCodec := frame.NewMyFrameCodec()
	for {
		framePayload, err := frameCodec.Decode(c)
		if err != nil {
			fmt.Println("handleConn:frame decode error:", err)
			return
		}
		ackFramePayload, err := handlePacket(framePayload)
		if err != nil {
			fmt.Println("handleConn: handle packet error:", err)
			return
		}

		// write ack frame to the connection
		err = frameCodec.Encode(c, ackFramePayload)
		if err != nil {
			fmt.Println("handleConn: frame encode error:", err)
			return
		}
	}
}

func handlePacket(framePayload []byte) (ackFramePayload []byte, err error) {
	var p packet.Packet
	p, err = packet.Decode(framePayload)
	if err != nil {
		fmt.Println("handleConn: packet decode error:", err)
		return
	}
	switch p.(type) {
	case *packet.Submit:
		submit := p.(*packet.Submit)
		fmt.Printf("recv submit: id=%s,payload=%s\n", submit.ID, string(submit.Payload))
		submitAck := packet.SubmitAck{
			ID:     submit.ID,
			Result: 0,
		}
		ackFramePayload, err = packet.Encode(&submitAck)
		if err != nil {
			fmt.Println("handleConn:packet encode error:", err)
			return nil, err
		}
		return ackFramePayload, nil

	default:
		return nil, fmt.Errorf("unkonw packet type")
	}
}
