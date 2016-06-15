package govnc

import (
	"fmt"
	//"net"
	//"time"
	//"log"
	"errors"
	"github.com/googollee/go-socket.io"
	"reflect"
	//govnc "github.com/kward/go-vnc"
	//gvc "github.com/mitchellh/go-vnc"
	//"golang.org/x/net/context"
	gcc "github.com/megamsys/vertice/libvncclient"
)

type Vnc struct {
	Host     string
	Port     string
	Password string
}

func (v *Vnc) Connect(so socketio.Socket) {
	gc := gcc.RfbGetClient(8, 3, 4)
	gc.SetServer("136.243.49.217", 6663)
	width, height := 640, 480
	bpp := 4
	gc.SetFrameBuffer(width, height, bpp)
	bol := gc.InitClient(0, nil)
	fmt.Println("----------------------")
	fmt.Println(bol)

	for {
		retWait := gc.WaitForMessage(500)
		if retWait < 0 {
			fmt.Printf("Error waiting for message\n")
			return
		} else {
			if ok := gc.HandleRFBServerMessage(); ok != true {
				fmt.Printf("Error handling server message\n %v", ok)
				return
			}
		}
	}

}

/*func (v *Vnc) Connect(so socketio.Socket) {
	fmt.Println("uhfbberijberignerjgknregj")
	nc, err := net.Dial("tcp", "136.243.49.217:6663")
	if err != nil {
		so.Emit("error", fmt.Sprintf("Error connecting to VNC host : %v", err))
		return
	}
	fmt.Println(nc)
	// Negotiate connection with the server.
	//vcc := &gvc.ClientConfig{Exclusive: false}
	vcc := &gvc.ClientConfig{
		Auth: []gvc.ClientAuth{
		},
		ServerMessages: []gvc.ServerMessage{
			&gvc.FramebufferUpdateMessage{},
			&gvc.SetColorMapEntriesMessage{},
			&gvc.ServerCutTextMessage{},
		},
	}
	cc, err1 := gvc.Client(nc, vcc)
	if err1 != nil {
		so.Emit("error", fmt.Sprintf("Error negotiating connection to VNC host : %v", err))
		return
	}
	fmt.Println(cc)
	fmt.Println(cc.FrameBufferWidth)
	fmt.Println(cc.FrameBufferHeight)
	fmt.Println(cc.DesktopName)

	go func() {
		w, h := cc.FrameBufferWidth, cc.FrameBufferHeight
		for {
			if err := cc.FramebufferUpdateRequest(true, 0, 0, w, h); err != nil {
				so.Emit("error", fmt.Sprintf("Error vnc update : %v", err))
				return
			}
			time.Sleep(1 * time.Second)
		}
	}()
	for {
		msg := vcc.ServerMessageCh

		fmt.Println("=============================")
		fmt.Println(reflect.TypeOf(msg))
		fmt.Println(msg)
	}
}*/

/*func (v *Vnc) Connect(so socketio.Socket) {
	// Establish TCP connection to VNC server.
	nc, err := net.Dial("tcp", "136.243.49.217:6663")
	if err != nil {
		so.Emit("error", fmt.Sprintf("Error connecting to VNC host : %v", err))
		return
	}

	// Negotiate connection with the server.
	vcc := govnc.NewClientConfig(v.Password)
	vc, err := govnc.Connect(context.Background(), nc, vcc)
	if err != nil {
		so.Emit("error", fmt.Sprintf("Error negotiating connection to VNC host : %v", err))
		return
	}
  fmt.Println(vc.FramebufferWidth())
	fmt.Println(vc.FramebufferHeight())
	// Periodically request framebuffer updates.
	go func() {
		w, h := vc.FramebufferWidth(), vc.FramebufferHeight()
		for {
			if err := vc.FramebufferUpdateRequest(govnc.RFBTrue, 0, 0, w, h); err != nil {
				so.Emit("error", fmt.Sprintf("Error vnc update : %v", err))
				return
			}
			time.Sleep(1 * time.Second)
		}
	}()

	// Listen and handle server messages.
	go vc.ListenAndHandle()
  fmt.Println("============================")
	// Process messages coming in on the ServerMessage channel.

	for {
		msg := vcc.ServerMessageCh
		fmt.Println("+++++++++++++++++++++++++++++")
		fmt.Println(reflect.TypeOf(msg))
		fmt.Println(msg)
		break
	}

}*/

func SetField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		invalidTypeError := errors.New("Provided value type didn't match obj field type")
		return invalidTypeError
	}

	structFieldValue.Set(val)
	return nil
}

func (s *Vnc) FillStruct(m map[string]interface{}) error {
	for k, v := range m {
		err := SetField(s, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}
