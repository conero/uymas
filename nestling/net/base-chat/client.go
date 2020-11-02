package main

import (
	"bufio"
	"fmt"
	"github.com/conero/uymas/bin"
	chat2 "github.com/conero/uymas/nestling/net/base-chat/chat"
	"log"
	"net"
	"net/url"
	"os"
	"strings"
	"time"
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
	username := input("请输入您的姓名：", true)
	pwsd := input("请输入密码（123456）：", true)
	if pwsd != "123456" {
		log.Fatal("您的用户密码错误！！")
		return
	}

	NewChat(conn, username)
	/*inputReader := bufio.NewReader(os.Stdin)
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
	}*/
}

func input(tip string, required bool) string {
	fmt.Print(tip)
	inputReader := bufio.NewReader(os.Stdin)
	var input string
	for {
		input, _ = inputReader.ReadString('\n')
		input = strings.TrimSpace(input)
		if required && "" == input {
			fmt.Println("  数据不可为空！")
			fmt.Print(tip)
			continue
		}
		break
	}
	return input
}

//聊天器
type Chat struct {
	conn net.Conn
}

//构造函数
func NewChat(conn net.Conn, username string) *Chat {
	chat := &Chat{
		conn: conn,
	}
	_, err := conn.Write([]byte(fmt.Sprintf("native-message://authorization?username=%v", username)))
	if err != nil {
		log.Fatalf("用户请求认证错误，Error: %v.", err)
	}

	timer := chat2.Timer(60 * time.Second)
	for {
		if timer() {
			log.Fatal("等待服务响应超时")
		}
		addr := chat2.RespondContent(conn)
		log.Printf("server>>\r\n%v", addr.Content)
		if addr.Action == "authorization" {
			v := addr.URL.Query()
			if "true" == v.Get("success") {
				log.Println("您成功登录服务器！")
				break
			}
		}
	}

	cCli := bin.NewCLI()

	//help
	cCli.RegisterFunc(func(cc *bin.CliCmd) {
		fmt.Println(" exit              退户程序")
		fmt.Println(" broadcast,bc      广播信息")
		fmt.Println(" user,us <user>    与用户进行聊天")
	}, "help", "?")

	//bc
	cCli.RegisterFunc(func(cc *bin.CliCmd) {
		//命令执行
		var input = bufio.NewScanner(os.Stdin)
		fmt.Println("驻留式命令行程序")
		fmt.Printf("$ broadcast/%v>", username)
		for input.Scan() {
			text := input.Text()
			text = strings.TrimSpace(text)
			switch text {
			case "exit":
				fmt.Println("您将退出系统！")
				break
			default:
				if "" != text {
					uV := &url.Values{}
					uV.Add("message", text)
					send := chat2.Address{
						Protocol: "native-message",
					}
					_, er := conn.Write([]byte(send.Send("broadcast", uV)))
					if er != nil {
						log.Printf("发送广播服务错误！信息: %v.\r\n", er.Error())
					}
				}
			}
			fmt.Println()
			fmt.Printf("$ broadcast/%v>", username)
		}
	}, "bc", "broadcast")

	//user
	cCli.RegisterFunc(func(cc *bin.CliCmd) {
		toUser := cc.SubCommand
		if toUser != "" {
			var input = bufio.NewScanner(os.Stdin)
			fmt.Println("驻留式命令行程序")
			fmt.Printf("$ %v>", username)
			for input.Scan() {
				text := input.Text()
				text = strings.TrimSpace(text)

				switch text {
				case "exit":
					fmt.Println("您将退出系统！")
					break
				default:
					if "" != text {
						uV := &url.Values{}
						uV.Add("message", text)
						uV.Add("username", toUser)
						send := chat2.Address{
							Protocol: "native-message",
						}
						_, er := conn.Write([]byte(send.Send("send-message", uV)))
						if er != nil {
							log.Printf("发送信息到用户(%v)，失败！错误：%v.", toUser, er.Error())
						}
					}
				}
				fmt.Println()
				fmt.Printf("$ %v>", username)
			}
		}
	}, "user", "us")

	//命令执行
	var input = bufio.NewScanner(os.Stdin)
	fmt.Println("驻留式命令行程序")
	fmt.Printf("$ %v>", username)
	for input.Scan() {
		text := input.Text()
		text = strings.TrimSpace(text)

		switch text {
		case "exit":
			fmt.Println("您将退出系统！")
			os.Exit(0)
		default:
			cCli.Run(strings.Split(text, " ")...)
		}

		fmt.Println()
		fmt.Printf("$ %v>", username)
	}
	return chat
}
