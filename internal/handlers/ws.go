package handlers

import (
	"net/http"
	"log"
	"bytes"
	// "strconv"

	"github.com/go-chi/chi/v5"
	"github.com/lesismal/nbio/nbhttp/websocket"
	"github.com/orsanawwad/htmxdemo/internal/templates"
)

type WSService interface {}

type WSHandler struct {}

func NewWSHandler() *WSHandler {
	return &WSHandler{}
}

func onWebsocket(hrw http.ResponseWriter, req *http.Request) {
	upgrader := websocket.NewUpgrader()
	upgrader.OnMessage(func(conn *websocket.Conn, messageType websocket.MessageType, data []byte) {
		// echo
		conn.WriteMessage(messageType, data)
	})
	upgrader.OnOpen(func(conn *websocket.Conn) {
		// conn.WriteMessage(websocket.TextMessage, []byte("Hello"))
		// conn.WriteMessage(websocket.TextMessage, )
		buf := &bytes.Buffer{}
		templates.Notifications().Render(req.Context(), buf)
		conn.WriteMessage(websocket.TextMessage, buf.Bytes())
		// templates.Notifications().Render(conn.Engine.Context .Context(), w)
	})

	upgrader.OnMessage(func(conn *websocket.Conn, messageType websocket.MessageType, data []byte) {
		// c.WriteMessage(messageType, data)
		buf := &bytes.Buffer{}
		templates.Notifications().Render(req.Context(), buf)
		conn.WriteMessage(websocket.TextMessage, buf.Bytes())
	})

	conn, err := upgrader.Upgrade(hrw, req, nil)
	if err != nil {
		panic(err)
	}
	conn.OnClose(func(c *websocket.Conn, err error) {
		log.Println("OnClose:", c.RemoteAddr().String(), err)
	})
}

func (h *WSHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/connect", onWebsocket)

	return r
}
