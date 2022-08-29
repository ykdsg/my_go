package main

import (
	"bufio"
	"fmt"
	"github.com/ykdsg/tcp-server-demo2/frame"
	"github.com/ykdsg/tcp-server-demo2/metrics"
	"github.com/ykdsg/tcp-server-demo2/packet"
	"net"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() { http.ListenAndServe(":6060", nil) }()

	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}
	fmt.Println("server start ok(on *.8888)")
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
	metrics.ClientConnected.Inc() // 连接建立，ClientConnected加1
	defer func() {
		metrics.ClientConnected.Dec() // 连接断开，ClientConnected减1
		c.Close()
	}()

	frameCodec := frame.NewMyFrameCodec()
	rbuf := bufio.NewReader(c)
	wbuf := bufio.NewWriter(c)

	defer wbuf.Flush()
	for {
		framePayload, err := frameCodec.Decode(rbuf)
		if err != nil {
			fmt.Println("handleConn:frame decode error:", err)
			return
		}
		metrics.ReqRecvTotal.Add(1) // 收到并解码一个消息请求，ReqRecvTotal消息计数加1
		ackFramePayload, err := handlePacket(framePayload)
		if err != nil {
			fmt.Println("handleConn: handle packet error:", err)
			return
		}

		// write ack frame to the connection
		err = frameCodec.Encode(wbuf, ackFramePayload)
		if err != nil {
			fmt.Println("handleConn: frame encode error:", err)
			return
		}
		metrics.RspSendTotal.Add(1) // 返回响应后，RspSendTotal消息计数器加1
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
		submitAck := &packet.SubmitAck{
			ID:     submit.ID,
			Result: 0,
		}
		ackFramePayload, err = packet.Encode(submitAck)
		if err != nil {
			fmt.Println("handleConn:packet encode error:", err)
			return nil, err
		}
		return ackFramePayload, nil

	default:
		return nil, fmt.Errorf("unkonw packet type")
	}
}
