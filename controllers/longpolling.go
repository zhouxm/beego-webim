package controllers

import (
	"chat/models"

	"github.com/beego/beego/v2/server/web"
)

// LongPollingController handles long polling requests.
type LongPollingController struct {
	web.Controller
}

// Join method handles GET requests for LongPollingController.
func (c *LongPollingController) Join() {
	// Safe check.
	uname := c.GetString("uname")
	if len(uname) == 0 {
		c.Redirect("/", 302)
		return
	}

	// Join chat room.
	Join(uname, nil)

	c.TplName = "longpolling.html"
	c.Data["IsLongPolling"] = true
	c.Data["UserName"] = uname
}

// Post method handles receive messages requests for LongPollingController.
func (c *LongPollingController) Post() {
	c.TplName = "longpolling.html"

	uname := c.GetString("uname")
	content := c.GetString("content")
	if len(uname) == 0 || len(content) == 0 {
		return
	}

	publish <- newEvent(models.EVENT_MESSAGE, uname, content)
}

// Fetch method handles fetch archives requests for LongPollingController.
func (c *LongPollingController) Fetch() {
	lastReceived, err := c.GetInt("lastReceived")
	if err != nil {
		return
	}

	events := models.GetEvents(int(lastReceived))
	if len(events) > 0 {
		c.Data["json"] = events
		c.ServeJSON()
		return
	}

	// Wait for new message(s).
	ch := make(chan bool)
	waitingList.PushBack(ch)
	<-ch

	c.Data["json"] = models.GetEvents(int(lastReceived))
	c.ServeJSON()
}
