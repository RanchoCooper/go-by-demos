package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// WebSocket 升级器配置
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// 允许所有来源（根据需求可修改）
		return true
	},
}

// 模拟的有效 Token
var validTokens = map[string]bool{
	"test_token": true,
}

// 定义一个简单的消息结构
type Message struct {
	Type    string `json:"type"`    // 消息类型
	Content string `json:"content"` // 消息内容
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)

	fmt.Println("WebSocket server is running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}

// 处理 WebSocket 连接
func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// 获取认证 Token
	token := r.URL.Query().Get("token")
	if !isValidToken(token) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// 升级为 WebSocket 连接
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("WebSocket upgrade failed:", err)
		return
	}
	defer conn.Close()

	// 认证成功后，开始处理连接
	fmt.Println("Client connected with token:", token)

	// 启动心跳检测
	go pingClient(conn)

	// 读取客户端消息
	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println("ReadJSON failed:", err)
			break
		}

		fmt.Printf("Received message: %+v\n", msg)

		// 简单回声逻辑：将消息返回给客户端
		if err := conn.WriteJSON(Message{
			Type:    "response",
			Content: "Echo: " + msg.Content,
		}); err != nil {
			fmt.Println("WriteJSON failed:", err)
			break
		}
	}
}

// 校验 Token
func isValidToken(token string) bool {
	_, valid := validTokens[token]
	return valid
}

// 心跳检测，定期发送 Ping 消息
func pingClient(conn *websocket.Conn) {
	for {
		err := conn.WriteMessage(websocket.PingMessage, nil)
		if err != nil {
			fmt.Println("Ping failed:", err)
			break
		}
		time.Sleep(10 * time.Second) // 每 10 秒发送一次心跳
	}
}
