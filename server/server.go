package server

import (
	"awesomeProject/cfg"
	"fmt"
)

func test() {
	fmt.Println(cfg.NAME)
	fmt.Println(cfg.Age)
}

type Func func(name string) string
