package service

import (
	cc "KosKita/features/chat"
	cd "KosKita/features/chat/data"
	"log"
	"strconv"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn    *websocket.Conn
	Message chan *cd.Chat
	ID      string `json:"id"`
	RoomID  string `json:"roomId"`
}

func (c *Client) WriteMessage() {
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

func (c *Client) ReadMessage(hub *Hub, chatService cc.ChatServiceInterface) {
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

		msg := &cd.Chat{
			Message:  string(m),
			RoomID:   c.RoomID,
		}

		coreMsg := cc.Core{
			ID:        msg.ID,
			Message:   msg.Message,
			RoomID:    msg.RoomID,
			UserID:    msg.UserID,
			CreatedAt: msg.CreatedAt,
			UpdatedAt: msg.UpdatedAt,
			User:      msg.User.ModelToCore(),
		}

		// Ubah c.ID menjadi integer
		userID, err := strconv.Atoi(c.ID)
		if err != nil {
			log.Printf("Error converting ID to integer: %v", err)
			continue
		}

		// Simpan pesan ke database
		_, err = chatService.CreateChat(userID, coreMsg)
		if err != nil {
			log.Printf("Error saving message: %v", err)
			continue
		}

		hub.Broadcast <- msg
	}
}
