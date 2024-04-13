package routers

import (
	"chat/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {

	// Register routers.
	web.Router("/", &controllers.AppController{})
	// Indicate AppController.Join method to handle POST requests.
	web.Router("/join", &controllers.AppController{}, "post:Join")

	// WebSocket.
	web.Router("/ws", &controllers.WebSocketController{})
	web.Router("/ws/join", &controllers.WebSocketController{}, "get:Join")

}
