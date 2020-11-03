package main

import (
	"bufio"
	"fmt"
	"github.com/conero/uymas/bin"
	"github.com/conero/uymas/nestling/net/base-chat/chat"
	"log"
	"net"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

//客户端
func main() {
	cli := bin.NewCLI()
	//call
	cli.RegisterFunc(func(cc *bin.CliCmd) {
		//端口号 [1-65535]; 一般不使用 0-1023
		port := cc.ArgRaw("port", "P")
		if "" == port {
			port = chat.DefChatPort
		}
		network := cc.ArgRaw("network", "N")
		if "" == network {
			network = chat.DefChatNetwork
		}
		host := cc.ArgRaw("host", "H")
		if "" == network {
			host = chat.DefChatHost
		}

		chat.Log.Info("正在连接……，网络类型 => %v, 端口号 => %v.\r\n", network, port)
		repl(network, port, host)
	}, "call", "c")
	//默认启动请求
	cli.RegisterFunc(func(cc *bin.CliCmd) {
		repl(chat.DefChatPort, chat.DefChatNetwork, chat.DefChatHost)
	})
	//默认服务
	cli.RegisterFunc(func(cc *bin.CliCmd) {
		fmt.Println("普通的 TCP 服务器，命令如下")
		fmt.Println("  call,cc 请求tcp服务器")
		fmt.Printf("     --port,-P=%v         设置端口号 \r\n", chat.DefChatPort)
		fmt.Printf("     --network,-N=%v     设置网络协议 tcp,tcp4,tcp6,unix，unixpacket \r\n", chat.DefChatNetwork)
	}, "help", "?")
	cli.Run()
}

//交互式网络请求
func repl(port, network, host string) {
	username := input("请输入您的姓名：", true)
	//@todo 开始注释密码
	//pwsd := input("请输入密码（123456）：", true)
	//if pwsd != "123456" {
	//	chat.Log.Error("您的用户密码错误！！")
	//	return
	//}

	hostAdrr := fmt.Sprintf("%v:%v", host, port)
	conn, err := net.Dial(network, hostAdrr)
	if err != nil {
		log.Fatal(err)
	}
	chat.Log.Info("服务（%v）已连接成功.\r\n", hostAdrr)
	defer conn.Close()
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
type ClientChat struct {
	conn net.Conn
}

//信息打印
func (c *ClientChat) ShowMessage() {
	for {
		addr, err := chat.RespondContent(c.conn)
		if err != nil {
			chat.Log.Error("读取服务器信息错误，信息: %v\r\n.", err.Error())
			break
		}
		switch addr.Protocol {
		case "native-message": //本地消息
			if u := addr.URL; u != nil {
				if uv := u.Query(); uv != nil {
					message := uv.Get("message")
					fromUser := uv.Get("from_user")
					fmt.Printf("\r\n ~~[Server] %v>> %v\r\n", fromUser, message)
				}
			}
		default:
			fmt.Printf("\r\n ~~[Server/Unhanlder]>> %v\r\n", addr.Content)
		}
	}
}

//构造函数
func NewChat(conn net.Conn, username string) *ClientChat {
	cc := &ClientChat{
		conn: conn,
	}
	addr := chat.Address{
		Protocol: "native-message",
		Action:   "authorization",
	}
	uV := &url.Values{}
	uV.Add("username", username)
	err := addr.SendValue(conn, uV)
	if err != nil {
		chat.Log.Error("用户请求认证错误，Error: %v.", err)
		return nil
	}

	//授权认证
	var subReplCmd = make(chan string)
	timer := chat.Timer(60 * time.Second)
	for {
		if timer() {
			chat.Log.Fatal("等待服务响应超时")
		}
		addr, err := chat.RespondContent(conn)
		if err != nil {
			chat.Log.Fatal("读取服务器数据错误，信息：%v", chat.Log)
		}
		if addr != nil {
			log.Printf("server>>\r\n%v", addr.Content)
			if addr.Action == "authorization" {
				v := addr.URL.Query()
				if "true" == v.Get("success") {
					chat.Log.Info("您成功登录服务器！")
					break
				}
			}
		} else {
			chat.Log.Error("服务器接收的数据为空！")
		}
	}

	//主线程命令集
	cCli := bin.NewCLI()

	//help
	cCli.RegisterFunc(func(cc *bin.CliCmd) {
		fmt.Println(" exit              退户程序")
		fmt.Println(" broadcast,bc      广播信息")
		fmt.Println(" chat,c <user>     与用户进行聊天")
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
			breakMk := false
			switch text {
			case "exit", ":q":
				fmt.Println("您将退出系统！")
				breakMk = true
			default:
				if "" != text {
					uV := &url.Values{}
					uV.Add("message", text)
					send := chat.Address{
						Protocol: "native-message",
						Action:   "broadcast",
					}
					er := send.SendValue(conn, uV)
					if er != nil {
						chat.Log.Error("发送广播服务错误！信息: %v.\r\n", er.Error())
					} else {
						chat.Log.Info("广播信息已发送")
					}
				}
			}
			if breakMk {
				break
			}
			fmt.Println()
			fmt.Printf("$ broadcast/%v>", username)
		}
	}, "bc", "broadcast")

	//chat
	cCli.RegisterFunc(func(cc *bin.CliCmd) {
		toUser := cc.SubCommand
		if toUser != "" {
			var input = bufio.NewScanner(os.Stdin)
			fmt.Println("驻留式命令行程序")
			fmt.Printf("$ chat/%v>", username)
			for input.Scan() {
				text := input.Text()
				text = strings.TrimSpace(text)

				breakMk := false
				switch text {
				case "exit", ":q":
					fmt.Println("您将退出系统！")
					breakMk = true
				default:
					if "" != text {
						uV := &url.Values{}
						uV.Add("message", text)
						uV.Add("username", toUser)
						send := chat.Address{
							Protocol: "native-message",
							Action:   "send-message",
						}
						er := send.SendValue(conn, uV)
						if er != nil {
							chat.Log.Error("发送信息到用户(%v)，失败！错误：%v.", toUser, er.Error())
						} else {
							chat.Log.Info("信息已发送")
						}
					}
				}
				if breakMk {
					break
				}
				fmt.Println()
				fmt.Printf("$ chat/%v>", username)
			}
		}
	}, "chat", "c")

	//显示数据
	go cc.ShowMessage()

	//命令执行
	var input = bufio.NewScanner(os.Stdin)
	fmt.Println("驻留式命令行程序")
	fmt.Printf("$ %v>", username)
	for input.Scan() {
		text := input.Text()
		text = strings.TrimSpace(text)

		switch text {
		case "exit", ":q":
			fmt.Println("您将退出系统！")
			os.Exit(0)
		default:
			var wg sync.WaitGroup
			wg.Add(1)
			go func(src chan string) {
				defer wg.Done()
				cCli.Inject("signal_chan", src)
				cCli.Run(strings.Split(text, " ")...)
				src <- "exit"
			}(subReplCmd)
			select {
			case cx := <-subReplCmd:
				if cx == "exit" {
					break
				}
			}
			wg.Wait()

			//for 循环继续
			//if isContinue {
			//	continue
			//}
		}

		fmt.Println()
		fmt.Printf("$ home/%v>", username)
	}
	return cc
}
