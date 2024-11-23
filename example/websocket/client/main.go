package main

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

// 消息结构
type Message struct {
	Type    string `json:"type"`    // 消息类型
	Content string `json:"content"` // 消息内容
}

func main() {
	// 服务端 WebSocket 地址
	// serverURL := "ws://localhost:8080/ws"
	token := "test_token" // 替换为正确的认证 Token

	// 构造 WebSocket URL
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws", RawQuery: "token=" + token}

	// 创建 WebSocket 连接
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatalf("Failed to connect to WebSocket server: %v", err)
	}
	defer conn.Close()
	fmt.Println("Connected to server")

	// 启动认证协程
	// authChan := make(chan error)
	go func() {
		time.Sleep(2 * time.Second)
		authenticate(conn)
	}()

	time.Sleep(4 * time.Second)
	// 读取消息
	for {
		var msg Message
		if err := conn.ReadJSON(&msg); err != nil {
			fmt.Printf("Failed to read message: %v\n", err)
			break
		}
		fmt.Printf("Received message: %+v\n", msg)
	}

	// 发送消息
	for {
		time.Sleep(5 * time.Second)
		err := conn.WriteJSON(Message{Type: "chat", Content: "Hello from client"})
		if err != nil {
			fmt.Printf("Failed to send message: %v\n", err)
			break
		}
		fmt.Println("Message sent")
	}
}

// authenticate 用于认证连接
func authenticate(conn *websocket.Conn) error {
	// 可以发送认证消息，视服务端实现决定是否需要额外认证逻辑
	// 示例中服务端仅使用 token，所以这里直接返回 nil 表示成功
	fmt.Println("Authenticating connection...")
	time.Sleep(1 * time.Second) // 模拟认证过程
	return nil
}
