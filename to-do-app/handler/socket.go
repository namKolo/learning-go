package handler

import (
	"learning-go/to-do-app/model"
	"log"
	"net/http"

	db "github.com/dancannon/gorethink"
	"github.com/gorilla/websocket"
)

type (
	SocketHandler struct {
		hub     *model.Hub
		session *db.Session
	}
)

func NewSocketHandler(hub *model.Hub, session *db.Session) *SocketHandler {
	return &SocketHandler{hub, session}
}

var upgrader = &websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}

func (sh SocketHandler) UpgradeWS(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	conn := model.NewConnection(make(chan interface{}, 256), ws)
	sh.hub.RegisterConnection(conn)
	defer func() { sh.hub.UnregisterConnection(conn) }()
	go conn.Write()
	conn.Read()
}

func (sh SocketHandler) GetAllItemsChanges() http.HandlerFunc {
	go func() {
		for {
			res, err := db.Table("items").Changes().Run(sh.session)
			if err != nil {
				log.Fatalln(err)
			}

			var response interface{}
			for res.Next(&response) {
				sh.hub.BroadcastMessage(response)
			}

			if res.Err() != nil {
				log.Println(res.Err())
			}
		}
	}()
	return sh.UpgradeWS
}
