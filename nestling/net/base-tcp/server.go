package main

import (
	"fmt"
	"github.com/conero/uymas/bin"
	"log"
	"net"
)

//tcp 服务器
func main() {
	cli := bin.NewCLI()
	cli.RegisterFunc(baseServer, "server", "sv")
	cli.RegisterEmpty(baseServer)
	cli.RegisterFunc(func(cc *bin.CliCmd) {
		fmt.Println("普通的 TCP 服务器，命令如下")
		fmt.Println("  server,sv 启动tcp服务器")
		fmt.Println("     --port,-P=7400       设置端口号 ")
		fmt.Println("     --network,-N=tcp     设置网络协议 tcp,tcp4,tcp6,unix，unixpacket ")
	}, "help", "?")
	cli.Run()
}

//服务器
func baseServer(cc *bin.CliCmd) {
	//端口号 [1-65535]; 一般不使用 0-1023
	port := cc.ArgRaw("port", "P")
	if "" == port {
		port = "7400"
	}
	network := cc.ArgRaw("network", "N")
	if "" == network {
		network = "tcp"
	}

	log.Printf("已启动服务器：网络类型 => %v, 端口号 => %v.\r\n", network, port)
	listener, err := net.Listen(network, fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	for {
		// Wait for a connection.
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		/*go func(c net.Conn) {
			// Echo all incoming data.
			io.Copy(c, c)
			// Shut down the connection.
			c.Close()
		}(conn)*/
		go handler(conn)
	}
}

//处理连接
func handler(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		n, err := conn.Read(buf)
		if err != nil {
			log.Println("read err:", err)
			return
		}
		content := string(buf[0:n])
		log.Printf("client>> %v\r\n", content)
	}
}
