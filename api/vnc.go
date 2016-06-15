package api

import (
	//"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	//"net/http"
	//"errors"
	"github.com/googollee/go-socket.io"
	"github.com/megamsys/libgo/cmd"
	"github.com/megamsys/vertice/govnc"

)

const (
	HOST     = "host"
	PORT     = "port"
	PASSWORD = "password"
)

func vncHandler(so socketio.Socket) {

	so.On("vncInit", func(config map[string]interface{}) {
		vnc := &govnc.Vnc{}
		err := vnc.FillStruct(config)
		if err != nil {
			so.Emit("error", fmt.Sprintf("%v", err))
			return
		}
		fmt.Println("0000000000000000000000000000")
		vnc.Connect(so)
	})

	so.On("disconnection", func() {
		log.Debugf(cmd.Colorfy("  > [socket] ", "blue", "", "bold") + fmt.Sprintf("Disconneted client : %s", so.Id()))
	})
}


/*var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func vnc(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("-------------------------------------------")
  conn, err := upgrader.Upgrade(w, r, nil)
  if err != nil {
    log.Errorf("Error in socket connection")
    return err
  }
  s := conn.Subprotocol()
	fmt.Println(s)
  messageType, p, err := conn.ReadMessage()

  if err != nil {
    return err
  }

  var entry govnc.VncHost

  _ = json.Unmarshal(p, &entry)
	govnc.Connect(&entry)
  l, _ := govnc.Connect(&entry)

  go func() {
    if _, _, err := conn.NextReader(); err != nil {
      conn.Close()
      l.Close()
      log.Debugf(cmd.Colorfy("  > [nsqd] unsub   ", "blue", "", "bold") + fmt.Sprintf("Unsubscribing from the Queue"))
    }
  }()

  for logbox := range l.B {
    logData, _ := json.Marshal(logbox)
    conn.WriteMessage(messageType, logData)
  }
conn.WriteMessage(messageType, []byte("hai"))
  return nil
}*/
