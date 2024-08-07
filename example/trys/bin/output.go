package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	//output1()
	//example1()
	wait(100)
}

// 覆盖原来得数据
// 保持输出
func example1() {
	for i := 0; i != 10; i = i + 1 {
		fmt.Fprintf(os.Stdout, "result is %d\r", i)
		time.Sleep(time.Second * 1)
	}
	fmt.Println("over")
}

// 尝试覆盖原来文件
// 输出测试
func output1() {
	fmt.Fprint(os.Stdout, time.Now(), 1, "\r")

	time.Sleep(2 * time.Second)
	fmt.Fprint(os.Stdout, "Joshua Conero", 2, "\r")

}

func wait(sec int) {
	for i := 0; i < sec; i = i + 1 {
		sign := "-"
		if i%2 == 0 {
			sign = "-"
		} else {
			sign = "/"
		}
		n := (sec - i) * 1000
		time.Sleep(time.Second)
		fmt.Printf("[%v] %v\r", sign, n)
	}
	fmt.Print("done")
}
