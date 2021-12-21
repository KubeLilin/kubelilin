package kubernetes

import (
	"encoding/json"
	"fmt"
	"github.com/fasthttp/websocket"
	"io"
	"k8s.io/client-go/tools/remotecommand"
	"log"
	"time"
)

//心跳检查的意义在于生产环境是lb  nginx ，有代理超时设置

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Maximum message size allowed from peer.
	maxMessageSize = 8192

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Time to wait before force close on connection.
	closeGracePeriod = 10 * time.Second

	// EndOfTransmission end
	EndOfTransmission = "\u0004"
)

type (
	/*
	   1 web container terminal
	   2 20分钟无输入要退出bash进程
	   3 需要支持心跳，不然websocket过腾讯云lb只能持续2分钟
	   4 程序ctrl c之后，断开所有容器的bash进程 //未能成功
	*/
	WebTerminal struct {
		wsConn   *websocket.Conn
		sizeChan chan remotecommand.TerminalSize
		doneChan chan struct{}
		tty      bool
	}

	//TerminalMessage Terminal Message
	TerminalMessage struct {
		Operation string `json:"operation"`
		Data      string `json:"data"`
		Rows      uint16 `json:"rows"`
		Cols      uint16 `json:"cols"`
	}
)

// NewWebTerminal web terminal的实现
func NewWebTerminal(conn *websocket.Conn) *WebTerminal {
	return &WebTerminal{
		wsConn:   conn,
		tty:      true,
		sizeChan: make(chan remotecommand.TerminalSize),
		doneChan: make(chan struct{}),
	}
}

func (t *WebTerminal) Done() {
	close(t.doneChan)
}

func (t *WebTerminal) Next() *remotecommand.TerminalSize {
	select {
	case size := <-t.sizeChan:
		return &size
	case <-t.doneChan:
		return nil
	}
}

// Stdin ...
func (t *WebTerminal) Stdin() io.Reader {
	return t
}

// Stdout ...
func (t *WebTerminal) Stdout() io.Writer {
	return t
}

// Stderr ...
func (t *WebTerminal) Stderr() io.Writer {
	return t
}

// 模拟stdout，stderr
func (t *WebTerminal) Write(p []byte) (n int, err error) {
	msg, err := json.Marshal(TerminalMessage{
		Operation: "stdout",
		Data:      string(p),
	})
	if err != nil {
		log.Printf("write parse message err: %v", err)
		return 0, err
	}
	if err := t.wsConn.WriteMessage(websocket.TextMessage, msg); err != nil {
		log.Printf("write message err: %v", err)
		return 0, err
	}
	return len(p), nil
}

// 模拟stdin
func (t *WebTerminal) Read(p []byte) (n int, err error) {
	_, message, err := t.wsConn.ReadMessage()
	if err != nil {
		log.Printf("read message err: %v", err)
		return copy(p, EndOfTransmission), err
	}
	var msg TerminalMessage
	if err := json.Unmarshal([]byte(message), &msg); err != nil {
		log.Printf("read parse message err: %v", err)
		// return 0, nil
		return copy(p, EndOfTransmission), err
	}
	switch msg.Operation {
	case "stdin":
		return copy(p, msg.Data), nil
	case "resize":
		t.sizeChan <- remotecommand.TerminalSize{Width: msg.Cols, Height: msg.Rows}
		return 0, nil
	case "ping":
		return 0, nil
	default:
		log.Printf("unknown message type '%s'", msg.Operation)
		// return 0, nil
		return copy(p, EndOfTransmission), fmt.Errorf("unknown message type '%s'", msg.Operation)
	}
}

// Close close session
func (t *WebTerminal) Close() error {
	return t.wsConn.Close()
}

func (t *WebTerminal) Tty() bool {
	return t.tty
}
