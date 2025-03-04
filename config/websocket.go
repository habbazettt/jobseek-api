package config

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WebSocketConnections menyimpan semua koneksi WebSocket
var WebSocketConnections = make(map[uint]*websocket.Conn)

// HandleWebSocket menangani koneksi WebSocket
func HandleWebSocket(w http.ResponseWriter, r *http.Request, userID uint) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket Upgrade Error:", err)
		return
	}
	WebSocketConnections[userID] = conn
	log.Println("WebSocket connected for user:", userID)
}

// SendNotificationToUser mengirimkan notifikasi real-time ke user
func SendNotificationToUser(userID uint, message string) {
	conn, exists := WebSocketConnections[userID]
	if exists {
		conn.WriteJSON(map[string]string{"notification": message})
	}
}
