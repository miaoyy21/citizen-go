package main

import (
	"citizen/tools"
	"log"
)

func main() {
	if err := tools.ParseAnimations(); err != nil {
		log.Fatalf("%s", err.Error())
	}
	log.Println("解析动作帧动画成功 ...")
}
