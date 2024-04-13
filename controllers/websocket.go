package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/beego/beego/v2/core/logs"

	"chat/models"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     checkOrigin,
	}
)

// 验证跨域问题
func checkOrigin(r *http.Request) bool {
	return r.Header.Get("Origin") == "http://"+r.Host
}

// WebSocketController handles WebSocket requests.
type WebSocketController struct {
	baseController
}

// Get method handles GET requests for WebSocketController.
func (c *WebSocketController) Get() {
	// Safe check.
	uname := c.GetString("uname")
	logs.Info("WebSocketController.Get() called uname: %s", uname)
	if len(uname) == 0 {
		c.Redirect("/", 302)
		return
	}

	c.TplName = "websocket.html"
	c.Data["IsWebSocket"] = true
	c.Data["UserName"] = uname
}

// Play method handles WebSocket requests for WebSocketController.
func (c *WebSocketController) Join() {
	uname := c.GetString("uname")
	logs.Info("WebSocketController.Join() called uname: [%s]", uname)
	if len(uname) == 0 {
		c.Redirect("/", 302)
		return
	}

	// Upgrade from http request to WebSocket.
	ws, err := upgrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(c.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		logs.Error("Cannot setup WebSocket connection:", err)
		return
	}

	// Join chat room.
	Join(uname, ws)
	defer Leave(uname)

	// Message receive loop.
	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			return
		}
		publish <- newEvent(models.EVENT_MESSAGE, uname, string(p))
	}
}

// broadcastWebSocket broadcasts messages to WebSocket users.
func broadcastWebSocket(event models.Event) {
	data, err := json.Marshal(event)
	if err != nil {
		logs.Error("Fail to marshal event:", err)
		return
	}
	logs.Info("Broadcasting WebSocketController	data:[%s]", data)
	for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
		// Immediately send event to WebSocket users.
		ws := sub.Value.(Subscriber).Conn
		if ws != nil {
			if ws.WriteMessage(websocket.TextMessage, data) != nil {
				// User disconnected.
				unsubscribe <- sub.Value.(Subscriber).Name
			}
		}
	}
}
