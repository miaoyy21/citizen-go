package main

import (
	"citizen/tools"
	"log"
)

func main() {
	srcAssets := "/Users/miaojingyi/Documents/dev/go/src/citizen/assets"
	dstAssets := "/Users/miaojingyi/Documents/dev/flutter/citizen/assets"
	if err := tools.ParseAnimations(srcAssets, dstAssets); err != nil {
		log.Fatalf("%s", err.Error())
	}
	log.Println("解析动作帧动画成功 ...")
}
