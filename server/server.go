package server

import (
	"awesomeProject/cfg"
	"fmt"
)

func test() {
	fmt.Println(cfg.NAME)
	fmt.Println(cfg.Age)
}

// 模拟投毒 加一行注释
type Func func(name string) string
type Callback func(name string) string
