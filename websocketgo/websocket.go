package websocketgo

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/habbazettt/jobseek-go/models"
	"github.com/habbazettt/jobseek-go/services"
)

type Client struct {
	Conn   *websocket.Conn
	UserID uint
	Send   chan []byte
}

type Hub struct {
	clients    map[uint]*Client
	register   chan *Client
	unregister chan *Client
	broadcast  chan []byte
	mu         sync.Mutex
}

var HubInstance = &Hub{
	clients:    make(map[uint]*Client),
	register:   make(chan *Client),
	unregister: make(chan *Client),
	broadcast:  make(chan []byte),
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			h.clients[client.UserID] = client
			h.mu.Unlock()
			log.Printf("🟢 User %d connected to WebSocket\n", client.UserID)

		case client := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[client.UserID]; ok {
				close(client.Send)
				delete(h.clients, client.UserID)
				log.Printf("❌ User %d disconnected\n", client.UserID)
			}
			h.mu.Unlock()

		case message := <-h.broadcast:
			h.mu.Lock()
			for _, client := range h.clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.clients, client.UserID)
					log.Printf("⚠️ User %d forcibly disconnected\n", client.UserID)
				}
			}
			h.mu.Unlock()
		}
	}
}

func (c *Client) ReadMessages(chatService services.ChatService) {
	defer func() {
		HubInstance.unregister <- c
		c.Conn.Close()
	}()

	for {
		var msg models.ChatMessage
		if err := c.Conn.ReadJSON(&msg); err != nil {
			log.Printf("❌ [WebSocket] Error reading message from User %d: %v", c.UserID, err)
			break
		}

		if msg.SenderID != c.UserID {
			log.Printf("⚠️ [Security] Invalid sender ID! User %d mencoba mengirim sebagai User %d", c.UserID, msg.SenderID)
			continue
		}

		if _, err := chatService.SendMessage(msg.SenderID, msg.ReceiverID, msg.Message); err != nil {
			log.Printf("❌ [Database] Error saving message: %v", err)
			continue
		}

		messageData, _ := json.Marshal(msg)

		HubInstance.mu.Lock()
		if receiverClient, exists := HubInstance.clients[msg.ReceiverID]; exists {
			receiverClient.Send <- messageData
		} else {
			log.Printf("🔴 [Offline] Receiver %d is not connected, message stored only", msg.ReceiverID)
		}
		HubInstance.mu.Unlock()
	}
}

func (c *Client) WriteMessages() {
	defer func() {
		HubInstance.unregister <- c
		c.Conn.Close()
	}()

	for message := range c.Send {
		if err := c.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Printf("❌ [WebSocket] Error sending message to User %d: %v", c.UserID, err)
			break
		}
	}
}

func HandleConnections(conn *websocket.Conn, userID uint, chatService services.ChatService) {
	log.Printf("🔵 [WebSocket] User %d is connecting...", userID)

	client := &Client{Conn: conn, UserID: userID, Send: make(chan []byte, 256)}

	HubInstance.register <- client

	log.Printf("✅ [WebSocket] User %d successfully connected!", userID)

	go client.WriteMessages()
	client.ReadMessages(chatService)
}
