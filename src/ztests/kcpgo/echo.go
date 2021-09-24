package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/xtaci/kcp-go"
)

/*
func main() {
	//key := pbkdf2.Key([]byte("demo pass"), []byte("demo salt"), 1024, 32, sha1.New)
	//block, _ := kcp.NewAESBlockCrypt(key)
	if listener, err := kcp.ListenWithOptions("127.0.0.1:12345", nil, 10, 0); err == nil {
		// spin-up the client
		go func() {
			go client(1)
			//go client(2)
		}()

		for {
			s, err := listener.AcceptKCP()
			s.SetNoDelay(1, 20, 2, 1)
			//关闭延迟发送，间隔，快速重传过确认包数，关闭拥塞控制
			if err != nil {
				log.Fatal(err)
			}
			go handleEcho(s)
		}
	} else {
		log.Fatal(err)
	}
}

// handleEcho send back everything it received
func handleEcho(conn *kcp.UDPSession) {
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Println(err)
			return
		}

		n, err = conn.Write(buf[:n])
		if err != nil {
			log.Println(err)
			return
		}
	}
}*/

func main() {
	client(1)
}

func client(id int) {
	//key := pbkdf2.Key([]byte("demo pass"), []byte("demo salt"), 1024, 32, sha1.New)
	//block, _ := kcp.NewAESBlockCrypt(key)

	// wait for server to become ready
	time.Sleep(time.Second)

	// dial to the echo server
	if sess, err := kcp.DialWithOptions("127.0.0.1:12345", nil, 10, 0); err == nil {
		//起多个协程执行
		sg := &sync.WaitGroup{}
		sess.SetNoDelay(1, 20, 2, 1)
		sess.SetWriteDelay(false)
		sg.Add(100)
		count = make([]int, 100)
		countreadbys = make([]int, 100)
		countwritebys = make([]int, 100)
		go printcount()
		for i := 0; i < 100; i++ {
			go clientRecv(id, sess, i)
		}
		for i := 0; i < 100; i++ {
			go clientSend(id, sess, sg, i)
		}
		sg.Wait()
		/*for {
			data := time.Now().String() + " " + strconv.Itoa(id)
			buf := make([]byte, len(data))
			log.Println("sent:", data)
			if _, err := sess.Write([]byte(data)); err == nil {
				// read back the data
				if _, err := io.ReadFull(sess, buf); err == nil {
					log.Println("recv:", string(buf))
				} else {
					log.Fatal(err)
				}
			} else {
				log.Fatal(err)
			}
			time.Sleep(time.Second)
		}*/
	} else {
		log.Fatal(err)
	}
	time.Sleep(time.Second * 20)
	by, _ := json.Marshal(kcp.DefaultSnmp.Copy())
	log.Println(string(by))
	var a int
	for _, u := range count {
		a += u
	}
	log.Println(a)
}

var count []int
var countreadbys []int
var countwritebys []int

func clientSend(id int, sess *kcp.UDPSession, sg *sync.WaitGroup, counti int) {
	data := time.Now().String() + " " + strconv.Itoa(id)
	//buf := make([]byte, len(data))
	for i := 0; i < 1; i++ {
		//sess.Write([]byte(data))
		/*n, _ := sess.Write([]byte(data))
		if n != 54 {
			log.Fatal(n)
		}*/
		//countwritebys[counti] += n
		if n, err := sess.Write([]byte(data)); err == nil {
			countwritebys[counti] += n
		} else {
			log.Fatal(err)
		}
		fmt.Println(counti, " send ", i)
	}
	sg.Done()
}

func clientRecv(id int, sess *kcp.UDPSession, counti int) {
	data := time.Now().String() + " " + strconv.Itoa(id)
	buf := make([]byte, len(data))
	for {
		if n, err := io.ReadFull(sess, buf); err == nil {
			log.Println(string(buf))
			countreadbys[counti] += n
		} else {
			log.Fatal(err)
		}
		count[counti]++
	}
}

func printcount() {
	time.Sleep(time.Second * 15)
	for {
		var a int
		for _, u := range count {
			a += u
		}
		log.Println("接收次数", a)
		a = 0
		for _, u := range countwritebys {
			a += u
		}
		log.Println("发送字节", a)
		a = 0
		for _, u := range countreadbys {
			a += u
		}
		log.Println("接收字节", a)
		time.Sleep(time.Second * 3)
	}

}

//TODO 解决应用层粘包问题
//TODO 增加数据压缩
//TODO 解决断线后将断线信息上报应用层的逻辑设计
