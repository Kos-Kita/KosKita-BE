package handler

import (
	ch "KosKita/features/chat"
	cd "KosKita/features/chat/data"
	hub "KosKita/features/chat/service"
	cu "KosKita/features/user"
	"KosKita/utils/responses"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type ChatHandler struct {
	chatService ch.ChatServiceInterface
	hub         *hub.Hub
	cu cu.UserServiceInterface
}

func New(cs ch.ChatServiceInterface, h *hub.Hub, cu cu.UserServiceInterface) *ChatHandler {
	return &ChatHandler{
		chatService: cs,
		hub:         h,
		cu: cu,
	}
}

func (ch *ChatHandler) CreateRoom(c echo.Context) error {
	var req CreateRoomReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	
	roomID := generateRoomID()
	ch.hub.Rooms[roomID] = &hub.Room{
		ID:      roomID,
		Clients: make(map[string]*hub.Client),
	}

	err := ch.chatService.CreateRoom(roomID, req.ReceiverID, req.SenderID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, RoomRes{ID: roomID})
}

func (ch *ChatHandler) JoinRoom(c echo.Context) error {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(c.Response().Writer, c.Request(), nil)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	roomID := c.Param("roomId")
	clientID := c.QueryParam("userId")

	cl := &hub.Client{
		Conn:    conn,
		Message: make(chan *cd.Chat, 10),
		ID:      clientID,
		RoomID:  roomID,
	}

	m := &cd.Chat{
		Message: "",
		RoomID:  roomID,
	}

	ch.hub.Register <- cl
	ch.hub.Broadcast <- m

	go cl.WriteMessage()
	cl.ReadMessage(ch.hub, ch.chatService, ch.cu)

	return nil
}

func (ch *ChatHandler) GetMessages(c echo.Context) error {
	roomID := c.Param("roomId")

	chats, errGet := ch.chatService.GetMessage(roomID)
	if errGet != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": errGet.Error()})
	}

	chatResult := CoreToGetChats(chats)

	return c.JSON(http.StatusOK, responses.WebResponse("success get message.", chatResult))

}

func (ch *ChatHandler) GetRooms(c echo.Context) error {
	rooms := make([]RoomRes, 0)

	for _, r := range ch.hub.Rooms {
		rooms = append(rooms, RoomRes{
			ID: r.ID,
		})
	}

	return c.JSON(http.StatusOK, rooms)
}
