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

	// Long polling.
	web.Router("/lp", &controllers.LongPollingController{}, "get:Join")
	web.Router("/lp/post", &controllers.LongPollingController{})
	web.Router("/lp/fetch", &controllers.LongPollingController{}, "get:Fetch")

	// WebSocket.
	web.Router("/ws", &controllers.WebSocketController{})
	web.Router("/ws/join", &controllers.WebSocketController{}, "get:Join")

}
