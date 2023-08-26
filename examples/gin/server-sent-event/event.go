package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	_ Event = (*event)(nil)
)

const (
	contextClientChanName = "clientChan"
)

type (
	// ClientChan 将新事件消息广播到所有已注册的客户端连接通道。
	ClientChan chan string

	// event 事件管理器。
	// 保留当前附加的客户机列表，并向这些客户机广播事件。
	event struct {
		// Message 要发送的消息都提交到这个通道。
		Message chan string

		// NewClients 新创建的连接。
		NewClients chan chan string

		// ClosedClients 关闭的连接。
		ClosedClients chan chan string

		// TotalClients 所有客户端。
		TotalClients map[chan string]bool
	}

	Event interface {
		Start(ctx context.Context)
		SendMessage(message string)
		ServeHTTP() gin.HandlerFunc
	}
)

func (e *event) ServeHTTP() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 初始化客户端连接通道。
		clientChan := make(ClientChan)

		// 发送新创建的连接消息。
		e.NewClients <- clientChan

		defer func() {
			// 发送关闭的连接消息。
			e.ClosedClients <- clientChan
		}()

		c.Set(contextClientChanName, clientChan)

		c.Next()
	}
}

// listen 监听所有消息。
// 客户端的传入、添加或删除客户端、广播消息。
func (e *event) listen(ctx context.Context) {
LOOP:
	for {
		select {
		// 添加活跃连接。
		case client := <-e.NewClients:
			e.TotalClients[client] = true
			log.Printf("Client added. %d registered clients", len(e.TotalClients))

		// 移除客户端。
		case client := <-e.ClosedClients:
			delete(e.TotalClients, client)
			close(client)
			log.Printf("Removed client. %d registered clients", len(e.TotalClients))

		// 向所有客户端广播。
		case eventMsg := <-e.Message:
			for clientMessageChan := range e.TotalClients {
				clientMessageChan <- eventMsg
			}
		case <-ctx.Done():
			break LOOP
		}

	}
}

func (e *event) Start(ctx context.Context) {
	go e.listen(ctx)
}

func (e *event) SendMessage(message string) {
	e.Message <- message
}

// newEvent 一个新的事件管理器。
func newEvent() Event {
	e := &event{
		Message:       make(chan string),
		NewClients:    make(chan chan string),
		ClosedClients: make(chan chan string),
		TotalClients:  make(map[chan string]bool),
	}

	return e
}
