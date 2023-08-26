package main

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	userName     = "admin"
	userPassword = "admin888"

	addr = "127.0.0.1:44444"

	indexHtml = `
<!doctype html>
<html lang="zh-CN">

<head>
	<meta charset="UTF-8">
	<title>Server Sent Event</title>
</head>

<body>
	<div class="event-data"></div>
</body>

<script src="https://code.jquery.com/jquery-1.11.1.js"></script>
<script>
	// EventSource object of javascript listens the streaming events from our go server and prints the message.
	var stream = new EventSource("/stream");
	stream.addEventListener("message", function (e) {
		$('.event-data').append(e.data + "</br>")
	});
</script>

</html>
`
)

func main() {
	router := gin.Default()

	ctx := context.Background()

	// Initialize new streaming server
	stream := newEvent()
	stream.Start(ctx)

	// We are streaming current time to clients in the interval 10 seconds
	go func() {
		for {
			time.Sleep(time.Second * 10)
			now := time.Now().Format("2006-01-02 15:04:05")
			currentTime := fmt.Sprintf("The Current Time Is %v", now)

			// Send current time to clients message channel
			stream.SendMessage(currentTime)

		}
	}()

	// 全站启用基本的身份认证。
	authorized := router.Group("/", gin.BasicAuth(gin.Accounts{
		userName: userPassword,
	}))

	// Authorized client can stream the event
	// Add event-streaming headers
	authorized.GET("/stream", HeadersMiddleware(), stream.ServeHTTP(), func(c *gin.Context) {
		v, ok := c.Get("clientChan")
		if !ok {
			return
		}
		clientChan, ok := v.(ClientChan)
		if !ok {
			return
		}
		c.Stream(func(w io.Writer) bool {
			// Stream message to client from message channel
			if msg, ok := <-clientChan; ok {
				c.SSEvent("message", msg)
				return true
			}
			return false
		})
	})

	router.GET("/", func(ctx *gin.Context) {
		// 直接输出而不是使用静态 HTML 本地文件，在调用 go run 时可以不需要进入当前目录。
		ctx.Writer.Header().Add("Content-Type", "text/html; charset=utf-8")
		ctx.Writer.WriteString(indexHtml)
	})

	router.Run(addr)
}

func HeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")
		c.Writer.Header().Set("Transfer-Encoding", "chunked")
		c.Next()
	}
}
