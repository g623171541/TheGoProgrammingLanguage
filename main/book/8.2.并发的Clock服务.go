package main

import (
	"net"
	"log"
	"io"
	"time"
)

// 时钟服务器，每隔一秒钟将当前时间写到客户端

func main()  {

	ClockService();
}

func ClockService()  {
	// Listen 函数创建了一个net.Listener的对象，这个对象会监听一个网络端口上到来的连接
	listener,err := net.Listen("tcp","localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	// for 代表一直循环下去，死循环
	for  {
		// listener对象的Accept方法会直接阻塞，直到一个新的连接被创建，然后会返回一个net.Conn 对象来标示这个连接。
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		handleConn(conn)
	}
}

func handleConn(c net.Conn)  {
	defer c.Close()
	for {
		// time.Now().Format 提供了一种格式化日期和时间信息的方式，参数是一个格式化模板。
		_,err := io.WriteString(c,time.Now().Format("15:04:02\n"))
		if err != nil {
			return
		}
		time.Sleep(1 *time.Second)
	}
}