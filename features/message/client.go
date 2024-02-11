package message

import (
	"KosKita/features/user/data"
	"log"

	"github.com/gorilla/websocket"
	"gorm.io/gorm"
)

type Client struct {
	Conn     *websocket.Conn
	Message  chan *Message
	ID       string `json:"id"`
	RoomID   string `json:"roomId"`
	Username string `json:"username"`
}

type Message struct {
	gorm.Model
	Message  string `json:"content"`
	RoomID   string `json:"roomId"`
	Username string `json:"username"`
	UserID   uint
	User     data.User
}

var db *gorm.DB

func (c *Client) writeMessage() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		message, ok := <-c.Message
		if !ok {
			return
		}

		c.Conn.WriteJSON(message)
	}
}

func (c *Client) readMessage(hub *Hub) {
	defer func() {
		hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, m, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		msg := &Message{
			Message:  string(m),
			RoomID:   c.RoomID,
			Username: c.Username,
		}

		tx := db.Create(&msg)
		if tx.Error != nil {
			log.Printf("error: %v", err)
			break
		}

		hub.Broadcast <- msg
	}
}
