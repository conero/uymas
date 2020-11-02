package main

import (
	"bufio"
	"fmt"
	"github.com/conero/uymas/bin"
	"log"
	"net"
	"os"
	"strings"
)

//客户端
func main() {
	cli := bin.NewCLI()
	defPort, defNetwork := "7400", "tcp"
	//call
	cli.RegisterFunc(func(cc *bin.CliCmd) {
		//端口号 [1-65535]; 一般不使用 0-1023
		port := cc.ArgRaw("port", "P")
		if "" == port {
			port = defPort
		}
		network := cc.ArgRaw("network", "N")
		if "" == network {
			network = defNetwork
		}
		fmt.Printf("正在连接……，网络类型 => %v, 端口号 => %v.\r\n", network, port)
		repl(network, port)
	}, "call", "c")
	//默认启动请求
	cli.RegisterFunc(func(cc *bin.CliCmd) {
		repl(defPort, defNetwork)
	})
	//默认服务
	cli.RegisterFunc(func(cc *bin.CliCmd) {
		fmt.Println("普通的 TCP 服务器，命令如下")
		fmt.Println("  call,cc 请求tcp服务器")
		fmt.Println("     --port,-P=7400       设置端口号 ")
		fmt.Println("     --network,-N=tcp     设置网络协议 tcp,tcp4,tcp6,unix，unixpacket ")
	}, "help", "?")
	cli.Run()
}

//交互式网络请求
func repl(port, network string) {
	hostAdrr := fmt.Sprintf("127.0.0.1:%v", port)
	conn, err := net.Dial(network, hostAdrr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("服务（%v）已连接成功.\r\n", hostAdrr)
	defer conn.Close()

	inputReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("client-me>> ")
		input, _ := inputReader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "exit" {
			return
		}

		_, err = conn.Write([]byte(input))
		if err != nil {
			return
		}
	}
}
