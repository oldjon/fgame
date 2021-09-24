package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/xtaci/kcp-go"
)

type PlayerTask struct {
	UId uint64
}

func main() {
	//key := pbkdf2.Key([]byte("demo pass"), []byte("demo salt"), 1024, 32, sha1.New)
	//block, _ := kcp.NewAESBlockCrypt(key)
	if listener, err := kcp.ListenWithOptions("127.0.0.1:12345", nil, 10, 0); err == nil {
		// spin-up the client
		/*go func() {
			go client(1)
			//go client(2)
		}()*/
		go PrintState()
		for {
			s, err := listener.AcceptKCP()
			s.SetNoDelay(1, 20, 2, 1)
			//s.Driver = nil
			//关闭延迟发送，间隔，快速重传过确认包数，关闭拥塞控制
			if err != nil {
				log.Fatal(err)
			}
			go handleEcho(s)
		}
	} else {
		log.Fatal(err)
	}
	time.Sleep(time.Second * 2000)
}

// handleEcho send back everything it received
func handleEcho(conn *kcp.UDPSession) {
	buf := make([]byte, 4096)
	for {
		select {
		case <-conn.Die:
			log.Println("session closed Do something")
			return
		default:
		}
		n, err := conn.Read(buf)
		if err != nil {
			log.Println(111, err)
			return
		}

		n, err = conn.Write(buf[:n])
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func PrintState() {
	for {
		time.Sleep(10 * time.Second)
		by, _ := json.Marshal(kcp.DefaultSnmp.Copy())
		log.Println(string(by))
	}
}
